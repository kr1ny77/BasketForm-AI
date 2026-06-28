package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Handler struct {
	uploadDir  string
	resultsDir string
}

func New(uploadDir, resultsDir string) *Handler {
	return &Handler{
		uploadDir:  uploadDir,
		resultsDir: resultsDir,
	}
}

func serveFile(w http.ResponseWriter, path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Failed to read %s: %v", path, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, "/upload", http.StatusFound)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "web/templates/login.html")
}

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "web/templates/signup.html")
}

func (h *Handler) UploadPageHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "web/templates/upload.html")
}

func (h *Handler) ResultsPageHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "web/templates/results.html")
}

func (h *Handler) ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "web/templates/profile.html")
}

func (h *Handler) ProgressPageHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "web/templates/progress.html")
}

func (h *Handler) ExportPageHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "web/templates/export.html")
}

func (h *Handler) FriendsPageHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "web/templates/friends.html")
}

func (h *Handler) SharedResultsPageHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "web/templates/shared.html")
}

func (h *Handler) ResultsFileServer() http.Handler {
	abs, err := filepath.Abs(h.resultsDir)
	if err != nil {
		abs = h.resultsDir
	}
	return http.StripPrefix("/results/", http.FileServer(http.Dir(abs)))
}
