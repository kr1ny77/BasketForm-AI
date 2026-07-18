# Product Roadmap

## Vision

BasketForm-AI is an AI-powered platform that helps basketball players improve their shooting technique through video analysis and personalized feedback.

## Product Goal

Deliver a trustworthy, production-ready basketball shooting-form coach: real biomechanical analysis of genuine basketball shots, credible and durable feedback, and sharing/progress features that keep players coming back.

## Current Status (as of 2026-07-09)

Deployment: http://80.74.30.14/ Latest release: v0.3.0 (Sprint 3 — Assignment 5, MVP v2). Sprint 4 is complete: all 7 Sprint 4 PBIs Done. Sprint 5 (Week 7) is in progress — stabilization, deployment automation, documentation finalization, and Demo Day preparation for MVP v3 (v0.5.0).

MVP v0 (v0.0.1) provided basic video upload. MVP v1 (v0.1.0) delivered core analysis with Go backend, Canvas frontend, and a mock ML pipeline. v0.2.0 (Sprint 2) added authentication, social features, a real ML pipeline with phase analysis, annotated output video, PDF export, and automated quality gates. v0.3.0 (Sprint 3, MVP v2) added LLM-powered personalized feedback, the Friends section with report sharing, progress tracking, pro comparison, RU translation, architecture documentation, and ADRs.

## Sprint Plan

### Sprint 1 — MVP v1 (Completed)

Milestone: [Sprint 1](https://github.com/kr1ny77/BasketForm-AI/milestone/1). Dates: 2026-06-16 to 2026-06-22. Sprint Goal: deliver core basketball shooting-form analysis with Go backend, Canvas frontend, and mock ML pipeline. Outcome: all MVP v1 user stories Done. Release: v0.1.0.

### Sprint 2 — Assignment 4 (Completed)

Milestone: [Sprint 2](https://github.com/kr1ny77/BasketForm-AI/milestone/2). Dates: 2026-06-23 to 2026-06-29. Sprint Goal: add authentication, social features, enhanced ML with phase analysis, and automated quality gates. Outcome: all Sprint 2 PBIs Done. Release: v0.2.0.

### Sprint 3 — Assignment 5 / MVP v2 (Completed)

Milestone: [Sprint 3](https://github.com/kr1ny77/BasketForm-AI/milestone/3). Dates: 2026-07-06 to 2026-07-12. Sprint Goal: deliver MVP v2 (v0.3.0) by connecting analysis to an LLM for personalized feedback, adding a Friends section with report sharing, adding pro comparison and progress tracking, completing the RU translation, and documenting architecture and ADRs. Outcome: all 11 Sprint 3 PBIs Done (64 Story Points). Release: v0.3.0 (MVP v2).

### Sprint 4 — Assignment 6 (Week 6: Trial Release and Transition Readiness) — Completed

Milestone: [Sprint 4](https://github.com/kr1ny77/BasketForm-AI/milestone/4). Dates: 2026-07-13 to 2026-07-19.

Sprint Goal: deliver a stable, customer-usable trial release (pre-release v0.4.0-rc.1) together with transition-readiness evidence: a timeboxed, regression-safe multi-throw video upload with automatic ball-release segmentation (falling back to a simpler manual multi-file upload if not stable within its timebox), a completed customer-facing documentation set, and re-verified deployment/access steps so the product could be handed over on short notice.

Outcome: all 7 Sprint 4 PBIs Done (28 Story Points). Multi-throw auto-segmentation was deferred due to complexity; person-detection validation was delivered instead. CONTRIBUTING.md, AGENTS.md, and docs/customer-handover.md created. Customer confirmed product is ready for independent use.

Sprint Backlog board: [BasketForm-AI Product Backlog — Sprint Backlog view](https://github.com/users/kr1ny77/projects/7/views/2)

### Sprint 5 — Assignment 6 (Week 7: Final Transition and MVP v3) — In Progress

Milestone: [Sprint 5](https://github.com/kr1ny77/BasketForm-AI/milestone/5). Dates: 2026-07-20 to 2026-07-26.

Sprint Goal: use the Week 6 customer trial feedback, remaining fixes, and documentation updates to deliver the final course version of BasketForm-AI, MVP v3 (v0.5.0): triage and address customer feedback from the Week 6 trial, finish or roll back any unstable part of the multi-throw auto-segmentation feature, execute a concrete and inspectable product transition (access, deployment ownership, and a customer-handover artifact verified against the live deployment), keep contributor, agent, and customer-facing documentation current, and prepare and rehearse the Demo Day presentation.

Selected Sprint Backlog (6 PBIs, 28 Story Points): PBI-042 (#134) Triage and address Week 6 customer trial feedback — marked as expected follow-up scope, to be refined at Sprint 5 planning once Week 6 feedback exists, Must Have, 8 SP. PBI-043 (#135) Final MVP v3 stabilization and regression pass, Must Have, 5 SP. PBI-044 (#136) Final documentation and handover package finalization, Must Have, 5 SP. PBI-045 (#137) Final Demo Day preparation and walkthrough script, Must Have, 3 SP. PBI-046 (#138) Automate trial-release deployment pipeline for reliability, Should Have, 5 SP. PBI-047 (#139) Final course retrospective and lessons-learned report, Should Have, 2 SP.

Sprint Backlog board: [BasketForm-AI Product Backlog — Sprint Backlog view](https://github.com/users/kr1ny77/projects/7/views/2)

## MVP v3 (v0.5.0) — End-of-Course Scope

MVP v3 is the final release of the course engagement, delivered at the close of Sprint 5. The scope is justified by customer value, transition quality, and course closure. Customer value: the Week 6 trial release lets the customer try the product before final delivery, and Sprint 5 reacts to that real feedback (PBI-042) instead of assumptions; the multi-throw auto-segmentation feature (PBI-023) is delivered or safely rolled back to its manual-upload fallback within a fixed timebox to protect product stability. Transition and quality: complete customer-facing documentation (README, docs/customer-handover.md, CONTRIBUTING.md, AGENTS.md), re-verified deployment and access (PBI-040), an automated deployment pipeline (PBI-046), and a final stabilization/regression pass (PBI-043) so the product can be run, verified, and maintained without the team. Course closure: a rehearsed Demo Day walkthrough (PBI-045) and a final retrospective and lessons-learned report (PBI-047) close out the course engagement.

Evidence of completeness: all 13 Sprint 4 and Sprint 5 PBIs have expected outcomes, acceptance criteria, story point estimates, an implementer, and a different reviewer assigned. Both Sprints have defined start and finish dates and an explicit Sprint Goal. Progress is inspectable on the Product Backlog board and the Sprint Backlog board: [BasketForm-AI Product Backlog](https://github.com/users/kr1ny77/projects/7).

## State Reached by the End of the Course

By the end of Sprint 5 (2026-07-26), BasketForm-AI is expected to reach MVP v3 (v0.5.0): a stable, customer-usable product that has been trialed with the customer, transitioned with verified access and deployment, documented for independent operation, and demonstrated at the course Demo Day. This roadmap covers the scope of the course engagement only; product evolution beyond MVP v3 is outside the scope of Assignment 6 and is intentionally not planned here.

## Links

User Stories Index: [user-stories.md](user-stories.md). Definition of Done: [definition-of-done.md](definition-of-done.md). Quality Requirements: [quality-requirements.md](quality-requirements.md). CHANGELOG: [../CHANGELOG.md](../CHANGELOG.md).
