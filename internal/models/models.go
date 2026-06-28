package models

import "time"

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Nickname     string    `json:"nickname"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

type Video struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Filename  string    `json:"filename"`
	Lang      string    `json:"lang"`
	Status    string    `json:"status"`
	Progress  int       `json:"progress"`
	Score     *int      `json:"score,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type PhaseScore struct {
	PhaseName string  `json:"phase_name"`
	Score     float64 `json:"score"`
}

type PhaseFeedback struct {
	PhaseName string `json:"phase_name"`
	Feedback  string `json:"feedback"`
}

type Result struct {
	ID        string          `json:"id"`
	UserID    string          `json:"user_id"`
	VideoID   string          `json:"video_id"`
	Filename  string          `json:"filename"`
	Score     int             `json:"score"`
	Feedback  string          `json:"feedback"`
	PoseData  []Point         `json:"pose_data"`
	Scores    []int           `json:"scores,omitempty"`
	Phases    []PhaseScore    `json:"phases,omitempty"`
	Feedbacks []PhaseFeedback `json:"feedbacks,omitempty"`
	CreatedAt string          `json:"created_at"`
}

type FriendRequest struct {
	ID         string    `json:"id"`
	FromUserID string    `json:"from_user_id"`
	ToUserID   string    `json:"to_user_id"`
	Status     string    `json:"status"` // "pending", "accepted", "rejected"
	CreatedAt  time.Time `json:"created_at"`
}

type Friend struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	FriendID  string    `json:"friend_id"`
	CreatedAt time.Time `json:"created_at"`
}

type SharedResult struct {
	ID         string    `json:"id"`
	ResultID   string    `json:"result_id"`
	FromUserID string    `json:"from_user_id"`
	ToUserID   string    `json:"to_user_id"`
	CreatedAt  time.Time `json:"created_at"`
}
