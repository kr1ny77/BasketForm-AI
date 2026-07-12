# Retrospective — Week 6

## What Went Well

- **Person-detection validation:** Successfully implemented and tested. The customer confirmed it works correctly and prevents invalid uploads.
- **Documentation completeness:** All customer-facing documentation (CONTRIBUTING.md, AGENTS.md, customer-handover.md) was created and approved.
- **Customer satisfaction:** The customer confirmed the product meets the Definition of Done and is ready for transition.
- **Core pipeline stability:** The single-throw analysis pipeline remained stable throughout the sprint with no regressions.
- **Transition readiness:** Deployment access, environment variables, and recovery steps were documented and verified.

## What Did Not Go Well

- **Multi-throw feature deferred:** The automatic video splitting feature was too complex for a single sprint. This was a planning gap — the feature should have been timeboxed or broken down further.
- **Regression testing was manual:** Mid-sprint regression testing took more time than expected because it was not automated. This slowed down other work.
- **Limited new features:** The sprint focused mostly on documentation and validation, which is appropriate for a trial release but felt less productive than feature-heavy sprints.

## Action Items

| Action | Owner | Deadline |
|--------|-------|----------|
| Automate regression tests for the core pipeline | Team | Sprint 5 |
| Timebox experimental features with explicit go/no-go criteria | Team | Sprint 5 planning |
| Involve customer in documentation review earlier | Team | Sprint 5 |

## Velocity and Process

- **Story Points Planned:** 28
- **Story Points Completed:** ~24 (multi-throw deferred, regression testing completed)
- **Sprint Goal Met:** Yes — trial release delivered with transition-readiness evidence

## Team Health

- **Morale:** Good — customer satisfaction boosted team confidence
- **Collaboration:** Effective — documentation and testing work was well-coordinated
- **Process:** Stable — git workflow, code review, and CI pipeline worked as expected
