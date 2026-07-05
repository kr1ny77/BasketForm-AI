# Week 5 Public Report — BasketForm-AI

Project: **BasketForm-AI** — AI-powered basketball shooting form analysis platform.
Deployment: http://80.74.30.14/

## Sprint Context

- **Sprint:** Sprint 3 — Assignment 5 (MVP v2)
- **Sprint Goal:** Deliver MVP v2 (v0.3.0): LLM-powered personalized feedback, Friends section, report sharing, progress tracking, pro comparison, RU translation, architecture documentation, and ADRs.
- **Sprint milestone:** [Sprint 3 - Assignment 5](https://github.com/kr1ny77/BasketForm-AI/milestone/3)
- **Product Backlog board:** [GitHub Projects board](https://github.com/users/kr1ny77/projects/7)
- **Sprint Backlog board:** [Sprint Backlog view](https://github.com/users/kr1ny77/projects/7/views/2)
- **Sprint dates:** 2026-07-06 to 2026-07-12
- **Total Sprint size:** 64 Story Points

## Delivered Product Changes

- LLM-powered personalized feedback via external API
- Friends section (add/find/manage friends)
- Share analysis reports with friends
- Progress tracking over time
- Compare shooting form against professional players
- Russian (RU) translation fix and completion
- Architecture documentation (static, dynamic, deployment views)
- Architecture Decision Records (ADRs)
- Development process and configuration management documentation
- Storage hardening (deadlock fix, batch-delete cleanup)

## Deployed Product

- **URL:** http://80.74.30.14/
- **Run instructions:** See [README.md](../../README.md)
- **Public sanitized demo video:** [Demo Video](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.3.0)

## Customer Feedback Response Table

| Feedback point | Resulting PBI or issue | Status | Response |
|---|---|---|---|
| Real ML analysis needed; mock pipeline is only acceptable for demo. | [#64](https://github.com/kr1ny77/BasketForm-AI/issues/64) (PBI-016) | Done | Replaced mock with real MediaPipe pipeline generating annotated video and phase scores (Sprint 2). |
| Authentication and data persistence across restarts. | [#66](https://github.com/kr1ny77/BasketForm-AI/issues/66) (PBI-018) | Done | Added JWT auth, bcrypt, JSON file persistence in data/ directory (Sprint 2). |
| Share results with coaches/friends. | [#67](https://github.com/kr1ny77/BasketForm-AI/issues/67) (PBI-019), [#96](https://github.com/kr1ny77/BasketForm-AI/issues/96) (PBI-033) | Done | Added friend system and result sharing with read-only access (Sprint 2). Extended with report sharing in Sprint 3. |
| PDF export for offline record-keeping. | [#70](https://github.com/kr1ny77/BasketForm-AI/issues/70) (PBI-022) | Done | Implemented jsPDF-based PDF export with full score breakdown (Sprint 2). |
| Batch upload for multiple shots. | [#71](https://github.com/kr1ny77/BasketForm-AI/issues/71) (PBI-023) | Deferred | Deferred because MVP v2 prioritized LLM feedback, progress tracking, and architecture hardening. Will be reconsidered in Sprint 4. |
| Consider adding progress tracking over time. | [#27](https://github.com/kr1ny77/BasketForm-AI/issues/27), [#88](https://github.com/kr1ny77/BasketForm-AI/issues/88) (PBI-025) | Done | Implemented progress tracking with historical data visualization (Sprint 3). |
| Compare form with professional players. | [#24](https://github.com/kr1ny77/BasketForm-AI/issues/24), [#87](https://github.com/kr1ny77/BasketForm-AI/issues/87) (PBI-024) | Done | Added pro comparison feature with reference player data (Sprint 3). |
| Personalized, coaching-style feedback. | [#94](https://github.com/kr1ny77/BasketForm-AI/issues/94) (PBI-031) | Done | Connected LLM analysis via external API for personalized feedback (Sprint 3). |
| Fix and complete Russian translation. | [#97](https://github.com/kr1ny77/BasketForm-AI/issues/97) (PBI-034) | Done | Fixed and completed RU translation for all UI elements (Sprint 3). |
| Improve deployment complexity (Python dependencies). | [#91](https://github.com/kr1ny77/BasketForm-AI/issues/91) (PBI-028) | Done | Documented deployment process and Python dependency installation (Sprint 3). |

## Feedback Not Addressed

- **Batch upload** (PBI-023): Deferred to Sprint 4. MVP v2 prioritized LLM-powered feedback (higher customer value), progress tracking (retention), and architecture documentation (maintainability). The batch upload depends on stable ML pipeline and storage, which are now hardened.

## Project Documentation Links

- [Hosted Documentation Site](https://kr1ny77.github.io/BasketForm-AI/)
- [Roadmap](../../docs/roadmap.md)
- [Definition of Done](../../docs/definition-of-done.md)
- [Quality Requirements](../../docs/quality-requirements.md)
- [Quality Requirement Tests](../../docs/quality-requirement-tests.md)
- [Testing Strategy](../../docs/testing.md)
- [User Acceptance Tests](../../docs/user-acceptance-tests.md)
- [Development Process](../../docs/development-process.md)
- [Architecture Documentation](../../docs/architecture/)
- [CHANGELOG](../../CHANGELOG.md)

The [Development Process](../../docs/development-process.md) document is also linked from the root README and from the [hosted documentation site](https://kr1ny77.github.io/BasketForm-AI/development-process/).

## Week 5 Report Documents

- Customer Review Summary — _to be added after Sprint Review_
- Customer Review Transcript — _to be added after Sprint Review_
- Reflection — _to be added after Sprint Review_
- Retrospective — _to be added after Sprint Review_
- LLM Usage Report — _to be added after Sprint Review_

## Architecture Documentation

Architecture documentation is maintained in `docs/architecture/` with diagrams-as-code (PlantUML):

- **Static View:** Component and package structure
- **Dynamic View:** Request flow and ML pipeline
- **Deployment View:** VM deployment and CI/CD

See [Architecture Documentation](../../docs/architecture/) for details.

## Quality Model

ISO/IEC 25010 sub-characteristics addressed:
- **QR-001:** Time Behaviour (API response < 2s)
- **QR-002:** Confidentiality (auth security, no cross-user access)
- **QR-003:** Testability (30% coverage for critical modules)
- **QR-004:** Usability (form guidance via placeholders/labels)

## Testing Status

| Test type | Status | Evidence |
|---|---|---|
| Unit tests | Passing locally | [internal/handlers](https://github.com/kr1ny77/BasketForm-AI/tree/main/internal/handlers), [internal/services](https://github.com/kr1ny77/BasketForm-AI/tree/main/internal/services) |
| Integration tests | Passing locally | [internal/handlers](https://github.com/kr1ny77/BasketForm-AI/tree/main/internal/handlers) |
| QRTs (QRT-001 to QRT-004) | Passing locally | [internal/qrt](https://github.com/kr1ny77/BasketForm-AI/tree/main/internal/qrt) |
| Coverage (services, handlers) | ≥30% target | See CI Coverage Check job |

### Test Locations

- **Unit tests:** [internal/handlers](https://github.com/kr1ny77/BasketForm-AI/tree/main/internal/handlers), [internal/services](https://github.com/kr1ny77/BasketForm-AI/tree/main/internal/services)
- **Integration tests:** [internal/handlers](https://github.com/kr1ny77/BasketForm-AI/tree/main/internal/handlers)
- **Automated quality requirement tests:** [internal/qrt](https://github.com/kr1ny77/BasketForm-AI/tree/main/internal/qrt)

## CI Pipeline

- **Workflow:** [CI workflow runs (main)](https://github.com/kr1ny77/BasketForm-AI/actions?query=branch%3Amain)
- **Latest protected-default-branch CI run:** [Link to latest CI run]
- Lint: golangci-lint
- Test: go test -race -coverprofile
- Coverage: 30% threshold for critical modules
- QRT: go test -tags=qrt
- QA Extra: govulncheck (security)
- Link Check: Lychee

## Branch Protection

The default branch `main` is protected by repository rulesets requiring pull requests, status checks, and review before merge.

- **Branch ruleset:** [Repository rules for main](https://github.com/kr1ny77/BasketForm-AI/rules)
- **Settings:** [Repository rulesets](https://github.com/kr1ny77/BasketForm-AI/rules)

## SemVer Release

- **Release:** [v0.3.0 — MVP v2](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.3.0)
- **Tag:** v0.3.0
- **Maps to:** Sprint 3 — Assignment 5 (linked to [Sprint milestone](https://github.com/kr1ny77/BasketForm-AI/milestone/3))

## UAT Results Summary

Week 5 UAT execution covered at least three active end-user scenarios from [docs/user-acceptance-tests.md](../../docs/user-acceptance-tests.md).

- **Passed scenarios:** LLM-powered feedback, progress tracking, pro comparison, friend sharing.
- **Scenarios needing changes:** none blocking.
- See Customer Review Summary (to be added after Sprint Review) for full details.

## Quality Gates Continuity

The automated tests, CI checks, quality requirement tests, and Definition of Done created in Assignment 4 are maintained project assets. Later PBIs must keep these gates passing (or extend them) and must not bypass, disable, or treat them as one-time submission evidence.

## Contribution Traceability

| Team member | Issues | PRs | Reviews | Testing | Quality | Automation | Documentation |
|---|---|---|---|---|---|---|---|
| [@kr1ny77](https://github.com/kr1ny77) | [#94](https://github.com/kr1ny77/BasketForm-AI/issues/94), [#90](https://github.com/kr1ny77/BasketForm-AI/issues/90) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3Akr1ny77) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3Akr1ny77) | Unit tests for LLM service | CI pipeline | CI workflow | CHANGELOG, README |
| [@Koomaz](https://github.com/Koomaz) | [#91](https://github.com/kr1ny77/BasketForm-AI/issues/91), [#92](https://github.com/kr1ny77/BasketForm-AI/issues/92) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3AKoomaz) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3AKoomaz) | Storage tests, architecture tests | Quality requirements | Architecture docs | Architecture documentation |
| [@romasntlv](https://github.com/romasntlv) | [#94](https://github.com/kr1ny77/BasketForm-AI/issues/94), [#87](https://github.com/kr1ny77/BasketForm-AI/issues/87) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3Aromasntlv) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3Aromasntlv) | ML pipeline testing | — | ML integration | Customer review docs |
| [@gimacorp](https://github.com/gimacorp) | [#93](https://github.com/kr1ny77/BasketForm-AI/issues/93) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3Agimacorp) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3Agimacorp) | — | — | Deployment scripts | Reports, roadmap, UAT docs, development process |
| [@mentalafffection](https://github.com/mentalafffection) | [#95](https://github.com/kr1ny77/BasketForm-AI/issues/95), [#96](https://github.com/kr1ny77/BasketForm-AI/issues/96), [#97](https://github.com/kr1ny77/BasketForm-AI/issues/97) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3Amentalafffection) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3Amentalafffection) | — | Quality requirements, Usability QR | Frontend QA | Quality requirements docs, i18n |

## Screenshots

> Sanitized evidence screenshots to be added after Sprint Review.

- Sprint milestone: _to be added_
- Latest protected-default-branch CI run: _to be added_
- Branch protection / rules evidence: _to be added_
- Coverage / test report: _to be added_
- Additional QA check result: _to be added_
- SemVer release: _to be added_
- Example reviewed issue-linked PR: _to be added_
- Hosted docs site ([kr1ny77.github.io/BasketForm-AI](https://kr1ny77.github.io/BasketForm-AI/)): _to be added_

## Current Product Status

BasketForm-AI now has a complete authentication system, social features (friends and result sharing), an enhanced ML pipeline with annotated video output and phase analysis, LLM-powered personalized feedback, progress tracking, pro comparison, PDF export, Russian translation, architecture documentation, ADRs, and automated quality gates in CI. The product is deployed at http://80.74.30.14/.

## Next Steps

- Sprint 4: Consolidate MVP v2 deferred work, grow test coverage, address any architecture or deployment issues.
- Additional user stories: US-009 leaderboard, US-010 drill recommendations.
- Performance optimization for large video files.
- Additional QA checks (accessibility, API contract testing).
