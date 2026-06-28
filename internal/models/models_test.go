package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestVideo_JSON(t *testing.T) {
	now := time.Now()
	v := Video{
		ID:        "abc123",
		Filename:  "shot.mp4",
		Status:    "done",
		Progress:  100,
		CreatedAt: now,
		UpdatedAt: now,
	}

	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var decoded Video
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if decoded.ID != "abc123" {
		t.Fatalf("expected abc123, got %s", decoded.ID)
	}
	if decoded.Status != "done" {
		t.Fatalf("expected done, got %s", decoded.Status)
	}
}

func TestVideo_ScoreOmittedWhenNil(t *testing.T) {
	v := Video{ID: "v1", Score: nil}
	data, _ := json.Marshal(v)
	if string(data) == "" {
		t.Fatal("unexpected empty json")
	}
	if contains(string(data), `"score"`) {
		t.Fatalf("score should be omitted when nil, got: %s", data)
	}
}

func TestVideo_ScorePresentWhenSet(t *testing.T) {
	score := 85
	v := Video{ID: "v1", Score: &score}
	data, _ := json.Marshal(v)
	if !contains(string(data), `"score":85`) {
		t.Fatalf("expected score:85 in json, got: %s", data)
	}
}

func TestResult_JSON(t *testing.T) {
	r := Result{
		ID:       "r1",
		VideoID:  "v1",
		Filename: "test.mp4",
		Score:    72,
		Feedback: "Good form.",
		PoseData: []Point{{X: 10, Y: 20}, {X: 30, Y: 40}},
		Scores:   []int{80, 70, 75, 68},
	}

	data, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var decoded Result
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if decoded.Score != 72 {
		t.Fatalf("expected 72, got %d", decoded.Score)
	}
	if len(decoded.PoseData) != 2 {
		t.Fatalf("expected 2 points, got %d", len(decoded.PoseData))
	}
	if len(decoded.Scores) != 4 {
		t.Fatalf("expected 4 scores, got %d", len(decoded.Scores))
	}
}

func TestResult_ScoresOmittedWhenNil(t *testing.T) {
	r := Result{ID: "r1"}
	data, _ := json.Marshal(r)
	if contains(string(data), `"scores"`) {
		t.Fatalf("scores should be omitted when nil, got: %s", data)
	}
}

func TestPoint_JSON(t *testing.T) {
	p := Point{X: 55.5, Y: 33.3}
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	var decoded Point
	json.Unmarshal(data, &decoded)
	if decoded.X != 55.5 || decoded.Y != 33.3 {
		t.Fatalf("expected (55.5,33.3), got (%f,%f)", decoded.X, decoded.Y)
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(s) > 0 && containsSubstr(s, sub))
}

func containsSubstr(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
