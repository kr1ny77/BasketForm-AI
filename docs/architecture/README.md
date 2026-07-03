# Architecture Documentation

This directory contains the maintained architecture documentation for BasketForm-AI, using diagrams-as-code (PlantUML).

## Views

### Static View — Component and Package Structure

The static view shows the system's structural organization: components, their responsibilities, and relationships.

```plantuml
@startuml
!theme plain
skinparam packageStyle rectangle
skinparam componentStyle rectangle

package "BasketForm-AI" {
  [cmd/server/main.go] as entrypoint
  
  package "internal/" {
    [handlers/] as handlers
    [services/] as services
    [models/] as models
    [qrt/] as qrt
  }
  
  package "web/" {
    [templates/] as templates
    [static/] as static
  }
  
  package "scripts/" {
    [Python ML scripts] as mlscripts
  }
  
  package "data/" as data
  
  package "uploads/" as uploads
  package "results/" as results
}

entrypoint --> handlers
handlers --> services
handlers --> templates
handlers --> static
services --> models
services --> data
services --> uploads
services --> results
services --> mlscripts
qrt --> services

note right of handlers
  HTTP handlers for pages,
  API endpoints, auth,
  friends, sharing
end note

note right of services
  Business logic:
  auth (bcrypt + JWT),
  storage (JSON files),
  processor (ML pipeline)
end note

note right of mlscripts
  Python scripts using
  MediaPipe Holistic + OpenCV
  for pose estimation
end note
@enduml
```

**Source:** [static-view.puml](static-view.puml)
**Rendered:** See [static-view.svg](static-view.svg)

### Dynamic View — Request Flow and ML Pipeline

The dynamic view shows how a user request flows through the system, from upload to feedback delivery.

```plantuml
@startuml
!theme plain
skinparam sequenceMessageAlign center

actor User
participant "Frontend\n(HTML/CSS/JS)" as FE
participant "Go HTTP Server" as Server
participant "Auth Middleware" as Auth
participant "Upload Handler" as Upload
participant "Storage Service" as Storage
participant "ML Processor" as ML
participant "LLM API" as LLM
participant "Results Handler" as Results

User -> FE : Upload video
FE -> Server : POST /api/upload
Server -> Auth : Validate JWT
Auth -> Upload : Process upload
Upload -> Storage : Save video file
Storage -> ML : Trigger processing
ML -> ML : Extract keypoints (MediaPipe)
ML -> ML : Analyze phases (Stance, Arm Angle, Release, Follow-through)
ML -> Storage : Save results + annotated video
ML -> LLM : Send analysis for personalized feedback
LLM -> Storage : Save feedback
Storage -> Server : Return status
Server -> FE : Return upload status
FE -> User : Show processing progress

User -> FE : View results
FE -> Server : GET /api/result/{id}
Server -> Auth : Validate JWT
Server -> Results : Fetch results
Results -> Storage : Load results + feedback
Storage -> Server : Return data
Server -> FE : Return results JSON
FE -> User : Display score, feedback, annotated video
@enduml
```

**Source:** [dynamic-view.puml](dynamic-view.puml)
**Rendered:** See [dynamic-view.svg](dynamic-view.svg)

### Deployment View — VM Deployment and CI/CD

The deployment view shows how the system is deployed and how changes flow from development to production.

```plantuml
@startuml
!theme plain
skinparam nodeStyle rectangle

node "Development" {
  [Developer Machine] as Dev
  [GitHub Repository] as GH
}

node "CI/CD (GitHub Actions)" {
  [Lint] as Lint
  [Test] as Test
  [Coverage] as Coverage
  [QRT] as QRT
  [govulncheck] as Sec
  [Lychee] as Link
}

node "Production VM (80.74.30.14)" {
  [Go Binary] as Binary
  [Python + MediaPipe] as Python
  [Static Files] as Static
  [Data Directory] as Data
  [Nginx/Reverse Proxy] as Proxy
}

Dev --> GH : Push branch
GH --> Lint : Trigger CI
GH --> Test : Trigger CI
GH --> Coverage : Trigger CI
GH --> QRT : Trigger CI
GH --> Sec : Trigger CI
GH --> Link : Trigger CI

GH --> Binary : Deploy on merge to main
Binary --> Python : exec ML scripts
Binary --> Static : Serve templates/static
Binary --> Data : Store JSON
Proxy --> Binary : Forward requests

note right of Proxy
  Production endpoint:
  http://80.74.30.14/
  Port 80 externally,
  Port 8080 internally
end note

note bottom of GH
  Protected branch:
  - Direct push disabled
  - Required reviews
  - Required status checks
end note
@enduml
```

**Source:** [deployment-view.puml](deployment-view.puml)
**Rendered:** See [deployment-view.svg](deployment-view.svg)

## Key Architecture Decisions

Architecture Decision Records (ADRs) are maintained in this directory. See individual ADR files for detailed rationale.

| ADR | Decision | Rationale |
|-----|----------|-----------|
| ADR-001 | Go standard library for backend | Simplicity, no framework dependencies, easy deployment |
| ADR-002 | JSON file storage | Simple persistence without database overhead for MVP |
| ADR-003 | MediaPipe for pose estimation | Pre-trained model, good accuracy, Python integration |
| ADR-004 | JWT with HttpOnly cookies | Security (XSS protection), stateless auth |
| ADR-005 | LLM API for feedback | Personalized coaching-style feedback without custom ML |

## How to Maintain

1. When code structure changes, update the static view
2. When request flows change, update the dynamic view
3. When deployment changes, update the deployment view
4. When key decisions change, create a new ADR
5. Keep PlantUML sources in the repository and render SVGs before merging
