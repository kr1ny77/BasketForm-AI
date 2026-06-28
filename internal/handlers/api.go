package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
	"github.com/kr1ny77/BasketForm-AI/internal/services"
)

type APIHandler struct {
	storage   *services.Storage
	processor *services.Processor
	uploadDir string
}

func NewAPI(storage *services.Storage, processor *services.Processor, uploadDir string) *APIHandler {
	return &APIHandler{
		storage:   storage,
		processor: processor,
		uploadDir: uploadDir,
	}
}

func (a *APIHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/upload", a.Upload)
	mux.HandleFunc("/api/status/", a.Status)
	mux.HandleFunc("/api/result/", a.Result)
	mux.HandleFunc("/api/videos", a.Videos)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func extractID(r *http.Request, prefix string) string {
	id := strings.TrimPrefix(r.URL.Path, prefix)
	if id == "" || strings.Contains(id, "/") {
		return ""
	}
	return id
}

func genID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func fileExt(filename string) string {
	ext := filepath.Ext(filename)
	if ext == "" {
		ext = ".mp4"
	}
	return ext
}

func (a *APIHandler) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID := r.Context().Value("user_id").(string)

	if err := r.ParseMultipartForm(100 << 20); err != nil {
		writeError(w, http.StatusBadRequest, "invalid form data")
		return
	}

	file, header, err := r.FormFile("video")
	if err != nil {
		writeError(w, http.StatusBadRequest, "missing video file")
		return
	}
	defer file.Close()

	allowed := map[string]bool{
		".mp4": true, ".mov": true, ".avi": true,
	}
	ext := strings.ToLower(fileExt(header.Filename))
	if !allowed[ext] {
		writeError(w, http.StatusBadRequest, "unsupported format: use MP4, MOV, or AVI")
		return
	}

	id := genID()
	dstPath := a.uploadDir + "/" + id + ext

	dst, err := os.Create(dstPath)
	if err != nil {
		log.Printf("Failed to create file %s: %v", dstPath, err)
		writeError(w, http.StatusInternalServerError, "failed to save file")
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		log.Printf("Failed to write file %s: %v", dstPath, err)
		writeError(w, http.StatusInternalServerError, "failed to save file")
		return
	}

	lang := r.FormValue("lang")
	if lang == "" {
		lang = "en"
	}

	video := a.storage.CreateVideo(id, header.Filename, userID)
	video.Lang = lang
	a.storage.SaveVideoJSON(video)

	go a.processor.ProcessVideo(id)

	writeJSON(w, http.StatusCreated, map[string]string{
		"id":       video.ID,
		"filename": video.Filename,
		"status":   video.Status,
	})
}

func (a *APIHandler) Status(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := extractID(r, "/api/status/")
	if id == "" {
		writeError(w, http.StatusBadRequest, "missing video id")
		return
	}

	video, ok := a.storage.GetVideo(id)
	if !ok {
		writeError(w, http.StatusNotFound, "video not found")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"id":       video.ID,
		"status":   video.Status,
		"progress": video.Progress,
	})
}

func (a *APIHandler) Result(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := extractID(r, "/api/result/")
	if id == "" {
		writeError(w, http.StatusBadRequest, "missing video id")
		return
	}

	result, err := a.storage.LoadResult(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "result not found")
		return
	}

	outputVideoURL := ""
	video, ok := a.storage.GetVideo(id)
	if ok && video.UserID != "" {
		_, err := os.Stat(filepath.Join("results", video.UserID, id, "output.mp4"))
		if err == nil {
			outputVideoURL = "/results-files/" + video.UserID + "/" + id + "/output.mp4"
		}
	}

	type resultResp struct {
		ID             string              `json:"id"`
		UserID         string              `json:"user_id"`
		VideoID        string              `json:"video_id"`
		Filename       string              `json:"filename"`
		Score          int                 `json:"score"`
		Feedback       string              `json:"feedback"`
		PoseData       []models.Point      `json:"pose_data"`
		Scores         []int               `json:"scores"`
		Phases         []models.PhaseScore `json:"phases"`
		OutputVideoURL string              `json:"output_video_url"`
		CreatedAt      string              `json:"created_at"`
	}
	writeJSON(w, http.StatusOK, resultResp{
		ID:             result.ID,
		UserID:         result.UserID,
		VideoID:        result.VideoID,
		Filename:       result.Filename,
		Score:          result.Score,
		Feedback:       result.Feedback,
		PoseData:       result.PoseData,
		Scores:         result.Scores,
		Phases:         result.Phases,
		OutputVideoURL: outputVideoURL,
		CreatedAt:      result.CreatedAt,
	})
}

func (a *APIHandler) Videos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID := r.Context().Value("user_id").(string)
	videos := a.storage.GetVideosByUserID(userID)
	if videos == nil {
		videos = []*models.Video{}
	}
	writeJSON(w, http.StatusOK, videos)
}
