package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
	"github.com/kr1ny77/BasketForm-AI/internal/services"
)

type ShareHandler struct {
	storage *services.Storage
}

func NewShareHandler(storage *services.Storage) *ShareHandler {
	return &ShareHandler{storage: storage}
}

func (h *ShareHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/share/result", h.ShareResult)
	mux.HandleFunc("/api/results/shared-with-me", h.SharedWithMe)
}

func (h *ShareHandler) ShareResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID := r.Context().Value("user_id").(string)

	var req struct {
		ResultID string `json:"result_id"`
		FriendID string `json:"friend_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.ResultID == "" || req.FriendID == "" {
		writeError(w, http.StatusBadRequest, "result_id and friend_id are required")
		return
	}

	if _, ok := h.storage.GetUserByID(req.FriendID); !ok {
		writeError(w, http.StatusNotFound, "friend not found")
		return
	}

	sr := &models.SharedResult{
		ID:         genID(),
		ResultID:   req.ResultID,
		FromUserID: userID,
		ToUserID:   req.FriendID,
		CreatedAt:  time.Now(),
	}

	if err := h.storage.SaveSharedResult(sr); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to share result")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"id": sr.ID})
}

func (h *ShareHandler) SharedWithMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID := r.Context().Value("user_id").(string)
	shared := h.storage.GetSharedWithMe(userID)

	type sharedResp struct {
		ID           string `json:"id"`
		ResultID     string `json:"result_id"`
		FromNickname string `json:"from_nickname"`
		Score        int    `json:"score"`
		VideoID      string `json:"video_id"`
		CreatedAt    string `json:"created_at"`
	}
	var results []sharedResp
	for _, s := range shared {
		result, err := h.storage.LoadResultByID(s.ResultID)
		if err != nil {
			continue
		}
		fromNick := ""
		if u, ok := h.storage.GetUserByID(s.FromUserID); ok {
			fromNick = u.Nickname
		}
		results = append(results, sharedResp{
			ID:           s.ID,
			ResultID:     s.ResultID,
			FromNickname: fromNick,
			Score:        result.Score,
			VideoID:      result.VideoID,
			CreatedAt:    s.CreatedAt.Format(time.RFC3339),
		})
	}
	if results == nil {
		results = []sharedResp{}
	}
	writeJSON(w, http.StatusOK, results)
}
