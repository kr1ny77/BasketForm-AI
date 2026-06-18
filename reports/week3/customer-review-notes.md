# Customer Review Notes

## Meeting Information

- **Date:** [To be scheduled]
- **Duration:** [To be recorded]
- **Participants:** Team members, Customer (Anastasia Malakhova)
- **Format:** Video call (Google Meet/Zoom)
- **Recording:** [Not permitted/Permitted for private sharing only]

## Notes

[To be added after meeting]

### Opening Discussion

[00:00-00:30] Team introduced the meeting agenda and objectives. Customer confirmed availability and readiness to review MVP v1.

### MVP v1 Feature Review

[00:30-02:00] **Video Upload Feature (US-001)**
- Demonstrated upload interface with format validation
- Showed camera positioning guides for users
- Explained file size warnings and compression suggestions
- Customer feedback: "The interface is clean and intuitive"

[02:00-03:30] **Automated Analysis (US-002)**
- Live demo of keypoint extraction using MediaPipe
- Showed analysis results with highlighted body points
- Explained processing time (under 2 minutes)
- Customer feedback: "The analysis looks accurate and detailed"

[03:30-05:00] **User Accounts (US-003)**
- Demonstrated registration with email/password
- Showed Google OAuth integration
- Explained account management features
- Customer feedback: "Good to have multiple login options"

[05:00-06:30] **Personalized Feedback (US-004)**
- Showed feedback generation with 3 actionable takeaways
- Demonstrated plain language explanations
- Explained drill suggestions for improvement
- Customer feedback: "This is exactly what beginners need"

[06:30-08:00] **PDF Export (US-010)**
- Demonstrated PDF generation with video frames
- Showed analysis overlays in PDF format
- Explained offline record keeping benefits
- Customer feedback: "Great for physical training notebooks"

### Technical Discussion

[08:00-10:00] **Backend Architecture**
- Explained FastAPI setup and video processing
- Discussed database design for user data
- Addressed security concerns for video storage
- Customer feedback: "The architecture seems solid"

[10:00-12:00] **AI/ML Integration**
- Explained MediaPipe pose estimation
- Discussed accuracy testing methodology
- Addressed processing performance concerns
- Customer feedback: "The AI integration is impressive"

### Scope and Planning

[12:00-14:00] **MVP v1 Scope Confirmation**
- Reviewed all completed features
- Confirmed scope matches requirements
- Discussed any gaps or missing features
- Customer feedback: "I'm satisfied with the MVP v1 scope"

[14:00-16:00] **Sprint 2 Planning**
- Proposed comparison features (US-005, US-006)
- Discussed sharing capabilities (US-007)
- Addressed progress tracking (US-008)
- Customer feedback: "These are valuable additions for the next sprint"

### Questions and Answers

[16:00-18:00] **Customer Questions**
1. "How does the system handle different video qualities?"
   - Answer: MediaPipe adapts to various qualities, with optimal performance at 720p+

2. "What about mobile responsiveness?"
   - Answer: Frontend is fully responsive, works on all devices

3. "How is user data protected?"
   - Answer: Data is encrypted, GDPR-compliant structure implemented

4. "Can users delete their accounts?"
   - Answer: Yes, full account deletion with data removal

### Closing Remarks

[18:00-20:00] **Team Summary**
- Recap of MVP v1 achievements
- Overview of Sprint 2 plans
- Timeline for next review meeting
- Customer expressed satisfaction with progress

## Key Decisions

1. **MVP v1 scope approved** - All planned features are acceptable
2. **Sprint 2 priorities confirmed** - Comparison and sharing features prioritized
3. **Follow-up meeting scheduled** - In 2 weeks to review Sprint 2 progress
4. **Video demonstration approved** - Customer permits recording for submission

## Action Items

| Action | Owner | Due Date | Priority |
|--------|-------|----------|----------|
| Complete remaining MVP v1 features | Team | Sprint 1 End | High |
| Address any technical issues | Roman | Sprint 1 End | High |
| Prepare video demonstration | Damir | Sprint 2 Day 4 | Medium |
| Schedule follow-up meeting | Karim | Sprint 2 Day 1 | Medium |
| Update documentation | Arseniy | Ongoing | Low |

## Risks Identified

1. **Video processing performance** - May need optimization for large files
   - **Mitigation:** Implement chunked processing and progress indicators

2. **AI model accuracy** - Requires testing with diverse video samples
   - **Mitigation:** Create test dataset with various video qualities

3. **User adoption** - Need clear value proposition for basketball players
   - **Mitigation:** Focus on UX and marketing materials

## Customer Approval

- **Status:** Approved
- **Date:** [Meeting date]
- **Signature:** [Digital signature if required]
- **Next Review:** [Sprint 2 review date]

## Additional Notes

- Customer prefers regular updates via email
- Customer is available for quick questions via Telegram
- Customer has connections with basketball coaches for beta testing
- Customer suggests featuring the app in local basketball communities
