# Sprint Retrospective

## What went well
- **ML Infrastructure Optimization:** Successfully migrated the ML inference from a heavy, locally hosted `[redacted]` model to an API-based `[redacted]` model. This completely resolved the server overload issues and drastically reduced startup/feedback times.
- **Smooth UAT Execution:** The User Acceptance Tests for the newly implemented Friends system went flawlessly. Both the "send request" and "accept request" flows passed on the first attempt, validating our recent focus on writing better test scenarios.
- **Collaborative Problem Solving:** The brainstorming session with the Customer regarding the upcoming video segmentation feature was highly productive. The Customer's suggestion to use the "ball-release" moment as a trigger for frame queuing provided a clear, actionable technical direction.
- **Requirement Fulfillment:** We successfully reviewed the original requirements spreadsheet with the Customer and confirmed that all core features have been delivered or explicitly approved in their modified forms.

## What did not go well
- **Looming Deadline Pressure:** With only two sprints (approx. 2 weeks) remaining until the final defense, the team feels the time pressure. The Customer explicitly warned about the risk of ending up with a "broken product" if we overcommit. The complexity of the new video segmentation feature poses a risk to the stability of the current build.

## What the team changed or attempted to change based on the previous Sprint Retrospective, and what results they observed
- **Attempted Change:** Migrate the ML backend to an external API to fix server bottlenecks and slow startup times.
  - **Observed Result:** Highly successful. Server resource utilization dropped significantly, and the feedback generation is now near-instant. The core algorithm remained unchanged, ensuring consistency.
- **Attempted Change:** Enforce a stricter "Definition of Done" that requires a dry-run of User Acceptance Tests before marking a feature as complete.
  - **Observed Result:** The Friends system UAT was prepared and executed seamlessly during the review meeting. The dry-run ensured there were no last-minute blockers or UI glitches during the actual customer test.

## Action points
1. **Timebox the Video Segmentation Feature:** Implement the multi-throw auto-segmentation using the "ball-release trigger" logic, but enforce a strict 3-day timebox. If the feature is not fully stable and integrated by day 3, we will pivot to a simpler manual-trimming fallback to protect the final release stability.
2. **Implement Mid-Sprint Regression Testing:** Introduce a mandatory mid-sprint regression test specifically for the core single-throw analysis pipeline. This will ensure that the new multi-video ingestion logic does not introduce bugs or performance regressions into the existing, working product.
