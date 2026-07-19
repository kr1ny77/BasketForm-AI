# Retrospective — Week 7

## What Went Well

- Version conflict caught before release: the Sprint 5 changes briefly reused the v0.4.0 tag already assigned to the Week 6 trial release; this was found and corrected to v0.5.0 before the final cut.
- Deployment automation delivered: PBI-046 produced a single deploy script with an automated post-deploy smoke test, documented in docs/customer-handover.md.
- Cross-document consistency restored: docs/customer-handover.md and reports/week7/README.md now state the same Handover Level and Customer-confirmation status, grounded in the actual Week 7 meeting record.
- Customer independently used the product: UAT-009 passed, with the customer uploading two videos in sequence and correctly reading the progress-tracking feature without live step-by-step guidance.

## What Did Not Go Well

- Backlog hygiene: five PBIs (#135–#139) remained marked "To Do" despite some of their work already being merged, and none could be closed outright because none had every acceptance criterion satisfied yet.
- Reviewer/implementer field mismatch: PBI-046 listed an implementer/reviewer assignment that did not match who actually authored the merged deployment-script PR; this was only caught during a manual audit, not the normal workflow.
- Review discipline: at least one merged Sprint 5 pull request did not have a recorded formal approval from a second team member before merge.
- Late-starting scope: PBI-045 (Demo Day prep) and PBI-047 (final course retrospective) had no recorded work partway through the sprint.

## Action Items

| Action | Owner | Deadline |
| --- | --- | --- |
| Reconcile the PBI-046 implementer/reviewer fields with actual PR authorship | Team | Before Sprint 5 close |
| Record a second-reviewer approval on outstanding Sprint 5 pull requests | Team | Before Sprint 5 close |
| Start and complete Demo Day prep and the final retrospective | mentalafffection / Team | Before Demo Day |
| Add the missing regression report and Should-Have triage for PBI-043 | Koomaz / Team | Before Sprint 5 close |

## Velocity and Process

- Story Points Planned: 28
- Story Points Completed: partial — PBI-042 delivered; PBI-043 and PBI-046 partially complete (fixes merged, some acceptance criteria still open); PBI-044 partially complete (handover-level and version-conflict fixes merged, remaining documentation checks open); PBI-045 and PBI-047 not started as of this audit
- Sprint Goal Met: In progress — the final transition outcome and MVP v3 stabilization are on track, but full Sprint 5 closure (all PBIs closed, final release cut) is not yet complete

## Team Health

- Morale: Good — the customer's final confirmation reinforced confidence in the product
- Collaboration: Mixed — cross-file documentation consistency required a manual audit rather than being caught earlier in the normal review flow
- Process: Needs improvement — board/status hygiene and review-before-merge discipline were the main gaps this sprint
