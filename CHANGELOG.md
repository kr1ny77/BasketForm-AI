# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/).

## [Unreleased]

### Fixed
- README.md release link updated from v0.3.0 to v0.4.0
- Added Week 7 report link to README.md Reports section

### Documentation
- CONTRIBUTING.md: added deployment section referencing `scripts/deploy.sh`
- AGENTS.md: added `scripts/deploy.sh` to key files table and deployment guidance
- Added regression report for PBI-043 fixes (`docs/regression-report-pbi-043.md`)
- Added Should-Have bug triage for PBI-043 (`docs/bug-triage-pbi-043.md`)

## [v0.4.0] — 2026-07-19

Sprint 4 — Assignment 6: Week 6 trial / handover-candidate release.

### Added
- YOLO-based human presence check: videos without a detected person are rejected with a clear error message instead of producing fake analysis results
- CONTRIBUTING.md, AGENTS.md, and docs/customer-handover.md (maintained customer-handover artifact)

### Changed  
- Friends section redesigned, navigation centered, profile page modals polished
- LLM prompt/update tweaks and documentation fixes for a smoother experience

### Documentation
- Finalized customer-facing documentation set: handover guide, Week 6 report, sprint review materials, and updated user acceptance tests

## [v0.5.0] — 2026-07-26

Sprint 5 — Assignment 6: Stabilization, deployment automation, and MVP v3.

### Fixed
- Removed hardcoded deployment path `/home/basketfrom-ai/BasketForm-AI` from processor — ML runner now uses the current working directory, fixing test failures and making the product portable
- Race condition in handler tests caused by hardcoded path in processor

### Changed
- Deploy script rewritten with ML dependency installation, smoke test, and error handling
- Processor log initialization removed (was writing to hardcoded path)

## [v0.3.0] — 2026-07-04

Sprint 3 — Assignment 5: Architecture documentation, ADRs, LLM-based scoring, UI/UX improvements.

### Added
- LLM-based scoring via OpenRouter API (GPT-4o-mini) with 8-frame vision analysis
- Follow-through scoring (4th phase with duration, elbow snap, arm stability)
- Nickname change endpoint with password verification
- Profile page modals for nickname/password changes
- Shared results: sent view for users
- Architecture documentation (component, sequence, deployment diagrams in PlantUML)
- ADR-001 (JSON storage), ADR-002 (exec ML), ADR-003 (JWT auth)
- Case-insensitive login, login by nickname
- Avatar display on Friends and Shared pages

### Fixed
- Data race in GetVideo (RLock while writing to map)
- Delete results (recompiled Go binary, removed duplicate storage.go)
- Friends page: t is not defined error in loadMyResults()
- Avatar upload: ProfileHandler returns avatar field

### Changed
- feedback_generator.py rewritten for OpenRouter API
- shot_analyzer.py: FOLLOW_THROUGH in initial scores, elbow_snap and arm_stability_std metrics
- Profile page: nickname/password buttons as modal popups

### Security
- API key in environment variable, not source code
- Data race fix prevents server crashes

## [v0.2.0] — 2026-06-28

Sprint 2 — Assignment 4: Authentication, social features, enhanced ML, quality automation.

### Added
- User registration and login with email, nickname, and password (bcrypt hashing)
- JWT-based session authentication with HttpOnly cookies
- Protected routes: all pages except /login, /register, /static require authentication
- User profile page with nickname, email display, and password change
- Friend system: search users by nickname, send/accept/reject friend requests
- Result sharing: share analysis results with friends (read-only access)
- Shared results page showing results shared by friends
- Enhanced ML agent: generates output video with pose overlay (keypoints, skeleton, HUD)
- Phase analysis: detailed scoring for Stance, Arm Angle, Release, Follow-through
- Phase feedback: personalized text feedback per phase
- PDF export: jsPDF-based report with user info, score, phases, and full feedback
- Share button in result modal with friend selection dropdown
- Quality requirements document (QR-001 to QR-004) following ISO/IEC 25010
- Quality requirement tests (QRT-001 to QRT-004) automated in CI
- Unit tests for authentication service (9 tests)
- Integration tests for API endpoints (5 tests: register, login, profile, share, friends)
- QRT tests for response time, auth security, coverage, form guidance
- Updated CI pipeline: lint, test, coverage (30%), QRT, govulncheck, build
- Updated Definition of Done with Assignment 4 quality gates
- User acceptance test scenarios (UAT-001 to UAT-004)

### Changed
- Video model now includes UserID field for per-user video ownership
- Results include UserID and phase data
- Storage layer uses JSON files in data/ directory for persistence
- Upload handler associates videos with authenticated user
- Results endpoint returns output_video_url and phases
- All page handlers and API handlers require JWT authentication
- Navigation updated: added Friends and Shared links

### Security
- All API endpoints require authentication (401 for unauthenticated requests)
- Passwords hashed with bcrypt before storage
- JWT tokens signed with HMAC-SHA256
- HttpOnly cookies prevent XSS token theft

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
