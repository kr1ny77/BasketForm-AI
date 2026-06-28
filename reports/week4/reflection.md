# Reflection — Week 4

## Learning Points

- Learned how to implement JWT-based authentication in Go with bcrypt password hashing and HttpOnly cookies.
- Understood the importance of per-user data isolation — videos and results are now tied to user IDs.
- Gained experience with ISO/IEC 25010 quality model and how to write measurable quality requirements.
- Learned how to structure QRT (Quality Requirement Tests) as automated CI checks that directly verify quality scenarios.
- Understood the difference between unit tests, integration tests, and QRTs in terms of scope and what they verify.

## Validated Assumptions

- **Assumption:** JSON file storage is sufficient for a course project. **Validated:** JSON files in `data/` directory work well for users, videos, friends, and shared results. No database needed for the current scale.
- **Assumption:** MediaPipe Holistic can provide useful basketball shot analysis. **Validated:** The pipeline extracts pose landmarks and scores stance, arm angle, release, and follow-through with reasonable accuracy.
- **Assumption:** jsPDF is sufficient for PDF export on the frontend. **Validated:** jsPDF generates clean PDF reports with score breakdown, phases, and full feedback text.

## Friction and Gaps

- The ML pipeline requires MediaPipe and OpenCV Python packages to be installed on the server, which adds deployment complexity compared to a pure Go solution.
- The current auth middleware applies to all routes uniformly; more granular permission control (e.g., admin vs. regular user) is not yet implemented.
- No rate limiting on API endpoints — potential for abuse if deployed publicly without a reverse proxy.
- Integration tests for sharing require creating result files manually since the ML pipeline is not run during tests.

## Planned Response

- **Next Sprint:** Add rate limiting middleware or deploy behind a reverse proxy (nginx) with rate limiting.
- **Next Sprint:** Consider a lightweight database (SQLite) if the JSON file approach shows performance issues with many users.
- **Later:** Implement admin roles and more granular permissions.
- **Later:** Add API contract testing (OpenAPI validation) as an additional QA check.
