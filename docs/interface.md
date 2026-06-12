# Interface Documentation - BasketForm-AI

## Overview

BasketForm-AI provides a web-based interface for basketball shot analysis. The system allows users to upload videos, analyze shooting technique, and receive personalized feedback.

## Interface Type

**Primary Interface:** Web Application (React/Next.js)
**Secondary Interface:** REST API (Node.js/Express)

## Target Users

1. **Basketball Players** - Upload videos and receive technique analysis
2. **Basketball Coaches** - Analyze player performance and provide guidance
3. **Sports Analysts** - Access detailed biomechanical data
4. **Basketball Enthusiasts** - Learn about shooting technique

## Core Workflows

### 1. Video Upload and Recording

**User Action:** Upload a basketball shot video
**System Response:** Process video and extract keypoints

**Steps:**
1. User navigates to upload page
2. User selects video file (MP4, MOV, AVI) or records directly
3. System validates file type and size
4. System uploads video to secure storage
5. System confirms upload success

**Success Example:**
```json
{
  "status": "success",
  "video_id": "vid_123456",
  "message": "Video uploaded successfully",
  "file_size": "45.2 MB",
  "format": "mp4"
}
```

**Error Examples:**
```json
{
  "status": "error",
  "error_code": "INVALID_FILE_TYPE",
  "message": "Unsupported video format. Please use MP4, MOV, or AVI."
}
```

```json
{
  "status": "error",
  "error_code": "FILE_TOO_LARGE",
  "message": "File size exceeds 100MB limit. Current size: 150MB"
}
```

### 2. Biomechanical Analysis

**User Action:** Request analysis of uploaded video
**System Response:** Extract keypoints and evaluate technique

**Steps:**
1. User selects video from library
2. User clicks "Analyze" button
3. System processes video frame-by-frame
4. System extracts pose landmarks using MediaPipe
5. System generates analysis report

**Success Example:**
```json
{
  "status": "success",
  "analysis_id": "ana_789012",
  "video_id": "vid_123456",
  "keypoints_extracted": 33,
  "frames_analyzed": 120,
  "processing_time": "8.5 seconds"
}
```

**Error Examples:**
```json
{
  "status": "error",
  "error_code": "INSUFFICIENT_KEYPOINTS",
  "message": "Could not detect enough keypoints in video. Please ensure good lighting and clear view of player."
}
```

### 3. Technique Evaluation

**User Action:** View technique evaluation results
**System Response:** Display scores and visual feedback

**Steps:**
1. System compares extracted keypoints to ideal form
2. System calculates scores for different aspects
3. System generates visual feedback overlay
4. System presents results to user

**Success Example:**
```json
{
  "status": "success",
  "evaluation": {
    "overall_score": 85,
    "aspects": {
      "stance": 90,
      "arm_angle": 80,
      "release_point": 85,
      "follow_through": 88
    },
    "strengths": ["Good balance", "Consistent release point"],
    "improvements": ["Elbow could be more tucked", "Follow-through could be higher"]
  }
}
```

### 4. Personalized Feedback

**User Action:** Receive personalized training recommendations
**System Response:** Generate AI-powered feedback and drills

**Steps:**
1. System analyzes evaluation results
2. System identifies areas for improvement
3. System generates personalized feedback
4. System suggests specific drills
5. System creates training plan

**Success Example:**
```json
{
  "status": "success",
  "feedback": {
    "summary": "Your shooting form is solid with good balance. Focus on elbow positioning for more consistency.",
    "drills": [
      {
        "name": "Elbow Tuck Drill",
        "description": "Practice shooting with a focus on keeping elbow tucked in",
        "repetitions": 50,
        "frequency": "Daily"
      },
      {
        "name": "Follow-Through Extension",
        "description": "Work on extending follow-through higher for better arc",
        "repetitions": 30,
        "frequency": "3x per week"
      }
    ],
    "training_plan": "Week 1: Focus on elbow positioning | Week 2: Add follow-through drills"
  }
}
```

## API Endpoints

### Base URL
- **Development:** http://localhost:4000
- **Production:** https://api.basketformai.com

### Authentication
- **Type:** JWT (JSON Web Tokens)
- **Header:** `Authorization: Bearer {{access_token}}`

### Endpoints

#### POST /api/videos/upload
Upload a video file for analysis.

