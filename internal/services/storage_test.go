package services

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
)

func tempDirs(t *testing.T) (string, string) {
	t.Helper()
	upload := t.TempDir()
	results := t.TempDir()
	t.Setenv("DATA_DIR", filepath.Join(upload, "data"))
	return upload, results
}

func TestStorage_CreateAndGetVideo(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	v := 	s.CreateVideo("id1", "test.mp4", "user1")
	if v.ID != "id1" {
		t.Fatalf("expected id1, got %s", v.ID)
	}
	if v.Filename != "test.mp4" {
		t.Fatalf("expected test.mp4, got %s", v.Filename)
	}
	if v.Status != "uploaded" {
		t.Fatalf("expected uploaded, got %s", v.Status)
	}
	if v.Progress != 0 {
		t.Fatalf("expected progress 0, got %d", v.Progress)
	}

	got, ok := s.GetVideo("id1")
	if !ok || got.ID != "id1" {
		t.Fatalf("GetVideo failed for id1")
	}
}

func TestStorage_GetVideo_NotFound(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	_, ok := s.GetVideo("nonexistent")
	if ok {
		t.Fatal("expected not found")
	}
}

func TestStorage_GetAllVideos_Sorted(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	s.CreateVideo("c", "c.mp4", "user1")
	time.Sleep(1100 * time.Millisecond)
	s.CreateVideo("a", "a.mp4", "user1")
	time.Sleep(1100 * time.Millisecond)
	s.CreateVideo("b", "b.mp4", "user1")

	list := s.GetAllVideos()
	if len(list) != 3 {
		t.Fatalf("expected 3 videos, got %d", len(list))
	}
	// Newest first (b created last)
	if list[0].ID != "b" {
		t.Fatalf("expected b first, got %s", list[0].ID)
	}
}

func TestStorage_GetAllVideos_Empty(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	list := s.GetAllVideos()
	if list != nil {
		t.Fatalf("expected nil, got %v", list)
	}
}

func TestStorage_UpdateStatus(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	s.CreateVideo("v1", "test.mp4", "user1")
	s.UpdateStatus("v1", "processing", 50)

	v, _ := s.GetVideo("v1")
	if v.Status != "processing" {
		t.Fatalf("expected processing, got %s", v.Status)
	}
	if v.Progress != 50 {
		t.Fatalf("expected 50, got %d", v.Progress)
	}
}

func TestStorage_UpdateStatus_NotFound(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	// Should not panic
	s.UpdateStatus("missing", "done", 100)
}

func TestStorage_SetScore(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	s.CreateVideo("v1", "test.mp4", "user1")
	s.SetScore("v1", 85)

	v, _ := s.GetVideo("v1")
	if v.Score == nil || *v.Score != 85 {
		t.Fatalf("expected score 85, got %v", v.Score)
	}
}

func TestStorage_SaveAndLoadResult(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	r := &models.Result{
		ID:       "r1",
		VideoID:  "v1",
		Filename: "test.mp4",
		Score:    75,
		Feedback: "Good form",
		PoseData: []models.Point{{X: 10, Y: 20}},
		Scores:   []int{80, 70, 75, 72},
	}

	if err := s.SaveResult(r); err != nil {
		t.Fatalf("SaveResult error: %v", err)
	}

	loaded, err := s.LoadResult("v1")
	if err != nil {
		t.Fatalf("LoadResult error: %v", err)
	}
	if loaded.Score != 75 {
		t.Fatalf("expected score 75, got %d", loaded.Score)
	}
	if loaded.Feedback != "Good form" {
		t.Fatalf("expected feedback 'Good form', got %s", loaded.Feedback)
	}
	if len(loaded.PoseData) != 1 {
		t.Fatalf("expected 1 pose point, got %d", len(loaded.PoseData))
	}
}

func TestStorage_LoadResult_NotFound(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	_, err := s.LoadResult("nonexistent")
	if err == nil {
		t.Fatal("expected error for missing result")
	}
}

func TestStorage_UploadPath(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	path := s.UploadPath("abc", ".mp4")
	expected := filepath.Join(upload, "abc.mp4")
	if path != expected {
		t.Fatalf("expected %s, got %s", expected, path)
	}
}

func TestStorage_ResultFileExists(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	r := &models.Result{VideoID: "v2", Score: 90}
	s.SaveResult(r)

	path := filepath.Join(results, "v2.json")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatal("result JSON file not created")
	}
}

func TestStorage_GetVideosByUserID(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	s.CreateVideo("v1", "a.mp4", "user1")
	s.CreateVideo("v2", "b.mp4", "user1")
	s.CreateVideo("v3", "c.mp4", "user2")

	list := s.GetVideosByUserID("user1")
	if len(list) != 2 {
		t.Fatalf("expected 2 videos for user1, got %d", len(list))
	}

	list = s.GetVideosByUserID("user2")
	if len(list) != 1 {
		t.Fatalf("expected 1 video for user2, got %d", len(list))
	}

	list = s.GetVideosByUserID("nobody")
	if len(list) != 0 {
		t.Fatalf("expected 0 videos for nobody, got %d", len(list))
	}
}

func TestStorage_LoadResultByID(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	r := &models.Result{ID: "r1", VideoID: "v1", Score: 75, Feedback: "Good"}
	s.SaveResult(r)

	loaded, err := s.LoadResultByID("r1")
	if err != nil {
		t.Fatalf("LoadResultByID error: %v", err)
	}
	if loaded.Score != 75 {
		t.Fatalf("expected score 75, got %d", loaded.Score)
	}
}

func TestStorage_LoadResultByID_NotFound(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	_, err := s.LoadResultByID("nonexistent")
	if err == nil {
		t.Fatal("expected error for missing result")
	}
}

