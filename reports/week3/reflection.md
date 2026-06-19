# Week 3 Reflection

## Learning Points

### Product Backlog Migration
- Migrated all active user stories from Assignment 2 to GitHub Issues with stable IDs preserved
- Learned the importance of traceability between requirements, issues, and PRs

### Go Backend Development
- Switched from FastAPI/React to Go standard library — simpler deployment, no frontend framework overhead
- Learned html/template rendering with shared layouts and block inheritance
- Built REST API without external routers by parsing URL paths manually

### Canvas Animation
- Implemented 35-object basketball animation with mouse repulsion physics
- Learned requestAnimationFrame loop optimization and canvas resize handling

### Mock ML Pipeline
- Designed async processing with goroutines and incremental progress updates
- Created JSON-based result storage for simplicity

### Testing
- Wrote 47 unit + integration tests using httptest
- Learned to test multipart file uploads and concurrent state updates

## Validated Assumptions

- **Go standard library is sufficient** — No external dependencies needed for MVP v1
- **Canvas animations work well** — Smooth 60fps with 35+ objects on modern browsers
- **Mock pipeline is viable** — Simulated processing demonstrates full user flow
- **JSON file storage works** — Sufficient for MVP, can migrate to database later
- **html/template is adequate** — Server-side rendering with JS enhancements works well

## Friction and Gaps

- **No real ML integration** — Mock pipeline needs replacement with MediaPipe or Python exec
- **No authentication** — Profile page is UI-only, needs session management
- **No persistence** — In-memory storage loses data on server restart
- **Video processing is simulated** — Real analysis requires ML model deployment
- **Mobile Canvas performance** — May need optimization for low-end devices

## Planned Response

### Next Sprint (Sprint 2)
1. Integrate real ML pipeline via Python exec
2. Add user authentication with sessions
3. Implement professional player comparison (US-005)
4. Add sharing capabilities (US-006, US-007)

### Process Improvements
1. Create estimation reference stories from Sprint 1 velocity data
2. Establish daily standup cadence
3. Increase code review participation

## Links to Affected PBIs

- [Sprint 1 Milestone](https://github.com/kr1ny77/BasketForm-AI/milestone/1)
- [v0.1.0 Release](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.1.0)
- [Roadmap](../../docs/roadmap.md)
