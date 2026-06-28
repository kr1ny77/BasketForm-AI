# Sprint Retrospective — Sprint 2

## What Went Well

1. **ML integration completed** — Replaced the mock Go processor with a real Python MediaPipe pipeline via `exec.Command` without breaking existing API contracts
2. **All existing tests passed** — The 47 unit and integration tests from Sprint 1 continued to pass after the ML integration

## What Did Not Go Well

1. **No tests for the ML script** — The Python script lacks automated tests; correctness is verified only manually
2. **Documentation not updated in same PR** — README and CHANGELOG updates required a separate follow-up commit

## What Changed Compared to Previous Sprint

1. **ML integration delivered** — Sprint 1 retrospective identified "No real ML integration" as the top gap; Sprint 2 addressed it
2. **Estimation improved** — More conservative estimates based on Sprint 1 velocity; no estimation surprises this sprint

## Process Improvements for Next Sprint

1. **Add tests for the ML script** — Write integration tests for `scripts/process_video.py` to cover the ML path in CI
2. **Bundle docs with feature PRs** — Enforce in PR template that CHANGELOG and README changes ship with code changes
