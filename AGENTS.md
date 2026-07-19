# Agent Guidance for BasketForm-AI

This document provides guidance for AI agents and automated tools working with the BasketForm-AI codebase.

## Repository Overview

BasketForm-AI is a web application for analyzing basketball shooting form from video. The tech stack:

- **Backend:** Go 1.22+ (standard library `net/http`)
- **Frontend:** HTML + CSS + JavaScript (vanilla, no build step)
- **ML Pipeline:** Python (MediaPipe, OpenCV, YOLO, LLM via OpenRouter API)
- **Storage:** Local JSON files in `data/` directory
- **Deployment:** Binary on VM with Nginx reverse proxy

## Project Structure

```
BasketForm-AI/
├── cmd/server/main.go          # Application entry point
├── internal/
│   ├── handlers/               # HTTP handlers (api, auth, friends, share, middleware)
│   ├── models/                 # Data structures
│   ├── services/               # Business logic (auth, storage, processor)
│   └── qrt/                    # Quality requirement tests
├── ML/                         # Python ML pipeline
├── web/
│   ├── templates/              # HTML templates
│   └── static/                 # CSS, JS, images
├── data/                       # JSON user/video/friend storage
├── uploads/                    # Uploaded video files
├── results/                    # Analysis result files
├── docs/                       # Project documentation
├── reports/                    # Course reports
└── scripts/                    # Deployment and utility scripts
```

## Development Rules

### Code Style

- Go code: follow standard Go conventions (`gofmt`, `go vet`)
- Use `golangci-lint` for linting
- No comments unless explicitly requested
- Follow existing patterns in the codebase

### Testing

- Write unit tests for new handlers and services
- Integration tests use `httptest` package
- Quality requirement tests use `go test -tags=qrt`
- Target 30% coverage for critical modules

### Git Workflow

- Branch from issues: `<issue-number>-<short-description>`
- All changes via Pull Requests
- At least one review approval required
- CI checks must pass before merge
- Update CHANGELOG.md for user-visible changes

### Security

- Never commit secrets, API keys, or credentials
- Use environment variables for configuration
- All API endpoints require JWT authentication
- Passwords hashed with bcrypt

## Key Files

| File | Purpose |
|------|---------|
| `cmd/server/main.go` | Application entry point |
| `internal/handlers/*.go` | HTTP request handlers |
| `internal/services/auth.go` | Authentication (bcrypt + JWT) |
| `internal/services/storage.go` | JSON file storage |
| `internal/services/processor.go` | ML video processing orchestration |
| `ML/main.py` | ML pipeline entry point |
| `ML/feedback_generator.py` | LLM scoring via OpenRouter API |
| `docs/` | All project documentation |
| `reports/` | Course weekly reports |
| `scripts/deploy.sh` | Production deployment script (build, upload, restart, smoke test) |

## Common Tasks

### Adding a New API Endpoint

1. Define handler in `internal/handlers/`
2. Add route in `cmd/server/main.go`
3. Write unit tests
4. Update API documentation if applicable

### Modifying ML Pipeline

1. Edit Python files in `ML/`
2. Ensure `requirements.txt` is updated if new dependencies are added
3. Test with sample video
4. Verify Go server still calls ML correctly via `exec.Command`

### Updating Documentation

1. Keep `docs/` files current with code changes
2. Update `README.md` if setup or access instructions change
3. Update `CHANGELOG.md` for user-visible changes

### Deploying to Production

1. Run `./scripts/deploy.sh` from repository root
2. The script builds a Linux binary, uploads to VM at 80.74.30.14, installs ML deps, restarts the service, and runs a smoke test
3. Set `SKIP_ML_DEPS=1` to skip ML dependency installation on repeated deploys
4. Verify deployment at http://80.74.30.14/

## CI Pipeline

The CI pipeline runs on every PR and push to `main`:

1. `golangci-lint` — Go code quality
2. `go test -race` — Unit and integration tests
3. Coverage check — 30% minimum for critical modules
4. QRT — Quality requirement tests
5. `govulncheck` — Security vulnerability scanning
6. Lychee — Markdown link validation

## Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `OPENROUTER_API_KEY` | Yes | API key for LLM vision scoring |
| `JWT_SECRET` | Yes | Secret for JWT token signing |
| `PORT` | No | Server port (default: 8080) |
| `UPLOAD_DIR` | No | Video upload directory (default: uploads) |
| `RESULTS_DIR` | No | Analysis results directory (default: results) |
| `DATA_DIR` | No | JSON storage directory (default: data) |

## Constraints

- JSON file storage limits concurrent users to ~10
- ML processing runs synchronously per request (10-30 seconds)
- No horizontal scaling without refactoring
- Python ML pipeline requires MediaPipe, OpenCV, YOLO dependencies
