# Week 4 Public Report — BasketForm-AI

Project: **BasketForm-AI** — AI-powered basketball shooting form analysis platform.
Deployment: http://80.74.30.14/

## Sprint Context

- **Sprint:** Sprint 2 — Assignment 4 (Auth, Social, ML, Quality)
- **Sprint Goal:** Add authentication, social features, enhanced ML with phase analysis, and automated quality gates.
- **Sprint milestone:** [Sprint 2 - Assignment 4](https://github.com/kr1ny77/BasketForm-AI/milestone/2)
- **Sprint Backlog:** [Issues view](https://github.com/kr1ny77/BasketForm-AI/issues/views/775)
- **Sprint dates:** 2026-06-29 to 2026-07-05
- **Total Sprint size:** 34 Story Points

## Delivered Product Changes

- User registration and login with email, nickname, password (bcrypt + JWT)
- Friend system: search, request, accept/reject
- Result sharing with friends (read-only)
- Enhanced ML agent with annotated output video and phase analysis (Stance, Arm Angle, Release, Follow-through)
- PDF export with full score breakdown
- Quality requirements (QR-001 to QR-004) following ISO/IEC 25010
- Automated QRT tests (QRT-001 to QRT-004)
- Updated CI pipeline with lint, test, coverage, QRT, govulncheck
- Unit tests (auth: 9 tests) and integration tests (5 tests)

## Deployed Product

- **URL:** http://80.74.30.14/
- **Run instructions:** See [README.md](../../README.md)

## Customer Feedback Response Table

| Feedback point | Resulting PBI or issue | Status | Response |
|---|---|---|---|
| Real ML analysis needed; mock pipeline is only acceptable for demo. | #64 (PBI-016) | Done | Replaced mock with real MediaPipe pipeline generating annotated video and phase scores. |
| Reject non-basketball videos instead of scoring them. | #65 (PBI-017) | Done | Added player detection; non-basketball videos get rejected. |
| Authentication and data persistence across restarts. | #66 (PBI-018) | Done | Added JWT auth, bcrypt, JSON file persistence in data/ directory. |
| Share results with coaches/friends. | #67 (PBI-019) | Done | Added friend system and result sharing with read-only access. |
| PDF export for offline record-keeping. | #70 (PBI-022) | Done | Implemented jsPDF-based PDF export with full score breakdown. |
| Batch upload for multiple shots. | #71 (PBI-023) | Not planned | Deferred; depends on real ML and persistence landing first. |

## Feedback Not Addressed

- Batch upload (PBI-023): Deferred to a later Sprint. Depends on real ML pipeline stability and persistence being fully validated.

## Links

- [Roadmap](../../docs/roadmap.md)
- [Definition of Done](../../docs/definition-of-done.md)
- [Quality Requirements](../../docs/quality-requirements.md)
- [Quality Requirement Tests](../../docs/quality-requirement-tests.md)
- [Testing Strategy](../../docs/testing.md)
- [User Acceptance Tests](../../docs/user-acceptance-tests.md)
- [CHANGELOG](../../CHANGELOG.md)

## Quality Model

ISO/IEC 25010 sub-characteristics addressed:
- **QR-001:** Time Behaviour (API response < 2s)
- **QR-002:** Confidentiality (auth security, no cross-user access)
- **QR-003:** Testability (30% coverage for critical modules)
- **QR-004:** Usability (form guidance via placeholders/labels)

## Testing Status

| Test type | Status | Evidence |
|---|---|---|
| Unit tests | Passing | CI run |
| Integration tests | Passing | CI run |
| QRTs (QRT-001 to QRT-004) | Passing | CI run |
| Coverage (services, handlers) | ≥30% | CI run |

## CI Pipeline

- Lint: golangci-lint
- Test: go test -race -coverprofile
- Coverage: 30% threshold for critical modules
- QRT: go test -tags=qrt
- QA Extra: govulncheck
- Link Check: Lychee

## SemVer Release

- **Release:** [v0.2.0](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.2.0)
- **Tag:** v0.2.0
- **Maps to:** Sprint 2 — Assignment 4

## Current Product Status

BasketForm-AI now has a complete authentication system, social features (friends and result sharing), an enhanced ML pipeline with annotated video output and phase analysis, PDF export, and automated quality gates in CI. The product is production-ready at http://80.74.30.14/.

## Next Steps

- Sprint 3: Comparison with professional players, progress tracking over time
- Performance optimization for large video files
- Additional QA checks (accessibility, API contract testing)
