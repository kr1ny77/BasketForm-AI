package services

import (
	"testing"
)

func TestExtOf(t *testing.T) {
	tests := []struct{ input, want string }{
		{"video.mp4", ".mp4"},
		{"clip.MOV", ".mov"},
		{"noext", ".mp4"},
		{"path/to/file.avi", ".avi"},
	}
	for _, tt := range tests {
		got := extOf(tt.input)
		if got != tt.want {
			t.Errorf("extOf(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestProcessor_VideoNotFound_SkipsGracefully(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	p := NewProcessor(s)

	p.ProcessVideo("nonexistent")
}
