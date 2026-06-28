# Week 4 Report — Assignment 4

## Project Information

**Project Name:** BasketForm-AI
**Description:** AI-powered basketball shooting form analysis platform
**License:** [MIT License](../../LICENSE)

## Team Members

| Name | GitHub Username | Role |
|------|-----------------|------|
| Roman Santalov | romasntlv | Backend / ML Engineer |
| Kamil Nizamov | Koomaz | Backend / ML Engineer |
| Arseniy Fedorov | arsafedo | Frontend Developer |
| Karim Gimadiev | kr1ny77 | Project Manager |
| Damir Galiev | DamirGaliev | Testing / Documentation |

---

## 1. Product Backlog Board

- [GitHub Projects Board](https://github.com/users/Koomaz/projects/1)

## 2. Sprint Backlog Board

- [GitHub Projects Board — Sprint View](https://github.com/users/Koomaz/projects/1)

## 3. Assignment 4 Sprint Milestone

- [Sprint 2 Milestone](https://github.com/kr1ny77/BasketForm-AI/milestone/2)
- **Sprint Goal:** Make the analysis trustworthy and durable — replace the mock pipeline with real basketball-shot analysis, validate that uploaded video contains a basketball shot, and persist results across restarts
- **Sprint Dates:** 2026-06-29 — 2026-07-05
- **Sprint Scope:**
  - PBI-016: Integrate real Python ML pipeline (replace mock processor)
  - PBI-017: Basketball + player detection and non-basketball video validation
  - PBI-018: Persist analysis results across server restarts
  - PBI-019: Automated tests for real ML integration, validation, and persistence
  - PBI-020: Deploy real ML runtime to production (Docker + server)
  - PBI-021: Track the basketball and refine the release-phase algorithm ✅

## 4. Total Sprint Size

- **Total Story Points:** TBD (see milestone issues for individual estimates)

## 5. Delivered Product Changes

- Replaced mock Go processor with real Python MediaPipe ML pipeline
- Go server now delegates to `scripts/process_video.py` via `exec.Command`
- ML script auto-detects Python virtualenv for zero-config setup
- Structured JSON output (score, feedback, pose) maintained for API compatibility

## 6. Deployed Product

- **URL:** http://80.74.30.14/
- **Run Instructions:** [README — Local Development Setup](../../README.md#local-development-setup)
- **Docker:** `docker build -t basketform-ai . && docker run -p 8080:8080 basketform-ai`

## 7. Customer Feedback Response

| Customer Feedback | Resulting PBI | Status |
|-------------------|---------------|--------|
| "I want real analysis, not mock data" | PBI-016: Integrate real Python ML pipeline | In Progress |
| "Detect if video actually contains basketball" | PBI-017: Basketball + player detection | In Progress |
| "Results should not disappear on restart" | PBI-018: Persist analysis results | In Progress |
| "Needs automated tests for ML" | PBI-019: Automated ML tests | In Progress |
| "Deploy real ML to production" | PBI-020: Deploy ML runtime | In Progress |

### Feedback Not Addressed

- US-005 (Compare with professional players) — deferred to Sprint 3+
- US-006 (Share with coach) — deferred to Sprint 3+
- US-008 (Track progress over time) — deferred to Sprint 3+

## 8. Documentation Links

- [Roadmap](../../docs/roadmap.md)
- [Definition of Done](../../docs/definition-of-done.md)
- [Quality Requirements](../../docs/quality-requirements.md)
- [Quality Requirement Tests](../../docs/quality-requirement-tests.md)
- [Testing Strategy](../../docs/testing.md)
- [User Acceptance Tests](../../docs/user-acceptance-tests.md)

## 9. Quality Model

Based on ISO/IEC 25010, the following sub-characteristics are selected:

| Sub-characteristic | Justification |
|--------------------|---------------|
| Functional Suitability | ML must correctly detect basketball shots and provide accurate feedback |
| Performance Efficiency | Video processing must complete within acceptable time on CPU |
| Usability | Upload, results, and PDF export must be intuitive |
| Reliability | ML pipeline must handle invalid input without crashing |
| Maintainability | Go + Python separation allows independent modification |

See [Quality Requirements](../../docs/quality-requirements.md) for full details.

## 10. Testing Status

| Metric | Value |
|--------|-------|
| Total tests | 47 (passing) |
| Line coverage | ~70% |
| Critical modules covered | API endpoints, UUID generation, file validation, storage, processor |
| ML integration tests | Pending (PBI-019) |

### Unit Tests

- Location: `*_test.go` files in project root
- Runner: `go test -race ./...`

### Integration Tests

- Location: `*_test.go` files using `httptest`
- Endpoints tested: `/api/upload`, `/api/status/{id}`, `/api/result/{id}`, `/api/videos`

### Automated Quality Requirement Tests

- API contract compliance: covered by existing 47 tests (QR-005)
- ML output structure validation: pending (QR-001)
- Non-basketball rejection: pending (QR-002)
- Result persistence: pending (QR-003)

## 11. CI Pipeline

- **CI Workflow:** [`.github/workflows/ci.yml`](../../.github/workflows/ci.yml)
- **Jobs:** `test` (go test with coverage) → `build` (go build)
- **Trigger:** Push/PR to `main`

### Latest CI Run

Pending — link to latest protected-default-branch CI run will be added after merge to main.

### Branch Protection

- Default branch (`main`) is protected
- Direct pushes disabled
- Required approvals: 1
- All CI checks must pass before merge

### Screenshots

Screenshots of CI run, coverage report, and branch protection will be added to `reports/week4/images/` before final submission.

## 12. Test Governance

The Assignment 4 tests, CI checks, quality requirement tests, and Definition of Done will continue to govern later project work as follows:

- **CI gates:** Every push/PR to `main` runs `go test` and `go build`; failing code cannot merge
- **Definition of Done:** All PBIs must meet the checklist in `docs/definition-of-done.md` before marking Done
- **Quality requirements:** QR-001 through QR-005 define the minimum quality bar for ML integration
- **Future sprints:** New features must include corresponding tests and pass all existing CI checks

## 13. SemVer Release

- **Sprint 1 Release:** [v0.1.0](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.1.0) — MVP v1
- **Sprint 2 Release:** Pending — will be created after Sprint 2 work is merged to main

## 14. CHANGELOG

- [CHANGELOG.md](../../CHANGELOG.md) — Updated with v0.1.0 entries; Sprint 2 entries pending

## 15. Demo Video

- [Video Demo (< 2 min)](https://youtu.be/To_1ZwAMe4M) — Sprint 1 demo
- Sprint 2 demo video: pending

## 16. Presentation

Optional: [reports/week4/presentation.pdf](presentation.pdf) — pending

## 17. UAT Results

See [User Acceptance Tests](../../docs/user-acceptance-tests.md) for full scenarios.

| Scenario | Result |
|----------|--------|
| Upload MP4 and receive analysis | ✅ Pass |
| Non-supported format rejected | ✅ Pass |
| PDF export works | ✅ Pass |
| Charts render correctly | ✅ Pass |
| Real ML pipeline integration | ✅ Pass |
| Non-basketball video rejection | Pending |

## 18. Customer Review

- [Customer Review Transcript](customer-review-transcript.md) — pending upload
- [Customer Review Summary](customer-review-summary.md)

## 19. Reflection

- [Week 4 Reflection](reflection.md)

## 20. Retrospective

- [Sprint Retrospective](retrospective.md)

## 21. LLM Report

- [LLM Usage Report](llm-report.md)

## 22. Current Product Status

- **MVP v0:** Deployed (v0.0.1)
- **MVP v1:** Completed and deployed (v0.1.0) — mock ML pipeline
- **MVP v2 (Sprint 2):** In Progress — real MediaPipe ML pipeline replacing mock
- **Backend:** Go 1.21 + Python (MediaPipe)
- **Frontend:** HTML/CSS/JS with Canvas, Chart.js, jsPDF
- **Tests:** 47 passing (unit + integration)

## 23. Next Steps

1. Complete Sprint 2 PBIs (PBI-017, PBI-018, PBI-019, PBI-020)
2. Deploy real ML runtime to production
3. Merge Sprint 2 branch to main and create v0.2.0 release
4. Begin Sprint 3: comparison features (US-005), sharing (US-006, US-007)

## 24. Contribution Traceability

| Team Member | Issues | PRs | Reviews |
|-------------|--------|-----|---------|
| Roman Santalov | [Issues](https://github.com/kr1ny77/BasketForm-AI/issues?q=assignee%3Aromasntlv) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3Aromasntlv) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3Aromasntlv) |
| Kamil Nizamov | [Issues](https://github.com/kr1ny77/BasketForm-AI/issues?q=assignee%3AKoomaz) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3AKoomaz) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3AKoomaz) |
| Arseniy Fedorov | [Issues](https://github.com/kr1ny77/BasketForm-AI/issues?q=assignee%3Aarsafedo) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3Aarsafedo) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3Aarsafedo) |
| Karim Gimadiev | [Issues](https://github.com/kr1ny77/BasketForm-AI/issues?q=assignee%3Akr1ny77) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3Akr1ny77) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3Akr1ny77) |
| Damir Galiev | [Issues](https://github.com/kr1ny77/BasketForm-AI/issues?q=assignee%3ADamirGaliev) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3ADamirGaliev) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3ADamirGaliev) |

## 25. Screenshots

Screenshots will be added to `reports/week4/images/` before final submission:

- Sprint milestone
- Latest CI run
- Branch protection settings
- Coverage / test report
- SemVer release
- Example reviewed PR
