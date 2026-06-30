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

var debugLog *os.File

func init() {
	f, err := os.OpenFile("/home/basketfrom-ai/BasketForm-AI/processor.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		debugLog = f
		log.SetOutput(f)
	}
}

type Processor struct {
	storage *Storage
}

func NewProcessor(storage *Storage) *Processor {
	return &Processor{storage: storage}
}

type mlReport struct {
	Scores     map[string]float64 `json:"scores"`
	Metrics    map[string]float64 `json:"metrics"`
	AIFeedback string             `json:"ai_feedback"`
}

func (p *Processor) ProcessVideo(id string) {
	p.storage.UpdateStatus(id, "processing", 0)
	log.Printf("[PROCESS] Start processing video %s", id)

	video, ok := p.storage.GetVideo(id)
	if !ok {
		log.Printf("[PROCESS] Video %s not found during processing", id)
		return
	}

	videoPath := filepath.Join(p.storage.UploadDir(), id+extOf(video.Filename))
	reportPath := filepath.Join(p.storage.UploadDir(), id+"_report.json")

	log.Printf("[PROCESS] videoPath=%s reportPath=%s", videoPath, reportPath)

	// Check file exists
	if _, err := os.Stat(videoPath); os.IsNotExist(err) {
		log.Printf("[PROCESS] Video file not found: %s", videoPath)
		p.storage.UpdateStatus(id, "error", 0)
		return
	}

	p.storage.UpdateStatus(id, "processing", 10)

	err := runML(videoPath, reportPath, video.Lang)
	if err != nil {
		log.Printf("[PROCESS] ML script failed for %s: %v", id, err)
		p.storage.UpdateStatus(id, "error", 0)
		return
	}

	p.storage.UpdateStatus(id, "processing", 70)

	mlReportData, err := loadMLReport(reportPath)
	if err != nil {
		log.Printf("No ML report for %s, using fallback: %v", id, err)
		mlReportData = &mlReport{
			Scores:     map[string]float64{"DIP": 50, "ASCENT": 50, "RELEASE": 50},
			Metrics:    map[string]float64{},
			AIFeedback: "Analysis completed. Scores are estimated.",
		}
	}

	p.storage.UpdateStatus(id, "processing", 90)

	phaseNames := []string{"Stance (DIP)", "Ascent (ASCENT)", "Release (RELEASE)", "Follow-through"}
	mlKeys := []string{"DIP", "ASCENT", "RELEASE", "FOLLOW_THROUGH"}

	var phases []models.PhaseScore
	var scoreList []int
	totalScore := 0.0
	count := 0.0

	for i, name := range phaseNames {
		s := 0.0
		if i < len(mlKeys) {
			s = mlReportData.Scores[mlKeys[i]]
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

	feedback := mlReportData.AIFeedback
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

	os.Remove(reportPath)

	p.storage.SetScore(id, overall)
	p.storage.UpdateStatus(id, "done", 100)
	log.Printf("Video %s processed: score=%d", id, overall)

	// Delete uploaded video file to save disk space
	os.Remove(videoPath)
}

func findPython() string {
	venvPython := filepath.Join("venv", "bin", "python3")
	if _, err := os.Stat(venvPython); err == nil {
		return venvPython
	}
	return "python3"
}

func runML(inputPath, reportPath, lang string) error {
	python := findPython()
	script := filepath.Join("ML", "main.py")

	if lang == "" {
		lang = "en"
	}

	log.Printf("ML run: python=%s script=%s input=%s report=%s lang=%s", python, script, inputPath, reportPath, lang)

	cmd := exec.Command(python, script, inputPath, reportPath, "--lang", lang)
	cmd.Dir = "/home/basketfrom-ai/BasketForm-AI"

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("ML FAILED: %v\nOUTPUT: %s", err, string(output))
		return err
	}
	log.Printf("ML OK: %s", string(output))
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
