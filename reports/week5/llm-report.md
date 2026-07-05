# LLM Usage Report — Week 5

This document records which AI tools were used during Sprint 3 (Assignment 5), for what purpose, and how the results were verified by the team.

## Tools Used

### Claude (Anthropic)

**Purpose:** Assistance with setting up the hosted documentation site (Assignment 5, architecture documentation hosting requirement).

**Specific tasks:**
- Reviewing the existing `docs/architecture/` structure against the assignment and artifact requirements.
- Identifying issues in the architecture documentation: unresolved git merge conflicts in `docs/architecture/README.md`, duplicate/stale PlantUML source files located outside the required `static-view/`, `dynamic-view/`, `deployment-view/` directories, and a syntax typo in the component diagram source.
- Rendering the PlantUML diagrams (component, sequence, deployment) to SVG images and embedding them in `docs/architecture/README.md` so they display as images rather than as links to raw diagram source.
- Setting up `mkdocs.yml`, `docs/index.md`, and a GitHub Actions workflow (`.github/workflows/docs.yml`) to automatically build and deploy the documentation site to GitHub Pages on every push to `main`.
- Adding links to the hosted documentation site from the root `README.md` and the Week 5 report.

### Qwen

**Purpose:** Support during User Acceptance Testing (UAT).

**Specific tasks:**
- Assistance during preparation and execution of UAT scenarios for the Sprint 3 increment (MVP v2).

## Verification

All AI-assisted output was reviewed and verified by the team before being accepted into the project:

- Every file and configuration change proposed by Claude was reviewed by the team, checked against the assignment requirements, and validated by building the documentation site locally and inspecting the rendered diagrams and pages before merging.
- The final hosted documentation site was manually checked by the team after deployment (navigation, architecture diagrams, and diagram-as-code rendering) to confirm it matches the maintained documentation in the repository.
- UAT scenarios and results supported by Qwen were reviewed and executed with the customer during the Sprint Review meeting; the team confirmed the outcomes directly with the customer rather than relying solely on AI output.

No AI-generated content (code, configuration, or documentation) was merged into the repository without team review.
