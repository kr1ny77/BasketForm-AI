# Quality Requirement Tests

## Automated Quality Requirement Tests

| QR ID | Test Description | Status | Test Location |
|-------|-----------------|--------|---------------|
| QR-001 | ML output contains required fields (score, feedback, pose) | Pending | To be added |
| QR-002 | Non-basketball video returns rejection error | Pending | To be added |
| QR-003 | Results survive server restart (file-based persistence) | Pending | To be added |
| QR-005 | API endpoints return correct status codes | Covered | `*_test.go` (47 existing tests) |

## Test Execution

- **Runner:** `go test -coverprofile=coverage.out ./...`
- **CI:** GitHub Actions (`ci.yml`) on push/PR to `main`
- **Local:** `go test -race ./...`
