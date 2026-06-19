package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Handler struct {
	uploadDir  string
	resultsDir string
.templates  *template.Template
}

func New(uploadDir, resultsDir string) *Handler {
	tmpl, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}
	return &Handler{
		uploadDir:  uploadDir,
		resultsDir: resultsDir,
		templates:  tmpl,
	}
}

func (h *Handler) render(w http.ResponseWriter, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.templates.ExecuteTemplate(w, name, data); err != nil {
		log.Printf("Template error (%s): %v", name, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, "/upload", http.StatusFound)
}

func (h *Handler) UploadPageHandler(w http.ResponseWriter, r *http.Request) {
	h.render(w, "upload.html", nil)
}

func (h *Handler) ResultsPageHandler(w http.ResponseWriter, r *http.Request) {
	h.render(w, "results.html", nil)
}

func (h *Handler) ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	h.render(w, "profile.html", nil)
}

func (h *Handler) ProgressPageHandler(w http.ResponseWriter, r *http.Request) {
	h.render(w, "progress.html", nil)
}

func (h *Handler) ExportPageHandler(w http.ResponseWriter, r *http.Request) {
	h.render(w, "export.html", nil)
}

func (h *Handler) resultsPath(id string) string {
	return filepath.Join(h.resultsDir, id+".json")
}
