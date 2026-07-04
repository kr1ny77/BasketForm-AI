# ADR-002: exec-based ML Integration

## Status
Accepted

## Context
<<<<<<< HEAD
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
=======
The Go server needs to invoke the Python ML pipeline for video analysis. The ML pipeline uses MediaPipe, OpenCV, and YOLO — libraries that are difficult to use from Go.

Options considered:
- gRPC microservice (Python server)
- REST API between Go and Python
- Direct exec.Command call
- CGo bindings to Python

## Decision
Use `exec.Command` from Go to invoke the Python script as a subprocess. The Go server spawns a goroutine for each video, runs `python ML/main.py input.mp4 report.json --lang en`, and waits for completion. Results are communicated via JSON file on disk.

## Consequences and Tradeoffs

**Positive:**
- Simplest integration — no network overhead, no service discovery
- Python code runs in its own process (memory isolation)
- Easy to debug — can run the Python script independently
- No need for gRPC/REST framework in Python

**Negative:**
- No hot-reloading — Python changes require server restart (if binary includes ML code)
- Process spawning overhead (~100ms per video)
- No streaming progress back to client during processing
- Harder to scale — each video processing blocks a goroutine

## Quality Requirements Addressed
- **QR-001 (Performance)**: Async goroutine prevents blocking; processing is background task
- **QR-004 (ML Accuracy)**: Full Python ecosystem available (MediaPipe, OpenCV, YOLO)
- **QR-005 (Concurrency)**: Each video gets its own goroutine; limited by CPU cores
>>>>>>> a0ea7bf844a20294cb3ee0403c4363ceca8235c0
