# Week 4 Public Report — BasketForm-AI

This is the public report for Week 4 (Assignment 4). It records how the team responded to customer feedback on the MVP v1 increment and how that feedback was turned into traceable backlog work for Sprint 2.

Private submission details (identity, recording links, credentials, access details) are intentionally kept out of this public report and are provided only in the private Moodle PDF wrapper.

## Sprint Context

- **Current Sprint:** Sprint 2 — Assignment 4 (Real ML & Trustworthy Analysis)
- - **Sprint milestone:** [Sprint 2 - Assignment 4](https://github.com/kr1ny77/BasketForm-AI/milestone/2)
  - - **Sprint Backlog board:** [Assignment 4 Sprint Backlog (GitHub Projects Kanban)](https://github.com/users/kr1ny77/projects/7/views/2)
    - - **Roadmap:** [docs/roadmap.md](https://github.com/kr1ny77/BasketForm-AI/blob/main/docs/roadmap.md)
     
      - The Sprint Goal for Sprint 2 is value-focused: deliver a more trustworthy analysis increment by replacing the mock pipeline with real ML, rejecting non-basketball videos, persisting results across restarts, and protecting all of this with automated tests and CI quality gates.
     
      - ## Customer Feedback Source
     
      - Feedback was gathered at the MVP v1 customer review (Anastasia Malakhova, customer). The source records are public in the repository:
     
      - - [Customer review summary (Week 3)](https://github.com/kr1ny77/BasketForm-AI/blob/main/reports/week3/customer-review-summary.md)
        - - [Customer review notes (Week 3)](https://github.com/kr1ny77/BasketForm-AI/blob/main/reports/week3/customer-review-notes.md)
          - - [Customer review transcript (Week 3)](https://github.com/kr1ny77/BasketForm-AI/blob/main/reports/week3/customer-review-transcript.md)
           
            - ## Customer Feedback Response Table
           
            - Every feedback point below is traced to a backlog item and given a clear response. Feedback that is not planned for this Sprint still has an explanation and a linked backlog item.
           
            - | Feedback point | Resulting PBI or issue | Status | Response |
            - |---|---|---|---|
            - | Real ML analysis is needed for production; the mock pipeline is only acceptable for a demo. | [#64](https://github.com/kr1ny77/BasketForm-AI/issues/64) (PBI-016) | In Sprint 2 (Ready) | Replacing the mock processor with a real Python ML pipeline so scores reflect actual shot analysis. Selected as a Sprint 2 must-have. |
            - | The product should verify the video really shows a basketball player/ball, and reject non-basketball clips (e.g. soccer, TikTok) instead of giving a misleading score. | [#65](https://github.com/kr1ny77/BasketForm-AI/issues/65) (PBI-017) | In Sprint 2 (Ready) | Adding basketball + player detection. If a ball or player is not detected, the user gets a clear "not a basketball shot / ball not detected" message instead of a meaningless score. |
            - | Authentication and data should persist across sessions; in-memory storage loses results on server restart. | [#66](https://github.com/kr1ny77/BasketForm-AI/issues/66) (PBI-018) | In Sprint 2 (Ready) | Adding a persistence layer so analysis results survive server restarts, moving the product toward production readiness. |
            - | The shot-phase / release analysis is still rough and needs to track the ball and refine the algorithm. | [#62](https://github.com/kr1ny77/BasketForm-AI/issues/62) (PBI-021) | Done | Ball tracking and the release-phase algorithm were refined and merged; this PBI is complete and feeds into the real ML pipeline work. |
            - | PDF export was planned for MVP v1 (US-010) but was not delivered. | [#70](https://github.com/kr1ny77/BasketForm-AI/issues/70) (PBI-022) | Not planned for this Sprint | Deferred because Sprint 2 prioritises trustworthy analysis, persistence, and automated quality gates over new feature surface. Tracked for a later Sprint and depends on the real ML pipeline so exports reflect real data. |
            - | The customer suggested batch upload for multiple shots in one session. | [#71](https://github.com/kr1ny77/BasketForm-AI/issues/71) (PBI-023) | Not planned for this Sprint | Deferred; depends on real ML, validation, and persistence landing first. Tracked as a backlog item and a candidate for a later Sprint. |
            - | Positive: clean and intuitive upload interface, engaging Canvas animation, responsive on mobile, PDF meets offline record-keeping needs. | No issue (positive feedback) | Acknowledged | No change required. These strengths are preserved; Sprint 2 work must not regress the existing UX, which is protected by the automated tests and CI gates. |
           
            - ## How Quality and Automation Continue
           
            - The quality requirements, automated tests, CI checks, coverage expectations, and Definition of Done updates created in Assignment 4 are maintained project assets. Later PBIs must maintain or extend these gates rather than bypass, disable, or treat them as one-time submission evidence. New work added to the backlog (including the deferred PBIs #70 and #71) explicitly requires its own automated test coverage in the acceptance criteria so the quality bar is carried forward.
           
            - ## Sprint Milestone vs Release Evidence
           
            - The Sprint 2 milestone is planning and scope evidence: it defines what the team committed to and which PBIs are in scope. The SemVer release is separate packaged-increment evidence created after the Sprint work is Done. The two are deliberately kept distinct, in line with the Process Requirements and Product Repository Requirements.
           
            - ## Links
           
            - - [Sprint 2 milestone](https://github.com/kr1ny77/BasketForm-AI/milestone/2)
              - - [Sprint Backlog board (GitHub Projects Kanban)](https://github.com/users/kr1ny77/projects/7/views/2)
                - - [Product roadmap](https://github.com/kr1ny77/BasketForm-AI/blob/main/docs/roadmap.md)
                  - - [Definition of Done](https://github.com/kr1ny77/BasketForm-AI/blob/main/docs/definition-of-done.md)
