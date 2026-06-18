# Assignment 3 - Progress Tracker

## Current Status: 🟡 In Progress

**Last Updated:** 2026-06-18

---

## Summary of Completed Work

### ✅ Issues Created (GitHub)
| Issue | Title | Type | Status |
|-------|-------|------|--------|
| #20 | US-001: Upload shooting form video | User Story | Open |
| #21 | US-002: View automated form analysis | User Story | Open |
| #22 | US-003: Create and manage a user account | User Story | Open |
| #23 | US-004: Receive simplified actionable feedback | User Story | Open |
| #24 | US-005: Compare form with professional players | User Story | Open |
| #25 | US-006: Share analysis report with a coach | User Story | Open |
| #26 | US-007: Share progress with friends | User Story | Open |
| #27 | US-008: Track shooting progress over time | User Story | Open |
| #28 | US-010: Export analysis report as PDF | User Story | Open |
| #29 | PBI-001: Set up FastAPI backend | PBI | Open |
| #30 | PBI-002: Implement MediaPipe pose estimation | PBI | Open |
| #31 | PBI-003: Implement shooting technique evaluation | PBI | Open |
| #32 | PBI-004: Implement personalized feedback generation | PBI | Open |
| #33 | PBI-005: Implement PDF export | PBI | Open |
| #34 | PBI-006: Create React frontend with video upload | PBI | Open |
| #35 | PBI-007: Create analysis results display page | PBI | Open |
| #36 | PBI-008: Set up PostgreSQL database schema | PBI | Open |
| #37 | PBI-009: Implement user authentication system | PBI | Open |
| #38 | PBI-010: Deploy MVP v1 to production | PBI | Open |
| #39 | PBI-011: Create API documentation | PBI | Open |
| #40 | PBI-012: Write unit tests | PBI | Open |
| #41 | PBI-013: Create user account management page | PBI | Open |
| #42 | PBI-014: Create video upload progress page | PBI | Open |
| #43 | PBI-015: Create PDF export page | PBI | Open |

### ✅ Documentation Created
| File | Description |
|------|-------------|
| `docs/user-stories.md` | Authoritative user-story index |
| `docs/definition-of-done.md` | Team completion standards |
| `docs/roadmap.md` | Product roadmap |
| `reports/week3/README.md` | Week 3 report index |
| `reports/week3/customer-review-summary.md` | Customer review summary |
| `reports/week3/customer-review-transcript.md` | Transcript template |
| `reports/week3/customer-review-notes.md` | Notes template |
| `reports/week3/reflection.md` | Week 3 reflection |
| `reports/week3/retrospective.md` | Sprint retrospective |
| `reports/week3/llm-report.md` | LLM usage report |

### ✅ Pull Requests Created
| PR | Title | Status |
|----|-------|--------|
| #44 | feat: Add Assignment 3 documentation and project structure | ❌ Closed (not merged) |
| #45 | docs: Update CHANGELOG with Assignment 3 documentation | ✅ Merged |
| #46 | docs: Add Assignment 3 progress tracker | ✅ Merged |
| #47 | fix: Resolve Lychee link check errors | ✅ Merged |
| #48 | feat: Add Assignment 3 documentation and report files | 🟡 Open |

### ✅ Milestone Created
- **Sprint 1 - MVP v1** (Due: 2026-07-02)

---

## Step-by-Step Plan to Complete Assignment 3

### Phase 1: Repository Setup & Workflow (Days 1-2)
**Owner:** Karim Gimadiev (Project Manager)

- [x] **1.1** Merge PR #45 (CHANGELOG)
- [x] **1.2** Merge PR #46 (progress tracker)
- [x] **1.3** Merge PR #47 (Lychee fixes)
- [ ] **1.4** Review and merge PR #48 (documentation)
- [ ] **1.5** Verify issue templates are working correctly
- [ ] **1.6** Verify PR template is complete
- [ ] **1.7** Set up branch protection rules (if not already)
- [ ] **1.8** Create GitHub Project board for Product Backlog
- [ ] **1.9** Create Sprint Backlog view in GitHub Projects
- [ ] **1.10** Verify Lychee CI is passing

### Phase 2: Product Backlog Refinement (Days 2-3)
**Owner:** Product Owner (Anastasia + Team)

