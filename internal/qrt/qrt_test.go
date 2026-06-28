//go:build qrt

package qrt

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/kr1ny77/BasketForm-AI/internal/handlers"
	"github.com/kr1ny77/BasketForm-AI/internal/services"
)

func repoRoot() string {
	_, f, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filepath.Dir(f)))
}

func setupTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	root := repoRoot()
	upload := t.TempDir()
	results := t.TempDir()
	t.Setenv("DATA_DIR", filepath.Join(upload, "data"))
	storage := services.NewStorage(upload, results)
	processor := services.NewProcessor(storage)
	auth := services.NewAuthService(storage)

	h := handlers.New(upload, results)
	api := handlers.NewAPI(storage, processor, upload)
	authHandler := handlers.NewAuthHandler(auth, storage)
	friendsHandler := handlers.NewFriendsHandler(storage)
	shareHandler := handlers.NewShareHandler(storage)
	middleware := handlers.NewMiddleware(auth)

	mux := http.NewServeMux()
	loginPath := filepath.Join(root, "web", "templates", "login.html")
	registerPath := filepath.Join(root, "web", "templates", "signup.html")
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile(loginPath)
		if err != nil {
			http.Error(w, "not found", 404)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile(registerPath)
		if err != nil {
			http.Error(w, "not found", 404)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})
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
	return httptest.NewServer(wrappedMux)
}

func TestQRT_APIResponseTime(t *testing.T) {
	ts := setupTestServer(t)
	defer ts.Close()

	registerAndLogin := func() string {
		resp, _ := http.Post(ts.URL+"/api/register", "application/json",
			strings.NewReader(`{"email":"qrt1@test.com","nickname":"qrt1user","password":"testpass123"}`))
		resp.Body.Close()
		resp, _ = http.Post(ts.URL+"/api/login", "application/json",
			strings.NewReader(`{"email":"qrt1@test.com","password":"testpass123"}`))
		defer resp.Body.Close()
		for _, c := range resp.Cookies() {
			if c.Name == "token" {
				return c.Value
			}
		}
		return ""
	}

	token := registerAndLogin()
	if token == "" {
		t.Fatal("failed to get auth token")
	}

	endpoints := []string{"/api/videos", "/api/profile"}
	for _, ep := range endpoints {
		start := time.Now()
		req, _ := http.NewRequest("GET", ts.URL+ep, nil)
		req.AddCookie(&http.Cookie{Name: "token", Value: token})
		resp, err := http.DefaultClient.Do(req)
		elapsed := time.Since(start)
		if err != nil {
			t.Errorf("endpoint %s: %v", ep, err)
			continue
		}
		resp.Body.Close()
		if elapsed > 2*time.Second {
			t.Errorf("endpoint %s: response time %v exceeds 2s threshold", ep, elapsed)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("endpoint %s: expected 200, got %d", ep, resp.StatusCode)
		}
	}
}

func TestQRT_AuthSecurity(t *testing.T) {
	ts := setupTestServer(t)
	defer ts.Close()

	noRedirectClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	protectedPages := []string{"/upload", "/results", "/profile", "/friends", "/shared"}
	for _, page := range protectedPages {
		resp, err := noRedirectClient.Get(ts.URL + page)
		if err != nil {
			t.Errorf("page %s: %v", page, err)
			continue
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusFound {
			t.Errorf("page %s: expected redirect (302), got %d", page, resp.StatusCode)
		}
		location := resp.Header.Get("Location")
		if location != "/login" {
			t.Errorf("page %s: redirect to %s, expected /login", page, location)
		}
	}

	protectedAPIs := []string{"/api/videos", "/api/profile"}
	for _, api := range protectedAPIs {
		resp, err := noRedirectClient.Get(ts.URL + api)
		if err != nil {
			t.Errorf("api %s: %v", api, err)
			continue
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("api %s: expected 401, got %d", api, resp.StatusCode)
		}
	}
}

func TestQRT_CriticalModuleCoverage(t *testing.T) {
	root := repoRoot()
	criticalModules := []string{
		"internal/services",
		"internal/handlers",
	}
	for _, mod := range criticalModules {
		testDir := filepath.Join(root, mod)
		if _, err := os.Stat(testDir); os.IsNotExist(err) {
			t.Errorf("critical module directory %s does not exist", testDir)
			continue
		}
		entries, err := os.ReadDir(testDir)
		if err != nil {
			t.Errorf("cannot read %s: %v", testDir, err)
			continue
		}
		hasTests := false
		for _, e := range entries {
			if strings.HasSuffix(e.Name(), "_test.go") {
				hasTests = true
				break
			}
		}
		if !hasTests {
			t.Errorf("critical module %s has no test files", testDir)
		}
	}
}

func TestQRT_FormGuidance(t *testing.T) {
	root := repoRoot()
	templateDir := filepath.Join(root, "web", "templates")
	formPages := []string{"login.html", "signup.html", "profile.html", "upload.html"}

	for _, page := range formPages {
		path := filepath.Join(templateDir, page)
		data, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("cannot read template %s: %v", page, err)
			continue
		}
		content := string(data)

		hasPlaceholder := strings.Contains(content, "placeholder=")
		hasLabel := strings.Contains(content, "<label")
		hasDropzoneHint := strings.Contains(content, "dropzone-hint") || strings.Contains(content, "data-i18n")

		if !hasPlaceholder && !hasLabel && !hasDropzoneHint {
			t.Errorf("template %s: no placeholder, label, or dropzone hint found for form guidance", page)
		}
	}
}
