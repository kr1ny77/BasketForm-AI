# User Acceptance Tests

End-user-facing scenarios for BasketForm-AI. These are executed by the customer during Sprint Review.

## UAT-001: Video Upload and Analysis

**Status:** Active

**User goal:** Upload a basketball shooting video and view the analysis result.

**Preconditions:** Application is running and accessible in a browser. User has a supported video file (MP4, MOV, or AVI).

**Step-by-step instructions:**
1. Navigate to the main page.
2. Find the upload section and drag the video file into the area, or upload it using the file button.
3. Wait for the video to process (processing progress bar should display, ~5-7 seconds).
4. Review the analysis results.
5. Verify the equipment ratings and recommendations for improvement are displayed.

**Expected outcome:** Video is uploaded successfully, processing progress bar is shown, and the results page displays score, phase analysis (Stance, Arm Angle, Release, Follow-through), and feedback.

**Execution result (Sprint 2 UAT):** Passed — customer confirmed video uploaded, processed, and results displayed correctly.

## UAT-002: PDF Export

**Status:** Active

**User goal:** Export the analysis report as a PDF for offline viewing.

**Preconditions:** User has at least one completed analysis result.

**Step-by-step instructions:**
1. Navigate to the "Export" section.
2. Select one analysis result.
3. Export it as PDF.

**Expected outcome:** The browser successfully initiates the file download without errors. The downloaded PDF can be opened, viewed, and contains the throw report with score breakdown and feedback.

**Execution result (Sprint 2 UAT):** Passed — customer confirmed PDF downloaded and contained the analysis report.

## UAT-003: Canvas Animation and UI Interaction

**Status:** Active

**User goal:** Verify the interactive user interface and canvas animation work correctly without breaking functionality.

**Preconditions:** Application is running and accessible in a browser.

**Step-by-step instructions:**
1. Navigate to the main page.
2. Observe the basketball-themed background animation (floating basketball objects).
3. Move the mouse cursor around the screen.
4. Click on the main elements (buttons, upload area, navigation links).
5. Verify that foreground elements remain fully clickable, readable, and unaffected by the animation.

**Expected outcome:** The background animation loads smoothly, dynamically reacts to mouse movement without FPS drops or browser freezes, and all foreground elements remain fully functional.

**Execution result (Sprint 2 UAT):** Passed — customer confirmed animation works well, UI elements are not affected by the background animation.

## UAT-004: Account Creation and Login

**Status:** Active

**User goal:** Register a new account and log in to access the application.

**Preconditions:** Application is running and accessible in a browser.

**Step-by-step instructions:**
1. Navigate to `/register`.
2. Enter a nickname (at least 3 characters), email, and password (at least 6 characters).
3. Click "Sign Up".
4. Verify redirect to `/login`.
5. Enter the registered email and password.
6. Click "Sign In".
7. Verify redirect to `/upload` page with navigation links visible.

**Expected outcome:** User is registered and can log in. Protected pages are accessible after login.
