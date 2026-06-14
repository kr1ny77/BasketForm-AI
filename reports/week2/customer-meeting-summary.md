# Customer Meeting Summary - BasketForm AI

**Meeting Date:** 13.06.2026
**Duration:** 13min 20sec
**Platform:** KTalk
**Recording Permission:** [Yes]
**Transcript Publication Permission:** [Yes]

## Participants or Roles

- **[Customer]**: Customer
- **[ML Engineer 1]**: Developer (BasketForm-AI)
- **[ML Engineer 2]**: Developer (BasketForm-AI)

---

## Artifacts Demonstrated

- **User Stories List**: Document detailing user stories with MoSCoW priorities.
- **MVP v0 Prototype**: Visual prototype of the application in dark mode.

## Discussion Points

- **User Stories & MoSCoW Priorities**: Reviewed 10 user stories covering video upload, automated analysis, personal accounts, simplified feedback, sharing with coaches/friends, pro comparison, social feed, and PDF export.
- **MVP v1 Scope**: Proposed focusing on a working neural network model tuned for basketball and PDF export.
- **Data Collection & Model Training**: Discussed the lack of pre-trained models specifically for sports mechanics. Need to use skeleton recognition models and fine-tune them on custom datasets of basketball poses, jumps, and throws.
- **Video Upload & Protection**: Confirmed abandoning instant video recording in favor of file uploads. Discussed the need for file type validation to prevent non-video uploads (e.g. images).
- **UI/UX for Feedback**: Discussed how to present analysis results without overwhelming the user with technical jargon.

## Decisions

- Approved the MoSCoW prioritization for all user stories.
- Approved MVP v1 scope: Must-have stories + PDF export (prioritized over the friend/buddy system to prevent data loss).
- Dropped instant video recording; proceeding with video file uploads.
- Decided that data persistence is critical; PDF export chosen for MVP v1 to ensure users don't lose data if the account system isn't fully ready.
- Agreed to push the public social feed to a future release.

## Action Points

- Update user stories to reflect the UI decision for feedback (short description first, expandable "read more").
- Implement file type protection/validation for the video upload feature.
- Start collecting a custom dataset of basketball videos (poses, jumps, throws) immediately.
- Search for and evaluate existing skeleton recognition models/weights (**[Customer]** to provide resource links).
- Schedule a follow-up meeting once a specific ML model is selected.
- Develop MVP v1 with the working model and PDF export functionality.

## Risks

- **Model Availability**: No fully trained models exist specifically for sports mechanics; reliance on general skeleton recognition models that require extensive custom training.
- **Dataset Collection**: Gathering and trimming sufficient video data for training might take longer or require more resources than anticipated.
- **Implementation Complexity**: Sharing reports with a coach via a secure link was noted as quite difficult to implement within the current month.

## Feedback

- **Prototype Visuals**: **[Customer]** liked the MVP v0 prototype, specifically the dark color theme.
- **Feedback Presentation**: Suggested showing a short, simple comment initially, with an option to read more detailed technical feedback.
- **Pro Comparison**: Validated the "Could" priority for side-by-side pro comparison, noting that every player has a unique style and shouldn't just copy pros.
- **Data Persistence**: Emphasized the importance of not losing user data, suggesting PDF export as a fallback if the account system isn't ready.

## Customer Approvals

- Formally approved all user stories and their MoSCoW priorities (with discussed adjustments).
- Approved the MVP v1 scope (Must-have stories + PDF export).
- Approved the visual design of the MVP v0 prototype.
- Granted permission for meeting recording and public transcript publication.

## Resulting Changes

- **User Stories**: Added UI requirement for analysis feedback to have a brief summary with an expandable detailed view.
- **MVP v1 Scope**: Adjusted to include PDF export instead of the social/friend sharing system.
- **Feature Removal**: Removed instant video recording from the scope; focusing entirely on video file uploads.
- **Validation**: Added requirement for file type protection on uploads.

## Linked Artifacts & Evidence

- **User Stories**: 
  - US-01: Upload shooting video (Must)
  - US-02: Automated analysis (Must)
  - US-03: Personal account (Must)
  - US-04: Simplified feedback (Must)
  - US-05: Share with coach (Should)
  - US-06: Share milestones with friends (Should)
  - US-07: Compare with pro (Could)
  - US-08: Track shooting progress over time (Should)
  - US-08: Public social feed (Could)
  - US-09: Export as PDF (Could -> prioritized for MVP v1)
- **Prototype Screens**: MVP v0 Prototype (Dark Mode UI)
- **API/Interface Artifacts**: Video Upload Endpoint (requires file type validation/protection)

---

## Sanitization Notes

- Replaced real names with GitHub/GitLab usernames
- Removed any PII or confidential business information
- Preserved enough context for evaluation
- Used [inaudible] where audio was unclear
- Used [redacted] for sensitive information

---

**Summary Prepared by:** Roman
**Date:** 13.06.2026
**Status:** Sanitized and ready for publication
