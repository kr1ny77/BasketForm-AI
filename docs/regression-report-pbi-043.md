# Regression Report — PBI-043

## Summary

PBI-043 addressed two production issues discovered during Sprint 5 stabilization: a data race in `GetVideo` and a hardcoded deployment path in the ML processor. This report documents the regression analysis performed after the fixes were merged.

## Issues Fixed

### 1. Data Race in GetVideo (PR #160)

**Root cause:** `GetVideo` held an `RLock` while writing to the videos map, causing concurrent map read/write panics under load.

**Fix:** Changed to `Lock` for the duration of the operation. Verified with `-race` flag.

**Regression check:**
- `go test -race -count=10 ./internal/services/...` — **PASS** (0 races detected across 10 iterations)
- `go test -race -count=10 ./internal/handlers/...` — **PASS** (0 races detected across 10 iterations)
- Manual concurrent upload test (5 parallel uploads) — **PASS** (no panics, all results returned)

### 2. Hardcoded ML Path (PR #161)

**Root cause:** `processor.go` used hardcoded `/home/basketfrom-ai/BasketForm-AI` as the working directory for `exec.Command`, causing test failures locally and portability issues.

**Fix:** ML runner now uses the current working directory. Removed hardcoded log initialization.

**Regression check:**
- `go test ./internal/services/...` — **PASS**
- Local build and run on different directory — **PASS** (ML pipeline found correctly)
- Production deploy via `scripts/deploy.sh` — **PASS** (smoke test HTTP 200)

### 3. Deploy Script Rewrite (PR #161)

**Changes:** Added ML dependency installation step, smoke test with HTTP status check, proper error handling with `set -euo pipefail`.

**Regression check:**
- Deploy from clean state — **PASS**
- Deploy with `SKIP_ML_DEPS=1` — **PASS**
- Smoke test validates both `/login` and `/api/videos` endpoints — **PASS**

## Test Results

| Test | Result | Notes |
|------|--------|-------|
| Unit tests (full suite) | PASS | 0 failures |
| Integration tests | PASS | 0 failures |
| `-race` detector (10 iterations) | PASS | 0 data races |
| QRT tests | PASS | All quality requirement tests pass |
| Manual concurrent upload | PASS | No panics under load |
| Production smoke test | PASS | HTTP 200 on /login |

## Conclusion

All three fixes have been verified with automated tests and manual checks. No regressions detected. The fixes improve portability, test reliability, and deployment robustness without changing user-facing behavior.
