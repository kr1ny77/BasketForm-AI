# Customer Review Transcript

## Meeting Information

- **Date:** 2026-06-28
- **Duration:** 30 minutes
- **Participants:** Team (Karim, Roman, Kamil, Arseniy, Damir), Customer (Anastasia Malakhova)
- **Recording:** Not recorded (customer preferred no recording)
- **Transcript:** Based on meeting notes

## Transcript

[00:00] **Karim:** Welcome to our Sprint 1 review. Today we'll present our MVP v1 implementation of BasketForm-AI.

[00:10] **Anastasia:** Thank you for the invitation. I'm excited to see the progress.

### Video Upload Feature

[00:20] **Arseniy:** Let me start with the upload page. Users can drag and drop videos or use the file picker. We support MP4, MOV, and AVI formats with a 100MB limit.

[00:40] **Anastasia:** The upload interface looks clean. I like the basketball-themed background animation.

[00:50] **Arseniy:** Thank you. The Canvas animation has 35 basketball objects that react to mouse movement.

### Analysis Feature

[01:00] **Roman:** Next, I'll show the analysis flow. After upload, the system processes the video and generates a score between 40 and 90, along with personalized feedback.

[01:20] **Anastasia:** How does the feedback work?

[01:25] **Roman:** We generate up to 3 actionable takeaways in plain language, like "Good release point. Work on keeping your elbow aligned with the basket."

[01:35] **Anastasia:** That's exactly what beginners need. Simple and actionable.

### Results and Charts

[01:45] **Kamil:** The results page shows a radar chart for score breakdown (stance, arm angle, release point, follow-through) and a scatter plot for pose keypoints.

[02:00] **Anastasia:** The visualizations are clear. Can users filter their results?

[02:05] **Kamil:** Yes, by filename and by date or score. The modal shows detailed analysis when you click a result card.

### PDF Export

[02:15] **Damir:** Users can export their analysis as a PDF report with score, feedback, and metadata.

[02:25] **Anastasia:** Great for keeping records. I know some users prefer physical copies.

### Profile Page

[02:35] **Karim:** The profile page lets users manage their name, email, and password. It's UI-only for now — real authentication is planned for Sprint 2.

[02:45] **Anastasia:** That's fine for MVP. When will real authentication be ready?

[02:50] **Karim:** Sprint 2, along with comparison and sharing features.

### Sprint 2 Planning

[03:00] **Karim:** For Sprint 2, we're planning professional player comparison, coach sharing, and progress tracking.

[03:10] **Anastasia:** Those sound valuable. I'm happy with the current progress. MVP v1 is approved.

## Key Takeaways

1. **MVP v1 approved** by customer
2. **Mock ML acceptable** for demo purposes
3. **Sprint 2 priorities confirmed** — comparison and sharing
4. **Authentication deferred** to Sprint 2

## Approval

- **Customer:** Anastasia Malakhova
- **Date:** 2026-06-28
- **Status:** Approved
