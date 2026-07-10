# Release v0.3.0 — MVP v2

**Release Date:** July 4, 2026
**Sprint:** 3 (Assignment 5)
**Milestone:** MVP v2

## What's New

### LLM-Powered Video Analysis
- Integrated OpenRouter API with GPT-4o-mini vision model
- LLM analyzes 8 key frames from the video alongside biomechanical metrics
- Generates accurate phase scores (DIP, ASCENT, RELEASE, FOLLOW-THROUGH) and detailed feedback
- Automatic fallback to rule-based scoring if LLM is unavailable

### Follow-Through Scoring
- 4th phase now fully scored with 3 components: duration (0-30pts), elbow snap (0-35pts), arm extension consistency (0-35pts)
- Tracks wrist angles, elbow stability, and forearm extension during follow-through

### Profile Improvements
- Change Nickname with password confirmation (modal popup)
- Change Password (modal popup)
- Avatar upload and display across the application
- Login by email OR nickname (case-insensitive)

### Social Features Enhancement
- Shared results page now shows both received AND sent results
- "Sent to Friends" section groups results by recipient

### Architecture Documentation
- Component diagram (PlantUML)
- Sequence diagram for video upload → analysis → result flow
- Deployment diagram showing VM architecture
- 3 Architecture Decision Records (ADR-001, ADR-002, ADR-003)

## Bug Fixes
- Fixed data race in `GetVideo` that caused server crashes with 3+ concurrent users
- Fixed avatar display on Friends and Shared pages (missing `onload` handler)
- Fixed `t is not defined` error on Friends page when loading results
- Fixed ProfileHandler not returning `avatar` field
- Removed duplicate `handlers/storage.go` that blocked Go compilation

## Quality Requirements
- QR-001 to QR-004 linked to ADRs in `docs/quality-requirements.md`

## Breaking Changes
- `feedback_generator.py` rewritten to use OpenRouter API instead of local Qwen model
- `llm_scorer.py` removed — scoring handled by existing `feedback_generator.py`
- Environment variable `OPENROUTER_API_KEY` required for LLM scoring

## Known Issues
- LLM scoring requires internet connection and valid OpenRouter API key
- Concurrent video processing limited by CPU cores on VM
- JSON file storage not suitable for >10 concurrent users
