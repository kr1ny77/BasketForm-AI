package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kr1ny77/BasketForm-AI/internal/services"
)

type AuthHandler struct {
	auth    *services.AuthService
	storage *services.Storage
}

func NewAuthHandler(auth *services.AuthService, storage *services.Storage) *AuthHandler {
	return &AuthHandler{auth: auth, storage: storage}
}

func (h *AuthHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/register", h.RegisterHandler)
	mux.HandleFunc("/api/login", h.LoginHandler)
	mux.HandleFunc("/api/profile", h.ProfileHandler)
	mux.HandleFunc("/api/profile/update", h.ProfileUpdateHandler)
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Email == "" || req.Nickname == "" || req.Password == "" {
		writeError(w, http.StatusBadRequest, "email, nickname and password are required")
		return
	}
	if len(req.Password) < 6 {
		writeError(w, http.StatusBadRequest, "password must be at least 6 characters")
		return
	}

	user, err := h.auth.Register(req.Email, req.Nickname, req.Password)
	if err != nil {
		switch err {
		case services.ErrEmailTaken:
			writeError(w, http.StatusConflict, "email already taken")
		case services.ErrNicknameTaken:
			writeError(w, http.StatusConflict, "nickname already taken")
		default:
			writeError(w, http.StatusInternalServerError, "internal error")
		}
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{
		"id":       user.ID,
		"email":    user.Email,
		"nickname": user.Nickname,
	})
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	token, user, err := h.auth.Login(req.Email, req.Password)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400,
	})

	writeJSON(w, http.StatusOK, map[string]string{
		"id":       user.ID,
		"email":    user.Email,
		"nickname": user.Nickname,
		"token":    token,
	})
}

func (h *AuthHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID := r.Context().Value("user_id").(string)
	user, ok := h.storage.GetUserByID(userID)
	if !ok {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"id":       user.ID,
		"email":    user.Email,
		"nickname": user.Nickname,
	})
}

func (h *AuthHandler) ProfileUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID := r.Context().Value("user_id").(string)

	var req struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		writeError(w, http.StatusBadRequest, "old_password and new_password are required")
		return
	}
	if len(req.NewPassword) < 6 {
		writeError(w, http.StatusBadRequest, "new password must be at least 6 characters")
		return
	}

	if err := h.auth.ChangePassword(userID, req.OldPassword, req.NewPassword); err != nil {
		switch err {
		case services.ErrInvalidCredentials:
			writeError(w, http.StatusUnauthorized, "incorrect current password")
		case services.ErrUserNotFound:
			writeError(w, http.StatusNotFound, "user not found")
		default:
			writeError(w, http.StatusInternalServerError, "internal error")
		}
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "password updated"})
}
