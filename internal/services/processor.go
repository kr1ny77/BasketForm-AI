package services

import (
	"encoding/json"
	"log"
	"os"
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

type mlReport struct {
	Scores    map[string]float64 `json:"scores"`
	Metrics   map[string]float64 `json:"metrics"`
	AIFeedback string            `json:"ai_feedback"`
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

	outputDir := filepath.Join("results", video.UserID, id)
	os.MkdirAll(outputDir, 0o755)
	outputVideo := filepath.Join(outputDir, "output.mp4")

	p.storage.UpdateStatus(id, "processing", 10)

	err := runML(videoPath, outputVideo)
	if err != nil {
		log.Printf("ML script failed for %s: %v", id, err)
		p.storage.UpdateStatus(id, "error", 0)
		return
	}

	p.storage.UpdateStatus(id, "processing", 70)

	reportPath := filepath.Join(outputDir, "output.json")
	mlReport, err := loadMLReport(reportPath)
	if err != nil {
		log.Printf("Failed to load ML report for %s: %v", id, err)
		p.storage.UpdateStatus(id, "error", 0)
		return
	}

	p.storage.UpdateStatus(id, "processing", 90)

	phaseNames := []string{"Stance", "Arm Angle", "Release", "Follow-through"}
	mlScores := []float64{
		mlReport.Scores["DIP"],
		mlReport.Scores["ASCENT"],
		mlReport.Scores["RELEASE"],
		mlReport.Scores["RELEASE"],
	}

	var phases []models.PhaseScore
	var scoreList []int
	totalScore := 0.0
	count := 0.0

	for i, name := range phaseNames {
		s := 0.0
		if i < len(mlScores) {
			s = mlScores[i]
		}
		phases = append(phases, models.PhaseScore{
			PhaseName: name,
			Score:     s,
		})
		scoreList = append(scoreList, int(s))
		totalScore += s
		count++
	}

	overall := 0
	if count > 0 {
		overall = int(totalScore / count)
	}

	feedback := mlReport.AIFeedback
	if feedback == "" {
		feedback = "No feedback available."
	}

	result := &models.Result{
		ID:        id,
		UserID:    video.UserID,
		VideoID:   id,
		Filename:  video.Filename,
		Score:     overall,
		Feedback:  feedback,
		PoseData:  nil,
		Scores:    scoreList,
		Phases:    phases,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if err := p.storage.SaveResult(result); err != nil {
		log.Printf("Failed to save result for %s: %v", id, err)
		p.storage.UpdateStatus(id, "error", 0)
		return
	}

	p.storage.SetScore(id, overall)
	p.storage.UpdateStatus(id, "done", 100)
	log.Printf("Video %s processed: score=%d", id, overall)
}

func findPython() string {
	venvPython := filepath.Join("venv", "bin", "python3")
	if _, err := os.Stat(venvPython); err == nil {
		return venvPython
	}
	return "python3"
}

func runML(inputPath, outputPath string) error {
	python := findPython()
	script := filepath.Join("ML", "main.py")

	cmd := exec.Command(python, script, inputPath, outputPath)
	cmd.Dir = "."

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("ML output: %s", string(output))
		return err
	}
	log.Printf("ML output: %s", string(output))
	return nil
}

func loadMLReport(reportPath string) (*mlReport, error) {
	data, err := os.ReadFile(reportPath)
	if err != nil {
		return nil, err
	}
	var report mlReport
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, err
	}
	return &report, nil
}

func extOf(filename string) string {
	ext := filepath.Ext(filename)
	if ext == "" {
		return ".mp4"
	}
	return ext
}
