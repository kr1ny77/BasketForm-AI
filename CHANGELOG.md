# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/).

## [Unreleased]

## [v0.1.0] — 2026-06-19

MVP v1 — Core basketball shooting form analysis with Go backend, Canvas frontend, and mock ML pipeline.

### Added
- Go HTTP server with `html/template` rendering and shared layout
- Video upload page with drag-and-drop, format validation (MP4/MOV/AVI), and progress indicator
- Basketball-themed Canvas background animation (35 objects, mouse repulsion, orange waves)
- REST API: `POST /api/upload`, `GET /api/status/{id}`, `GET /api/result/{id}`, `GET /api/videos`
- Mock ML pipeline: random score (40–90), feedback generation, 12-point pose skeleton
- Background processing with incremental progress (0→100%) and goroutine-based async
- Results page with filterable/sortable grid, radar chart (Chart.js), pose scatter plot
- PDF export using jsPDF with score breakdown and feedback
- Profile page with account settings UI (demo mode)
- Progress page with real-time status polling
- Dark theme with glass-morphism, orange accents, mobile-first responsive design
- Unit tests for UUID generation, file extension validation, storage, processor, models
- Integration tests for all API endpoints using `httptest`
- Manual test checklist with 50+ scenarios
- GitHub Actions CI: golangci-lint, go test, go build
- Lychee link-checking CI for all Markdown files
- Issue templates: User Story, Bug Report, Other PBI, Course Task
- PR template with related issue, acceptance criteria, and changelog checklist
- Definition of Done, Roadmap, User Stories index
- Dockerfile for containerized deployment
- Python mock processing script (`scripts/process_video.py`)

### Changed
- Migrated project from FastAPI/React/PostgreSQL to Go standard library stack
- Updated all documentation to reflect Go-based architecture
- Revised MVP v1 scope based on customer feedback (PDF export replaces social sharing)

### Removed
- US-009: Public social feed of user shots (removed due to privacy concerns)

## [v0.0.1] — 2026-06-12

MVP v0 — Initial deployment with basic video upload.

### Added
- Initial project structure and repository setup
- MIT License
- Basic README with project overview
- Assignment 2 reports
- MVP v0 deployment at http://80.74.30.14/
- Figma interactive prototype
- OpenAPI specification and Postman collection
