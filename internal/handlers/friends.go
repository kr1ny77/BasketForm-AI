package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
	"github.com/kr1ny77/BasketForm-AI/internal/services"
)

type FriendsHandler struct {
	storage *services.Storage
}

func NewFriendsHandler(storage *services.Storage) *FriendsHandler {
	return &FriendsHandler{storage: storage}
}

func (h *FriendsHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/friends/search", h.SearchUsers)
	mux.HandleFunc("/api/friends/request", h.SendRequest)
	mux.HandleFunc("/api/friends/requests", h.ListRequests)
	mux.HandleFunc("/api/friends/accept", h.AcceptRequest)
	mux.HandleFunc("/api/friends/reject", h.RejectRequest)
	mux.HandleFunc("/api/friends", h.ListFriends)
}

func (h *FriendsHandler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	query := r.URL.Query().Get("q")
	if query == "" {
		writeJSON(w, http.StatusOK, []interface{}{})
		return
	}

	userID := r.Context().Value("user_id").(string)
	users := h.storage.SearchUsersByNickname(query)

	type userResp struct {
		ID       string `json:"id"`
		Nickname string `json:"nickname"`
	}
	var results []userResp
	for _, u := range users {
		if u.ID == userID {
			continue
		}
		results = append(results, userResp{ID: u.ID, Nickname: u.Nickname})
	}
	if results == nil {
		results = []userResp{}
	}
	writeJSON(w, http.StatusOK, results)
}

func (h *FriendsHandler) SendRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID := r.Context().Value("user_id").(string)

	var req struct {
		ToUserID string `json:"to_user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.ToUserID == "" || req.ToUserID == userID {
		writeError(w, http.StatusBadRequest, "invalid target user")
		return
	}

	if _, ok := h.storage.GetUserByID(req.ToUserID); !ok {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	if h.storage.HasFriendship(userID, req.ToUserID) {
		writeError(w, http.StatusConflict, "already friends")
		return
	}

	if h.storage.HasPendingRequest(userID, req.ToUserID) {
		writeError(w, http.StatusConflict, "request already sent")
		return
	}

	fr := &models.FriendRequest{
		ID:         genID(),
		FromUserID: userID,
		ToUserID:   req.ToUserID,
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	if err := h.storage.SaveFriendRequest(fr); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to send request")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"id": fr.ID, "status": "pending"})
}

func (h *FriendsHandler) ListRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID := r.Context().Value("user_id").(string)
	requests := h.storage.GetFriendRequestsByToUser(userID)

	type reqResp struct {
		ID         string `json:"id"`
		FromUserID string `json:"from_user_id"`
		FromNick   string `json:"from_nickname"`
		Status     string `json:"status"`
	}
	var results []reqResp
	for _, fr := range requests {
		nick := ""
		if u, ok := h.storage.GetUserByID(fr.FromUserID); ok {
			nick = u.Nickname
		}
		results = append(results, reqResp{
			ID:         fr.ID,
			FromUserID: fr.FromUserID,
			FromNick:   nick,
			Status:     fr.Status,
		})
	}
	if results == nil {
		results = []reqResp{}
	}
	writeJSON(w, http.StatusOK, results)
}

func (h *FriendsHandler) AcceptRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID := r.Context().Value("user_id").(string)

	var req struct {
		RequestID string `json:"request_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	fr, ok := h.storage.GetFriendRequestByID(req.RequestID)
	if !ok || fr.ToUserID != userID || fr.Status != "pending" {
		writeError(w, http.StatusNotFound, "request not found")
		return
	}

	fr.Status = "accepted"
	if err := h.storage.UpdateFriendRequest(fr); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to accept request")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "accepted"})
}

func (h *FriendsHandler) RejectRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID := r.Context().Value("user_id").(string)

	var req struct {
		RequestID string `json:"request_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	fr, ok := h.storage.GetFriendRequestByID(req.RequestID)
	if !ok || fr.ToUserID != userID || fr.Status != "pending" {
		writeError(w, http.StatusNotFound, "request not found")
		return
	}

	fr.Status = "rejected"
	if err := h.storage.UpdateFriendRequest(fr); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to reject request")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "rejected"})
}

func (h *FriendsHandler) ListFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID := r.Context().Value("user_id").(string)
	friends := h.storage.GetFriends(userID)

	type friendResp struct {
		ID       string `json:"id"`
		Nickname string `json:"nickname"`
	}
	var results []friendResp
	for _, f := range friends {
		results = append(results, friendResp{ID: f.ID, Nickname: f.Nickname})
	}
	if results == nil {
		results = []friendResp{}
	}
	writeJSON(w, http.StatusOK, results)
}
