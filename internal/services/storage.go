package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
)

type Storage struct {
	mu         sync.RWMutex
	videos     map[string]*models.Video
	uploadDir  string
	resultsDir string
}

func NewStorage(uploadDir, resultsDir string) *Storage {
	return &Storage{
		videos:     make(map[string]*models.Video),
		uploadDir:  uploadDir,
		resultsDir: resultsDir,
	}
}

func (s *Storage) CreateVideo(id, filename string) *models.Video {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	v := &models.Video{
		ID:        id,
		Filename:  filename,
		Status:    "uploaded",
		Progress:  0,
		CreatedAt: now,
		UpdatedAt: now,
	}
	s.videos[id] = v
	return v
}

func (s *Storage) GetVideo(id string) (*models.Video, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.videos[id]
	return v, ok
}

func (s *Storage) GetAllVideos() []*models.Video {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var list []*models.Video
	for _, v := range s.videos {
		list = append(list, v)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.After(list[j].CreatedAt)
	})
	return list
}

func (s *Storage) UpdateStatus(id, status string, progress int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if v, ok := s.videos[id]; ok {
		v.Status = status
		v.Progress = progress
		v.UpdatedAt = time.Now()
	}
}

func (s *Storage) SetScore(id string, score int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if v, ok := s.videos[id]; ok {
		v.Score = &score
		v.UpdatedAt = time.Now()
	}
}

func (s *Storage) SaveResult(r *models.Result) error {
	path := filepath.Join(s.resultsDir, r.VideoID+".json")
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal result: %w", err)
	}
	return os.WriteFile(path, data, 0o644)
}

func (s *Storage) LoadResult(videoID string) (*models.Result, error) {
	path := filepath.Join(s.resultsDir, videoID+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read result: %w", err)
	}
	var r models.Result
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, fmt.Errorf("unmarshal result: %w", err)
	}
	return &r, nil
}

func (s *Storage) UploadDir() string {
	return s.uploadDir
}

func (s *Storage) UploadPath(id, ext string) string {
	return filepath.Join(s.uploadDir, id+ext)
}
