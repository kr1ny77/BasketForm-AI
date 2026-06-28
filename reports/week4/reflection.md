# Week 4 Reflection

## Learning Points

- **Customer feedback** — Customer review from Sprint 1 emphasized the need for real analysis over mock data; this shaped the Sprint 2 goal of integrating MediaPipe
- **Quality requirements** — Established that ML output JSON must include `score`, `feedback`, and `pose` fields to maintain API compatibility with the frontend
- **Automated tests** — Existing 47 Go tests validate API contract compliance after ML integration; Python ML script has no automated tests yet
- **CI** — CI pipeline (golangci-lint + go test + go build) remained stable; no changes needed since ML runs as a subprocess
- **UAT** — Tested full flow: upload video → processing → results display → PDF export with real MediaPipe output
- **Release** — ML integration committed on `feat/replace-mock-nn-with-python-ml` branch; pending merge to main
- **Sprint Review** — Demonstrated real biomechanical analysis to the customer; feedback was positive

## Validated Assumptions

- **Python subprocess execution works** — `exec.Command` with JSON parsing is a viable integration pattern between Go and Python
- **MediaPipe runs without GPU** — Pipeline executes on CPU within acceptable time for MVP
- **API contract stability** — Changing backend from mock to real ML required zero frontend changes
- **Virtualenv auto-detection** — Script correctly finds the Python environment without manual configuration

## Friction and Gaps

- **No ML unit tests** — Python script has no automated tests; regressions caught only manually
- **No processing benchmarks** — Unknown whether real ML processing meets performance expectations
- **Error handling gaps** — ML script assumes valid input; malformed videos could crash the subprocess
- **Documentation lag** — README and CHANGELOG not updated in ML integration PR
- **No structured logging in ML subprocess** — Debugging requires manual testing

## Planned Response

1. **Add Python ML tests** — Create tests for `scripts/process_video.py` covering valid input, edge cases, and errors
2. **Add error handling** — Implement try/except in ML script with structured error JSON output
3. **Update documentation** — Refresh README and CHANGELOG for ML pipeline
4. **Proceed with Sprint 3 scope** — Begin comparison features ([US-005](https://github.com/kr1ny77/BasketForm-AI/issues/24)), sharing ([US-006](https://github.com/kr1ny77/BasketForm-AI/issues/25), [US-007](https://github.com/kr1ny77/BasketForm-AI/issues/26)), and progress tracking ([US-008](https://github.com/kr1ny77/BasketForm-AI/issues/27))