**Request:**
```bash
curl -X POST https://api.basketformai.com/api/videos/upload \
  -H "Authorization: Bearer {{access_token}}" \
  -F "video=@basketball_shot.mp4"
```

**Response:** 201 Created
```json
{
  "status": "success",
  "video_id": "vid_123456",
  "message": "Video uploaded successfully"
}
```

#### GET /api/videos/{video_id}
Get video details and analysis status.

**Request:**
```bash
curl -X GET https://api.basketformai.com/api/videos/vid_123456 \
  -H "Authorization: Bearer {{access_token}}"
```

**Response:** 200 OK
```json
{
  "status": "success",
  "video": {
    "id": "vid_123456",
    "filename": "basketball_shot.mp4",
    "size": "45.2 MB",
    "uploaded_at": "2026-06-12T10:30:00Z",
    "analysis_status": "completed"
  }
}
```

#### POST /api/analysis/{video_id}
Start analysis of uploaded video.

**Request:**
```bash
curl -X POST https://api.basketformai.com/api/analysis/vid_123456 \
  -H "Authorization: Bearer {{access_token}}"
```

**Response:** 202 Accepted
```json
{
  "status": "success",
  "analysis_id": "ana_789012",
  "message": "Analysis started",
  "estimated_time": "10 seconds"
}
```

#### GET /api/analysis/{analysis_id}
Get analysis results.

**Request:**
```bash
curl -X GET https://api.basketformai.com/api/analysis/ana_789012 \
  -H "Authorization: Bearer {{access_token}}"
```

**Response:** 200 OK
```json
{
  "status": "success",
  "analysis": {
    "id": "ana_789012",
    "video_id": "vid_123456",
    "overall_score": 85,
    "keypoints": [...],
    "evaluation": {...},
    "feedback": {...}
  }
}
```

#### GET /api/users/{user_id}/history
Get user's analysis history.

**Request:**
```bash
curl -X GET https://api.basketformai.com/api/users/usr_123/history \
  -H "Authorization: Bearer {{access_token}}"
```

**Response:** 200 OK
```json
{
  "status": "success",
  "history": [
    {
      "analysis_id": "ana_789012",
      "video_id": "vid_123456",
      "overall_score": 85,
      "analyzed_at": "2026-06-12T10:35:00Z"
    }
  ],
  "total_analyses": 15,
  "average_score": 82
}
```

## Error Codes

| Code | Description | HTTP Status |
|------|-------------|-------------|
| `INVALID_FILE_TYPE` | Unsupported video format | 400 |
| `FILE_TOO_LARGE` | File exceeds size limit | 400 |
| `UNAUTHORIZED` | Invalid or missing authentication | 401 |
| `FORBIDDEN` | Insufficient permissions | 403 |
| `NOT_FOUND` | Resource not found | 404 |
| `INSUFFICIENT_KEYPOINTS` | Cannot detect enough pose landmarks | 422 |
| `ANALYSIS_FAILED` | Analysis processing failed | 500 |
| `RATE_LIMIT_EXCEEDED` | Too many requests | 429 |

## Configuration

### Environment Variables
- `{{access_token}}` - JWT authentication token
- `API_KEY` - API key for external integrations
- `AWS_S3_BUCKET` - S3 bucket for video storage

### Rate Limits
- **Free tier:** 10 analyses per day
- **Pro tier:** 100 analyses per day
- **Enterprise:** Custom limits

## Security Considerations

1. **Authentication:** All API requests require JWT authentication
2. **File Validation:** Videos are validated for type and size before processing
3. **Data Encryption:** Videos are encrypted at rest and in transit
4. **Access Control:** Users can only access their own videos and analyses
5. **Rate Limiting:** API requests are rate-limited to prevent abuse

## Implementation Status

### MVP v0 (Current)
- ✅ Video upload endpoint
- ✅ Basic authentication
- ✅ Mock analysis response
- ⬜ Real-time analysis
- ⬜ Advanced feedback generation

### MVP v1 (Planned)
- ✅ All MVP v0 features
- ✅ Real biomechanical analysis
- ✅ AI-powered feedback
- ✅ Progress tracking
- ✅ Team management

## Notes

- Placeholder URLs will be replaced with actual endpoints before submission
- API documentation will be enhanced with Swagger/OpenAPI specification
- Postman collection will be created for testing
- All credentials use `{{placeholder}}` format in public documentation