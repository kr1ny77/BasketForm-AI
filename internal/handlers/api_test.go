package handlers

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
	"github.com/kr1ny77/BasketForm-AI/internal/services"
)

func setupTest(t *testing.T) (*APIHandler, string, string) {
	t.Helper()
	upload := t.TempDir()
	results := t.TempDir()
	storage := services.NewStorage(upload, results)
	processor := services.NewProcessor(storage)
	api := NewAPI(storage, processor, upload)
	return api, upload, results
}

func newMux(api *APIHandler) *http.ServeMux {
	mux := http.NewServeMux()
	api.Register(mux)
	return mux
}

func TestAPI_Upload_Success(t *testing.T) {
	api, upload, _ := setupTest(t)
	mux := newMux(api)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("video", "test.mp4")
	part.Write([]byte("fake video content"))
	writer.Close()

	req := httptest.NewRequest("POST", "/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", w.Code, w.Body.String())
	}

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["id"] == "" {
		t.Fatal("expected id in response")
	}
	if resp["filename"] != "test.mp4" {
		t.Fatalf("expected test.mp4, got %s", resp["filename"])
	}
	if resp["status"] != "uploaded" {
		t.Fatalf("expected uploaded, got %s", resp["status"])
	}

	// Verify file was saved
	entries, _ := os.ReadDir(upload)
	if len(entries) != 1 {
		t.Fatalf("expected 1 file in upload dir, got %d", len(entries))
	}
	if !strings.HasSuffix(entries[0].Name(), ".mp4") {
		t.Fatalf("expected .mp4 file, got %s", entries[0].Name())
	}
}

func TestAPI_Upload_UnsupportedFormat(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("video", "video.mkv")
	part.Write([]byte("fake content"))
	writer.Close()

	req := httptest.NewRequest("POST", "/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", w.Code, w.Body.String())
	}
}

func TestAPI_Upload_NoFile(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req := httptest.NewRequest("POST", "/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", w.Code, w.Body.String())
	}
}

func TestAPI_Upload_WrongMethod(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	req := httptest.NewRequest("GET", "/api/upload", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", w.Code)
	}
}

func TestAPI_Status_Found(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	api.storage.CreateVideo("vid1", "test.mp4")

	req := httptest.NewRequest("GET", "/api/status/vid1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["id"] != "vid1" {
		t.Fatalf("expected vid1, got %v", resp["id"])
	}
	if resp["status"] != "uploaded" {
		t.Fatalf("expected uploaded, got %v", resp["status"])
	}
}

func TestAPI_Status_NotFound(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	req := httptest.NewRequest("GET", "/api/status/missing", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestAPI_Status_MissingID(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	req := httptest.NewRequest("GET", "/api/status/", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestAPI_Result_Found(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	api.storage.CreateVideo("vid2", "test.mp4")
	api.storage.SaveResult(&models.Result{
		ID:       "vid2",
		VideoID:  "vid2",
		Filename: "test.mp4",
		Score:    85,
		Feedback: "Great shot!",
		PoseData: []models.Point{{X: 50, Y: 30}},
	})

	req := httptest.NewRequest("GET", "/api/result/vid2", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["score"] != float64(85) {
		t.Fatalf("expected score 85, got %v", resp["score"])
	}
	if resp["feedback"] != "Great shot!" {
		t.Fatalf("expected feedback 'Great shot!', got %v", resp["feedback"])
	}
}

func TestAPI_Result_NotFound(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	req := httptest.NewRequest("GET", "/api/result/missing", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestAPI_Videos_Empty(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	req := httptest.NewRequest("GET", "/api/videos", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp []any
	json.Unmarshal(w.Body.Bytes(), &resp)
	if len(resp) != 0 {
		t.Fatalf("expected empty list, got %d items", len(resp))
	}
}

func TestAPI_Videos_WithEntries(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	api.storage.CreateVideo("v1", "a.mp4")
	api.storage.CreateVideo("v2", "b.mp4")

	req := httptest.NewRequest("GET", "/api/videos", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp []map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	if len(resp) != 2 {
		t.Fatalf("expected 2 videos, got %d", len(resp))
	}
}

func TestAPI_Upload_MOV(t *testing.T) {
	api, upload, _ := setupTest(t)
	mux := newMux(api)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("video", "clip.mov")
	part.Write([]byte("fake mov content"))
	writer.Close()

	req := httptest.NewRequest("POST", "/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", w.Code, w.Body.String())
	}

	entries, _ := os.ReadDir(upload)
	found := false
	for _, e := range entries {
		if filepath.Ext(e.Name()) == ".mov" {
			found = true
		}
	}
	if !found {
		t.Fatal("expected .mov file in upload dir")
	}
}

func TestAPI_Upload_EmptyFile(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("video", "empty.mp4")
	part.Write([]byte{})
	writer.Close()

	req := httptest.NewRequest("POST", "/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	// Empty file should still be accepted (0-byte upload)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 for empty file, got %d: %s", w.Code, w.Body.String())
	}
}

func TestAPI_Status_AfterUpload(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	// Upload
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("video", "test.mp4")
	part.Write([]byte("content"))
	writer.Close()

	req := httptest.NewRequest("POST", "/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	var uploadResp map[string]string
	json.Unmarshal(w.Body.Bytes(), &uploadResp)
	id := uploadResp["id"]

	// Check status
	req2 := httptest.NewRequest("GET", "/api/status/"+id, nil)
	w2 := httptest.NewRecorder()
	mux.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w2.Code)
	}

	var statusResp map[string]any
	json.Unmarshal(w2.Body.Bytes(), &statusResp)
	status := statusResp["status"].(string)
	if status != "uploaded" && status != "processing" {
		t.Fatalf("expected uploaded or processing, got %s", status)
	}
}

func TestAPI_VideosJSONFormat(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	api.storage.CreateVideo("v1", "a.mp4")

	req := httptest.NewRequest("GET", "/api/videos", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	ct := w.Header().Get("Content-Type")
	if !strings.Contains(ct, "application/json") {
		t.Fatalf("expected JSON content type, got %s", ct)
	}
}

func TestAPI_ErrorJSONFormat(t *testing.T) {
	api, _, _ := setupTest(t)
	mux := newMux(api)

	req := httptest.NewRequest("GET", "/api/status/missing", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	var resp map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("error response is not valid JSON: %v", err)
	}
	if resp["error"] == "" {
		t.Fatal("expected error message in response")
	}
}
