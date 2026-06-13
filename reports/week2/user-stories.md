# User Stories - BasketForm AI

## User Roles/Personas

1. **Basketball Player** - Amateur or professional player looking to improve shooting technique
2. **Basketball Coach** - Trainer analyzing player performance
3. **Sports Analyst** - Data-driven analyst studying biomechanics
4. **Basketball Enthusiast** - Casual fan interested in technique analysis

## User Stories

## US-01: Upload shooting form video

Requirement status: Active
MoSCoW priority: Must Have

As a beginner basketball player,
I want to upload a video of my shooting form,
so that the system can analyze my mechanics and provide feedback.

### Notes and constraints
- Supported video formats should be limited to common web formats (e.g., MP4, MOV) to ensure compatibility.
- The upload process must include clear visual guides on camera positioning (e.g., side profile, full body in frame).
- Client-side compression or size warnings should be implemented to improve upload reliability.

---

## US-02: View automated form analysis

Requirement status: Active
MoSCoW priority: Must Have

As a beginner basketball player,
I want to receive an automated analysis of my uploaded shooting video,
so that I can identify specific areas for improvement (e.g., elbow alignment, follow-through, release point).

### Notes and constraints
- The analysis should highlight key biomechanical checkpoints using visual overlays or timestamped markers on the video.
- Processing time should remain under 2 minutes for standard mobile recordings.
- Open question: Will the analysis run server-side via cloud inference or on-device via WebAssembly? (To be decided in Assignment 3).

---

## US-03: Create and manage a user account

Requirement status: Active
MoSCoW priority: Must Have

As a new user,
I want to create a personal account,
so that I can securely save my uploaded videos, analysis history, and progress over time.

### Notes and constraints
- Account creation should support email/password and at least one OAuth provider (Google or Apple) for low friction.
- Basic data privacy compliance (GDPR/CCPA-ready structure) should be considered from day one.
- Users must be able to delete their account and associated data per privacy requirements.

---

## US-04: Receive simplified, actionable feedback

Requirement status: Active
MoSCoW priority: Must Have

As a beginner basketball player,
I want to receive simplified, easy-to-understand feedback on my shot,
so that I am not overwhelmed by technical jargon and can make immediate, practical corrections.

### Notes and constraints
- Feedback should be presented in plain language with a maximum of 3 prioritized, actionable takeaways per analysis.
- A glossary or "tap-to-learn-more" toggle can be added later if technical terms are necessary.

---

## US-05: Compare form with professional players

Requirement status: Active
MoSCoW priority: Could Have

As a beginner basketball player,
I want to compare my shooting form side-by-side with a professional player's form,
so that I can visually understand the ideal mechanics I should be aiming for.

### Notes and constraints
- Requires a curated, legally cleared library of professional shooting clips.
- Video players must support synchronized playback or independent frame-by-frame scrubbing.
- Aspect ratio normalization will be needed to ensure accurate visual comparison.

---

## US-06: Share analysis report with a coach

Requirement status: Active
MoSCoW priority: Should Have

As a beginner basketball player,
I want to share my video and analysis report with my coach via a secure link,
so that they can review my form and provide personalized, expert guidance.

### Notes and constraints
- Shared links should grant view-only access to the specific video/analysis pair.
- Links should expire after 30 days by default for security and storage hygiene.
- Coach account creation should not be required to view shared reports.

---

## US-07: Share progress with friends

Requirement status: Active
MoSCoW priority: Should Have

As a beginner basketball player,
I want to share my shooting progress and improvement milestones with friends,
so that I can celebrate achievements and maintain accountability through social support.

### Notes and constraints
- Sharing should generate a clean, non-editable summary card containing a before/after thumbnail, date range, and key metric improvements.
- Requires explicit user consent per

This scope provides the core functionality for analyzing basketball shooting technique while remaining small enough to prototype and discuss meaningfully with the customer.
shared item to maintain privacy controls.
- External sharing (email, social media) and in-app sharing (if friend lists are implemented later) should be architected separately.

---

## US-08: Track shooting progress over time

Requirement status: Active
MoSCoW priority: Should Have

As a beginner basketball player,
I want to track my shooting progress over time with visual charts,
so that I can stay motivated and see tangible evidence of my improvement.

### Notes and constraints
- Progress should be visualized using simple line or bar charts mapping key metrics (e.g., elbow angle consistency, release timing) across dates.
- Requires a minimum of 3-4 successful analyses to render meaningful trends.
- Users should be able to filter progress by specific mechanical categories.

---

## US-09: Public social feed of user shots

Requirement status: Removed
Previous MoSCoW priority: Could Have

As a beginner basketball player,
I want to browse a public social feed of other users' shots and analyses,
so that I can learn from the community and stay motivated.

Reason: Out of scope for the core value proposition of individual form analysis. It introduces significant privacy, content moderation, community management, and scope creep issues that are not required for the initial product or course timeline.

---

## US-10: Export analysis report as PDF

Requirement status: Active
MoSCoW priority: Should Have

As a beginner basketball player,
I want to export my analysis report as a PDF,
so that I can keep an offline record or print it for my physical training notebook.

### Notes and constraints
- The PDF should include a static snapshot of the video frame with analysis overlays, the key actionable takeaways, and the generation date.
- PDF generation will rely on server-side templating to ensure consistent formatting across devices.
## Initial proposed MVP v1 scope

Based on the active Must Have user stories, the initial proposed MVP v1 scope includes:

1. **US-01**: Video Upload and Recording
2. **US-02**: Biomechanical Keypoint Extraction
3. **US-03**: Shooting Technique Evaluation
4. **US-04**: Personalized Feedback Generation
5. **US-10**: PDF Export of the Feedback

This scope provides the core functionality for analyzing basketball shooting technique while remaining small enough to prototype and discuss meaningfully with the customer.
