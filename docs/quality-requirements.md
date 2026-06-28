# Quality Requirements

Quality requirements for BasketForm-AI, structured by ISO/IEC 25010 quality sub-characteristics.

## QR-001: API Response Time

**ISO/IEC 25010 sub-characteristic:** Time Behaviour

**Scenario:** When an authenticated user submits a request to load the results list or a single result under normal load, the API shall return a response within 2 seconds for 95% of requests.

**Why this matters:** Users need fast feedback when browsing their analyses. Slow API responses degrade the user experience and block the main workflow of reviewing shooting form results.

**Linked quality requirement tests:** [QRT-001](quality-requirement-tests.md#qrt-001-api-response-time)

## QR-002: Authentication Security

**ISO/IEC 25010 sub-characteristic:** Confidentiality

**Scenario:** When an unauthenticated user attempts to access a protected endpoint (e.g., `/upload`, `/results`, `/api/videos`), the system shall reject the request with HTTP 401/403 and redirect to `/login`. A user shall not be able to access another user's results or data through direct URL manipulation.

**Why this matters:** User data (video uploads, analysis results, friend lists) must remain private. Unauthorized access would expose personal data and violate user trust.

**Linked quality requirement tests:** [QRT-002](quality-requirement-tests.md#qrt-002-authentication-security)

## QR-003: Critical Module Testability

**ISO/IEC 25010 sub-characteristic:** Testability

**Scenario:** When a developer changes a critical product module (e.g., `internal/services`, `internal/handlers`) under the standard CI environment, the module shall have automated unit tests that achieve at least 30% line coverage for that module.

**Why this matters:** Core business logic (authentication, storage, video processing) must be directly verifiable so defects can be detected before merge. Low coverage leads to undetected regressions in critical user workflows.

**Linked quality requirement tests:** [QRT-003](quality-requirement-tests.md#qrt-003-critical-module-unit-coverage)

## QR-004: Usability — Form Guidance

**ISO/IEC 25010 sub-characteristic:** Usability

**Scenario:** When a user opens any form page (login, register, upload, profile), all input fields shall contain placeholder text or visible labels that explain what data is expected, within the initial page render.

**Why this matters:** Clear guidance reduces user errors, lowers support needs, and makes the application accessible to new users without external documentation.

**Linked quality requirement tests:** [QRT-004](quality-requirement-tests.md#qrt-004-usability-form-guidance)
