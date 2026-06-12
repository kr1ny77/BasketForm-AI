# BasketForm-AI

Website with a backend service for recording and uploading basketball shot videos, extracting biomechanical keypoints, evaluating shooting technique, and generating personalized feedback.

## Project Overview

BasketForm-AI is an AI-powered platform that helps basketball players improve their shooting technique through video analysis and personalized feedback. The system uses computer vision and machine learning to extract biomechanical keypoints from video footage and provide detailed analysis of shooting form.

## Key Features

- **Video Upload & Recording:** Record or upload basketball shot videos
- **Biomechanical Analysis:** Automatic extraction of key body points
- **Technique Evaluation:** Scoring of stance, arm angle, release point, and follow-through
- **Personalized Feedback:** AI-generated recommendations and drills
- **Progress Tracking:** Monitor improvement over time

## Project Structure

```
BasketForm-AI/
├── LICENSE                    # MIT License
├── README.md                 # This file
├── CHANGELOG.md              # Version history
├── .gitignore               # Git ignore rules
├── .env.example             # Environment variables template
├── .lycheeignore            # Link checking exclusions
├── reports/
│   └── week2/
│       ├── README.md         # Assignment 2 public index
│       ├── user-stories.md   # User stories with MoSCoW priorities
│       ├── mvp-v0-report.md  # MVP v0 deployment report
│       ├── customer-meeting-summary.md
│       ├── customer-meeting-transcript.md
│       ├── customer-meeting-notes.md
│       ├── analysis.md       # Week 2 analysis
│       ├── llm-report.md     # LLM usage report
│       └── images/           # Screenshots and images
├── docs/
│   └── interface.md          # Interface documentation
└── .github/
    └── pull_request_template.md
```

## Local Development Setup

### Prerequisites

- Node.js (v18 or later)
- npm or yarn
- Python (v3.8 or later) - for backend services
- PostgreSQL (v14 or later)
- Git

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/kr1ny77/BasketForm-AI.git
   cd BasketForm-AI
   ```

2. **Install dependencies:**
   ```bash
   # Frontend
   cd frontend
   npm install
   
   # Backend
   cd ../backend
   pip install -r requirements.txt
   ```

3. **Set up environment variables:**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Database setup:**
   ```bash
   cd backend
   npm run db:create
   npm run db:migrate
   npm run db:seed
   ```

5. **Start development servers:**
   ```bash
   # Terminal 1 - Backend
   cd backend
   npm run dev
   
   # Terminal 2 - Frontend
   cd frontend
   npm run dev
   ```

6. **Access the application:**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:4000
   - API Documentation: http://localhost:4000/docs

## Documentation

### Assignment 2
- [Assignment 2 Report Index](reports/week2/README.md)
- [User Stories](reports/week2/user-stories.md)
- [MVP v0 Report](reports/week2/mvp-v0-report.md)
- [Customer Meeting Summary](reports/week2/customer-meeting-summary.md)
- [Week 2 Analysis](reports/week2/analysis.md)
- [LLM Usage Report](reports/week2/llm-report.md)

### Technical Documentation
- [Interface Documentation](docs/interface.md)
- [API Reference](docs/interface.md#rest-api)
- [Environment Configuration](.env.example)

### Project Management
- [Changelog](CHANGELOG.md)
- [License](LICENSE)

## Deployment

### MVP v0
- **Production URL:** [https://mvp-v0.basketformai.com](https://mvp-v0.basketformai.com)
- **Staging URL:** [https://staging.mvp-v0.basketformai.com](https://staging.mvp-v0.basketformai.com)
- **Video Demonstration:** [https://youtu.be/PLACEHOLDER_VIDEO_ID](https://youtu.be/PLACEHOLDER_VIDEO_ID)

For detailed deployment instructions, see [MVP v0 Report](reports/week2/mvp-v0-report.md).

### Smoke Check
To verify the deployment is working:
```bash
curl https://mvp-v0.basketformai.com/api/health
# Expected: {"status": "healthy", "version": "0.1.0"}
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### PR Requirements
- At least one approval from another team member
- All CI checks must pass
- Link to related issue (if applicable)
- Update CHANGELOG.md if user-visible change

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, please contact the development team via:
- GitHub Issues: [https://github.com/kr1ny77/BasketForm-AI/issues](https://github.com/kr1ny77/BasketForm-AI/issues)
- Email: [team@basketformai.com](mailto:team@basketformai.com)

## Acknowledgments

- Customer/Product Owner for valuable feedback and guidance
- Course instructors for project requirements and evaluation
- Open source community for tools and libraries used in this project