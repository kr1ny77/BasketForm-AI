package services

import (
	"encoding/json"
	"log"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
)

type Processor struct {
	storage *Storage
}

func NewProcessor(storage *Storage) *Processor {
	return &Processor{storage: storage}
}

type mlResult struct {
	Score    int            `json:"score"`
	Scores   []int          `json:"scores"`
	Feedback string         `json:"feedback"`
	PoseData []models.Point `json:"pose_data"`
	Error    string         `json:"error,omitempty"`
}

func (p *Processor) ProcessVideo(id string) {
	p.storage.UpdateStatus(id, "processing", 0)
	log.Printf("Processing video %s", id)

	video, ok := p.storage.GetVideo(id)
	if !ok {
		log.Printf("Video %s not found during processing", id)
		return
	}

	videoPath := filepath.Join(p.storage.UploadDir(), id+extOf(video.Filename))

	for progress := 10; progress <= 40; progress += 10 {
		time.Sleep(500 * time.Millisecond)
		p.storage.UpdateStatus(id, "processing", progress)
	}

	mlOut, err := runML(videoPath)
	if err != nil {
		log.Printf("ML script failed for %s: %v", id, err)
		p.storage.UpdateStatus(id, "error", 0)
		return
	}

	for progress := 50; progress <= 90; progress += 10 {
		time.Sleep(200 * time.Millisecond)
		p.storage.UpdateStatus(id, "processing", progress)
	}

	if mlOut.Error != "" {
		log.Printf("ML analysis error for %s: %s", id, mlOut.Error)
		p.storage.UpdateStatus(id, "error", 0)
		return
	}

	if len(mlOut.Scores) != 4 {
		mlOut.Scores = []int{0, 0, 0, 0}
	}
	if mlOut.Feedback == "" {
		mlOut.Feedback = "No feedback available."
	}

	result := &models.Result{
		ID:        id,
		VideoID:   id,
		Filename:  video.Filename,
		Score:     mlOut.Score,
		Feedback:  mlOut.Feedback,
		PoseData:  mlOut.PoseData,
		Scores:    mlOut.Scores,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if err := p.storage.SaveResult(result); err != nil {
		log.Printf("Failed to save result for %s: %v", id, err)
		p.storage.UpdateStatus(id, "error", 0)
		return
	}

	p.storage.SetScore(id, mlOut.Score)
	p.storage.UpdateStatus(id, "done", 100)
	log.Printf("Video %s processed: score=%d", id, mlOut.Score)
}

func runML(videoPath string) (*mlResult, error) {
	script := filepath.Join("scripts", "analyze_video.py")
	cmd := exec.Command("python3", script, videoPath)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var r mlResult
	if err := json.Unmarshal(out, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func extOf(filename string) string {
	ext := filepath.Ext(filename)
	if ext == "" {
		return ".mp4"
	}
	return ext
}
