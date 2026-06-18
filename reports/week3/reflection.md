# Week 3 Reflection

## Learning Points

### Product Backlog Migration
- Learned how to migrate user stories from documentation to GitHub Issues
- Understood the importance of preserving stable IDs for traceability
- Practiced creating well-structured issue bodies with acceptance criteria

### Product Backlog Refinement
- Learned to split large user stories into smaller, manageable PBIs
- Understood the DEEP principles (Detailed appropriately, Estimated, Emergent, Prioritized)
- Practiced adding supporting technical and infrastructure PBIs

### Estimation
- Learned story point estimation techniques
- Understood how to break down complex features into estimable tasks
- Practiced relative estimation compared to known reference stories

### Sprint Planning
- Learned to create Sprint milestones with clear goals
- Understood how to select PBIs for Sprint execution
- Practiced assigning story points and tracking Sprint velocity

### MVP v1 Delivery
- Learned to prioritize features based on customer value
- Understood the difference between Must Have, Should Have, and Could Have features
- Practiced delivering a minimum viable product with core functionality

### Workflow Enforcement
- Learned to use GitHub Issues, PRs, and milestones for project management
- Understood the importance of branch protection and code review
- Practiced maintaining a clean git history with meaningful commits

## Validated Assumptions

### Technical Assumptions
- **MediaPipe Pose estimation** works well for basketball shooting analysis
- **FastAPI** provides good performance for video processing endpoints
- **React** offers sufficient flexibility for the frontend interface

### Process Assumptions
- **GitHub Issues** are effective for tracking user stories and PBIs
- **Sprint milestones** provide clear boundaries for delivery
- **Story points** help in estimating and tracking progress

### User Assumptions
- **Beginners** need simplified, actionable feedback rather than technical details
- **PDF export** is valuable for offline record keeping
- **User accounts** are necessary for saving analysis history

## Friction and Gaps

### Technical Gaps
- **Video processing performance** may need optimization for large files
- **AI model accuracy** requires extensive testing with diverse video samples
- **Database design** needs refinement for scalability

### Process Gaps
- **Story point estimation** lacks historical data for accurate velocity tracking
- **Code review process** needs more team member participation
- **Testing coverage** may be insufficient for production readiness

### Requirements Gaps
- **Professional video library** for comparison feature needs legal clearance
- **Sharing mechanisms** require security review
- **Progress tracking** needs clear metric definitions

### Blocked Work
- **AI model training** requires labeled dataset of basketball shots
- **Production deployment** needs server access and configuration
- **Video demonstration** requires working MVP v1 features

## Planned Response

### Next Sprint (Sprint 2)
1. **Focus on comparison features** (US-005, US-006)
2. **Improve video processing performance** with optimization
3. **Enhance testing coverage** with integration tests
4. **Refine database design** for better scalability

### Process Improvements
1. **Establish regular code review sessions** to increase team participation
2. **Create estimation reference stories** for better velocity tracking
3. **Implement automated testing in CI pipeline** for quality assurance

### Technical Improvements
1. **Optimize MediaPipe integration** for faster processing
2. **Implement caching** for repeated analyses
3. **Add monitoring and logging** for production debugging

### Documentation Updates
1. **Update API documentation** with new endpoints
2. **Create user guides** for MVP v1 features
3. **Document deployment procedures** for production environment

## Links to Affected PBIs

- [US-001: Upload shooting form video](https://github.com/kr1ny77/BasketForm-AI/issues/20)
- [US-002: View automated form analysis](https://github.com/kr1ny77/BasketForm-AI/issues/21)
- [US-003: Create and manage a user account](https://github.com/kr1ny77/BasketForm-AI/issues/22)
- [US-004: Receive simplified actionable feedback](https://github.com/kr1ny77/BasketForm-AI/issues/23)
- [US-010: Export analysis report as PDF](https://github.com/kr1ny77/BasketForm-AI/issues/28)
- [Sprint 1 Milestone](https://github.com/kr1ny77/BasketForm-AI/milestone/1)
- [Roadmap](../../docs/roadmap.md)
