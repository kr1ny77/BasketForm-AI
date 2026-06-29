package services

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
)

type Storage struct {
	mu         sync.RWMutex
	videos     map[string]*models.Video
	uploadDir  string
	resultsDir string
	dataDir    string
}

func genID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func NewStorage(uploadDir, resultsDir string) *Storage {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = "data"
	}
	s := &Storage{
		videos:     make(map[string]*models.Video),
		uploadDir:  uploadDir,
		resultsDir: resultsDir,
		dataDir:    dataDir,
	}
	s.ensureDataDirs()
	return s
}

func (s *Storage) ensureDataDirs() {
	dirs := []string{
		filepath.Join(s.dataDir, "users"),
		filepath.Join(s.dataDir, "videos"),
		filepath.Join(s.dataDir, "friends"),
		filepath.Join(s.dataDir, "shared"),
	}
	for _, d := range dirs {
		os.MkdirAll(d, 0o755)
	}
}

func (s *Storage) CreateVideo(id, filename, userID string) *models.Video {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	v := &models.Video{
		ID:        id,
		UserID:    userID,
		Filename:  filename,
		Status:    "uploaded",
		Progress:  0,
		CreatedAt: now,
		UpdatedAt: now,
	}
	s.videos[id] = v
	s.saveVideoJSON(v)
	return v
}

func (s *Storage) GetVideo(id string) (*models.Video, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.videos[id]
	if !ok {
		v, ok = s.loadVideoJSON(id)
	}
	return v, ok
}

func (s *Storage) GetAllVideos() []*models.Video {
	s.loadAllVideosFromDisk()

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

func (s *Storage) GetVideosByUserID(userID string) []*models.Video {
	s.loadAllVideosFromDisk()

	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*models.Video
	for _, v := range s.videos {
		if v.UserID == userID {
			list = append(list, v)
		}
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
		s.saveVideoJSON(v)
	}
}

func (s *Storage) SetScore(id string, score int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if v, ok := s.videos[id]; ok {
		v.Score = &score
		v.UpdatedAt = time.Now()
		s.saveVideoJSON(v)
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

func (s *Storage) LoadResultByID(resultID string) (*models.Result, error) {
	entries, err := os.ReadDir(s.resultsDir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.resultsDir, entry.Name()))
		if err != nil {
			continue
		}
		var r models.Result
		if err := json.Unmarshal(data, &r); err != nil {
			continue
		}
		if r.ID == resultID {
			return &r, nil
		}
	}
	return nil, fmt.Errorf("result not found")
}

func (s *Storage) UploadDir() string {
	return s.uploadDir
}

func (s *Storage) UploadPath(id, ext string) string {
	return filepath.Join(s.uploadDir, id+ext)
}

func (s *Storage) saveVideoJSON(v *models.Video) {
	path := filepath.Join(s.dataDir, "videos", v.ID+".json")
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(path, data, 0o644)
}

func (s *Storage) SaveVideoJSON(v *models.Video) {
	s.saveVideoJSON(v)
}

func (s *Storage) DeleteVideo(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.videos, id)
	os.Remove(filepath.Join(s.dataDir, "videos", id+".json"))
	os.Remove(filepath.Join(s.resultsDir, id+".json"))
	os.Remove(filepath.Join(s.uploadDir, id+".mp4"))
	os.Remove(filepath.Join(s.uploadDir, id+".MP4"))
	os.Remove(filepath.Join(s.uploadDir, id+".mov"))
	os.Remove(filepath.Join(s.uploadDir, id+".avi"))
}

func (s *Storage) loadVideoJSON(id string) (*models.Video, bool) {
	path := filepath.Join(s.dataDir, "videos", id+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}
	var v models.Video
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, false
	}
	s.videos[id] = &v
	return &v, true
}

func (s *Storage) loadAllVideosFromDisk() {
	s.mu.Lock()
	defer s.mu.Unlock()

	entries, err := os.ReadDir(filepath.Join(s.dataDir, "videos"))
	if err != nil {
		return
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		id := strings.TrimSuffix(entry.Name(), ".json")
		if _, loaded := s.videos[id]; loaded {
			continue
		}
		s.loadVideoJSON(id)
	}
}

// --- Users ---

func (s *Storage) SaveUser(user *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	path := filepath.Join(s.dataDir, "users", user.ID+".json")
	data, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func (s *Storage) GetUserByID(id string) (*models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	path := filepath.Join(s.dataDir, "users", id+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}
	var u models.User
	if err := json.Unmarshal(data, &u); err != nil {
		return nil, false
	}
	return &u, true
}

func (s *Storage) GetUserByEmail(email string) (*models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	entries, err := os.ReadDir(filepath.Join(s.dataDir, "users"))
	if err != nil {
		return nil, false
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dataDir, "users", entry.Name()))
		if err != nil {
			continue
		}
		var u models.User
		if err := json.Unmarshal(data, &u); err != nil {
			continue
		}
		if u.Email == email {
			return &u, true
		}
	}
	return nil, false
}

func (s *Storage) GetUserByNickname(nickname string) (*models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	entries, err := os.ReadDir(filepath.Join(s.dataDir, "users"))
	if err != nil {
		return nil, false
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dataDir, "users", entry.Name()))
		if err != nil {
			continue
		}
		var u models.User
		if err := json.Unmarshal(data, &u); err != nil {
			continue
		}
		if u.Nickname == nickname {
			return &u, true
		}
	}
	return nil, false
}

