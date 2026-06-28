# Product Roadmap

## Vision

BasketForm-AI is an AI-powered platform that helps basketball players improve their shooting technique through video analysis and personalized feedback.

## Product Goal

Deliver a trustworthy, production-ready basketball shooting-form coach: real biomechanical analysis of genuine basketball shots, credible and durable feedback, and sharing/progress features that keep players coming back.

## Current Status (as of 2026-06-28)

Deployment: http://80.74.30.14/
Latest release: v0.2.0 (Sprint 2 — Assignment 4)

MVP v0 (v0.0.1) provided basic video upload. MVP v1 (v0.1.0) delivered core analysis with Go backend, Canvas frontend, and a mock ML pipeline. v0.2.0 (Sprint 2) adds authentication, social features, real ML pipeline with phase analysis, annotated output video, PDF export, and automated quality gates.

## Sprint Plan

### Sprint 1 — MVP v1 (Completed)

Milestone: [Sprint 1](https://github.com/kr1ny77/BasketForm-AI/milestone/1). Dates: 2026-06-16 to 2026-06-22. Sprint Goal: deliver core basketball shooting-form analysis with Go backend, Canvas frontend, and mock ML pipeline. Outcome: all MVP v1 user stories Done. Release: v0.1.0.

### Sprint 2 — Assignment 4 (Completed)

Milestone: [Sprint 2](https://github.com/kr1ny77/BasketForm-AI/milestone/2). Dates: 2026-06-23 to 2026-06-29. Sprint Goal: add authentication, social features, enhanced ML with phase analysis, and automated quality gates. Outcome: all Sprint 2 PBIs Done. Release: v0.2.0.

### Sprint 3 — Expected Next Sprint (Planned)

Focus: build value features on top of the now-trustworthy analysis — comparison and progress tracking. Candidate items: US-005 compare with professional players, US-008 track progress over time. Any deferred work carries over.

## Quality and Automation Work That Must Continue

This work is ongoing across later Sprints and must not be dropped: keep the GitHub Actions CI green on every PR; maintain and grow automated test coverage as features land and do not let it fall below the current baseline; keep the quality requirements and QRTs current; keep the non-basketball rejection and ball/player validation covered by tests; keep deployment reproducible so each increment can be released; and keep traceability in sync across docs, milestones, and CHANGELOG.

## Links

User Stories Index: [user-stories.md](user-stories.md). Definition of Done: [definition-of-done.md](definition-of-done.md). Quality Requirements: [quality-requirements.md](quality-requirements.md). CHANGELOG: [../CHANGELOG.md](../CHANGELOG.md).
