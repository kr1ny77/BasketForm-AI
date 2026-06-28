# LLM Usage Report — Week 4

## AI Tools Used

- **MiMo Code Agent** (mimo-auto model) — used as the primary development assistant throughout Sprint 2.

## How AI Was Used

### Code Generation
- Generated authentication service (`auth.go`) with bcrypt hashing and JWT token generation.
- Generated storage layer (`storage.go`) with JSON file persistence for users, videos, friends, and shared results.
- Generated auth, friends, and share HTTP handlers with proper error handling and JSON responses.
- Generated JWT middleware for route protection.
- Generated HTML templates for login, signup, friends, shared results, and profile pages.
- Generated the enhanced ML Python script with pose overlay, HUD, and phase analysis output.
- Generated integration tests, unit tests, and QRT tests.

### Architecture Decisions
- AI recommended using bcrypt for password hashing (industry standard).
- AI suggested JWT with HttpOnly cookies for session management (prevents XSS token theft).
- AI recommended storing data in JSON files for the current project scale (no database needed).
- AI suggested ISO/IEC 25010 sub-characteristics for structuring quality requirements.

### Documentation
- Generated quality requirements document (QR-001 to QR-004) following ISO/IEC 25010 format.
- Generated quality requirement tests document (QRT-001 to QRT-004).
- Generated testing strategy document with critical modules and coverage targets.
- Generated user acceptance test scenarios (UAT-001 to UAT-004).
- Generated week4 report files.

### CI and Quality
- Generated the updated CI pipeline configuration with lint, test, coverage, QRT, and govulncheck steps.
- Generated the coverage threshold check script for critical modules.

## Human Review

All AI-generated code was reviewed by the team before merging. The AI provided a strong starting point, but human judgment was applied to:
- Security decisions (JWT secret management, cookie settings)
- UI/UX design (template layouts, CSS styling)
- Quality requirement selection (which ISO/IEC 25010 sub-characteristics to address)
- Test scenarios (which edge cases to cover)

## Impact

AI assistance significantly accelerated development velocity, allowing the team to complete authentication, social features, enhanced ML, quality automation, and documentation within a single Sprint. The AI-generated code required minimal modifications after review.