- [ ] **2.1** Review all created issues for completeness
- [ ] **2.2** Add acceptance criteria to all MVP v1 issues
- [ ] **2.3** Estimate story points for all PBIs
- [ ] **2.4** Assign team members to Sprint 1 issues
- [ ] **2.5** Add labels (mvp-v1, must-have, should-have, could-have)
- [ ] **2.6** Create additional PBIs if needed
- [ ] **2.7** Verify 15+ qualifying PBIs in Product Backlog
- [ ] **2.8** Update `docs/user-stories.md` with issue links

### Phase 3: Sprint Planning (Day 3)
**Owner:** Karim Gimadiev (Project Manager)

- [ ] **3.1** Conduct Sprint Planning meeting
- [ ] **3.2** Define Sprint Goal in milestone description
- [ ] **3.3** Select PBIs for Sprint 1
- [ ] **3.4** Assign reviewers to each PBI
- [ ] **3.5** Calculate total Sprint story points
- [ ] **3.6** Decompose user stories into technical tasks
- [ ] **3.7** Verify all Sprint PBIs have acceptance criteria
- [ ] **3.8** Update Sprint Backlog board

### Phase 4: MVP v1 Implementation (Days 4-10)
**Owner:** Development Team

#### Backend - AI/ML (Roman Santalov - ML Engineer)
- [ ] **4.1** Set up FastAPI project structure (#29)
- [ ] **4.2** Integrate MediaPipe pose estimation (#30)
- [ ] **4.3** Implement shooting technique evaluation (#31)
- [ ] **4.4** Implement personalized feedback generation (#32)
- [ ] **4.5** Implement PDF export endpoint (#33)

#### Backend - API & Database (Kamil Nizamov - ML Engineer)
- [ ] **4.6** Create PostgreSQL database schema (#36)
- [ ] **4.7** Implement user authentication (#37)
- [ ] **4.8** Create API documentation (#39)

#### Frontend (Arseniy Fedorov - Frontend Developer)
- [ ] **4.9** Create React project structure (#34)
- [ ] **4.10** Build video upload UI (#34)
- [ ] **4.11** Build analysis results page (#35)
- [ ] **4.12** Build user account management page (#41)
- [ ] **4.13** Build video upload progress page (#42)
- [ ] **4.14** Build PDF export page (#43)
- [ ] **4.15** Implement responsive design
- [ ] **4.16** Deploy MVP v1 to production (#38)

#### Testing & Documentation (Damir Galiev - ML Analyst)
- [ ] **4.17** Write unit tests for backend (#40)
- [ ] **4.18** Write unit tests for frontend
- [ ] **4.19** Perform manual testing
- [ ] **4.20** Test edge cases
- [ ] **4.21** Update documentation

### Phase 5: Code Review & Merge (Days 8-12)
**Owner:** All Team Members

- [ ] **5.1** Create feature branches for each PBI
- [ ] **5.2** Submit PRs with issue links
- [ ] **5.3** Review and approve PRs (each member reviews at least 1)
- [ ] **5.4** Leave meaningful review comments
- [ ] **5.5** Verify acceptance criteria in PRs
- [ ] **5.6** Merge approved PRs
- [ ] **5.7** Update CHANGELOG for each user-visible change

### Phase 6: Deployment & Testing (Days 12-14)
**Owner:** Arseniy Fedorov (Frontend Developer)

- [ ] **6.1** Deploy MVP v1 to production (#38)
- [ ] **6.2** Verify deployment at http://80.74.30.14/
- [ ] **6.3** Test all MVP v1 features in production
- [ ] **6.4** Fix any production issues
- [ ] **6.5** Create SemVer release (v1.0.0)
- [ ] **6.6** Tag release on main branch

### Phase 7: Documentation & Reports (Days 13-14)
**Owner:** Damir Galiev (ML Analyst)

- [ ] **7.1** Update README.md with MVP v1 instructions
- [ ] **7.2** Update CHANGELOG.md for release
- [ ] **7.3** Record video demonstration (<2 minutes)
- [ ] **7.4** Upload video to YouTube
- [ ] **7.5** Take screenshots for Week 3 report
- [ ] **7.6** Fill in Week 3 README.md with all links
- [ ] **7.7** Update contribution traceability table

### Phase 8: Customer Review (Day 14)
**Owner:** Product Owner (Anastasia + Team)

- [ ] **8.1** Schedule Sprint Review meeting
- [ ] **8.2** Prepare demo of MVP v1
- [ ] **8.3** Conduct Sprint Review
- [ ] **8.4** Record meeting (if permitted)
- [ ] **8.5** Get customer approval
- [ ] **8.6** Write customer review transcript/notes
- [ ] **8.7** Write customer review summary
- [ ] **8.8** Update Product Backlog based on feedback

### Phase 9: Final Reports (Day 14)
**Owner:** All Team Members

- [ ] **9.1** Write Week 3 reflection
- [ ] **9.2** Write Sprint retrospective
- [ ] **9.3** Write LLM usage report
- [ ] **9.4** Create PDF for Moodle submission
- [ ] **9.5** Verify all links are accessible
- [ ] **9.6** Final review of all artifacts

### Phase 10: Submission (Day 14)
**Owner:** Karim Gimadiev (Project Manager)

- [ ] **10.1** Commit all changes to main branch
- [ ] **10.2** Create final PR and merge
- [ ] **10.3** Verify commit on protected main branch
- [ ] **10.4** Submit PDF to Moodle
- [ ] **10.5** Verify all public links work
- [ ] **10.6** Confirm submission

---

## Team Member Responsibilities

| Member | Role | Responsibilities |
|--------|------|------------------|
| **Roman Santalov** | ML Engineer | Backend AI/ML implementation, MediaPipe integration, technique evaluation |
| **Kamil Nizamov** | ML Engineer | Backend API development, database design, authentication |
| **Arseniy Fedorov** | Frontend Developer | Frontend UI, React components, deployment to server |
| **Karim Gimadiev** | Project Manager | Sprint planning, customer communication, coordination |
| **Damir Galiev** | ML Analyst | Testing, documentation, quality assurance |
| **Anastasia Malakhova** | Customer | Product owner, requirements, acceptance |

---

## Key Links

### Repository
- **Main Repo:** https://github.com/kr1ny77/BasketForm-AI
- **MVP v1 Deployment:** http://80.74.30.14/
- **Figma Prototype:** https://silk-sadly-37202889.figma.site

### Issues & PRs
- **Product Backlog:** https://github.com/kr1ny77/BasketForm-AI/issues
- **Sprint 1 Milestone:** https://github.com/kr1ny77/BasketForm-AI/milestone/1
- **Open PRs:** https://github.com/kr1ny77/BasketForm-AI/pulls

### Documentation
- **User Stories Index:** docs/user-stories.md
- **Definition of Done:** docs/definition-of-done.md
- **Roadmap:** docs/roadmap.md
- **CHANGELOG:** CHANGELOG.md

### Reports
- **Week 3 Report:** reports/week3/README.md
- **Customer Review:** reports/week3/customer-review-summary.md
- **Reflection:** reports/week3/reflection.md
- **Retrospective:** reports/week3/retrospective.md

---

## Estimation Summary

### Product Backlog Size
| PBI | Story Points | Priority | Sprint |
|-----|--------------|----------|--------|
| US-001 | 5 | Must Have | Sprint 1 |
| US-002 | 8 | Must Have | Sprint 1 |
| US-003 | 5 | Must Have | Sprint 1 |
| US-004 | 5 | Must Have | Sprint 1 |
| US-005 | 5 | Could Have | Sprint 2 |
| US-006 | 3 | Should Have | Sprint 2 |
| US-007 | 3 | Should Have | Sprint 2 |
| US-008 | 5 | Should Have | Sprint 2 |
| US-010 | 3 | Should Have | Sprint 1 |
| PBI-001 | 5 | Must Have | Sprint 1 |
| PBI-002 | 8 | Must Have | Sprint 1 |
| PBI-003 | 8 | Must Have | Sprint 1 |
| PBI-004 | 5 | Must Have | Sprint 1 |
| PBI-005 | 3 | Should Have | Sprint 1 |
| PBI-006 | 5 | Must Have | Sprint 1 |
| PBI-007 | 5 | Must Have | Sprint 1 |
| PBI-008 | 3 | Must Have | Sprint 1 |
| PBI-009 | 5 | Must Have | Sprint 1 |
| PBI-010 | 3 | Must Have | Sprint 1 |
| PBI-011 | 2 | Must Have | Sprint 1 |
| PBI-012 | 3 | Must Have | Sprint 1 |
| PBI-013 | 3 | Must Have | Sprint 1 |
| PBI-014 | 2 | Must Have | Sprint 1 |
| PBI-015 | 2 | Should Have | Sprint 1 |

**Total Product Backlog:** 102 Story Points
**Sprint 1 (MVP v1):** 76 Story Points

---

## Risks & Mitigations

| Risk | Impact | Mitigation | Status |
|------|--------|------------|--------|
| Video processing performance | High | Optimize algorithms, chunk processing | 🟡 Open |
| AI model accuracy | High | Test with diverse video samples | 🟡 Open |
| User adoption | High | Focus on UX, gather feedback | 🟡 Open |
| Time constraints | Medium | Prioritize MVP v1 features | 🟡 Open |
| Team coordination | Medium | Daily standups, clear communication | 🟡 Open |

---

## Notes for Next Session

1. **PR #48 needs review** - Merge to complete documentation setup
2. **Issues need assignment** - Assign team members to Sprint 1 issues
3. **GitHub Project board** - Create Product Backlog and Sprint Backlog views
4. **Start implementation** - Begin with backend setup (#29)
5. **Daily standups** - Coordinate team progress
6. **Customer meeting** - Schedule Sprint Review for end of Sprint 1

---

## Assignment 3 Requirements Checklist

### Part 1: Migrate User Stories ✅
- [x] Preserve reports/week2/user-stories.md
- [x] Create issues for active user stories
- [x] Preserve stable IDs in issue titles
- [x] Copy information into migrated issues
- [x] Create docs/user-stories.md index
- [x] Maintain requirement status

### Part 2: Repository Workflow 🟡
- [x] Issue templates created
- [x] PR template extended
- [ ] Issue-linked branches and PRs
- [ ] Acceptance-criteria verification before merge
- [ ] Merge-commit workflow
- [ ] SemVer releases
- [ ] CHANGELOG.md workflow
- [ ] Video demonstration

### Part 3: Product Backlog 🟡
- [x] 15+ qualifying PBIs created
- [x] Clear titles and descriptions
- [x] Types assigned
- [x] MoSCoW priorities
- [x] Story points estimated
- [ ] GitHub Project board created

### Part 4: Acceptance Criteria 🟡
- [x] Acceptance criteria in MVP v1 issues
- [ ] Review criteria as team
- [ ] Clarify ambiguous PBIs

### Part 5: Estimate Product Backlog ✅
- [x] Story points recorded
- [x] Total backlog size calculated

### Part 6: Definition of Done ✅
- [x] docs/definition-of-done.md created
- [x] Minimum content included
- [x] CHANGELOG requirement included

### Part 7: Sprint Backlog & MVP v1 Plan 🟡
- [x] Sprint 1 milestone created
- [ ] Sprint Goal defined
- [ ] PBIs assigned to Sprint
- [ ] Reviewers assigned
- [ ] Total Sprint points calculated

### Part 8: Implement & Verify MVP v1 ⏳
- [ ] All MVP v1 features implemented
- [ ] Acceptance criteria verified
- [ ] Delivered increment accessible
- [ ] Video demonstration recorded

### Part 9: Sprint Review ⏳
- [ ] Customer meeting scheduled
- [ ] MVP v1 demonstrated
- [ ] Customer feedback obtained
- [ ] Transcript/notes written

### Part 10: Roadmap ✅
- [x] docs/roadmap.md created
- [x] Current and next Sprint included

### Part 11: Reflection ✅
- [x] reports/week3/reflection.md created

### Part 12: Retrospective ✅
- [x] reports/week3/retrospective.md created

### Part 13: LLM Report ✅
- [x] reports/week3/llm-report.md created

### Assignment Report ✅
- [x] reports/week3/README.md created
- [ ] All links verified
- [ ] Screenshots added
- [ ] PDF created for Moodle

---

## Quick Commands

```bash
# Check PR status
curl -s -H "Authorization: token ghp_..." "https://api.github.com/repos/kr1ny77/BasketForm-AI/pulls" | python3 -c "import sys, json; [print(f'#{p[\"number\"]}: {p[\"title\"]} ({p[\"state\"]})') for p in json.load(sys.stdin)]"

# Check issues
curl -s -H "Authorization: token ghp_..." "https://api.github.com/repos/kr1ny77/BasketForm-AI/issues?state=open" | python3 -c "import sys, json; [print(f'#{i[\"number\"]}: {i[\"title\"]}') for i in json.load(sys.stdin) if 'pull_request' not in i]"

# Check milestones
curl -s -H "Authorization: token ghp_..." "https://api.github.com/repos/kr1ny77/BasketForm-AI/milestones"
```
