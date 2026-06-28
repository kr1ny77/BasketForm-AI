# Testing Strategy

Testing strategy for BasketForm-AI.

## Critical Modules and Coverage

| Critical module | Why critical | Required line coverage | Current line coverage | Evidence |
|---|---|---|---|---|
| `internal/services/auth.go` | Authentication and password hashing — security-critical. | 30% | ~70% | CI run |
| `internal/services/storage.go` | Data persistence for users, videos, friends, shared results. | 30% | ~60% | CI run |
| `internal/services/processor.go` | ML integration and video processing pipeline. | 30% | ~40% | CI run |
| `internal/handlers/api.go` | Core API endpoints: upload, status, result, videos. | 30% | ~55% | CI run |
| `internal/handlers/auth.go` | Registration, login, profile — auth-critical. | 30% | ~50% | CI run |
| `internal/handlers/friends.go` | Social features — friend requests and search. | 30% | ~45% | CI run |

## Automated Test Status

| Test type | Scope | Command or CI check | Latest result | Evidence |
|---|---|---|---|---|
| Unit tests | Critical product logic | `go test ./...` | Passing | CI run |
| Integration tests | API + storage interaction | `go test -tags=integration ./...` | Passing | CI run |
| Automated QRTs | QR-001, QR-002, QR-003, QR-004 | `go test -tags=qrt ./internal/qrt/...` | Passing | QRT report |

## CI and QA Check Status

| Gate or check | Required for Done? | Latest protected-branch status | Evidence |
|---|---|---|---|
| Linting (golangci-lint) | Yes | Passing | CI run |
| Unit tests | Yes | Passing | CI run |
| Integration tests | Yes | Passing | CI run |
| Coverage check (30%) | Yes | Passing | CI run |
| QRTs | Yes | Passing | CI run |
| Additional QA (govulncheck) | Yes | Passing | CI run |
| Link checking (Lychee) | Yes | Passing | CI run |

## Additional QA Check Rationale

| QA objective or risk | Additional QA check | Scope | Latest result | Evidence | Limitations or follow-up |
|---|---|---|---|---|---|
| Known vulnerabilities in Go dependencies may expose users to security risks. | govulncheck | Product Go dependencies. | Passing | CI run | Some vulnerabilities may require upstream fixes. |

## Test Locations

- Unit tests: `internal/services/*_test.go`, `internal/handlers/*_test.go`, `internal/models/*_test.go`
- Integration tests: `internal/handlers/api_integration_test.go`
- QRTs: `internal/qrt/qrt_test.go`

## CI Pipeline

CI runs on every PR and push to `main`:
1. **Lint** — `golangci-lint run`
2. **Test** — `go test -coverprofile=coverage.out ./...`
3. **Coverage** — check 30% threshold for critical modules
4. **QRT** — `go test -tags=qrt ./internal/qrt/...`
5. **QA Extra** — `govulncheck ./...`
6. **Link Check** — Lychee on Markdown files

## Manual Evidence That Does Not Count as QRT

| Evidence | Scope | Result | Follow-up PBI or issue |
|---|---|---|---|
| Customer UAT observation | Registration, video upload, sharing | Passed | — |
