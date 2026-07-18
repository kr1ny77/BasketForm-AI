# Reflection — Week 7

## Learning points

Backlog hygiene needs to be continuous, not a sprint-end activity. Auditing the five outstanding Sprint 5 PBIs (#135–#139) found that several already had merged work, yet their Work Status fields still read "To Do — selected for Sprint 5 (Week 7)." The lesson is that issue and Project board status should be updated as pull requests merge, not batched at the end of the sprint.

Implementer/reviewer fields must be reconciled against actual PR authorship. The audit surfaced a mismatch on PBI-046: the issue's listed implementer/reviewer did not match who actually authored the merged deployment-script PR. Assigning people at planning time does not guarantee the assignment still reflects reality once the work happens, so this needs an explicit check before closing an issue.

Cross-document consistency requires a deliberate side-by-side check. docs/customer-handover.md still carried Week 6-era "Handover level" and "Customer confirmation status" wording even after reports/week7/README.md had already been updated with the Week 7 outcome. One updated document does not guarantee the other stays in sync automatically.

A version-numbering collision is easy to introduce and easy to catch late. Sprint 5 work briefly reused the v0.4.0 tag already assigned to the Week 6 trial release. Catching this before the final cut avoided shipping two different releases under the same version number.

CI signal quality matters even for documentation-only changes. Adding new report cross-links before their target files exist produces an avoidable broken-link failure; that kind of failure is the responsibility of whoever added the link, and is separate from unrelated pre-existing code-level CI failures.

## Validated assumptions

- The deployment automation delivered for PBI-046 (single deploy script, automated smoke test, documentation in docs/customer-handover.md) is real and merged, though a live redeploy verification and a recorded second-reviewer approval are still outstanding.
- The hardcoded-path and data-race fixes delivered for PBI-043 are genuinely merged and address real defects, though a written regression report and Should-Have bug triage remain open.
- The customer's Week 7 confirmation (independent upload of two videos, correct progress-tracking read-out) is grounded in sprint-review-summary.md and sprint-review-transcript.md, not assumed.

## Friction and gaps

- Sprint 5 board hygiene lagged behind the work that had actually been merged for most of the five audited PBIs.
- PBI-045 (Demo Day prep) and PBI-047 (final course retrospective) had not been started as of this audit.
- Formal second-reviewer approval was not consistently recorded before merge on Sprint 5 pull requests.

## Planned response

| Action | Owner | Target |
| --- | --- | --- |
| Update Work Status and Project board fields at merge time, not sprint end | Team | Ongoing |
| Resolve the PBI-046 implementer/reviewer discrepancy | Team | Before Sprint 5 close |
| Complete the regression report and Should-Have triage for PBI-043 | Koomaz / Team | Before Sprint 5 close |
| Complete Demo Day prep (PBI-045) and the final retrospective (PBI-047) | mentalafffection / Team | Before Demo Day |
| Add the recommended camera-angle UI hint and automatic video splitting | Team | Post-course |
