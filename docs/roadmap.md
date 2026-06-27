# Product Roadmap

## Vision

BasketForm-AI is an AI-powered platform that helps basketball players improve their shooting technique through video analysis and personalized feedback.

## Product Goal

Deliver a trustworthy, production-ready basketball shooting-form coach: real biomechanical analysis of genuine basketball shots, credible and durable feedback, and sharing/progress features that keep players coming back.

## Current Status (as of 2026-06-28)

MVP v0 (v0.0.1) provides basic video upload, deployed at http://80.74.30.14/. MVP v1 (v0.1.0) delivers core analysis with a Go backend, Canvas frontend, and a mock ML pipeline; it is completed and released (see the v0.1.0 release). The current direction is to harden MVP v1 into a trustworthy product by replacing the mock pipeline with real basketball-shot analysis and adding result persistence. The customer accepted the mock ML for the demo only and explicitly asked for real analysis, basketball/player detection, and rejection of non-basketball videos.

## Sprint Plan

### Sprint 1 - MVP v1 (Completed)

Milestone: https://github.com/kr1ny77/BasketForm-AI/milestone/1 . Dates: 2026-06-16 to 2026-06-29. Sprint Goal: deliver core basketball shooting-form analysis with Go backend, Canvas frontend, and mock ML pipeline. Outcome: all MVP v1 user stories Done (US-001, US-002, US-003, US-004, US-010), packaged as SemVer release v0.1.0.

### Sprint 2 - Assignment 4 (Current Sprint)

Milestone (authoritative Sprint container): https://github.com/kr1ny77/BasketForm-AI/milestone/2 . Dates: 2026-06-29 (Mon) to 2026-07-05 (Sun). Sprint Goal: make the analysis trustworthy and durable - replace the mock pipeline with real basketball-shot analysis that detects and tracks the ball, clearly rejects non-basketball videos instead of scoring them, and persists results across server restarts. Sprint Backlog board (GitHub repository Issues view): https://github.com/kr1ny77/BasketForm-AI/issues/views/775 - it shows the issues in Milestone 2 with priority, MVP version, milestone, and assignee. Planned items (see each issue for expected outcome, acceptance criteria, Story Points, implementer, reviewer, and Work Status): PBI-016 (#64) real Python ML pipeline; PBI-017 (#65) basketball + player detection and non-basketball validation; PBI-018 (#66) result persistence across restarts; PBI-019 (#67) automated tests for ML, validation, persistence; PBI-020 (#68) deploy real ML runtime to production; PBI-021 (#62) track the basketball and refine the release-phase algorithm. Packaged increment: a SemVer release (planned v0.2.0) is created after the Sprint work is Done. The milestone is planning and scope evidence; the SemVer release is the packaged increment evidence.

### Sprint 3 - Expected Next Sprint (Planned)

Focus: build value features on top of the now-trustworthy analysis - comparison and sharing. Candidate items: US-005 compare with professional players (#24), US-006 share report with a coach (#25), US-007 share progress with friends (#26), US-008 track progress over time (#27). Any Sprint 2 should-have work not finished (for example deployment polish) carries over here.

## Quality and Automation Work That Must Continue

This work is ongoing across later Sprints and must not be dropped: keep the GitHub Actions CI green (golangci-lint, go test -race, go build, link check) on every PR; maintain and grow automated test coverage as features land and do not let it fall below the current baseline, with new ML, validation, and persistence behaviour covered by tests (PBI-019); keep the non-basketball rejection and ball/player validation covered by positive and negative sample tests; keep deployment reproducible (Docker plus documented run instructions) so each increment can be released; and keep traceability in sync across docs/user-stories.md, this roadmap, the milestones, and CHANGELOG.

## Links

User Stories Index: user-stories.md . Definition of Done: definition-of-done.md . CHANGELOG: ../CHANGELOG.md . Sprint 1 milestone: https://github.com/kr1ny77/BasketForm-AI/milestone/1 . Sprint 2 (Assignment 4) milestone: https://github.com/kr1ny77/BasketForm-AI/milestone/2 . Repository: https://github.com/kr1ny77/BasketForm-AI
