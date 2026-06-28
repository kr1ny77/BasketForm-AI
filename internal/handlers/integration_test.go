package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
	"github.com/kr1ny77/BasketForm-AI/internal/services"
)

func setupIntegrationTest(t *testing.T) (*httptest.Server, string, *services.Storage) {
	t.Helper()
	upload := t.TempDir()
	results := t.TempDir()
	t.Setenv("DATA_DIR", filepath.Join(upload, "data"))
	storage := services.NewStorage(upload, results)
	processor := services.NewProcessor(storage)
	auth := services.NewAuthService(storage)

	h := New(upload, results)
	api := NewAPI(storage, processor, upload)
	authHandler := NewAuthHandler(auth, storage)
	friendsHandler := NewFriendsHandler(storage)
	shareHandler := NewShareHandler(storage)
	middleware := NewMiddleware(auth)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", h.LoginHandler)
	mux.HandleFunc("/register", h.RegisterHandler)
	mux.HandleFunc("/", h.IndexHandler)
	mux.HandleFunc("/upload", h.UploadPageHandler)
	mux.HandleFunc("/results", h.ResultsPageHandler)
	mux.HandleFunc("/profile", h.ProfilePageHandler)
	mux.HandleFunc("/progress", h.ProgressPageHandler)
	mux.HandleFunc("/export", h.ExportPageHandler)
	mux.HandleFunc("/friends", h.FriendsPageHandler)
	mux.HandleFunc("/shared", h.SharedResultsPageHandler)
	authHandler.Register(mux)
	api.Register(mux)
	friendsHandler.Register(mux)
	shareHandler.Register(mux)

	wrappedMux := middleware.AuthRequired(mux)
	ts := httptest.NewServer(wrappedMux)
	t.Cleanup(ts.Close)
	return ts, results, storage
}

func integrationDo(t *testing.T, ts *httptest.Server, method, path, body string, cookies []*http.Cookie) *httptest.ResponseRecorder {
	t.Helper()
	var reqBody *bytes.Buffer
	if body != "" {
		reqBody = bytes.NewBufferString(body)
	} else {
		reqBody = bytes.NewBuffer(nil)
	}
	req, err := http.NewRequest(method, ts.URL+path, reqBody)
	if err != nil {
		t.Fatalf("create request error: %v", err)
	}
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	for _, c := range cookies {
		req.AddCookie(c)
	}
	w := httptest.NewRecorder()
	ts.Config.Handler.ServeHTTP(w, req)
	return w
}

func integrationLogin(t *testing.T, ts *httptest.Server, email, password string) []*http.Cookie {
	t.Helper()
	w := integrationDo(t, ts, "POST", "/api/login", `{"email":"`+email+`","password":"`+password+`"}`, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("login failed for %s: %d %s", email, w.Code, w.Body.String())
	}
	return w.Result().Cookies()
}

