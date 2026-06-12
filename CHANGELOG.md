# Changelog

All notable changes to the BasketForm-AI project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project structure and repository setup
- MIT License
- Basic README with project description
- Assignment 2 documentation and reports
- User stories with MoSCoW priorities
- Interactive Figma prototype
- MVP v0 deployment with basic functionality
- Customer meeting documentation
- Week 2 analysis and LLM usage report
- Lychee link checking configuration
- GitHub PR template
- Interface documentation

### Changed
- Updated project description for clarity
- Improved repository structure for better organization

### Deprecated
- Nothing deprecated yet

### Removed
- US-09: Social Sharing (removed due to privacy concerns)

### Fixed
- Initial setup and configuration

### Security
- Added .gitignore for sensitive files
- Created .env.example for environment configuration
- Implemented basic authentication system

## [0.1.0] - 2026-06-12

### Added
- **US-01: Video Upload and Recording**
  - Basic video upload functionality
  - Support for MP4, MOV, AVI formats
  - File size validation (100MB limit)
  - Upload progress indicator

- **US-02: Biomechanical Keypoint Extraction**
  - Backend service structure ready
  - API endpoints defined
  - Placeholder for MediaPipe integration

- **US-03: Shooting Technique Evaluation**
  - API endpoints for analysis
  - Mock analysis results
  - Technique scoring system

- **US-04: Personalized Feedback Generation**
  - Feedback display interface
  - Placeholder for AI feedback generation
  - Basic drill recommendations

- **US-05: User Account Management**
  - User registration and login
  - JWT authentication
  - Basic profile management

### Technical Foundation
- Frontend: React/Next.js
- Backend: Node.js/Express
- Database: PostgreSQL
- File Storage: AWS S3
- Authentication: JWT
- CI/CD: GitHub Actions

### Documentation
- Project README with setup instructions
- Assignment 2 reports and documentation
- User stories and prioritization
- Interface documentation
- MVP v0 deployment report

---

## Versioning Strategy

- **Major versions (X.0.0):** Significant changes to core functionality
- **Minor versions (0.X.0):** New features and enhancements
- **Patch versions (0.0.X):** Bug fixes and minor improvements

## Release Process

1. Update CHANGELOG.md with new version entry
2. Update version in package.json
3. Create Git tag with version number
4. Create GitHub release with release notes
5. Deploy to production environment

## Links

- [Keep a Changelog](https://keepachangelog.com/)
- [Semantic Versioning](https://semver.org/)
- [GitHub Releases](https://github.com/kr1ny77/BasketForm-AI/releases)