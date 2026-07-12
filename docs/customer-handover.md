# Customer Handover

This document describes the current handover state of BasketForm-AI, including what has been transferred, what the customer needs to know, and what support remains necessary.

## Product Summary

BasketForm-AI is an AI-powered web platform for analyzing basketball shooting form from video. Users upload shooting videos, the system extracts biomechanical keypoints using MediaPipe and YOLO, evaluates stance, arm angle, release point, and follow-through, then generates personalized AI feedback using GPT-4o-mini vision analysis.

**Live product:** [http://80.74.30.14/](http://80.74.30.14/)
**Source repository:** [github.com/kr1ny77/BasketForm-AI](https://github.com/kr1ny77/BasketForm-AI)
**Latest release:** [v0.3.0 — MVP v2](https://github.com/kr1ny77/BasketForm-AI/releases/tag/v0.3.0)

## What Has Been Transferred

| Item | Status | Details |
|------|--------|---------|
| Source code | Transferred | Full repository access at github.com/kr1ny77/BasketForm-AI |
| Live deployment | Transferred | Running at http://80.74.30.14/ on team-managed VM |
| Documentation | Transferred | Complete docs/ directory, hosted documentation site, README |
| User accounts | Transferred | Customer account created and tested during Week 6 trial |

## What Has Been Delegated

| Item | Status | Details |
|------|--------|---------|
| VM hosting | Delegated | VM at 80.74.30.14 remains with team for course duration |
| DNS / domain | Not applicable | Using IP address directly |
| SSL / HTTPS | Not configured | HTTP only — may need configuration for production use |

## What Has Been Retained by the Team

| Item | Reason |
|------|--------|
| GitHub repository admin access | Required for course grading and submission |
| VM root access | Required for deployment during course |
| OpenRouter API key | External service credential — should not be shared |

## Environment Variables and Configuration

The following environment variables are required to run the application:

| Variable | Required | Description |
|----------|----------|-------------|
| `OPENROUTER_API_KEY` | Yes | API key for LLM vision scoring (GPT-4o-mini). Obtain from [OpenRouter](https://openrouter.ai/) |
| `JWT_SECRET` | Yes | Secret key for JWT token signing. Any random string works |
| `PORT` | No | Server port (default: 8080) |
| `UPLOAD_DIR` | No | Video upload directory (default: uploads) |
| `RESULTS_DIR` | No | Analysis results directory (default: results) |
| `DATA_DIR` | No | JSON storage directory (default: data) |

**Important:** Never commit API keys or secrets to the repository. Set them as environment variables on the deployment server.

## Setup and Deployment Steps

### Prerequisites

- Go 1.22 or later
- Python 3.12+ with MediaPipe, OpenCV, YOLO dependencies
- Nginx (for reverse proxy in production)

### Local Setup

```bash
# 1. Clone the repository
git clone https://github.com/kr1ny77/BasketForm-AI.git
cd BasketForm-AI

# 2. Install Python dependencies
pip install -r requirements.txt

# 3. Set environment variables
export OPENROUTER_API_KEY="your-api-key"
export JWT_SECRET="your-secret"

# 4. Run the server
go run ./cmd/server/

# 5. Open in browser
# http://localhost:8080
```

### Production Deployment

```bash
# 1. Build the binary
go build -o bin/server ./cmd/server/

# 2. Set environment variables
export OPENROUTER_API_KEY="your-key"
export JWT_SECRET="your-secret"

# 3. Run
PORT=8080 ./bin/server
```

Nginx reverse proxy: port 80 → 8080.

### Docker Deployment

```bash
docker build -t basketform-ai .
docker run -p 8080:8080 \
  -e OPENROUTER_API_KEY=your-key \
  -e JWT_SECRET=your-secret \
  basketform-ai
```

## Recovery and Verification Steps

If the service goes down:

1. Check if the Go server process is running: `ps aux | grep server`
2. Check Nginx status: `sudo systemctl status nginx`
3. Restart the server: `PORT=8080 ./bin/server &`
4. Verify the service: `curl http://localhost:8080/`
5. Check logs for errors

If ML analysis fails:

1. Verify Python dependencies are installed: `pip list | grep mediapipe`
2. Check that `ML/best.pt` model weights exist
3. Verify `OPENROUTER_API_KEY` is set correctly
4. Test ML pipeline manually: `python ML/main.py <video_path>`

## Documentation Entry Points

| Document | Purpose | Link |
|----------|---------|------|
| README.md | Main repository entry point, setup instructions | [README.md](../README.md) |
| Architecture Overview | System design and diagrams | [docs/architecture/README.md](architecture/README.md) |
| Quality Requirements | ISO/IEC 25010 quality criteria | [docs/quality-requirements.md](quality-requirements.md) |
| Testing Strategy | Test approach and coverage | [docs/testing.md](testing.md) |
| User Acceptance Tests | End-user test scenarios | [docs/user-acceptance-tests.md](user-acceptance-tests.md) |
| Development Process | Git workflow, branching, CI | [docs/development-process.md](development-process.md) |
| Roadmap | Product direction and Sprint plan | [docs/roadmap.md](roadmap.md) |
| CHANGELOG | Version history | [CHANGELOG.md](../CHANGELOG.md) |
| Hosted Documentation Site | Browsable documentation | [https://kr1ny77.github.io/BasketForm-AI/](https://kr1ny77.github.io/BasketForm-AI/) |

## Known Limitations

1. **Single-VM deployment** — no horizontal scaling, single point of failure
2. **JSON file storage** — limits concurrent users to ~10 due to file-level locking
3. **No HTTPS** — HTTP only; consider adding SSL for production use
4. **No automated backup** — data is stored locally on the VM
5. **Multi-throw video upload** — automatic video splitting was deferred; manual single-file upload is the current workflow
6. **ML model weights** — `best.pt` must be present in the ML directory for analysis to work

## Current Handover Level

**Handover level:** Ready for independent use

The customer has been able to register an account, upload videos, view analysis results, use the friends feature, and share results during the Week 6 trial. The product is deployed and accessible at the provided URL. The customer can use the product independently without team assistance.

**Customer confirmation status:** Accepted with follow-up items

The customer confirmed the product works and meets the Definition of Done. Follow-up items identified during the Week 6 trial:
- Automatic video splitting feature (deferred to post-course)
- General site optimization and bug fixes (planned for Sprint 5)

## Support Expectations

During the course (through Week 7):
- Team available for bug fixes and final stabilization
- Customer can report issues via the weekly meeting or direct communication

After course completion:
- Source code and documentation remain available in the public repository
- VM hosting may be discontinued after course evaluation
- Customer responsible for own deployment if continued use is desired
