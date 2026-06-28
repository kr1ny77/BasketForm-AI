# Week 4 Reflection

## Learning Points

### Responding to Customer Feedback
- Customer feedback from Sprint 1 Review emphasized the need for real analysis over mock data; this directly shaped the Sprint 2 goal of integrating MediaPipe
- Learned to translate qualitative customer feedback ("I want real analysis") into a concrete technical PBI with clear acceptance criteria

### Defining Quality Requirements
- Established that output JSON from the ML script must include `score`, `feedback`, and `pose` fields to maintain API compatibility
- Defined "quality" for the ML pipeline as: deterministic structure, non-empty results, and graceful failure on invalid input

### Automating Quality Requirement Tests
- The existing 47 Go tests continue to validate API contract compliance after ML integration
- Identified a gap: no automated tests exercise the Python ML script directly; this is flagged as a next-sprint priority

### Configuring CI
- CI pipeline (golangci-lint + go test + go build) remained stable throughout Sprint 2
- No CI changes were needed since the ML script runs as a subprocess, not as part of the Go test suite

### Running UAT
- UAT was performed by running the full flow: upload video → processing → results display → PDF export
- Verified that real MediaPipe output integrates correctly with the existing Chart.js radar and scatter visualizations

### Releasing the Sprint Increment
- The ML integration was committed on the `feat/replace-mock-nn-with-python-ml` branch
- Release is pending merge to main; the increment is functional and testable locally

### Reviewing the Increment with the Customer
- Sprint 2 Review demonstrated real biomechanical analysis to the customer
- Customer confirmed that pose visualization and feedback quality meet expectations

## Validated Assumptions

- **Python subprocess execution works reliably** — `exec.Command` with JSON parsing is a viable integration pattern between Go and Python
- **MediaPipe runs without GPU** — The pipeline executes on CPU within acceptable time for MVP purposes
- **API contract stability** — Changing the backend from mock to real ML required zero frontend changes, confirming clean interface design
- **Virtualenv auto-detection** — The script correctly finds and activates the Python environment without manual configuration

## Friction and Gaps

- **No ML unit tests** — The Python script has no automated tests; regressions in pose detection or scoring logic would only be caught manually
- **No processing time benchmarks** — Unknown whether real ML processing meets performance expectations for production use
- **Error handling in ML script** — The script assumes valid input; malformed videos could cause unhandled exceptions that crash the Go subprocess
- **Missing documentation updates** — README and CHANGELOG were not updated in the ML integration PR
- **No logging in ML subprocess** — Debugging ML failures requires manual testing; no structured logs are captured from the Python process
- **Uncertain MediaPipe version stability** — MediaPipe API changes could break the pipeline without warning

## Planned Response

### Next Sprint (Sprint 3)

1. **Add Python ML tests** — Create `tests/test_process_video.py` with 3–5 integration tests covering valid input, edge cases, and error paths → affects the ML processing PBI
2. **Add subprocess error handling** — Implement try/except in `process_video.py` and return structured error JSON to Go → affects API reliability
3. **Benchmark processing time** — Measure and document MediaPipe processing latency for different video lengths → affects performance baseline
4. **Update documentation** — Refresh README setup instructions for the ML pipeline and update CHANGELOG with Sprint 2 changes → affects docs/roadmap.md and CHANGELOG.md
5. **Proceed with Sprint 3 scope** — Begin comparison features (US-005), sharing (US-006, US-007), and progress tracking (US-008) as planned in the roadmap

### Links to Affected Artifacts

- [Sprint 2 Milestone](https://github.com/kr1ny77/BasketForm-AI/milestone/1)
- [ML Integration PR](https://github.com/kr1ny77/BasketForm-AI/pull) — `feat/replace-mock-nn-with-python-ml`
- [Roadmap](../../docs/roadmap.md)
- [User Stories](../../docs/user-stories.md)
- [CHANGELOG](../../CHANGELOG.md)
