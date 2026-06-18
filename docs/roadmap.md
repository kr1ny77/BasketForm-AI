# Product Roadmap

## Vision

BasketForm-AI is an AI-powered platform that helps basketball players improve their shooting technique through video analysis and personalized feedback.

## Current Status

- **MVP v0:** Deployed with basic video upload functionality
- **MVP v1:** In development - core analysis and feedback features

## MVP Versions

### MVP v0 (Completed)
- Basic video upload functionality
- Simple web interface
- Deployment at http://80.74.30.14/

### MVP v1 (Current Sprint)
**Goal:** Deliver core basketball shooting form analysis with AI-powered feedback

**Scope:**
- US-001: Upload shooting form video
- US-002: View automated form analysis
- US-003: Create and manage a user account
- US-004: Receive simplified actionable feedback
- US-010: Export analysis report as PDF

**Key Features:**
- Video upload with format validation
- MediaPipe pose estimation for keypoint extraction
- Biomechanical analysis algorithm
- Personalized feedback generation
- User account management
- PDF export of reports

### MVP v2 (Future)
**Goal:** Enhanced comparison and sharing features

**Planned Features:**
- US-005: Compare form with professional players
- US-006: Share analysis report with a coach
- US-007: Share progress with friends
- US-008: Track shooting progress over time

### MVP v3 (Future)
**Goal:** Community and social features

**Planned Features:**
- Advanced analytics dashboard
- Coach portal
- Team management
- Competition features

## Timeline

### Sprint 1 (Current)
- **Duration:** 2 weeks
- **Focus:** MVP v1 implementation
- **Deliverables:** Core analysis features, user accounts, PDF export

### Sprint 2 (Planned)
- **Duration:** 2 weeks
- **Focus:** Comparison and sharing features
- **Deliverables:** Professional comparison, coach sharing, progress tracking

### Sprint 3 (Future)
- **Duration:** 2 weeks
- **Focus:** Community features
- **Deliverables:** Social features, advanced analytics

## Technical Roadmap

### Phase 1: Foundation (MVP v0-v1)
- FastAPI backend setup
- MediaPipe integration
- React frontend
- PostgreSQL database
- Basic authentication

### Phase 2: Enhancement (MVP v2)
- Professional video library
- Sharing mechanisms
- Progress tracking
- Advanced analytics

### Phase 3: Scale (MVP v3)
- Coach portal
- Team features
- Competition system
- Mobile app

## Success Metrics

- **MVP v1:** 100+ registered users, 500+ analyses performed
- **MVP v2:** 500+ registered users, coach adoption
- **MVP v3:** 1000+ users, team adoption

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| AI model accuracy | High | Extensive testing with diverse videos |
| Processing performance | Medium | Optimize algorithms, consider GPU acceleration |
| User adoption | High | Focus on UX, gather feedback early |
| Privacy concerns | High | Implement robust data protection |

## Links

- [User Stories Index](user-stories.md)
- [Definition of Done](definition-of-done.md)
- [CHANGELOG](../CHANGELOG.md)
- [Repository](https://github.com/kr1ny77/BasketForm-AI)
