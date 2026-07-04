# BasketForm-AI

Web service for analyzing basketball shooting form via video upload, biomechanical keypoint extraction, technique scoring, and personalized AI feedback.

**Live:** http://80.74.30.14/ · [v0.3.0 Release](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.3.0)

## Project Overview

BasketForm-AI is an AI-powered platform that helps basketball players improve their shooting technique through video analysis and personalized feedback. Users upload shooting videos, the system extracts biomechanical keypoints, evaluates stance, arm angle, release point, and follow-through, then generates actionable recommendations using LLM vision analysis.

## Key Features

- **User Accounts:** Registration with email/nickname/password, JWT-based login, case-insensitive email, login by nickname
- **Video Upload:** Drag-and-drop or file picker for basketball shot videos (MP4, MOV, AVI)
- **Biomechanical Analysis:** Automatic extraction of key body points using MediaPipe
- **LLM Vision Analysis:** GPT-4o-mini analyzes 8 key frames from the video for accurate scoring
- **Technique Evaluation:** Scoring of stance, arm angle, release point, and follow-through (0–100)
- **Phase Analysis:** Detailed per-phase scoring with personalized AI feedback
- **Social Features:** Search users, friend requests, share results with friends
- **PDF Export:** Download analysis reports with full score breakdown and feedback
- **Profile Management:** Avatar upload, nickname change (with password confirmation), password change
- **Dark Theme:** Glass-morphism design, orange accents, Canvas background animation

## Tech Stack

- **Backend:** Go 1.22+ (standard library `net/http`)
- **Frontend:** HTML + CSS + JavaScript, Canvas animations, Chart.js, jsPDF
- **ML:** MediaPipe Holistic + OpenCV + YOLO (Python scripts via `exec`)
- **LLM:** OpenRouter API — GPT-4o-mini (vision model for frame analysis)
- **Auth:** bcrypt password hashing, JWT tokens (HMAC-SHA256), HttpOnly cookies
- **Storage:** Local JSON files in `data/` directory
- **Testing:** Go `testing`, `httptest` (unit + integration + QRT tests)
- **CI:** GitHub Actions — golangci-lint, test, coverage, QRT, govulncheck, Lychee
- **Deployment:** Binary on VM with Nginx reverse proxy
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

- Go 1.22 or later
- Python 3.12+ with MediaPipe, OpenCV, YOLO (for ML analysis)
- `golangci-lint` (for linting, optional)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/kr1ny77/BasketForm-AI.git
   cd BasketForm-AI
   ```

2. **Install Python dependencies:**
   ```bash
   pip install -r requirements.txt
   ```

3. **Set environment variables:**
   ```bash
   export OPENROUTER_API_KEY="your-api-key"
   export JWT_SECRET="your-secret"
   ```

4. **Run the server:**
   ```bash
   go run ./cmd/server/
   ```

5. **Open in browser:**
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
# Unit and integration tests
go test -race -coverprofile=coverage.out ./...

# Quality requirement tests
go test -tags=qrt -v ./internal/qrt/...
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
| `DATA_DIR` | `data` | User/video/friend JSON storage |
| `JWT_SECRET` | default | JWT signing secret |
| `OPENROUTER_API_KEY` | — | OpenRouter API key for LLM scoring |

## Deployment

### Production (VM at http://80.74.30.14/)

```bash
# Build
go build -o bin/server ./cmd/server/

# Run with env vars
export OPENROUTER_API_KEY="your-key"
PORT=8080 ./bin/server
```

Nginx reverse proxy configured on port 80 → 8080.

### Docker

```bash
docker build -t basketform-ai .
docker run -p 8080:8080 -e OPENROUTER_API_KEY=your-key basketform-ai
```

## Project Structure

```
BasketForm-AI/
├── cmd/server/main.go              # Application entry point
├── internal/
│   ├── handlers/                   # HTTP handlers
│   │   ├── handlers.go             # Page handlers + AvatarHandler
│   │   ├── api.go                  # API handlers (upload, status, result, videos, delete)
│   │   ├── auth.go                 # Auth handlers (register, login, profile, nickname, avatar)
│   │   ├── friends.go              # Friends handlers
│   │   ├── share.go                # Share handlers (with-me, by-me)
│   │   ├── middleware.go           # JWT auth middleware
│   │   └── *_test.go               # Tests
│   ├── models/                     # Data structures
│   ├── services/                   # Business logic
│   │   ├── auth.go                 # Auth service (bcrypt + JWT)
│   │   ├── storage.go              # JSON file storage
│   │   └── processor.go            # ML video processing
│   └── qrt/                        # Quality requirement tests
├── ML/                             # Python ML pipeline
│   ├── main.py                     # Entry point for video analysis
│   ├── ball_tracker.py             # YOLO-based ball detection
│   ├── pose_tracker.py             # MediaPipe pose estimation
│   ├── shot_analyzer.py            # State machine for shot phases
│   ├── feedback_generator.py       # LLM scoring via OpenRouter API
│   ├── custom_feedback.py          # Rule-based fallback scoring
│   └── best.pt                     # YOLO model weights
├── web/
│   ├── templates/                  # HTML templates
│   │   ├── login.html, signup.html # Auth pages
│   │   ├── upload.html             # Video upload
│   │   ├── results.html            # Analysis results
│   │   ├── friends.html            # Friend management
│   │   ├── shared.html             # Shared results
│   │   └── profile.html            # User profile (with modals)
│   └── static/                     # CSS, JS, images
├── data/                           # JSON user/video/friend storage
├── uploads/                        # Uploaded video files
├── results/                        # Analysis result files
├── docs/                           # Project documentation
│   ├── architecture/               # Architecture docs + diagrams
│   │   ├── README.md               # Architecture overview
│   │   ├── static-view/            # Component diagram
│   │   ├── dynamic-view/           # Sequence diagram
│   │   ├── deployment-view/        # Deployment diagram
│   │   └── adr/                    # Architecture Decision Records
│   ├── quality-requirements.md     # QR-001 to QR-004
│   └── ...
├── reports/                        # Course reports
├── Dockerfile                      # Container build
├── CHANGELOG.md                    # Version history
└── LICENSE                         # MIT License
```

## Documentation

### Architecture
- [Architecture Overview](docs/architecture/README.md)
- [Component Diagram](docs/architecture/static-view/component-diagram.puml)
- [Sequence Diagram](docs/architecture/dynamic-view/sequence-diagram.puml)
- [Deployment Diagram](docs/architecture/deployment-view/deployment-diagram.puml)
- [ADR-001: JSON Storage](docs/architecture/adr/ADR-001-use-json-storage.md)
- [ADR-002: exec ML Integration](docs/architecture/adr/ADR-002-exec-ml-script.md)
- [ADR-003: JWT Authentication](docs/architecture/adr/ADR-003-jwt-auth.md)

### Reports
- [Week 4 Report](reports/week4/README.md)
- [Customer Review Summary](reports/week4/customer-review-summary.md)
- [Reflection](reports/week4/reflection.md)
- [Retrospective](reports/week4/retrospective.md)
- [LLM Usage Report](reports/week4/llm-report.md)
- [User Acceptance Tests](docs/user-acceptance-tests.md)

### Project Management
- [User Stories Index](docs/user-stories.md)
- [Roadmap](docs/roadmap.md)
- [Definition of Done](docs/definition-of-done.md)
- [Quality Requirements](docs/quality-requirements.md)
- [Quality Requirement Tests](docs/quality-requirement-tests.md)
- [Testing Strategy](docs/testing.md)
- [Changelog](CHANGELOG.md)

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
