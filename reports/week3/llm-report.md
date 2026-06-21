# LLM Usage Report

## Tools Used

- **Claude Code (Anthropic)** — Code generation, documentation, project planning
- **GitHub Copilot** — Code completion during development

## Usage Scenarios

### 1. Go Backend Architecture
- **Tool:** Claude Code
- **Usage:** Designed Go project structure (`cmd/server`, `internal/handlers`, `internal/services`, `internal/models`), handler wiring, API endpoint design, and service layer separation

### 2. Canvas Animation
- **Tool:** Claude Code
- **Usage:** Generated the basketball-themed Canvas animation with 35 objects, mouse repulsion physics, trail lines, and wave rendering

### 3. API Handler Implementation
- **Tool:** Claude Code
- **Usage:** Created REST API handlers for upload (multipart), status, result, and video list endpoints with proper error handling and JSON responses

### 4. Mock ML Pipeline
- **Tool:** Claude Code
- **Usage:** Designed the async processor with goroutines, incremental progress, score generation, feedback selection, and pose skeleton generation

### 5. Test Suite
- **Tool:** Claude Code + GitHub Copilot
- **Usage:** Generated 47 unit and integration tests covering UUID generation, file validation, storage operations, API endpoints, and processor behavior

### 6. Documentation and Reports
- **Tool:** Claude Code
- **Usage:** Created CHANGELOG, roadmap, reflection, retrospective, and Week 3 report following course requirements

### 7. Dockerfile
- **Tool:** Claude Code
- **Usage:** Multi-stage Docker build with Alpine-based final image

## Impact

- **Positive:** Accelerated boilerplate creation (handlers, tests, templates), reduced time on documentation formatting
- **Challenge:** Required careful review of generated code for correctness, especially template syntax and error handling paths

## Best Practices

1. Always review AI-generated code before committing
2. Use AI for boilerplate and structure, human judgment for architecture decisions
3. Verify tests actually test the right behavior, not just pass
