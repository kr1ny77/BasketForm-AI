# Reflection — Week 6

## What Went Well

- Person-detection validation was successfully implemented and tested, preventing invalid video uploads
- Customer-facing documentation (CONTRIBUTING.md, AGENTS.md, customer-handover.md) was created and approved by the customer
- The core single-throw analysis pipeline remained stable throughout the sprint
- Customer trial confirmed the product meets the Definition of Done
- Transition-readiness evidence was collected and documented

## What Did Not Go Well

- Multi-throw video upload with automatic ball-release segmentation was deferred due to technical complexity exceeding the one-sprint timebox
- Some time was spent on regression testing that could have been automated earlier

## Lessons Learned

1. **Timeboxing is essential:** Features that exceed a sprint's capacity should be explicitly timeboxed with a fallback plan (manual upload in this case)
2. **Person-detection validation adds significant value:** Preventing invalid uploads improves user experience and data quality
3. **Documentation review with the customer early:** Getting customer feedback on documentation during the sprint (not just at review) catches issues earlier
4. **Regression testing should be automated:** Manual regression testing is time-consuming; investing in automated tests pays off across sprints

## What We Would Do Differently

- Start regression test automation earlier in the sprint
- Timebox experimental features more aggressively with clear go/no-go criteria
- Involve the customer in documentation review earlier in the sprint

## Sprint 4 Outcome

The sprint achieved its primary goal: a stable trial release with transition-readiness evidence. The customer confirmed the product works and documentation is clear. The deferred multi-throw feature does not block the final delivery.
