# BasketForm-AI Documentation

BasketForm-AI is a web application that analyzes basketball shooting form from video using computer vision (MediaPipe, YOLO) and LLM-powered feedback (OpenRouter, GPT-4o-mini).

This site is the hosted, browsable version of the maintained documentation kept in [`docs/`](https://github.com/kr1ny77/BasketForm-AI/tree/main/docs) in the product repository. The repository remains the source of truth; this site exists so reviewers can browse the same content as readable, navigable pages instead of raw repository paths.

**Live product:** [http://80.74.30.14/](http://80.74.30.14/)
**Source repository:** [github.com/kr1ny77/BasketForm-AI](https://github.com/kr1ny77/BasketForm-AI)
**Latest release:** [v0.3.0 — MVP v2](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.3.0)

## Product

- [Roadmap](roadmap.md) — current direction, current Sprint, and what comes next
- [User Stories](user-stories.md)
- [Interface Documentation](interface.md)
- [Release Notes — v0.3.0 (MVP v2)](release-notes-v0.3.0.md)

## Process

- [Development Process & Configuration Management](development-process.md) — git workflow, branching model, configuration management
- [Definition of Done](definition-of-done.md)

## Quality & Testing

- [Quality Requirements](quality-requirements.md)
- [Quality Requirement Tests](quality-requirement-tests.md)
- [Testing Strategy](testing.md)
- [User Acceptance Tests](user-acceptance-tests.md)

## Architecture

- [Architecture Overview](architecture/README.md) — static, dynamic, and deployment views
- [ADR-001: Use JSON File Storage](architecture/adr/ADR-001-use-json-storage.md)
- [ADR-002: exec-based ML Integration](architecture/adr/ADR-002-exec-ml-script.md)
- [ADR-003: JWT Authentication](architecture/adr/ADR-003-jwt-auth.md)

## Weekly Reports

Weekly Sprint reports (reviews, retrospectives, reflections) are published in the repository under [`reports/`](https://github.com/kr1ny77/BasketForm-AI/tree/main/reports) and linked from each week's public report README.
