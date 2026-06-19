# Sprint Retrospective

## What Went Well

1. **Clean tech stack pivot** — Switching from FastAPI/React to Go standard library simplified deployment, eliminated frontend framework complexity, and reduced dependency count to zero
2. **End-to-end flow working** — From video upload through mock processing to results display and PDF export, the full user journey is functional and testable
3. **Test coverage from day one** — Writing 47 unit + integration tests alongside implementation caught bugs early and provided confidence for refactoring

## What Did Not Go Well

1. **No real ML integration** — The mock pipeline demonstrates the flow but doesn't provide actual analysis; real MediaPipe integration was deferred
2. **In-memory storage** — Data loss on server restart; no persistence layer means the product isn't production-ready for real users
3. **Estimation was off** — Initial story point estimates were optimistic; the Go pivot took longer than expected due to manual URL routing and template debugging

## Action Points

1. **Sprint 2: Prioritize ML integration** — Connect Python script via exec.Command to replace mock processor with real analysis
2. **Sprint 2: Add persistence** — Migrate from in-memory map to JSON file-based state that survives restarts, or add SQLite

## Metrics

- **Sprint velocity:** 76 Story Points completed
- **Tests:** 47 passing (unit + integration)
- **Test coverage:** ~70% line coverage
- **Bugs found:** 3 (all fixed before merge)
