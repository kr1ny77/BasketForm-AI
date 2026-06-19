# BasketForm-AI

Web service for analyzing basketball shooting form via video upload, biomechanical keypoint extraction, technique scoring, and personalized AI feedback.

## Project Overview

BasketForm-AI is an AI-powered platform that helps basketball players improve their shooting technique through video analysis and personalized feedback. Users upload shooting videos, the system extracts biomechanical keypoints, evaluates stance, arm angle, release point, and follow-through, then generates simplified actionable recommendations.

## Key Features

- **Video Upload:** Drag-and-drop or file picker for basketball shot videos
- **Biomechanical Analysis:** Automatic extraction of key body points from video
- **Technique Evaluation:** Scoring of stance, arm angle, release point, and follow-through
- **Personalized Feedback:** AI-generated recommendations and drills (max 3 actionable takeaways)
- **Progress Tracking:** Monitor improvement over time
- **PDF Export:** Download analysis reports for offline reference

## Tech Stack

- **Backend:** Go 1.21+ (standard library + html/template)
- **Frontend:** HTML + CSS + JavaScript, Canvas animations (basketball theme)
- **Storage:** Local `uploads/` and `results/` directories
- **ML:** Mock data stub or Python script via `exec`
- **Testing:** Go `testing` package, `httptest`
- **CI:** GitHub Actions (golangci-lint, go test, go build, Lychee link check)
- **License:** MIT

## Project Structure

```
BasketForm-AI/
├── cmd/server/main.go          # Application entry point
├── internal/
│   ├── handlers/               # HTTP handlers
│   ├── models/                 # Data structures
│   └── services/               # Business logic
├── web/
│   ├── templates/              # HTML templates (html/template)
│   └── static/                 # CSS, JS, images
├── uploads/                    # Uploaded video files (gitignored)
├── results/                    # Analysis result JSON files (gitignored)
├── reports/                    # Course reports and evidence
├── docs/                       # Project documentation
├── scripts/                    # Helper scripts (e.g., process_video.py)
├── tests/                      # Test files
├── CHANGELOG.md                # Version history
├── LICENSE                     # MIT License
├── .gitignore                  # Git ignore rules
├── .env.example                # Environment variables template
└── .github/                    # CI workflows, issue/PR templates
```

## Local Development Setup

### Prerequisites

- Go 1.21 or later
- `golangci-lint` (for linting)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/kr1ny77/BasketForm-AI.git
   cd BasketForm-AI
   ```

2. **Initialize Go module (if needed):**
   ```bash
   go mod init github.com/kr1ny77/BasketForm-AI
   go mod tidy
   ```

3. **Create required directories:**
   ```bash
   mkdir -p uploads results
   ```

4. **Run the server:**
   ```bash
   go run ./cmd/server/
   ```
   The server starts on `http://localhost:8080` by default.

5. **Build a binary:**
   ```bash
   go build -o bin/server ./cmd/server/
   ./bin/server
   ```

6. **Run tests:**
   ```bash
   go test -race -coverprofile=coverage.out ./...
   ```

7. **Run linter:**
   ```bash
   golangci-lint run
   ```

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `UPLOAD_DIR` | `uploads` | Video upload directory |
| `RESULTS_DIR` | `results` | Analysis results directory |

Copy `.env.example` to `.env` and adjust as needed:
```bash
cp .env.example .env
```

## Documentation

### Assignment Reports
- [Assignment 3 Report Index](reports/week3/README.md)
- [Customer Review Summary](reports/week3/customer-review-summary.md)
- [Reflection](reports/week3/reflection.md)
- [Retrospective](reports/week3/retrospective.md)
- [LLM Usage Report](reports/week3/llm-report.md)

### Assignment 2 (Historical)
- [Assignment 2 Report Index](reports/week2/README.md)
- [User Stories](reports/week2/user-stories.md)
- [MVP v0 Report](reports/week2/mvp-v0-report.md)

### Project Management
- [User Stories Index](docs/user-stories.md)
- [Roadmap](docs/roadmap.md)
- [Definition of Done](docs/definition-of-done.md)
- [Changelog](CHANGELOG.md)
- [License](LICENSE)

## Deployment

### Local
```bash
go build -o bin/server ./cmd/server/
PORT=8080 ./bin/server
```

### Docker
```bash
docker build -t basketform-ai .
docker run -p 8080:8080 basketform-ai
```

### Production
See the [Week 3 Report](reports/week3/README.md) for deployment instructions and the live URL.

## Contributing

1. Fork the repository
2. Create a feature branch: `<issue-number>-short-description` (e.g., `42-add-login-form`)
3. Commit changes
4. Push and open a Pull Request linked to the related issue
5. Obtain at least one review approval before merging

### PR Requirements
- At least one approval from another team member
- All CI checks must pass (golangci-lint, go test, go build)
- Link to related issue
- Update `CHANGELOG.md` if user-visible change
- Verify acceptance criteria

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

## Contact

- GitHub Issues: [https://github.com/kr1ny77/BasketForm-AI/issues](https://github.com/kr1ny77/BasketForm-AI/issues)
