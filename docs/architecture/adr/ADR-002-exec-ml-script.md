# ADR-002: exec-based ML Integration

## Status
Accepted

## Context
The Go server needs to invoke the Python ML pipeline for video analysis. The ML pipeline uses MediaPipe, OpenCV, and YOLO — libraries difficult to use from Go.

Options considered: gRPC microservice, REST API between Go and Python, Direct exec.Command call, CGo bindings.

## Decision
Use exec.Command from Go to invoke the Python script as a subprocess. Go server spawns a goroutine for each video, runs python ML/main.py, waits for completion. Results communicated via JSON file on disk.

## Consequences and Tradeoffs

**Positive:** Simplest integration, no network overhead, Python runs in own process, easy to debug independently.

**Negative:** No hot-reloading, process spawning overhead, no streaming progress, harder to scale.

## Quality Requirements Addressed
- QR-001 (Performance): Async goroutine prevents blocking
- QR-004 (ML Accuracy): Full Python ecosystem available
- QR-005 (Concurrency): Each video gets own goroutine, limited by CPU
