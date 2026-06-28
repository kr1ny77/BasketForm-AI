# Testing Strategy

## Test Types

| Type | Framework | Coverage | Location |
|------|-----------|----------|----------|
| Unit tests | `go test` | UUID generation, file validation, storage, processor, models | `*_test.go` |
| Integration tests | `httptest` | All API endpoints (upload, status, result, videos) | `*_test.go` |
| Manual tests | Checklist | 50+ scenarios | `reports/week3/manual-test-checklist.md` |
| ML integration tests | Pending | Python script output validation | To be added |

## Current Status

- **Total tests:** 47 (passing)
- **Coverage:** ~70% line coverage
- **CI:** GitHub Actions runs `go test -coverprofile=coverage.out ./...` on every push/PR to `main`
- **Race detection:** `go test -race ./...` available locally

## Gaps

- No automated tests for Python ML script (`scripts/process_video.py`)
- No performance/load tests
- No browser-based E2E tests
