# Week 3 Report — Assignment 3

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

## 1. User-Story and PBI Scope Summary

Since Assignment 2, user stories were migrated from `reports/week2/user-stories.md` to GitHub Issues with stable IDs preserved. The Product Backlog contains 15+ qualifying PBIs including user stories, technical infrastructure, testing, and documentation items. See [docs/user-stories.md](../../docs/user-stories.md) for the full index.

**MVP v1 scope:** US-001, US-002, US-003, US-004, US-010 — all marked Done.

## 2. Customer Feedback Addressed in MVP v1

Based on the customer meeting from Assignment 2:

1. **Video upload with format guidance** — Drag-and-drop with MP4/MOV/AVI validation and size limits
2. **Simplified feedback** — Max 3 actionable takeaways per analysis in plain language
3. **PDF export** — jsPDF-based report download for offline record keeping
4. **User accounts** — Profile page with name, email, password management (UI/demo mode)

## 3. Historical User Stories

- [reports/week2/user-stories.md](../week2/user-stories.md) — Assignment 2 artifact (preserved)

## 4. Current User Stories Index

- [docs/user-stories.md](../../docs/user-stories.md) — Authoritative registry with issue links, MoSCoW, Work Status, Sprint

## 5. Product Backlog Board

- [GitHub Projects Board](https://github.com/users/Koomaz/projects/1)

## 6. Sprint Backlog Board

- [GitHub Projects Board — Sprint View](https://github.com/users/Koomaz/projects/1)

## 7. Sprint Milestone

- [Sprint 1 Milestone](https://github.com/kr1ny77/BasketForm-AI/milestone/1)
- **Sprint Goal:** Deliver MVP v1 — core basketball shooting form analysis with Go backend, Canvas frontend, and mock ML pipeline
- **Dates:** 2026-06-16 — 2026-06-29

## 8. Product Backlog Size

- **Total Product Backlog Size:** 102 Story Points
- **Current Sprint (Sprint 1) Size:** 76 Story Points

## 9. MVP v1 Scope

- [MVP v1 filtered view](https://github.com/kr1ny77/BasketForm-AI/issues?q=milestone%3A1+label%3Amvp-v1)

**MVP v1 includes:**
1. **US-001** — Video upload with drag-and-drop and format validation
2. **US-002** — Automated analysis with score, feedback, and pose visualization
3. **US-003** — User account management (profile page)
4. **US-004** — Simplified actionable feedback (max 3 takeaways)
5. **US-010** — PDF export of analysis reports

Plus supporting PBIs: Go backend setup, Canvas animation, API endpoints, mock ML processor, unit/integration tests, CI pipeline, documentation.

## 10. PBI Types, Statuses, and Priorities

- **User Stories:** US-001 to US-010 with MoSCoW priorities (Must/Should/Could Have)
- **Technical PBIs:** Backend, API, animation, processing pipeline
- **Testing PBIs:** Unit tests, integration tests, manual checklist
- **Documentation PBIs:** Reports, roadmap, DoD, CHANGELOG
- **Work Status:** To Do → Ready → In Progress → Review → Done
- **MVP Version:** Tracked via `mvp-v1` label
- **Task Decomposition:** User stories split into linked technical PBIs for implementability

## 11. Verification Evidence

All MVP v1 PBIs were verified through:
- Acceptance criteria in issue descriptions
- 47 unit + integration tests (`go test -race ./...`)
- Manual test checklist (50+ scenarios in `reports/week3/manual-test-checklist.md`)
- API endpoint verification via curl/Postman
- Cross-browser testing (Chrome, Firefox, Safari)

## 12. Current Product Status

- **MVP v0:** Deployed at http://80.74.30.14/
- **MVP v1:** Completed, release [v0.1.0](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.1.0)
- **Backend:** Go 1.21, standard library HTTP server
- **Frontend:** HTML/CSS/JS with Canvas animation, Chart.js, jsPDF
- **ML Pipeline:** Mock processor (score + feedback + pose), Python script available
- **Tests:** 47 passing (unit + integration)

## 13. Next Steps

1. Deploy v0.1.0 to production server
2. Conduct Sprint Review with customer
3. Begin Sprint 2: comparison features (US-005), sharing (US-006, US-007)
4. Integrate real ML pipeline via Python exec
5. Add user authentication with sessions

## 14. Contribution Traceability

| Team Member | Issues | PRs | Reviews |
|-------------|--------|-----|---------|
| Roman Santalov | [Issues](https://github.com/kr1ny77/BasketForm-AI/issues?q=assignee%3Aromasntlv) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3Aromasntlv) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3Aromasntlv) |
| Kamil Nizamov | [Issues](https://github.com/kr1ny77/BasketForm-AI/issues?q=assignee%3AKoomaz) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3AKoomaz) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3AKoomaz) |
| Arseniy Fedorov | [Issues](https://github.com/kr1ny77/BasketForm-AI/issues?q=assignee%3Aarsafedo) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3Aarsafedo) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3Aarsafedo) |
| Karim Gimadiev | [Issues](https://github.com/kr1ny77/BasketForm-AI/issues?q=assignee%3Akr1ny77) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3Akr1ny77) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3Akr1ny77) |
| Damir Galiev | [Issues](https://github.com/kr1ny77/BasketForm-AI/issues?q=assignee%3ADamirGaliev) | [PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=author%3ADamirGaliev) | [Reviews](https://github.com/kr1ny77/BasketForm-AI/pulls?q=reviewed-by%3ADamirGaliev) |

## 15. SemVer Release

- **MVP v1 Release:** [v0.1.0](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.1.0)
- Tag on `main` branch, linked to Sprint 1 milestone

## 16. CHANGELOG

- [CHANGELOG.md](../../CHANGELOG.md) — Updated with v0.1.0 entries (Added/Changed/Removed)

## 17. Documentation Links

- [docs/roadmap.md](../../docs/roadmap.md)
- [docs/definition-of-done.md](../../docs/definition-of-done.md)

## 18. Issue and PR Templates

- [Issue Templates](https://github.com/kr1ny77/BasketForm-AI/tree/main/.github/ISSUE_TEMPLATE) — User Story, Bug Report, Other PBI, Course Task
- [PR Template](https://github.com/kr1ny77/BasketForm-AI/blob/main/.github/pull_request_template.md) — Summary, testing, related issue, acceptance criteria, changelog

## 19. Reviewed Issue-Linked PRs

- [Merged PRs](https://github.com/kr1ny77/BasketForm-AI/pulls?q=is%3Apr+is%3Amerged)

## 20. MVP v1 Deployment

- **URL:** http://80.74.30.14/
- **Run Instructions:** [README — Local Development Setup](../../README.md#local-development-setup)
- **Docker:** `docker build -t basketform-ai . && docker run -p 8080:8080 basketform-ai`

## 21. Video Demonstration

- [Video Demo (< 2 min)](https://youtu.be/To_1ZwAMe4M)

## 22. Screenshots

*Add screenshots from `reports/week3/images/` after deployment:*
- Product Backlog view
- Sprint Backlog view
- Sprint milestone
- MVP version view
- SemVer release
- Delivered MVP v1
- Example reviewed PR

## 23. Customer Review

- [Customer Review Transcript](customer-review-transcript.md)
- [Customer Review Summary](customer-review-summary.md)

## 24. Reflection

- [Week 3 Reflection](reflection.md)

## 25. Retrospective

- [Sprint Retrospective](retrospective.md)

## 26. LLM Report

- [LLM Usage Report](llm-report.md)
