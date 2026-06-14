# Assignment 2: BasketForm-AI

## Project Information

- **Project Name:** BasketForm-AI
- **Description:** Website with a backend service for recording and uploading basketball shot videos, extracting biomechanical keypoints, evaluating shooting technique, and generating personalized feedback.
- **License:** [MIT License](../../LICENSE)
- **Repository:** [https://github.com/kr1ny77/BasketForm-AI](https://github.com/kr1ny77/BasketForm-AI)

## Required Artifacts

### User Stories
- [User Stories and MoSCoW Priorities](user-stories.md)

### Interface Design
- **Interactive Prototype:** [Figma Prototype](https://silk-sadly-37202889.figma.site)
- **Interface Documentation:** [docs/interface.md](../../docs/interface.md)
- **OpenAPI Specification:** [api/openapi.yaml](../../api/openapi.yaml)
- **Postman Collection:** [api/postman_collection.json](../../api/postman_collection.json)

### MVP v0
- [MVP v0 Report](mvp-v0-report.md)
- **Deployment URL:** [http://80.74.30.14/](http://80.74.30.14/)
- **Video Demonstration:** [Video Link](https://youtu.be/To_1ZwAMe4M)

### Customer Review
- [Customer Meeting Summary](customer-meeting-summary.md)
- [Customer Meeting Transcript](customer-meeting-transcript.md)

### Analysis and Reports
- [Week 2 Analysis](analysis.md)
- [LLM Usage Report](llm-report.md)

## Repository Setup

### Branch Protection
- Default branch (`main`) is protected
- Direct pushes disabled
- Required approvals: 1
- PR template: [`.github/pull_request_template.md`](../../.github/pull_request_template.md)

### Lychee Link Checking
- Configuration: [`.lycheeignore`](../../.lycheeignore)
- CI Workflow: [`.github/workflows/lychee.yml`](../../.github/workflows/lychee.yml)
- Latest successful run: [Link to CI run](https://github.com/kr1ny77/BasketForm-AI/actions)

### Excluded Links
- `http://80.74.30.14` — Local/dev server IP, not publicly routable domain; verified accessible in browser
- `mvp-v0.example.com` — Placeholder URL, not yet replaced with actual deployment
- `mvp-v0.basketformai.com` — Custom domain not yet configured
- `staging.mvp-v0.basketformai.com` — Custom domain not yet configured
- ~~`https://youtu.be/PLACEHOLDER_VIDEO_ID`~~ — Video recorded and linked

All excluded links have been manually verified in a browser where applicable.

## Screenshots

> Screenshots to be added before final submission:
> - Protected branch settings
> - Example reviewed PR
> - Figma prototype
> - MVP v0 deployment

## Coverage

### Prototype Coverage
The interactive prototype covers the following user stories:
- US-01: Video upload and recording
- US-02: Biomechanical keypoint extraction
- US-03: Shooting technique evaluation
- US-04: Personalized feedback generation

### MVP v0 Coverage
MVP v0 establishes the technical foundation for:
- US-01: Video upload and recording (partial implementation)
- US-02: Biomechanical keypoint extraction (backend service)
- US-03: Shooting technique evaluation (API endpoint)
- US-04: Personalized feedback generation (placeholder)

For detailed smoke-check scenario, see [MVP v0 Report](mvp-v0-report.md).

## Team Contributions

| Team Member | Role | Responsibilities | Contributions |
|-------------|------|------------------|---------------|
| Karim Gimadiev | Product Owner | Product management | User stories, customer review, sprint planning |
| Roman Santalov | Scrum Master | Process management, ML Engineering | Documentation, ML research, customer meeting |
| Kamil Nizamov | Developer | ML Engineering | Backend development, MVP v0 deployment |
| Arseniy Fedorov | Developer | ML Analysis | Data analysis, prototype review |
| Damir Galiev | Developer | ML Analysis | Research, documentation support |

## Links

- [Root README](../../README.md)
