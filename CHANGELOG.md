# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/).

## [Unreleased]

## [v0.3.0] — 2026-07-04

Sprint 3 — Assignment 5: Architecture documentation, ADRs, LLM-based scoring, UI/UX improvements.

### Added
- **LLM-based scoring**: integrated OpenRouter API (GPT-4o-mini) for video analysis — LLM receives 8 key frames + biomechanical metrics and generates scores + detailed feedback
- **Vision analysis**: LLM analyzes actual video frames, not just numerical metrics
- **Follow-through scoring**: 4th phase now fully scored with duration, elbow snap, and arm stability metrics
- **Nickname change**: POST /api/profile/nickname endpoint with password verification and uniqueness check
- **Profile page modals**: Change Nickname and Change Password as popup modals with large action buttons
- **Shared results — sent view**: users can now see results they shared with friends (not just received)
- **Architecture documentation** (Part 4): component diagram, sequence diagram, deployment diagram (PlantUML)
- **Architecture Decision Records** (Part 5): ADR-001 (JSON storage), ADR-002 (exec ML), ADR-003 (JWT auth)
- **Case-insensitive login**: email comparison uses `strings.EqualFold`; login accepts both email and nickname
- **Avatar display**: avatars now visible on Friends and Shared pages (fixed missing `onload` handler)
- **Quality requirements updated**: linked ADRs to QR-001, QR-002

### Fixed
- **Data race in GetVideo**: `RLock` while writing to map caused panic with 3+ concurrent users — changed to `Lock`
- **Delete results**: recompiled Go binary; removed duplicate `handlers/storage.go` with wrong package declaration
- **Friends page**: fixed `t is not defined` error in `loadMyResults()` — Share button now works
- **Avatar upload**: `ProfileHandler` now returns `avatar` field in JSON response
- **Login field**: i18n updated to "Email or Nickname" / "Эл. почта или никнейм"
- **Checkbox styling**: enlarged to 20px, centered checkmark, rounded corners
- **Score display**: gradient matches other pages, `/100` repositioned with proper spacing
- **Delete bar**: correctly stays visible in selection mode

### Changed
- ML pipeline: `feedback_generator.py` rewritten to use OpenRouter API instead of local Qwen model
- ML pipeline: removed `llm_scorer.py` — scoring now handled by existing `feedback_generator.py`
- Profile page: nickname/password buttons moved to modal popups below profile card
- Server: removed duplicate `internal/handlers/storage.go` that blocked Go compilation
- `shot_analyzer.py`: `FOLLOW_THROUGH` included in initial scores dict; added `elbow_snap` and `arm_stability_std` metrics

### Security
- API key stored in environment variable (`OPENROUTER_API_KEY`), not in source code
- Data race fix prevents server crashes under concurrent load

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