func TestStorage_UploadDir(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	if s.UploadDir() != upload {
		t.Fatalf("expected %s, got %s", upload, s.UploadDir())
	}
}

func TestStorage_DeleteVideo(t *testing.T) {
	upload, results := tempDirs(t)
	t.Setenv("DATA_DIR", filepath.Join(upload, "data"))
	s := NewStorage(upload, results)

	s.CreateVideo("del1", "test.mp4", "user1")
	if _, ok := s.GetVideo("del1"); !ok {
		t.Fatal("video should exist before delete")
	}

	s.DeleteVideo("del1")
	if _, ok := s.GetVideo("del1"); ok {
		t.Fatal("video should not exist after delete")
	}
}

func TestStorage_SearchUsersByNickname(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)

	s.SaveUser(&models.User{ID: "u1", Email: "a@test.com", Nickname: "alice"})
	s.SaveUser(&models.User{ID: "u2", Email: "b@test.com", Nickname: "bob"})
	s.SaveUser(&models.User{ID: "u3", Email: "c@test.com", Nickname: "alice_wonderland"})

	users := s.SearchUsersByNickname("alice")
	if len(users) != 2 {
		t.Fatalf("expected 2 users matching 'alice', got %d", len(users))
	}
}

func TestStorage_FriendRequests(t *testing.T) {
	upload, results := tempDirs(t)
	t.Setenv("DATA_DIR", filepath.Join(upload, "data"))
	s := NewStorage(upload, results)

	fr := &models.FriendRequest{
		ID:         "fr1",
		FromUserID: "u1",
		ToUserID:   "u2",
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	if err := s.SaveFriendRequest(fr); err != nil {
		t.Fatalf("SaveFriendRequest error: %v", err)
	}

	got, ok := s.GetFriendRequestByID("fr1")
	if !ok || got.FromUserID != "u1" {
		t.Fatal("GetFriendRequestByID failed")
	}

	requests := s.GetFriendRequestsByToUser("u2")
	if len(requests) != 1 {
		t.Fatalf("expected 1 request for u2, got %d", len(requests))
	}

	if !s.HasPendingRequest("u1", "u2") {
		t.Fatal("expected pending request from u1 to u2")
	}
	if s.HasPendingRequest("u1", "u3") {
		t.Fatal("expected no pending request from u1 to u3")
	}

	if s.HasFriendship("u1", "u2") {
		t.Fatal("should not be friends while status is pending")
	}

	got.Status = "accepted"
	s.UpdateFriendRequest(got)

	if !s.HasFriendship("u1", "u2") {
		t.Fatal("expected friendship after acceptance")
	}
	if !s.HasFriendship("u2", "u1") {
		t.Fatal("expected friendship in reverse direction")
	}
}

func TestStorage_HasFriendship(t *testing.T) {
	upload, results := tempDirs(t)
	t.Setenv("DATA_DIR", filepath.Join(upload, "data"))
	s := NewStorage(upload, results)

	fr := &models.FriendRequest{
		ID:         "fr2",
		FromUserID: "u1",
		ToUserID:   "u2",
		Status:     "accepted",
		CreatedAt:  time.Now(),
	}
	s.SaveFriendRequest(fr)

	if !s.HasFriendship("u1", "u2") {
		t.Fatal("expected friendship between u1 and u2")
	}
	if !s.HasFriendship("u2", "u1") {
		t.Fatal("expected friendship between u2 and u1")
	}
	if s.HasFriendship("u1", "u3") {
		t.Fatal("expected no friendship between u1 and u3")
	}
}

func TestStorage_GetFriends(t *testing.T) {
	upload, results := tempDirs(t)
	t.Setenv("DATA_DIR", filepath.Join(upload, "data"))
	s := NewStorage(upload, results)

	s.SaveUser(&models.User{ID: "u1", Email: "a@test.com", Nickname: "alice"})
	s.SaveUser(&models.User{ID: "u2", Email: "b@test.com", Nickname: "bob"})
	s.SaveUser(&models.User{ID: "u3", Email: "c@test.com", Nickname: "carol"})

	fr := &models.FriendRequest{
		ID:         "fr3",
		FromUserID: "u1",
		ToUserID:   "u2",
		Status:     "accepted",
		CreatedAt:  time.Now(),
	}
	s.SaveFriendRequest(fr)

	friends := s.GetFriends("u1")
	if len(friends) != 1 {
		t.Fatalf("expected 1 friend for u1, got %d", len(friends))
	}
	if friends[0].Nickname != "bob" {
		t.Fatalf("expected friend bob, got %s", friends[0].Nickname)
	}
}

func TestStorage_SharedResults(t *testing.T) {
	upload, results := tempDirs(t)
	t.Setenv("DATA_DIR", filepath.Join(upload, "data"))
	s := NewStorage(upload, results)

	s.SaveUser(&models.User{ID: "u1", Email: "a@test.com", Nickname: "alice"})
	s.SaveUser(&models.User{ID: "u2", Email: "b@test.com", Nickname: "bob"})

	sr := &models.SharedResult{
		ID:         "sr1",
		ResultID:   "r1",
		FromUserID: "u1",
		ToUserID:   "u2",
		CreatedAt:  time.Now(),
	}
	s.SaveSharedResult(sr)

	withMe := s.GetSharedWithMe("u2")
	if len(withMe) != 1 {
		t.Fatalf("expected 1 shared with u2, got %d", len(withMe))
	}
	if withMe[0].FromUserID != "u1" {
		t.Fatalf("expected from u1, got %s", withMe[0].FromUserID)
	}

	byMe := s.GetSharedByMe("u1")
	if len(byMe) != 1 {
		t.Fatalf("expected 1 shared by u1, got %d", len(byMe))
	}
}
