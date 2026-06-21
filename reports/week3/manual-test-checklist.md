# Manual Test Checklist — BasketForm-AI MVP v1

## Preconditions

- Server running: `go run ./cmd/server/`
- Browser: Chrome, Firefox, or Safari (latest)
- Test videos: one valid MP4 (<100MB), one MOV, one AVI, one MKV (invalid), one empty file (0 bytes)

---

## 1. Upload Page

### Positive Scenarios

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 1.1 | Open `/upload` | Page loads with dropzone, basketball animation in background | |
| 1.2 | Click the dropzone | File browser opens | |
| 1.3 | Select valid MP4 file | Progress bar appears, uploads to 100%, redirects to `/progress?id=<uuid>` | |
| 1.4 | Drag valid MOV file onto dropzone | Dropzone highlights on dragover, file uploads on drop | |
| 1.5 | Upload valid AVI file | Same as 1.3 | |
| 1.6 | After redirect, check `uploads/` directory | File `<uuid>.mp4` exists with correct content | |

### Negative / Edge Cases

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 1.7 | Select MKV file | Alert: "Unsupported format. Use MP4, MOV, or AVI." | |
| 1.8 | Select JPG/PNG image | Alert: "Unsupported format" | |
| 1.9 | Select file >100MB | Alert: "File too large. Maximum size is 100 MB." | |
| 1.10 | Upload empty file (0 bytes) | Accepted (201), file created in uploads/ | |
| 1.11 | Upload while server is stopped | XHR error, "Upload failed." message shown | |
| 1.12 | Click dropzone, cancel file picker | No change, no error | |
| 1.13 | Rapidly click dropzone multiple times | No duplicate uploads or errors | |

---

## 2. Progress Page

### Positive Scenarios

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 2.1 | Navigate to `/progress?id=<valid-uuid>` | Status shows "Analyzing video…", progress bar fills | |
| 2.2 | Wait ~5 seconds | Progress reaches 100%, status shows "Analysis complete!" | |
| 2.3 | Click "View Results" button | Navigates to `/results?id=<uuid>` | |

### Negative / Edge Cases

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 2.4 | Navigate to `/progress` (no id) | "No video ID provided. Go to Upload to start." | |
| 2.5 | Navigate to `/progress?id=nonexistent` | "Cannot reach server" or "video not found" after retries | |
| 2.6 | Start polling, then stop server | "Cannot reach server. Retrying…" appears, no crash | |
| 2.7 | Navigate to `/progress?id=<uuid>` for already-done video | Immediately shows "Analysis complete!" | |

---

## 3. Results Page

### Positive Scenarios

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 3.1 | Navigate to `/results` with no data | "No analyses yet" empty state | |
| 3.2 | Upload video, wait for processing, go to `/results` | Card shows filename, ID (truncated), status badge, date, score | |
| 3.3 | Click a result card | Modal opens with score, feedback, radar chart, pose scatter plot | |
| 3.4 | Press ESC | Modal closes | |
| 3.5 | Click outside modal | Modal closes | |
| 3.6 | Click X button on modal | Modal closes | |
| 3.7 | Type in filter input | Results filter by filename in real-time | |
| 3.8 | Change sort dropdown | Results reorder (newest, oldest, score desc, score asc) | |

### Negative / Edge Cases

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 3.9 | Filter with no matches | "No results found" shown | |
| 3.10 | Upload multiple videos, check sort order | Default: newest first | |

---

## 4. Profile Page

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 4.1 | Navigate to `/profile` | Form with name, email, password fields | |
| 4.2 | Change name, click Save | Alert: "Profile saved (demo)." | |
| 4.3 | Enter mismatched passwords, click Save | Alert: "Passwords do not match." | |
| 4.4 | Leave password blank, click Save | Profile saved without password change | |

---

## 5. Export Page

### Positive Scenarios

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 5.1 | Navigate to `/export` with no completed analyses | Dropdown shows "No completed analyses" | |
| 5.2 | After a completed analysis, navigate to `/export` | Dropdown shows the completed video | |
| 5.3 | Select a video, click Download PDF | PDF downloads with filename, score, feedback | |
| 5.4 | Open downloaded PDF | Contains: title, score, feedback, date, video ID | |

### Negative / Edge Cases

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 5.5 | Don't select any video, click Download | Button is disabled | |
| 5.6 | Select video that is still processing | Not shown in dropdown (only "done" status) | |

---

## 6. Navigation

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 6.1 | Click each nav link | Correct page loads | |
| 6.2 | Check active state | Current page link is highlighted in orange | |
| 6.3 | Resize to mobile width | Nav wraps correctly, all links accessible | |

---

## 7. Canvas Animation

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 7.1 | Open any page | Canvas animation visible behind content | |
| 7.2 | Move mouse across the screen | Balls repel away from cursor | |
| 7.3 | Stop moving mouse | Balls continue falling with gravity | |
| 7.4 | Check object count | At least 30 animated objects visible | |
| 7.5 | Orange waves at bottom | Visible, animated, moving | |

---

## 8. Responsive Design

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 8.1 | View at 1200px width | Desktop layout | |
| 8.2 | View at 768px width | Tablet layout, cards stack | |
| 8.3 | View at 375px width (iPhone) | Mobile layout, nav wraps, forms usable | |
| 8.4 | Upload on mobile | Dropzone works with touch/file picker | |

---

## 9. API Endpoints (curl / Postman)

| # | Request | Expected | Pass |
|---|---------|----------|------|
| 9.1 | `POST /api/upload` with valid file | 201 + `{"id":"...","filename":"...","status":"uploaded"}` | |
| 9.2 | `GET /api/status/<id>` | 200 + `{"id":"...","status":"...","progress":N}` | |
| 9.3 | `GET /api/result/<id>` (after processing) | 200 + full result JSON | |
| 9.4 | `GET /api/videos` | 200 + array of video objects | |
| 9.5 | `GET /api/status/nonexistent` | 404 + `{"error":"video not found"}` | |
| 9.6 | `POST /api/upload` without file | 400 + `{"error":"missing video file"}` | |
| 9.7 | `POST /api/upload` with .txt file | 400 + `{"error":"unsupported format"}` | |
| 9.8 | `GET /api/upload` (wrong method) | 405 + `{"error":"method not allowed"}` | |

---

## 10. Error Handling & Resilience

| # | Steps | Expected | Pass |
|---|-------|----------|------|
| 10.1 | Start server, upload video, kill server, restart | Video still in uploads/, result may be missing | |
| 10.2 | Fill disk space, try upload | 500 error, no crash | |
| 10.3 | Send malformed multipart request | 400 error, no crash | |
| 10.4 | Send concurrent uploads (5 tabs) | All complete without corruption | |

---

## 11. Cross-Browser

| # | Browser | Pass |
|---|---------|------|
| 11.1 | Chrome (latest) | |
| 11.2 | Firefox (latest) | |
| 11.3 | Safari (latest) | |
| 11.4 | Mobile Chrome (Android) | |
| 11.5 | Mobile Safari (iOS) | |

---

## Sign-off

| Tester | Date | Result |
|--------|------|--------|
| | | Pass / Fail |
