package handlers

import (
	"net/http"
	"strings"
	"testing"
)

func TestGenID_Unique(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		id := genID()
		if seen[id] {
			t.Fatalf("duplicate ID generated: %s", id)
		}
		seen[id] = true
	}
}

func TestGenID_Length(t *testing.T) {
	id := genID()
	if len(id) != 32 {
		t.Fatalf("expected 32 hex chars, got %d: %s", len(id), id)
	}
}

func TestGenID_HexOnly(t *testing.T) {
	id := genID()
	for _, c := range id {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			t.Fatalf("non-hex character in ID: %c", c)
		}
	}
}

func TestFileExt_MP4(t *testing.T) {
	if ext := fileExt("video.mp4"); ext != ".mp4" {
		t.Fatalf("expected .mp4, got %s", ext)
	}
}

func TestFileExt_MOV(t *testing.T) {
	if ext := fileExt("clip.mov"); ext != ".mov" {
		t.Fatalf("expected .mov, got %s", ext)
	}
}

func TestFileExt_AVI(t *testing.T) {
	if ext := fileExt("recording.avi"); ext != ".avi" {
		t.Fatalf("expected .avi, got %s", ext)
	}
}

func TestFileExt_NoExtension_DefaultsToMP4(t *testing.T) {
	if ext := fileExt("noext"); ext != ".mp4" {
		t.Fatalf("expected .mp4 default, got %s", ext)
	}
}

func TestFileExt_Unsupported(t *testing.T) {
	ext := fileExt("video.mkv")
	if ext != ".mkv" {
		t.Fatalf("expected .mkv, got %s", ext)
	}
}

func TestExtractID_Valid(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/status/abc123", nil)
	id := extractID(req, "/api/status/")
	if id != "abc123" {
		t.Fatalf("expected abc123, got %s", id)
	}
}

func TestExtractID_Empty(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/status/", nil)
	id := extractID(req, "/api/status/")
	if id != "" {
		t.Fatalf("expected empty, got %s", id)
	}
}

func TestExtractID_NestedPath(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/status/abc/def", nil)
	id := extractID(req, "/api/status/")
	if id != "" {
		t.Fatalf("expected empty for nested path, got %s", id)
	}
}

func TestExtractID_VeryLong(t *testing.T) {
	long := strings.Repeat("a", 256)
	req, _ := http.NewRequest("GET", "/api/status/"+long, nil)
	id := extractID(req, "/api/status/")
	if id != long {
		t.Fatalf("expected long id, got %s", id)
	}
}
