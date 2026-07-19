# Bug Triage — PBI-043 (Should-Have)

## Context

PBI-043 covered Sprint 5 stabilization fixes: data race in `GetVideo`, hardcoded ML path, and deploy script improvements. This document triages the bugs discovered during Sprint 5 and classifies them by priority.

## Discovered Issues

### BUG-001: Data Race in GetVideo — Severity: Critical → Resolved

| Field | Value |
|-------|-------|
| **Status** | Resolved (PR #160) |
| **Priority** | Should-Have (Sprint 5) |
| **Description** | `GetVideo` used `RLock` while writing to the map, causing concurrent map read/write panics |
| **Reproduction** | Upload 5+ videos simultaneously |
| **Fix** | Changed `RLock`/`RUnlock` to `Lock`/`Unlock` in `GetVideo` |
| **Regression** | Verified with `-race` flag (10 iterations, 0 races) |

### BUG-002: Hardcoded ML Path — Severity: High → Resolved

| Field | Value |
|-------|-------|
| **Status** | Resolved (PR #161) |
| **Priority** | Should-Have (Sprint 5) |
| **Description** | `processor.go` hardcoded `/home/basketfrom-ai/BasketForm-AI` as working directory for ML execution |
| **Reproduction** | Run tests from any directory other than `/home/basketfrom-ai/BasketForm-AI` |
| **Fix** | Use current working directory; removed hardcoded log initialization |
| **Regression** | All tests pass; production deploy works |

### BUG-003: Deploy Script Missing ML Deps & Smoke Test — Severity: Medium → Resolved

| Field | Value |
|-------|-------|
| **Status** | Resolved (PR #161) |
| **Priority** | Should-Have (Sprint 5) |
| **Description** | Deploy script did not install ML dependencies or verify deployment success |
| **Reproduction** | Deploy to fresh VM without pre-installed ML packages |
| **Fix** | Added `pip install` step, smoke test with HTTP status check, `set -euo pipefail` |
| **Regression** | Deploy from clean state succeeds; smoke test validates endpoints |

## Triage Summary

| Bug | Severity | Priority | Status | Sprint |
|-----|----------|----------|--------|--------|
| BUG-001: Data race in GetVideo | Critical | Should-Have | Resolved | Sprint 5 |
| BUG-002: Hardcoded ML path | High | Should-Have | Resolved | Sprint 5 |
| BUG-003: Deploy script gaps | Medium | Should-Have | Resolved | Sprint 5 |

## Remaining Known Limitations (Not Bugs)

These are documented constraints, not defects:

- JSON file storage limits concurrent users to ~10 (documented in README)
- ML processing runs synchronously per request (10–30 seconds)
- No horizontal scaling without refactoring

## Conclusion

All Sprint 5 Should-Have bugs have been resolved and verified. No outstanding defects remain for PBI-043.
