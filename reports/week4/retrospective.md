# Sprint Retrospective — Sprint 2

## What Went Well

1. **Smooth ML integration** — Replaced the mock Go processor with a real Python MediaPipe pipeline without breaking existing API contracts; the Go server now delegates to `scripts/process_video.py` via `exec.Command` and parses JSON output seamlessly
2. **Zero downtime migration** — The switch from mock to real ML was backwards-compatible; all 47 existing tests continued to pass after the integration
3. **Automatic environment detection** — The ML script auto-detects the Python virtualenv, removing manual setup friction for contributors
4. **Clean separation of concerns** — Go handles HTTP/routing, Python handles ML; each team member could work in their stack without conflicts

## What Did Not Go Well

1. **Limited test coverage for ML path** — The Python script lacks unit tests; we rely on manual verification of MediaPipe output correctness
2. **No performance benchmarking** — We did not measure processing time for real video analysis vs. the mock pipeline; potential latency regressions are undetected
3. **Documentation lag** — README and CHANGELOG were not updated in the same PR as the ML integration, requiring a follow-up commit

## What Changed Compared to Previous Sprint

Based on the Sprint 1 retrospective action points:

1. **ML integration prioritized and delivered** — Sprint 1 flagged "No real ML integration" as the top gap; Sprint 2 addressed it by connecting Python MediaPipe via `exec.Command`
2. **Improved estimation discipline** — Story point estimates were more conservative this sprint, based on Sprint 1 velocity data; no estimation-related surprises occurred

## Process Improvements for Next Sprint

1. **Add integration tests for the ML script** — Write at least 3–5 tests that invoke `process_video.py` with a sample video and assert on output JSON structure, ensuring the ML path is covered in CI
2. **Bundle documentation updates with feature PRs** — Enforce in the PR template that CHANGELOG and README changes ship in the same PR as code changes, preventing documentation debt
