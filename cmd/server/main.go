package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kr1ny77/BasketForm-AI/internal/handlers"
	"github.com/kr1ny77/BasketForm-AI/internal/services"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	uploadDir := os.Getenv("UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "uploads"
	}

	resultsDir := os.Getenv("RESULTS_DIR")
	if resultsDir == "" {
		resultsDir = "results"
	}

	ensureDir(uploadDir)
	ensureDir(resultsDir)

	storage := services.NewStorage(uploadDir, resultsDir)
	processor := services.NewProcessor(storage)
	auth := services.NewAuthService(storage)

	h := handlers.New(uploadDir, resultsDir)
	api := handlers.NewAPI(storage, processor, uploadDir)
	authHandler := handlers.NewAuthHandler(auth, storage)
	friendsHandler := handlers.NewFriendsHandler(storage)
	shareHandler := handlers.NewShareHandler(storage)
	middleware := handlers.NewMiddleware(auth)

	mux := http.NewServeMux()

	// Auth pages (no auth required)
	mux.HandleFunc("/login", h.LoginHandler)
	mux.HandleFunc("/register", h.RegisterHandler)

	// Auth API (no auth required for login/register)
	authHandler.Register(mux)

	// Protected page handlers
	mux.HandleFunc("/", h.IndexHandler)
	mux.HandleFunc("/upload", h.UploadPageHandler)
	mux.HandleFunc("/results", h.ResultsPageHandler)
	mux.HandleFunc("/profile", h.ProfilePageHandler)
	mux.HandleFunc("/progress", h.ProgressPageHandler)
	mux.HandleFunc("/export", h.ExportPageHandler)
	mux.HandleFunc("/friends", h.FriendsPageHandler)
	mux.HandleFunc("/shared", h.SharedResultsPageHandler)

	// Protected API routes
	api.Register(mux)
	friendsHandler.Register(mux)
	shareHandler.Register(mux)

	// Static files (no auth required)
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Results files (no auth for now - videos need to be accessible)
	mux.Handle("/results-files/", http.StripPrefix("/results-files/", h.ResultsFileServer()))

	wrappedMux := middleware.AuthRequired(mux)

	addr := fmt.Sprintf(":%s", port)
	log.Printf("BasketForm-AI server starting on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func ensureDir(dir string) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("Failed to resolve path %s: %v", dir, err)
	}
	if _, err := os.Stat(abs); os.IsNotExist(err) {
		if err := os.MkdirAll(abs, 0o755); err != nil {
			log.Fatalf("Failed to create directory %s: %v", abs, err)
		}
		log.Printf("Created directory: %s", abs)
	}
}
