# Quality Requirements

## Quality Model

Based on ISO/IEC 25010, the following sub-characteristics are selected for BasketForm-AI:

| Sub-characteristic | Justification |
|--------------------|---------------|
| **Functional Suitability** | The product must correctly detect basketball shots and provide accurate feedback |
| **Performance Efficiency** | Video processing must complete within acceptable time on CPU |
| **Usability** | Upload flow, results display, and PDF export must be intuitive |
| **Reliability** | ML pipeline must handle invalid input gracefully without crashing |
| **Maintainability** | Go + Python separation must allow independent modification of each layer |

## Quality Requirements

| ID | Requirement | Related PBI | Priority |
|----|-------------|-------------|----------|
| QR-001 | ML pipeline outputs structured JSON with score, feedback, and pose fields | PBI-016 | Must Have |
| QR-002 | Non-basketball videos are rejected with clear error message | PBI-017 | Must Have |
| QR-003 | Analysis results persist across server restarts | PBI-018 | Must Have |
| QR-004 | ML processing completes within 30 seconds for a 10-second video on CPU | PBI-016 | Should Have |
| QR-005 | All API endpoints return correct status codes and JSON structure | PBI-016 | Must Have |