func (s *Storage) SearchUsersByNickname(query string) []*models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var users []*models.User
	entries, err := os.ReadDir(filepath.Join(s.dataDir, "users"))
	if err != nil {
		return nil
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dataDir, "users", entry.Name()))
		if err != nil {
			continue
		}
		var u models.User
		if err := json.Unmarshal(data, &u); err != nil {
			continue
		}
		if contains(u.Nickname, query) {
			users = append(users, &u)
		}
	}
	return users
}

func contains(s, sub string) bool {
	return len(sub) > 0 && len(s) >= len(sub) && strings.Contains(strings.ToLower(s), strings.ToLower(sub))
}

// --- Friend Requests ---

func (s *Storage) SaveFriendRequest(fr *models.FriendRequest) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	path := filepath.Join(s.dataDir, "friends", fr.ID+".json")
	data, err := json.MarshalIndent(fr, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func (s *Storage) GetFriendRequestsByToUser(userID string) []*models.FriendRequest {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var requests []*models.FriendRequest
	entries, err := os.ReadDir(filepath.Join(s.dataDir, "friends"))
	if err != nil {
		return nil
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dataDir, "friends", entry.Name()))
		if err != nil {
			continue
		}
		var fr models.FriendRequest
		if err := json.Unmarshal(data, &fr); err != nil {
			continue
		}
		if fr.ToUserID == userID {
			requests = append(requests, &fr)
		}
	}
	return requests
}

func (s *Storage) GetFriendRequestByID(id string) (*models.FriendRequest, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	path := filepath.Join(s.dataDir, "friends", id+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}
	var fr models.FriendRequest
	if err := json.Unmarshal(data, &fr); err != nil {
		return nil, false
	}
	return &fr, true
}

func (s *Storage) UpdateFriendRequest(fr *models.FriendRequest) error {
	return s.SaveFriendRequest(fr)
}

func (s *Storage) HasPendingRequest(fromID, toID string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	entries, err := os.ReadDir(filepath.Join(s.dataDir, "friends"))
	if err != nil {
		return false
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dataDir, "friends", entry.Name()))
		if err != nil {
			continue
		}
		var fr models.FriendRequest
		if err := json.Unmarshal(data, &fr); err != nil {
			continue
		}
		if fr.FromUserID == fromID && fr.ToUserID == toID && fr.Status == "pending" {
			return true
		}
	}
	return false
}

func (s *Storage) HasFriendship(userID, friendID string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	entries, err := os.ReadDir(filepath.Join(s.dataDir, "friends"))
	if err != nil {
		return false
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dataDir, "friends", entry.Name()))
		if err != nil {
			continue
		}
		var fr models.FriendRequest
		if err := json.Unmarshal(data, &fr); err != nil {
			continue
		}
		if fr.Status == "accepted" &&
			((fr.FromUserID == userID && fr.ToUserID == friendID) ||
				(fr.FromUserID == friendID && fr.ToUserID == userID)) {
			return true
		}
	}
	return false
}

// --- Friends ---

func (s *Storage) GetFriends(userID string) []*models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var friends []*models.User
	entries, err := os.ReadDir(filepath.Join(s.dataDir, "friends"))
	if err != nil {
		return nil
	}
	seen := make(map[string]bool)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dataDir, "friends", entry.Name()))
		if err != nil {
			continue
		}
		var fr models.FriendRequest
		if err := json.Unmarshal(data, &fr); err != nil {
			continue
		}
		if fr.Status != "accepted" {
			continue
		}
		var friendID string
		if fr.FromUserID == userID {
			friendID = fr.ToUserID
		} else if fr.ToUserID == userID {
			friendID = fr.FromUserID
		}
		if friendID == "" || seen[friendID] {
			continue
		}
		seen[friendID] = true
		user, ok := s.GetUserByID(friendID)
		if ok {
			friends = append(friends, user)
		}
	}
	return friends
}

// --- Shared Results ---

func (s *Storage) SaveSharedResult(sr *models.SharedResult) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	path := filepath.Join(s.dataDir, "shared", sr.ID+".json")
	data, err := json.MarshalIndent(sr, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func (s *Storage) GetSharedWithMe(userID string) []*models.SharedResult {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var results []*models.SharedResult
	entries, err := os.ReadDir(filepath.Join(s.dataDir, "shared"))
	if err != nil {
		return nil
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dataDir, "shared", entry.Name()))
		if err != nil {
			continue
		}
		var sr models.SharedResult
		if err := json.Unmarshal(data, &sr); err != nil {
			continue
		}
		if sr.ToUserID == userID {
			results = append(results, &sr)
		}
	}
	return results
}

func (s *Storage) GetSharedByMe(userID string) []*models.SharedResult {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var results []*models.SharedResult
	entries, err := os.ReadDir(filepath.Join(s.dataDir, "shared"))
	if err != nil {
		return nil
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dataDir, "shared", entry.Name()))
		if err != nil {
			continue
		}
		var sr models.SharedResult
		if err := json.Unmarshal(data, &sr); err != nil {
			continue
		}
		if sr.FromUserID == userID {
			results = append(results, &sr)
		}
	}
	return results
}
