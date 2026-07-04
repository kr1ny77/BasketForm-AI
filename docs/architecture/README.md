# BasketForm-AI Architecture

This document describes the architecture of BasketForm-AI — a web application for analyzing basketball shooting form from video. It serves as the canonical maintained architecture index artifact.

## Overview

BasketForm-AI is a monolithic web application with three main layers:

1. **Frontend** — HTML/CSS/JS pages served as static files
2. **Backend** — Go HTTP server handling API, auth, and orchestration
3. **ML Pipeline** — Python scripts for video analysis (MediaPipe, OpenCV, YOLO, LLM)

## Architecture Views

| View | Description | Source File |
|------|-------------|-------------|
| Static (Component) | System components and their relationships | [component-diagram.puml](static-view/component-diagram.puml) |
| Dynamic (Sequence) | Video upload -> analysis -> result flow | [sequence-diagram.puml](dynamic-view/sequence-diagram.puml) |
| Deployment | How the system is deployed on VM | [deployment-diagram.puml](deployment-view/deployment-diagram.puml) |

---

## Static View — Component Diagram

The component diagram shows the main internal components of the system, their relationships, external systems, and communication protocols.

### What the Diagram Shows

The system has four major component groups:

1. **User (Browser)** — the external actor interacting with the application via HTTP.
2. **Frontend** — five HTML pages (Upload, Results, Profile, Friends, Shared) served as static files. All API calls use fetch() with JWT cookie authentication.
3. **Go Server** — the central orchestration layer with Auth Middleware, Handlers (api, auth, friends, share), Services (storage, auth, processor), and Models.
4. **ML Pipeline** — five Python modules: main.py (coordinator), ball_tracker.py (YOLO), pose_tracker.py (MediaPipe), shot_analyzer.py (state machine), feedback_generator.py (LLM via OpenRouter).

External dependencies: JSON Files (data/), File System (uploads/, results/), OpenRouter API (GPT-4o-mini).

### Coupling and Cohesion

- **Low coupling** between frontend and backend — separated by HTTP API boundary.
- **Moderate coupling** between Go server and ML pipeline — connected via exec.Command and file I/O.
- **Low coupling** between ML modules — each has a single responsibility.
- **High cohesion** within each handler (auth, friends, share each handle one domain).
- **High cohesion** within ML modules (each does one thing).

### Maintainability Implications

- JSON file storage is simple but limits scalability beyond ~10 concurrent users.
- exec-based ML integration means Python changes require no Go recompilation, but prevent hot-reloading.
- Frontend uses vanilla JavaScript — no build step, easy to modify.

### Quality Requirements Supported or Constrained

| QR | Status | Explanation |
|----|--------|-------------|
| QR-001 (API < 2s) | Supported | Go server is fast; ML runs async in goroutines |
| QR-002 (Auth security) | Supported | JWT + bcrypt; middleware enforces auth |
| QR-003 (UI hints) | Supported | i18n system with tooltips on all pages |
| QR-004 (ML accuracy) | Partially supported | Depends on pose tracker + LLM vision quality |
| QR-005 (Concurrency) | Constrained | JSON file storage + mutex limits to ~10 users |

---

## Dynamic View — Sequence Diagram

The sequence diagram shows the most important workflow: video upload, ML analysis, and result retrieval.

### What the Diagram Shows

Three phases:
1. **Upload Phase**: User selects video -> Browser sends POST to Go server -> Server validates JWT, saves video and metadata -> Spawns goroutine for async processing -> Returns 201.
2. **Processing Phase** (async): Go server calls Python ML via exec.Command -> ML runs ball detection (YOLO) and pose estimation (MediaPipe) -> Shot phase state machine tracks DIP/ASCENT/RELEASE/FOLLOW-THROUGH -> ML sends 8 key frames + metrics to OpenRouter LLM -> LLM returns scores and feedback -> Results saved as JSON.
3. **Result Retrieval**: Browser polls GET /api/videos -> User clicks result -> GET /api/result/{id} -> Server reads result JSON -> Returns score, phases, feedback -> Browser renders.

### Why This Scenario Is Important

This is the primary value-creation pipeline. Every user interaction depends on this sequence working correctly.

### Architecture Decisions and Quality Requirements

- **ADR-001 (JSON Storage)**: Results stored as JSON, read on demand.
- **ADR-002 (exec ML Integration)**: Async goroutine prevents blocking HTTP server.
- **ADR-003 (JWT Auth)**: Every API call passes through JWT validation.
- **QR-001**: Upload returns 201 quickly while processing runs in background.
- **QR-004**: ML accuracy depends on pose tracker + LLM vision quality.
- **QR-005**: Concurrent uploads handled by separate goroutines.

---

## Deployment View — Deployment Diagram

The deployment diagram shows how the system is deployed on a single VM.

### What the Diagram Shows

1. **Nginx** (port 80) — reverse proxy forwarding to Go server.
2. **Go Server** (port 8080) — application binary, handles all API requests and ML orchestration.
3. **Python ML Pipeline** — invoked as subprocess by Go server.
4. **File System** — data/, uploads/, results/, web/ for all persistent state.
5. **OpenRouter API** — external LLM service for vision scoring.

User access: Browser -> HTTPS (port 80) -> Nginx -> Go Server (port 8080).

### Why This Deployment Model Was Chosen

- **Simplicity**: One server, one deployment process, one set of logs.
- **Cost**: Single VM is cheapest for MVP.
- **Speed**: No container overhead; direct file access for ML processing.

### Deployment Considerations

- OPENROUTER_API_KEY must be set for LLM scoring.
- ML model weights (best.pt) must be present.
- Python venv with MediaPipe, OpenCV must be active.
- Go binary must be rebuilt after Go code changes.

---

## Architecture Decision Records

| ADR | Decision | Status |
|-----|----------|--------|
| [ADR-001](adr/ADR-001-use-json-storage.md) | JSON file storage instead of relational database | Accepted |
| [ADR-002](adr/ADR-002-exec-ml-script.md) | exec-based ML integration instead of microservice | Accepted |
| [ADR-003](adr/ADR-003-jwt-auth.md) | JWT authentication instead of server sessions | Accepted |

These ADRs are linked to quality requirements in [docs/quality-requirements.md](../quality-requirements.md).
