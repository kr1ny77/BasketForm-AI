# Interface Documentation - BasketForm-AI

## Overview

BasketForm-AI provides a web-based interface for basketball players to upload videos, receive AI-powered shooting form analysis, and get personalized feedback. The product also exposes a REST API for programmatic access.

## Interface Type

**Primary Interface:** Web Application (Graphical User Interface)
**Secondary Interface:** REST API

## Intended Users

- **Basketball Players:** Upload videos, view analysis, receive feedback
- **Basketball Coaches:** Review player analysis reports (via shared links)
- **Developers:** Integrate with the system via REST API

## Web Interface

### Screens and Navigation

1. **Home Page** (`/`)
   - Project overview and call-to-action
   - Navigation to upload, analysis, and account pages

2. **Video Upload Page** (`/upload`)
   - Drag-and-drop or file picker for video upload
   - Supported formats: MP4, MOV, AVI
   - Maximum file size: 50MB
   - Camera positioning guide overlay
   - File type validation (video files only)

3. **Analysis Results Page** (`/analysis/:id`)
   - Video player with biomechanical overlay
   - Keypoint visualization on frame
   - Summary feedback (brief, expandable for details)
   - Score breakdown (stance, arm angle, release point, follow-through)
   - PDF export button

4. **User Dashboard** (`/dashboard`)
   - List of uploaded videos and analyses
   - Progress tracking charts
   - Account settings

5. **Sign Up / Sign In** (`/auth/signup`, `/auth/login`)
   - Email/password registration
   - OAuth (Google) login option

### Key States

- **Empty State:** No videos uploaded yet, prompt to upload first video
- **Loading State:** Video processing in progress with progress indicator
- **Success State:** Analysis complete, results displayed
- **Error State:** Upload failure, processing error, or unsupported format

## REST API

### Base URL

```
http://80.74.30.14/api
```

### Authentication

- JWT token-based authentication
- Token passed via `Authorization: Bearer {{access_token}}` header
- Token obtained via `/auth/login` endpoint

### Endpoints

#### Health Check
```
GET /health
```
Returns service health status.

**Response (200):**
```json
{
  "status": "healthy",
  "version": "0.1.0"
}
```

#### Video Upload
```
POST /videos/upload
Content-Type: multipart/form-data
```
Upload a basketball shooting video for analysis.

**Request:**
- `file`: Video file (MP4, MOV, AVI; max 50MB)

**Response (201):**
```json
{
  "id": "video_abc123",
  "filename": "shot_video.mp4",
  "status": "uploaded",
  "created_at": "2026-06-14T10:00:00Z"
}
```

**Error (400):**
```json
{
  "error": "invalid_file_type",
  "message": "Only MP4, MOV, and AVI files are supported"
}
```

#### Get Analysis
```
GET /videos/{id}/analysis
```
Retrieve the biomechanical analysis for a video.

**Response (200):**
```json
{
  "video_id": "video_abc123",
  "status": "completed",
  "scores": {
    "stance": 85,
    "arm_angle": 72,
    "release_point": 90,
    "follow_through": 78
  },
  "keypoints": [...],
  "feedback": {
    "brief": "Good release point. Adjust elbow alignment.",
    "detailed": "..."
  }
}
```

**Error (404):**
```json
{
  "error": "not_found",
  "message": "Analysis not found for this video"
}
```

#### User Account
```
POST /auth/signup
POST /auth/login
GET /auth/me
```

### Implemented / Mocked in MVP v0

| Endpoint | Status |
|----------|--------|
| `GET /health` | Implemented |
| `POST /videos/upload` | Implemented (basic) |
| `GET /videos/{id}/analysis` | Mocked (returns sample data) |
| `POST /auth/signup` | Implemented (basic) |
| `POST /auth/login` | Implemented (basic) |

## Authentication and Configuration

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `DATABASE_URL` | PostgreSQL connection string | Yes |
| `JWT_SECRET` | Secret for JWT token signing | Yes |
| `AWS_S3_BUCKET` | S3 bucket for video storage | Yes |
| `API_KEY` | Third-party API key | No |

### Configuration

- CORS enabled for frontend origin
- Rate limiting: 100 requests/minute per user
- File upload limit: 50MB
- Session timeout: 24 hours

## Error Handling

All API errors follow a consistent format:

```json
{
  "error": "error_code",
  "message": "Human-readable error description"
}
```

Common error codes:
- `invalid_file_type`: Unsupported video format
- `file_too_large`: File exceeds 50MB limit
- `unauthorized`: Missing or invalid authentication
- `not_found`: Requested resource not found
- `processing_error`: Video analysis failed
