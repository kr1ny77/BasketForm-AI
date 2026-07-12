# Reflection — Week 6

## Learning points

1. **Timeboxing is essential for experimental features.** We attempted to implement multi-throw video upload with automatic ball-release segmentation in a single sprint, but the technical complexity exceeded our capacity. We learned that ambitious features must be explicitly timeboxed with a clear fallback plan (in our case, keeping manual upload) and go/no-go criteria defined before work begins.

2. **Input validation at the pipeline entry point creates compounding value.** Implementing person-detection validation — rejecting videos that don't contain a human — turned out to be more impactful than expected. It not only prevents nonsensical analysis results but also improves perceived product quality and saves compute resources. A small check at the boundary of the system prevented a whole class of bad user experiences.

3. **Customer-facing documentation should be reviewed during the sprint, not only at the review meeting.** We created CONTRIBUTING.md, AGENTS.md, and docs/customer-handover.md and got customer feedback within the same sprint. This caught potential misunderstandings early and gave us confidence that the transition-readiness artifacts are actually usable.

4. **Manual regression testing is a hidden cost that grows with the project.** Mid-sprint regression testing of the core pipeline took more time than anticipated because it was not automated. The lesson is that investing in automated regression tests early pays off across every subsequent sprint.

5. **LLM tools are most valuable for accelerating repetitive, structured work.** Using an LLM to generate and update documentation files (CONTRIBUTING.md, AGENTS.md, customer-handover.md, weekly reports) saved significant time. The key was that all technical decisions, scope selection, and code changes remained entirely the team's own work — the LLM was used strictly for synthesis and formatting.

## Validated assumptions

- **The core single-throw analysis pipeline is stable.** Regression testing confirmed that the scoring and analysis logic remains correct after multiple sprints of changes. No regressions were found.
- **Person-detection validation effectively filters invalid input.** UAT-007 (uploading a video without a human) passed — the customer confirmed that the error message is displayed correctly and the analysis is not attempted.
- **The product meets the Definition of Done.** The customer explicitly confirmed that all primary tasks are completed and the product is ready for transition.
- **The documentation is clear and intuitive.** The customer reviewed the updated User Guide, README, and the new "Share" button instructions and found them accurate and easy to follow.
- **The "Friends" sharing feature works correctly.** UAT-008 passed — the customer successfully shared analysis progress via the Friends tab.
- **The product is transition-ready.** Deployment access, environment variables, recovery steps, and source code access are all documented and verified.

## Friction and gaps

- **Multi-throw video segmentation was deferred.** The automatic splitting of a large video into separate clips upon upload proved too complex for a one-sprint timebox. This was a planning gap — the feature should have been broken down into smaller stories or explicitly deferred from the start.
- **Regression testing was not automated.** The manual regression pass consumed time that could have been spent on other improvements. This is a process debt that will need to be addressed.
- **The sprint felt less "productive" in terms of new features.** Because the focus shifted to documentation, validation, and transition readiness, the volume of visible new functionality was lower than in previous sprints. While this was the right call for a trial release, it required a mindset shift for the team.

## Planned response

| Action | Owner | Target |
|--------|-------|--------|
| Automate regression tests for the core single-throw pipeline | Team | Sprint 5 (Week 7) |
| Conduct thorough QA testing and fix any bugs found during customer's early-week review | Team | Sprint 5 (Week 7) |
| Optimize site performance and refactor code for maintainability | Team | Sprint 5 (Week 7) |
| Prepare for Demo Day presentation | Team | Sprint 5 (Week 7) |
| Final documentation updates based on any last-minute customer feedback | Team | Sprint 5 (Week 7) |
| Post-course: implement multi-throw video upload with automatic ball-release segmentation | Team | After course completion |
