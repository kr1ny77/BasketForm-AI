# User Stories - BasketForm-AI

## User Roles/Personas

1. **Basketball Player** - Amateur or professional player looking to improve shooting technique
2. **Basketball Coach** - Trainer analyzing player performance
3. **Sports Analyst** - Data-driven analyst studying biomechanics
4. **Basketball Enthusiast** - Casual fan interested in technique analysis

## User Stories

### US-01: Video Upload and Recording

**Requirement status:** Active
**MoSCoW priority:** Must Have

As a basketball player,
I want to record and upload basketball shot videos,
so that I can analyze my shooting technique.

#### Notes and constraints
- Support multiple video formats (MP4, MOV, AVI)
- Maximum file size: 100MB
- Mobile and desktop upload support
- Progress indicator during upload

### US-02: Biomechanical Keypoint Extraction

**Requirement status:** Active
**MoSCoW priority:** Must Have

As a basketball coach,
I want the system to automatically extract biomechanical keypoints from videos,
so that I can analyze shooting form without manual annotation.

#### Notes and constraints
- Use MediaPipe or similar pose estimation
- Extract key joints: shoulders, elbows, wrists, hips, knees, ankles
- Process videos in real-time or near real-time
- Provide confidence scores for keypoints

### US-03: Shooting Technique Evaluation

**Requirement status:** Active
**MoSCoW priority:** Must Have

As a basketball player,
I want to receive an evaluation of my shooting technique,
so that I can understand my strengths and areas for improvement.

#### Notes and constraints
- Compare against ideal shooting form
- Provide scores for different aspects: stance, arm angle, release point, follow-through
- Visual feedback with highlighted areas
- Historical comparison over time

### US-04: Personalized Feedback Generation

**Requirement status:** Active
**MoSCoW priority:** Must Have

As a basketball player,
I want to receive personalized feedback and drills,
so that I can specifically address my technique weaknesses.

#### Notes and constraints
- AI-generated recommendations based on analysis
- Suggest specific drills for improvement
- Progress tracking over time
- Customized training plans

### US-05: User Account Management

**Requirement status:** Active
**MoSCoW priority:** Must Have

As a basketball player,
I want to create and manage my account,
so that I can save my analysis history and track progress.

#### Notes and constraints
- Email/password authentication
- Social login options (Google, Apple)
- Profile management
- Password reset functionality

### US-06: Video Library Management

**Requirement status:** Active
**MoSCoW priority:** Should Have

As a basketball player,
I want to organize my videos into collections,
so that I can easily find and compare past analyses.

#### Notes and constraints
- Create folders/collections
- Tag videos with metadata
- Search and filter functionality
- Bulk operations

### US-07: Real-time Analysis

**Requirement status:** Active
**MoSCoW priority:** Should Have

As a basketball coach,
I want to analyze shots in real-time during practice,
so that I can provide immediate feedback to players.

#### Notes and constraints
- Live camera feed analysis
- Low latency processing
- Mobile-friendly interface
- Offline capability for basic analysis

### US-08: Team Management

**Requirement status:** Active
**MoSCoW priority:** Could Have

As a basketball coach,
I want to manage my team and compare player performance,
so that I can identify team-wide trends and individual needs.

#### Notes and constraints
- Create team rosters
- Compare players side-by-side
- Team performance dashboards
- Export reports

### US-09: Social Sharing

**Requirement status:** Removed
**Previous MoSCoW priority:** Could Have

As a basketball enthusiast,
I want to share my analysis results on social media,
so that I can show my progress and get feedback from friends.

**Reason:** Removed due to privacy concerns and limited technical value for core product goals.

### US-10: Mobile App

**Requirement status:** Active
**MoSCoW priority:** Should Have

As a basketball player,
I want to use the service on my mobile phone,
so that I can record and analyze shots on the court.

#### Notes and constraints
- Responsive web design or native app
- Camera integration
- Offline basic analysis
- Push notifications for completed analyses

### US-11: API Access

**Requirement status:** Active
**MoSCoW priority:** Could Have

As a sports analyst,
I want to access the analysis API programmatically,
so that I can integrate the service into my own workflows.

#### Notes and constraints
- RESTful API with documentation
- API key authentication
- Rate limiting
- Webhook support for completion notifications

### US-12: Data Export

**Requirement status:** Active
**MoSCoW priority:** Could Have

As a sports analyst,
I want to export my analysis data in various formats,
so that I can perform custom analysis in external tools.

#### Notes and constraints
- Export as CSV, JSON, PDF
- Include all keypoint data
- Historical data export
- Batch export functionality

## Initial proposed MVP v1 scope

Based on the active Must Have user stories, the initial proposed MVP v1 scope includes:

1. **US-01**: Video Upload and Recording
2. **US-02**: Biomechanical Keypoint Extraction
3. **US-03**: Shooting Technique Evaluation
4. **US-04**: Personalized Feedback Generation
5. **US-05**: User Account Management

This scope provides the core functionality for analyzing basketball shooting technique while remaining small enough to prototype and discuss meaningfully with the customer.