# BasketForm-AI Architecture

This document describes the architecture of BasketForm-AI — a web application for analyzing basketball shooting form from video.

## Overview

BasketForm-AI is a monolithic web application with three main layers:

1. **Frontend** — HTML/CSS/JS pages served as static files
2. **Backend** — Go HTTP server handling API, auth, and orchestration
3. **ML Pipeline** — Python scripts for video analysis (MediaPipe, OpenCV, YOLO)

## Architecture Views

| View | Description | File |
|------|-------------|------|
| Static (Component) | System components and their relationships | [static-view/component-diagram.puml](static-view/component-diagram.puml) |
| Dynamic (Sequence) | Video upload → analysis → result flow | [dynamic-view/sequence-diagram.puml](dynamic-view/sequence-diagram.puml) |
| Deployment | How the system is deployed on VM | [deployment-view/deployment-diagram.puml](deployment-view/deployment-diagram.puml) |

## Component Overview

### Go Server (`cmd/server/main.go`)
- HTTP server on port 8080
- JWT authentication middleware
- REST API for users, videos, friends, sharing
- Orchestrates ML processing via `exec.Command`

### ML Pipeline (`ML/`)
- `main.py` — entry point, coordinates analysis
- `ball_tracker.py` — YOLO-based ball detection
- `pose_tracker.py` — MediaPipe pose estimation
- `shot_analyzer.py` — state machine for shot phases
- `feedback_generator.py` — LLM-based scoring via OpenRouter API
- `custom_feedback.py` — rule-based fallback scoring

### Frontend (`web/`)
- `templates/` — HTML pages (upload, results, profile, friends, shared)
- `static/css/` — stylesheets
- `static/js/` — i18n, canvas animations, API calls

### Data Layer (`data/`)
- `users/` — user JSON files
- `videos/` — video metadata JSON files
- `friends/` — friend request JSON files
- `shared/` — shared result JSON files
- `avatars/` — user avatar images

## Architectural Quality

### Coupling
- **Low coupling** between frontend and backend (HTTP API boundary)
- **Moderate coupling** between Go server and ML pipeline (exec call + file I/O)
- **Low coupling** between ML components (each module has clear responsibility)

### Cohesion
- **High cohesion** within each handler (auth, friends, share each handle one domain)
- **High cohesion** within ML modules (each does one thing well)

### Maintainability
- JSON file storage is simple but limits scalability
- exec-based ML integration is easy to understand but prevents hot-reloading
- Frontend uses vanilla JS — no build step, easy to modify

## Quality Requirements Impact

| QR | Supported/Limited | Explanation |
|----|-------------------|-------------|
| QR-001 (API < 2s) | Supported | Go server is fast; ML runs async in goroutine |
| QR-002 (Auth security) | Supported | JWT + bcrypt password hashing |
| QR-003 (UI hints) | Supported | i18n system with tooltips on all pages |
| QR-004 (ML accuracy) | Limited | Depends on pose tracker quality; LLM adds intelligence |
| QR-005 (Concurrent users) | Limited | JSON file storage + mutex; fine for <10 users |

## Architecture Decision Records

See [adr/](adr/) directory for key architectural decisions:

- [ADR-001: JSON File Storage](adr/ADR-001-use-json-storage.md)
- [ADR-002: exec-based ML Integration](adr/ADR-002-exec-ml-script.md)
- [ADR-003: JWT Authentication](adr/ADR-003-jwt-auth.md)
