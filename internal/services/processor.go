package services

import (
	"log"
	"math/rand"
	"time"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
)

var feedbacks = []string{
	"Good release point. Work on keeping your elbow aligned with the basket.",
	"Strong stance. Try to extend your follow-through higher for better arc.",
	"Nice arm angle. Focus on squaring your shoulders to the basket before release.",
	"Consistent base. Keep your shooting hand under the ball for a cleaner release.",
	"Good rhythm. Work on a quicker release without sacrificing form.",
	"Solid footwork. Try to jump straight up rather than drifting forward.",
}

type Processor struct {
	storage *Storage
}

func NewProcessor(storage *Storage) *Processor {
	return &Processor{storage: storage}
}

func (p *Processor) ProcessVideo(id string) {
	p.storage.UpdateStatus(id, "processing", 0)
	log.Printf("Processing video %s", id)

	// Simulate incremental progress
	for progress := 10; progress <= 90; progress += 20 {
		time.Sleep(time.Second)
		p.storage.UpdateStatus(id, "processing", progress)
	}

	// Simulate final processing
	time.Sleep(2 * time.Second)

	score := 40 + rand.Intn(51) // 40–90
	feedback := feedbacks[rand.Intn(len(feedbacks))]

	scores := []int{
		clamp(score-5+rand.Intn(11), 0, 100),
		clamp(score-10+rand.Intn(21), 0, 100),
		clamp(score-3+rand.Intn(7), 0, 100),
		clamp(score-8+rand.Intn(17), 0, 100),
	}

	pose := generateMockPose()

	video, ok := p.storage.GetVideo(id)
	if !ok {
		log.Printf("Video %s not found during processing", id)
		return
	}

	result := &models.Result{
		ID:        id,
		VideoID:   id,
		Filename:  video.Filename,
		Score:     score,
		Feedback:  feedback,
		PoseData:  pose,
		Scores:    scores,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if err := p.storage.SaveResult(result); err != nil {
		log.Printf("Failed to save result for %s: %v", id, err)
		p.storage.UpdateStatus(id, "error", 0)
		return
	}

	p.storage.SetScore(id, score)
	p.storage.UpdateStatus(id, "done", 100)
	log.Printf("Video %s processed: score=%d", id, score)
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func generateMockPose() []models.Point {
	skeleton := [][2]float64{
		{50, 8},  // head
		{50, 22}, // neck
		{50, 40}, // torso
		{32, 52}, // left elbow
		{68, 52}, // right elbow
		{20, 62}, // left hand
		{80, 62}, // right hand
		{50, 40}, // hip center
		{38, 68}, // left knee
		{62, 68}, // right knee
		{36, 88}, // left foot
		{64, 88}, // right foot
	}
	var pts []models.Point
	for _, s := range skeleton {
		pts = append(pts, models.Point{
			X: s[0] + (rand.Float64()-0.5)*4,
			Y: s[1] + (rand.Float64()-0.5)*3,
		})
	}
	return pts
}
