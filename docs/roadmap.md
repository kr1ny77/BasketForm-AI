# Product Roadmap

## Vision

BasketForm-AI is an AI-powered platform that helps basketball players improve their shooting technique through video analysis and personalized feedback.

## Current Status

- **MVP v0:** Deployed with basic video upload functionality (FastAPI/Python)
- **MVP v1:** In development — core analysis features with Go backend and Canvas frontend

## MVP Versions

### MVP v0 (Completed)
- Basic video upload functionality
- Simple web interface
- Deployment at http://80.74.30.14/

### MVP v1 (Sprint 1 — Current)
**Goal:** Deliver core basketball shooting form analysis with Go backend, Canvas-animated frontend, and mock ML pipeline.

**Scope:**
- [US-001](https://github.com/kr1ny77/BasketForm-AI/issues/20): Upload shooting form video
- [US-002](https://github.com/kr1ny77/BasketForm-AI/issues/21): View automated form analysis
- [US-003](https://github.com/kr1ny77/BasketForm-AI/issues/22): Create and manage a user account
- [US-004](https://github.com/kr1ny77/BasketForm-AI/issues/23): Receive simplified actionable feedback
- [US-010](https://github.com/kr1ny77/BasketForm-AI/issues/28): Export analysis report as PDF

**Key Features:**
- Go HTTP server with html/template rendering
- Video upload with format validation and UUID naming
- Canvas basketball-themed background animation (30+ objects, mouse interaction)
- Mock ML pipeline (score, feedback, pose data)
- Progress polling via REST API
- Chart.js result visualization
- jsPDF export
- Dark theme with glass-morphism, mobile-first responsive design

### MVP v2 (Future)
**Goal:** Enhanced comparison and sharing features

**Planned Features:**
- [US-005](https://github.com/kr1ny77/BasketForm-AI/issues/24): Compare form with professional players
- [US-006](https://github.com/kr1ny77/BasketForm-AI/issues/25): Share analysis report with a coach
- [US-007](https://github.com/kr1ny77/BasketForm-AI/issues/26): Share progress with friends
- [US-008](https://github.com/kr1ny77/BasketForm-AI/issues/27): Track shooting progress over time

### MVP v3 (Future)
**Goal:** Community and social features

**Planned Features:**
- Advanced analytics dashboard
- Coach portal
- Team management
- Competition features

## Sprint Plan

### Sprint 1 (Current)
- **Duration:** 2 weeks (2026-06-16 — 2026-06-29)
- **Goal:** Deliver MVP v1 — core analysis with Go backend, Canvas frontend, mock ML
- **Focus:** Web server, all pages, upload/analysis API, tests, deployment
- **Milestone:** [Sprint 1](https://github.com/kr1ny77/BasketForm-AI/milestone/1)
- **PBIs:** US-001, US-002, US-003, US-004, US-010 + supporting technical PBIs

### Sprint 2 (Planned)
- **Duration:** 2 weeks
- **Focus:** Comparison features, sharing, progress tracking
- **Deliverables:** Professional comparison, coach sharing, social sharing, progress charts

### Sprint 3 (Future)
- **Duration:** 2 weeks
- **Focus:** Community features
- **Deliverables:** Social feed, advanced analytics, coach portal

## Technical Roadmap

### Phase 1: Foundation (MVP v0 → v1)
- Go HTTP server with standard library
- html/template rendering with shared layout
- Canvas basketball-themed animations
- REST API for upload, status, results
- Mock ML pipeline (score + feedback + pose)
- Unit and integration tests
- CI with golangci-lint + go test + go build

### Phase 2: Enhancement (MVP v2)
- Python ML script integration via exec
- Real biomechanical analysis
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
