package services

import (
	"testing"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
)

func TestClamp(t *testing.T) {
	tests := []struct {
		v, lo, hi, want int
	}{
		{5, 0, 10, 5},
		{-1, 0, 10, 0},
		{15, 0, 10, 10},
		{0, 0, 100, 0},
		{100, 0, 100, 100},
		{50, 40, 90, 50},
	}
	for _, tt := range tests {
		got := clamp(tt.v, tt.lo, tt.hi)
		if got != tt.want {
			t.Errorf("clamp(%d,%d,%d) = %d, want %d", tt.v, tt.lo, tt.hi, got, tt.want)
		}
	}
}

func TestGenerateMockPose_PointsCount(t *testing.T) {
	pts := generateMockPose()
	if len(pts) != 12 {
		t.Fatalf("expected 12 pose points, got %d", len(pts))
	}
}

func TestGenerateMockPose_Range(t *testing.T) {
	pts := generateMockPose()
	for i, p := range pts {
		if p.X < 0 || p.X > 100 {
			t.Fatalf("point %d X out of range: %f", i, p.X)
		}
		if p.Y < 0 || p.Y > 100 {
			t.Fatalf("point %d Y out of range: %f", i, p.Y)
		}
	}
}

func TestGenerateMockPose_Varies(t *testing.T) {
	a := generateMockPose()
	b := generateMockPose()
	same := true
	for i := range a {
		if a[i].X != b[i].X || a[i].Y != b[i].Y {
			same = false
			break
		}
	}
	if same {
		t.Fatal("two calls to generateMockPose returned identical results (very unlikely)")
	}
}

func TestFeedbacks_NonEmpty(t *testing.T) {
	if len(feedbacks) == 0 {
		t.Fatal("feedbacks slice is empty")
	}
	for i, f := range feedbacks {
		if f == "" {
			t.Fatalf("feedback %d is empty", i)
		}
	}
}

func TestProcessor_ProcessVideo_Completes(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	p := NewProcessor(s)

	s.CreateVideo("test1", "test.mp4")

	// Run processing in a goroutine (it sleeps ~5s)
	done := make(chan bool)
	go func() {
		p.ProcessVideo("test1")
		done <- true
	}()

	<-done

	v, ok := s.GetVideo("test1")
	if !ok {
		t.Fatal("video not found after processing")
	}
	if v.Status != "done" {
		t.Fatalf("expected status done, got %s", v.Status)
	}
	if v.Progress != 100 {
		t.Fatalf("expected progress 100, got %d", v.Progress)
	}
	if v.Score == nil {
		t.Fatal("expected score to be set")
	}
	if *v.Score < 40 || *v.Score > 90 {
		t.Fatalf("score out of expected range: %d", *v.Score)
	}

	result, err := s.LoadResult("test1")
	if err != nil {
		t.Fatalf("LoadResult error: %v", err)
	}
	if result.Feedback == "" {
		t.Fatal("feedback is empty")
	}
	if len(result.PoseData) != 12 {
		t.Fatalf("expected 12 pose points, got %d", len(result.PoseData))
	}
	if len(result.Scores) != 4 {
		t.Fatalf("expected 4 scores, got %d", len(result.Scores))
	}
}

func TestProcessor_VideoNotFound_SkipsGracefully(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	p := NewProcessor(s)

	// Should not panic
	p.ProcessVideo("nonexistent")
}
