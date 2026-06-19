package services

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
)

func tempDirs(t *testing.T) (string, string) {
	t.Helper()
	upload := t.TempDir()
	results := t.TempDir()
	return upload, results
}

func TestStorage_CreateAndGetVideo(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	v := s.CreateVideo("id1", "test.mp4")
	if v.ID != "id1" {
		t.Fatalf("expected id1, got %s", v.ID)
	}
	if v.Filename != "test.mp4" {
		t.Fatalf("expected test.mp4, got %s", v.Filename)
	}
	if v.Status != "uploaded" {
		t.Fatalf("expected uploaded, got %s", v.Status)
	}
	if v.Progress != 0 {
		t.Fatalf("expected progress 0, got %d", v.Progress)
	}

	got, ok := s.GetVideo("id1")
	if !ok || got.ID != "id1" {
		t.Fatalf("GetVideo failed for id1")
	}
}

func TestStorage_GetVideo_NotFound(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	_, ok := s.GetVideo("nonexistent")
	if ok {
		t.Fatal("expected not found")
	}
}

func TestStorage_GetAllVideos_Sorted(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	s.CreateVideo("c", "c.mp4")
	s.CreateVideo("a", "a.mp4")
	s.CreateVideo("b", "b.mp4")

	list := s.GetAllVideos()
	if len(list) != 3 {
		t.Fatalf("expected 3 videos, got %d", len(list))
	}
	// Newest first (c created last)
	if list[0].ID != "c" {
		t.Fatalf("expected c first, got %s", list[0].ID)
	}
}

func TestStorage_GetAllVideos_Empty(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	list := s.GetAllVideos()
	if list != nil {
		t.Fatalf("expected nil, got %v", list)
	}
}

func TestStorage_UpdateStatus(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	s.CreateVideo("v1", "test.mp4")
	s.UpdateStatus("v1", "processing", 50)

	v, _ := s.GetVideo("v1")
	if v.Status != "processing" {
		t.Fatalf("expected processing, got %s", v.Status)
	}
	if v.Progress != 50 {
		t.Fatalf("expected 50, got %d", v.Progress)
	}
}

func TestStorage_UpdateStatus_NotFound(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	// Should not panic
	s.UpdateStatus("missing", "done", 100)
}

func TestStorage_SetScore(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	s.CreateVideo("v1", "test.mp4")
	s.SetScore("v1", 85)

	v, _ := s.GetVideo("v1")
	if v.Score == nil || *v.Score != 85 {
		t.Fatalf("expected score 85, got %v", v.Score)
	}
}

func TestStorage_SaveAndLoadResult(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	r := &models.Result{
		ID:       "r1",
		VideoID:  "v1",
		Filename: "test.mp4",
		Score:    75,
		Feedback: "Good form",
		PoseData: []models.Point{{X: 10, Y: 20}},
		Scores:   []int{80, 70, 75, 72},
	}

	if err := s.SaveResult(r); err != nil {
		t.Fatalf("SaveResult error: %v", err)
	}

	loaded, err := s.LoadResult("v1")
	if err != nil {
		t.Fatalf("LoadResult error: %v", err)
	}
	if loaded.Score != 75 {
		t.Fatalf("expected score 75, got %d", loaded.Score)
	}
	if loaded.Feedback != "Good form" {
		t.Fatalf("expected feedback 'Good form', got %s", loaded.Feedback)
	}
	if len(loaded.PoseData) != 1 {
		t.Fatalf("expected 1 pose point, got %d", len(loaded.PoseData))
	}
}

func TestStorage_LoadResult_NotFound(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	_, err := s.LoadResult("nonexistent")
	if err == nil {
		t.Fatal("expected error for missing result")
	}
}

func TestStorage_UploadPath(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	path := s.UploadPath("abc", ".mp4")
	expected := filepath.Join(upload, "abc.mp4")
	if path != expected {
		t.Fatalf("expected %s, got %s", expected, path)
	}
}

func TestStorage_ResultFileExists(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	r := &models.Result{VideoID: "v2", Score: 90}
	s.SaveResult(r)

	path := filepath.Join(results, "v2.json")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatal("result JSON file not created")
	}
}
