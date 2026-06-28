package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/kr1ny77/BasketForm-AI/internal/services"
)

type Middleware struct {
	auth *services.AuthService
}

func NewMiddleware(auth *services.AuthService) *Middleware {
	return &Middleware{auth: auth}
}

func (m *Middleware) AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/login" || path == "/register" || path == "/api/login" || path == "/api/register" {
			next.ServeHTTP(w, r)
			return
		}
		if strings.HasPrefix(path, "/static/") {
			next.ServeHTTP(w, r)
			return
		}

		token, err := r.Cookie("token")
		if err != nil || token.Value == "" {
			if strings.HasPrefix(path, "/api/") {
				writeError(w, http.StatusUnauthorized, "authentication required")
			} else {
				http.Redirect(w, r, "/login", http.StatusFound)
			}
			return
		}

		userID, err := m.auth.ValidateToken(token.Value)
		if err != nil {
			if strings.HasPrefix(path, "/api/") {
				writeError(w, http.StatusUnauthorized, "invalid token")
			} else {
				http.Redirect(w, r, "/login", http.StatusFound)
			}
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
