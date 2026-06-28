# Quality Requirement Tests

Automated tests that verify quality requirements for BasketForm-AI.

## QRT-001: API Response Time

**Linked quality requirement:** QR-001

**Verification method:** Automated performance test.

**Test data, setup, or environment:** Standard CI build environment. Test server started locally with `go test`.

**Automated command or CI check:** `go test -tags=qrt -run TestQRT_APIResponseTime ./internal/qrt/...`

**Expected measurable result:** API endpoints (`/api/videos`, `/api/result/{id}`) respond within 2 seconds under normal load.

**Evidence link:** CI run showing QRT test pass.

## QRT-002: Authentication Security

**Linked quality requirement:** QR-002

**Verification method:** Automated security test.

**Test data, setup, or environment:** Standard CI build environment. Test server with middleware applied.

**Automated command or CI check:** `go test -tags=qrt -run TestQRT_AuthSecurity ./internal/qrt/...`

**Expected measurable result:** Protected endpoints return 302 redirect to `/login` for unauthenticated users. API endpoints return 401. Cross-user access is rejected.

**Evidence link:** CI run showing QRT test pass.

## QRT-003: Critical Module Unit Coverage

**Linked quality requirement:** QR-003

**Verification method:** Automated coverage check.

**Test data, setup, or environment:** Standard CI build environment.

**Automated command or CI check:** `go test -coverprofile=coverage.out ./internal/services/... ./internal/handlers/...` followed by coverage threshold check (30%).

**Expected measurable result:** Each critical module (`internal/services`, `internal/handlers`) achieves at least 30% line coverage.

**Evidence link:** CI run with coverage report.

## QRT-004: Usability — Form Guidance

**Linked quality requirement:** QR-004

**Verification method:** Automated HTML inspection test.

**Test data, setup, or environment:** Standard CI build environment. Template files checked via test.

**Automated command or CI check:** `go test -tags=qrt -run TestQRT_FormGuidance ./internal/qrt/...`

**Expected measurable result:** All form pages (login.html, signup.html, profile.html, upload.html) contain `<input>` elements with `placeholder` or associated `<label>` attributes.

**Evidence link:** CI run showing QRT test pass.
