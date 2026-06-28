# User Acceptance Tests

## UAT Scenarios

| ID | Scenario | Result | Notes |
|----|----------|--------|-------|
| UAT-001 | Upload MP4 video and receive analysis results | Pass | Score, feedback, and pose displayed correctly |
| UAT-002 | Upload non-MOV/AVI/MP4 file is rejected | Pass | Format validation works |
| UAT-003 | PDF export contains score breakdown and feedback | Pass | jsPDF generates correct output |
| UAT-004 | Results page shows radar chart and scatter plot | Pass | Chart.js renders correctly |
| UAT-005 | Profile page displays account settings | Pass | UI demo mode functional |
| UAT-006 | Upload processing shows progress indicator | Pass | Polling updates UI correctly |
| UAT-007 | Real MediaPipe analysis produces structured output | Pass | Sprint 2: ML integration verified |
| UAT-008 | Non-basketball video is rejected by ML pipeline | Pending | Sprint 2: PBI-017 |

## Execution

- **Date:** Sprint 2 UAT performed locally by team members
- **Environment:** Local development, Chrome/Firefox/Safari
- **Video samples:** Basketball shots (valid), non-basketball videos (invalid)
