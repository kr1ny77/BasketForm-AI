package models

import "time"

type Video struct {
	ID        string     `json:"id"`
	Filename  string     `json:"filename"`
	Status    string     `json:"status"`
	Progress  int        `json:"progress"`
	Score     *int       `json:"score,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Result struct {
	ID        string  `json:"id"`
	VideoID   string  `json:"video_id"`
	Filename  string  `json:"filename"`
	Score     int     `json:"score"`
	Feedback  string  `json:"feedback"`
	PoseData  []Point `json:"pose_data"`
	Scores    []int   `json:"scores"`
	CreatedAt string  `json:"created_at"`
}