func TestIntegration_RegisterAndLogin(t *testing.T) {
	ts, _, _ := setupIntegrationTest(t)

	w := integrationDo(t, ts, "POST", "/api/register", `{"email":"int@test.com","nickname":"intuser","password":"testpass123"}`, nil)
	if w.Code != http.StatusCreated {
		t.Fatalf("register: expected 201, got %d: %s", w.Code, w.Body.String())
	}

	w = integrationDo(t, ts, "POST", "/api/login", `{"email":"int@test.com","password":"testpass123"}`, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("login: expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var loginResp map[string]string
	json.Unmarshal(w.Body.Bytes(), &loginResp)
	if loginResp["token"] == "" {
		t.Fatal("expected token in login response")
	}
}

func TestIntegration_GetProfile(t *testing.T) {
	ts, _, _ := setupIntegrationTest(t)

	integrationDo(t, ts, "POST", "/api/register", `{"email":"prof@test.com","nickname":"profuser","password":"testpass123"}`, nil)
	cookies := integrationLogin(t, ts, "prof@test.com", "testpass123")

	w := integrationDo(t, ts, "GET", "/api/profile", "", cookies)
	if w.Code != http.StatusOK {
		t.Fatalf("profile: expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var profile map[string]string
	json.Unmarshal(w.Body.Bytes(), &profile)
	if profile["nickname"] != "profuser" {
		t.Fatalf("expected nickname profuser, got %s", profile["nickname"])
	}
}

func TestIntegration_UnauthorizedAccess(t *testing.T) {
	ts, _, _ := setupIntegrationTest(t)

	endpoints := []struct {
		method, path string
	}{
		{"GET", "/api/videos"},
		{"GET", "/api/profile"},
		{"POST", "/api/upload"},
	}

	for _, ep := range endpoints {
		w := integrationDo(t, ts, ep.method, ep.path, "", nil)
		if w.Code != http.StatusUnauthorized {
			t.Errorf("%s %s: expected 401, got %d", ep.method, ep.path, w.Code)
		}
	}
}

func TestIntegration_ShareResult(t *testing.T) {
	ts, resultsDir, _ := setupIntegrationTest(t)

	integrationDo(t, ts, "POST", "/api/register", `{"email":"share1@test.com","nickname":"share1","password":"testpass123"}`, nil)
	integrationDo(t, ts, "POST", "/api/register", `{"email":"share2@test.com","nickname":"share2","password":"testpass123"}`, nil)

	cookies1 := integrationLogin(t, ts, "share1@test.com", "testpass123")
	cookies2 := integrationLogin(t, ts, "share2@test.com", "testpass123")

	w := integrationDo(t, ts, "GET", "/api/profile", "", cookies1)
	var p1 map[string]string
	json.Unmarshal(w.Body.Bytes(), &p1)

	w = integrationDo(t, ts, "GET", "/api/profile", "", cookies2)
	var p2 map[string]string
	json.Unmarshal(w.Body.Bytes(), &p2)
	user2ID := p2["id"]

	// Create a result file for user1
	result := &models.Result{
		ID:       "testresult1",
		UserID:   p1["id"],
		VideoID:  "v1",
		Filename: "test.mp4",
		Score:    85,
		Feedback: "Great",
		Scores:   []int{80, 80, 80, 80},
	}
	// Save result directly to the results directory
	resultData, _ := json.Marshal(result)
	os.WriteFile(filepath.Join(resultsDir, "v1.json"), resultData, 0o644)

	// User1 shares a result with user2
	w = integrationDo(t, ts, "POST", "/api/share/result", `{"result_id":"testresult1","friend_id":"`+user2ID+`"}`, cookies1)
	if w.Code != http.StatusCreated {
		t.Fatalf("share: expected 201, got %d: %s", w.Code, w.Body.String())
	}

	// User2 checks shared results
	w = integrationDo(t, ts, "GET", "/api/results/shared-with-me", "", cookies2)
	if w.Code != http.StatusOK {
		t.Fatalf("shared-with-me: expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var shared []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &shared)
	if len(shared) != 1 {
		t.Fatalf("expected 1 shared result, got %d", len(shared))
	}
	if shared[0]["from_nickname"] != "share1" {
		t.Fatalf("expected from_nickname share1, got %v", shared[0]["from_nickname"])
	}
}

func TestIntegration_SearchAndFriendRequest(t *testing.T) {
	ts, _, _ := setupIntegrationTest(t)

	integrationDo(t, ts, "POST", "/api/register", `{"email":"search1@test.com","nickname":"searchable_user","password":"testpass123"}`, nil)
	integrationDo(t, ts, "POST", "/api/register", `{"email":"search2@test.com","nickname":"searcher2","password":"testpass123"}`, nil)

	cookies1 := integrationLogin(t, ts, "search1@test.com", "testpass123")
	cookies2 := integrationLogin(t, ts, "search2@test.com", "testpass123")

	// Search for user
	w := integrationDo(t, ts, "GET", "/api/friends/search?q=searchable", "", cookies2)
	if w.Code != http.StatusOK {
		t.Fatalf("search: expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var users []map[string]string
	json.Unmarshal(w.Body.Bytes(), &users)
	if len(users) != 1 {
		t.Fatalf("expected 1 search result, got %d", len(users))
	}

	targetID := users[0]["id"]

	// Send friend request
	w = integrationDo(t, ts, "POST", "/api/friends/request", `{"to_user_id":"`+targetID+`"}`, cookies2)
	if w.Code != http.StatusCreated {
		t.Fatalf("send request: expected 201, got %d: %s", w.Code, w.Body.String())
	}

	// User1 checks requests
	w = integrationDo(t, ts, "GET", "/api/friends/requests", "", cookies1)
	var reqs []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &reqs)
	if len(reqs) != 1 {
		t.Fatalf("expected 1 pending request, got %d", len(reqs))
	}

	// Accept request
	reqID := reqs[0]["id"].(string)
	w = integrationDo(t, ts, "POST", "/api/friends/accept", `{"request_id":"`+reqID+`"}`, cookies1)
	if w.Code != http.StatusOK {
		t.Fatalf("accept: expected 200, got %d: %s", w.Code, w.Body.String())
	}

	// Check friends list
	w = integrationDo(t, ts, "GET", "/api/friends", "", cookies2)
	var friends []map[string]string
	json.Unmarshal(w.Body.Bytes(), &friends)
	if len(friends) != 1 {
		t.Fatalf("expected 1 friend, got %d", len(friends))
	}
}
