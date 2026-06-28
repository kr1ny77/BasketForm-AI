# Retrospective — Week 4

## What Went Well

- Authentication system was implemented cleanly with bcrypt and JWT — secure by default.
- The ML pipeline integration worked well: MediaPipe produces useful phase scores and the annotated video is a strong visual feature.
- Quality requirements and QRTs were well-structured and caught real issues during development (e.g., response time testing revealed no bottlenecks).
- CI pipeline is comprehensive: lint, test, coverage, QRT, govulncheck all running on every PR.
- Social features (friends, sharing) were straightforward to implement with JSON file storage.

## What Did Not Go Well

- The QRT tests initially failed because of file path issues when running from a subdirectory — had to use `runtime.Caller` to find the repo root.
- Integration tests for sharing required manually creating result files since the ML pipeline is not run during tests — this creates a gap between test and production behavior.
- The `data/` directory was shared between tests initially, causing test pollution — fixed by making `DATA_DIR` configurable via environment variable.

## Changes from Previous Sprint

- Improved test isolation by using `t.TempDir()` and environment variables for test-specific directories.
- Added QRT tests as a first-class CI step, which was not present in Sprint 1.
- Structured quality requirements following ISO/IEC 25010 format from the start, instead of retrofitting.

## Concrete Process Improvements for Next Sprint

1. **Add API contract testing:** Use OpenAPI spec validation as an additional QA check to catch API breaking changes early.
2. **Test with real ML pipeline in integration tests:** Create a mock/stub for the Python script that returns deterministic results, so integration tests can exercise the full flow including result creation.
