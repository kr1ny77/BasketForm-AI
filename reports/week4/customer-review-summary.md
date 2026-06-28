# Customer Review Summary — Week 4

**Date:** 2026-06-28
**Participants:** Customer (Anastasia Malakhova), Development team

## Sprint Goal Reviewed

Sprint 2 — Assignment 4: Add authentication, social features, enhanced ML with phase analysis, and automated quality gates.

## Delivered Increment

The team demonstrated:
1. **User registration and login** with email, nickname, and password
2. **Friend system** — search users, send/accept friend requests
3. **Result sharing** — share analysis results with friends (read-only)
4. **Enhanced ML analysis** — annotated output video with pose overlay, phase scores (Stance, Arm Angle, Release, Follow-through), and personalized feedback
5. **PDF export** — downloadable report with full score breakdown
6. **Quality automation** — CI pipeline with lint, test, coverage, QRT, and security scanning

## UAT Results

- UAT-001 (Registration and Login): **Passed** — user can register and log in successfully
- UAT-002 (Video Upload and Analysis): **Passed** — video uploads, processes, and shows phases with annotated video
- UAT-003 (Share Result): **Passed** — result is shared and friend can view it
- UAT-004 (Friend Request): **Passed** — friend request workflow works end-to-end

## Quality Evidence Discussed

- Quality requirements (QR-001 to QR-004) following ISO/IEC 25010
- Automated QRT tests passing in CI
- Coverage ≥30% for critical modules
- Security scanning (govulncheck) passing

## Feedback

- Customer approved the authentication and social features
- Customer appreciated the phase analysis detail in the ML output
- Customer noted that the annotated video is a strong visual feature for coaches
- Customer requested batch upload for multiple shots (deferred to later Sprint)

## Approvals

- All delivered features approved by customer
- No blocking issues identified

## Risks

- ML pipeline requires Python dependencies on the server — deployment complexity
- No rate limiting on API endpoints — recommended to add reverse proxy

## Action Points

- Continue with Sprint 3 planning (comparison with professionals, progress tracking)
- Address deployment complexity by documenting Python dependency installation
