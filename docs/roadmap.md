# Product Roadmap

## Vision

BasketForm-AI is an AI-powered platform that helps basketball players improve their shooting technique through video analysis and personalized feedback.

## Current Status

- **MVP v0:** Deployed with basic video upload (v0.0.1)
- **MVP v1:** Completed and deployed (v0.1.0) — [Release](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.1.0)

## MVP Versions

### MVP v0 (Completed — v0.0.1)
- Basic video upload functionality
- Simple web interface
- Deployment at http://80.74.30.14/

### MVP v1 (Completed — v0.1.0)
**Goal:** Deliver core basketball shooting form analysis with Go backend, Canvas frontend, and mock ML pipeline.

**Scope (all Done):**
- [US-001](https://github.com/kr1ny77/BasketForm-AI/issues/20): Upload shooting form video
- [US-002](https://github.com/kr1ny77/BasketForm-AI/issues/21): View automated form analysis
- [US-003](https://github.com/kr1ny77/BasketForm-AI/issues/22): Create and manage a user account
- [US-004](https://github.com/kr1ny77/BasketForm-AI/issues/23): Receive simplified actionable feedback
- [US-010](https://github.com/kr1ny77/BasketForm-AI/issues/28): Export analysis report as PDF

**Delivered Features:**
- Go HTTP server with html/template rendering
- Video upload with drag-and-drop, format validation, UUID naming
- Canvas basketball-themed background animation (35 objects, mouse interaction)
- Mock ML pipeline (score, feedback, pose data)
- REST API with upload, status, results, and video list endpoints
- Progress polling with real-time UI updates
- Results page with Chart.js radar + scatter charts
- PDF export via jsPDF
- Profile page (demo mode)
- Dark glass-morphism theme, mobile-first responsive
- 47 unit + integration tests
- Manual test checklist with 50+ scenarios

### MVP v2 (Sprint 2 — Planned)
**Goal:** Enhanced comparison and sharing features

**Planned Features:**
- [US-005](https://github.com/kr1ny77/BasketForm-AI/issues/24): Compare form with professional players
- [US-006](https://github.com/kr1ny77/BasketForm-AI/issues/25): Share analysis report with a coach
- [US-007](https://github.com/kr1ny77/BasketForm-AI/issues/26): Share progress with friends
- [US-008](https://github.com/kr1ny77/BasketForm-AI/issues/27): Track shooting progress over time

### MVP v3 (Sprint 3 — Future)
**Goal:** Community and social features

**Planned Features:**
- Advanced analytics dashboard
- Coach portal
- Team management
- Competition features

## Sprint Plan

### Sprint 1 (Completed)
- **Duration:** 2026-06-16 — 2026-06-29
- **Goal:** Deliver MVP v1 — core analysis with Go backend, Canvas frontend, mock ML
- **Milestone:** [Sprint 1](https://github.com/kr1ny77/BasketForm-AI/milestone/1)
- **Release:** [v0.1.0](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.1.0)

### Sprint 2 (Planned)
- **Duration:** 2 weeks
- **Focus:** Comparison features, sharing, progress tracking
- **Deliverables:** Professional comparison, coach sharing, social sharing, progress charts

### Sprint 3 (Future)
- **Duration:** 2 weeks
- **Focus:** Community features
- **Deliverables:** Social feed, advanced analytics, coach portal

## Technical Roadmap

### Phase 1: Foundation (MVP v0 → v1) — DONE
- Go HTTP server with standard library
- html/template rendering with shared layout
- Canvas basketball-themed animations
- REST API for upload, status, results
- Mock ML pipeline (score + feedback + pose)
- Unit and integration tests
- CI with golangci-lint + go test + go build

### Phase 2: Enhancement (MVP v2)
- Python ML script integration via exec
- Real biomechanical analysis with MediaPipe
- Professional player comparison library
- Sharing mechanisms
- Progress tracking over time

### Phase 3: Scale (MVP v3)
- Coach portal
- Team features
- Competition system
- Potential mobile app

## Links

- [User Stories Index](user-stories.md)
- [Definition of Done](definition-of-done.md)
- [CHANGELOG](../CHANGELOG.md)
- [Repository](https://github.com/kr1ny77/BasketForm-AI)
