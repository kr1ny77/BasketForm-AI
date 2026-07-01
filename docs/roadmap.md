# Product Roadmap

## Vision

BasketForm-AI is an AI-powered platform that helps basketball players improve their shooting technique through video analysis and personalized feedback.

## Product Goal

Deliver a trustworthy, production-ready basketball shooting-form coach: real biomechanical analysis of genuine basketball shots, credible and durable feedback, and sharing/progress features that keep players coming back.

## Current Status (as of 2026-07-06)

Deployment: http://80.74.30.14/ Latest release: v0.2.0 (Sprint 2 — Assignment 4). Sprint 3 is now active.

MVP v0 (v0.0.1) provided basic video upload. MVP v1 (v0.1.0) delivered core analysis with Go backend, Canvas frontend, and a mock ML pipeline. v0.2.0 (Sprint 2) adds authentication, social features, real ML pipeline with phase analysis, annotated output video, PDF export, and automated quality gates. Sprint 3 is building MVP v2 (v0.3.0): LLM-powered personalized feedback, Friends + report sharing, progress tracking, pro comparison, RU translation, architecture documentation, and ADRs.

## Sprint Plan

### Sprint 1 — MVP v1 (Completed)

Milestone: [Sprint 1](https://github.com/kr1ny77/BasketForm-AI/milestone/1). Dates: 2026-06-16 to 2026-06-22. Sprint Goal: deliver core basketball shooting-form analysis with Go backend, Canvas frontend, and mock ML pipeline. Outcome: all MVP v1 user stories Done. Release: v0.1.0.

### Sprint 2 — Assignment 4 (Completed)

Milestone: [Sprint 2](https://github.com/kr1ny77/BasketForm-AI/milestone/2). Dates: 2026-06-23 to 2026-06-29. Sprint Goal: add authentication, social features, enhanced ML with phase analysis, and automated quality gates. Outcome: all Sprint 2 PBIs Done. Release: v0.2.0.

### Sprint 3 — Assignment 5 / MVP v2 (Active)

Milestone: [Sprint 3](https://github.com/kr1ny77/BasketForm-AI/milestone/3). Dates: 2026-07-06 to 2026-07-12.

Sprint Goal: Deliver MVP v2 (v0.3.0): make BasketForm-AI's feedback genuinely intelligent and social, and make the product explainable and maintainable. Concretely, the Sprint will (1) connect the shooting-form analysis to an LLM via an external API so users get personalized, coaching-style feedback; (2) add a Friends section and let users share their analysis reports with friends/coaches; (3) let users compare their form against professional reference players and track progress over time; (4) fix and complete the Russian (RU) translation; and (5) polish and harden the website-update experience and storage — all backed by documented architecture (static, dynamic, deployment views), Architecture Decision Records, and development-process / configuration-management documentation.

Selected Sprint Backlog (11 PBIs, 64 Story Points):

| PBI | Title | Priority | SP |
|-----|-------|----------|----|
| PBI-031 (#94) | Connect the LLM analysis via external API | Must Have | 8 |
| PBI-027 (#90) | Storage hardening — fix load deadlock, batch-delete cleanup, persistence reliability | Must Have | 5 |
| PBI-028 (#91) | Architecture documentation — static, dynamic & deployment views (PlantUML) | Must Have | 8 |
| PBI-029 (#92) | Architecture Decision Records (ADRs) for the key MVP v2 decisions | Must Have | 5 |
| PBI-032 (#95) | Build the "Friends" section (add / find / manage friends) | Should Have | 8 |
| PBI-033 (#96) | Share analysis reports with friends | Should Have | 5 |
| PBI-025 (#88) | Track shooting progress over time (US-008) | Should Have | 8 |
| PBI-026 (#89) | Website experience update — profile + avatar, localized UI (i18n), batch result management | Should Have | 5 |
| PBI-030 (#93) | Development-workflow & configuration-management documentation | Should Have | 3 |
| PBI-034 (#97) | Fix and complete the Russian (RU) translation | Should Have | 3 |
| PBI-024 (#87) | Compare shooting form against professional reference players (US-005) | Could Have | 8 |

Sprint Backlog board: [BasketForm-AI Product Backlog — Sprint Backlog view](https://github.com/users/kr1ny77/projects/7/views/2)

## MVP v2 (v0.3.0) Scope

MVP v2 is planned for release at the end of Sprint 3. The scope is justified by:

**Customer value**: Sprint 2 made the analysis real and credible — the customer can now trust a single result. The biggest remaining product value is personalization, retention, and social use: LLM-generated feedback makes the result feel personal; Friends + report sharing let players get feedback from a coach and compare with friends (US-006, US-007); progress tracking and pro comparison show whether you are improving (US-008, US-005).

**Quality and maintainability**: Architecture documentation (static, dynamic, deployment views as PlantUML diagrams), ADRs for key MVP v2 decisions, and development-workflow / configuration-management documentation reduce maintainability risk and support continued evolution. Storage hardening fixes a confirmed deadlock bug and cleanup gaps.

**Evidence of completeness**: All 11 PBIs have expected outcomes, acceptance criteria, story point estimates, implementers, and reviewers assigned. Sprint dates are defined (July 6–12, 2026). The total is 64 SP, consistent with Sprint 2 velocity.

## Next Expected Increment — Sprint 4 (Planned)

Focus: consolidate MVP v2 deferred work, grow test coverage, and address any architecture or deployment issues discovered in Sprint 3. Candidate items include: additional user stories (US-009 leaderboard, US-010 drill recommendations), expanding ML model accuracy, performance and load testing, and continued documentation of quality requirements.

## Architecture, Quality, and Process Work That Must Continue

This work is ongoing across later Sprints and must not be dropped: keep the GitHub Actions CI green on every PR; maintain and grow automated test coverage as features land and do not let it fall below the current baseline; keep the quality requirements and QRTs current; keep the non-basketball rejection and ball/player validation covered by tests; keep deployment reproducible so each increment can be released; keep architecture documentation (diagrams, ADRs) in sync with the code; and keep traceability in sync across docs, milestones, and CHANGELOG.

## Links

User Stories Index: user-stories.md. Definition of Done: definition-of-done.md. Quality Requirements: quality-requirements.md. CHANGELOG: ../CHANGELOG.md.
