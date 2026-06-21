# BasketForm-AI

Web service for analyzing basketball shooting form via video upload, biomechanical keypoint extraction, technique scoring, and personalized AI feedback.

**Live:** http://80.74.30.14/ · [Video Demo](https://youtu.be/To_1ZwAMe4M) · [v0.1.0 Release](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.1.0)

## Project Overview

BasketForm-AI is an AI-powered platform that helps basketball players improve their shooting technique through video analysis and personalized feedback. Users upload shooting videos, the system extracts biomechanical keypoints, evaluates stance, arm angle, release point, and follow-through, then generates simplified actionable recommendations.

## Key Features

- **Video Upload:** Drag-and-drop or file picker for basketball shot videos (MP4, MOV, AVI)
- **Biomechanical Analysis:** Automatic extraction of key body points from video
- **Technique Evaluation:** Scoring of stance, arm angle, release point, and follow-through (0–100)
- **Personalized Feedback:** AI-generated recommendations (max 3 actionable takeaways)
- **Progress Tracking:** Real-time processing status with auto-polling
- **PDF Export:** Download analysis reports via jsPDF
- **Canvas Animation:** Basketball-themed background (35 objects, mouse-reactive, orange waves)

## Tech Stack

- **Backend:** Go 1.21+ (standard library `net/http` + `html/template`)
- **Frontend:** HTML + CSS + JavaScript, Canvas animations, Chart.js, jsPDF
- **Storage:** Local `uploads/` and `results/` directories
- **ML:** Mock data pipeline (Python script available via `exec`)
- **Testing:** Go `testing`, `httptest` (47 unit + integration tests)
- **CI:** GitHub Actions — golangci-lint, go test, go build, Lychee link check
- **Deployment:** Binary or Docker
- **License:** MIT

## Quick Start

```bash
git clone https://github.com/kr1ny77/BasketForm-AI.git
cd BasketForm-AI
go run ./cmd/server/
# Open http://localhost:8080
```

## Local Development Setup

### Prerequisites

- Go 1.21 or later
- `golangci-lint` (for linting, optional)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/kr1ny77/BasketForm-AI.git
   cd BasketForm-AI
   ```

2. **Run the server:**
   ```bash
   go run ./cmd/server/
   ```

3. **Open in browser:**
   ```
   http://localhost:8080
   ```

### Build

```bash
go build -o bin/server ./cmd/server/
PORT=8080 ./bin/server
```

### Run Tests

```bash
go test -race -coverprofile=coverage.out ./...
```

### Run Linter

```bash
golangci-lint run
```

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `UPLOAD_DIR` | `uploads` | Video upload directory |
| `RESULTS_DIR` | `results` | Analysis results directory |

## Deployment

### Binary

```bash
go build -o bin/server ./cmd/server/
PORT=80 ./bin/server
```

### Docker

```bash
docker build -t basketform-ai .
docker run -p 8080:8080 basketform-ai
```

### Production (systemd)

```bash
# Build
go build -o /usr/local/bin/basketform-ai ./cmd/server/

# Create service file /etc/systemd/system/basketform-ai.service
# [Unit]
# Description=BasketForm-AI
# After=network.target
#
# [Service]
# Type=simple
# User=www-data
# WorkingDirectory=/opt/basketform-ai
# ExecStart=/usr/local/bin/basketform-ai
# Restart=always
# Environment=PORT=80
# Environment=UPLOAD_DIR=uploads
# Environment=RESULTS_DIR=results
#
# [Install]
# WantedBy=multi-user.target

sudo systemctl enable basketform-ai
sudo systemctl start basketform-ai
```

### Smoke Check

```bash
curl http://localhost:8080/
# Expected: redirects to /upload

curl http://localhost:8080/api/videos
# Expected: []
```

## Project Structure

```
BasketForm-AI/
├── cmd/server/main.go          # Application entry point
├── internal/
│   ├── handlers/               # HTTP handlers (pages + API)
│   │   ├── handlers.go         # Page handlers
│   │   ├── api.go              # API handlers
│   │   └── *_test.go           # Tests
│   ├── models/                 # Data structures
│   └── services/               # Business logic (storage, processor)
├── web/
│   ├── templates/              # HTML templates (html/template)
│   └── static/                 # CSS, JS, images
├── scripts/                    # Helper scripts (process_video.py)
├── uploads/                    # Uploaded video files (gitignored)
├── results/                    # Analysis result JSON files (gitignored)
├── reports/                    # Course reports
├── docs/                       # Project documentation
├── Dockerfile                  # Container build
├── CHANGELOG.md                # Version history
└── LICENSE                     # MIT License
```

## Documentation

### Assignment Reports
- [Assignment 3 Report](reports/week3/README.md)
- [Customer Review Summary](reports/week3/customer-review-summary.md)
- [Reflection](reports/week3/reflection.md)
- [Retrospective](reports/week3/retrospective.md)
- [LLM Usage Report](reports/week3/llm-report.md)
- [Manual Test Checklist](reports/week3/manual-test-checklist.md)

### Assignment 2 (Historical)
- [Assignment 2 Report](reports/week2/README.md)
- [User Stories](reports/week2/user-stories.md)
- [MVP v0 Report](reports/week2/mvp-v0-report.md)

### Project Management
- [User Stories Index](docs/user-stories.md)
- [Roadmap](docs/roadmap.md)
- [Definition of Done](docs/definition-of-done.md)
- [Changelog](CHANGELOG.md)
- [License](LICENSE)

## Contributing

1. Create a feature branch: `<issue-number>-short-description`
2. Commit changes with meaningful messages
3. Push and open a Pull Request linked to the related issue
4. Obtain at least one review approval before merging

### PR Requirements
- At least one approval from another team member
- All CI checks must pass
- Link to related issue
- Update `CHANGELOG.md` if user-visible change
- Verify acceptance criteria

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
