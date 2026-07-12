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

## UAT-005: Send Friend Request

**Status:** Active

**User goal:** Navigate to the "Friends" page and send a friend request to another user using their username.

**Preconditions:** User is logged in. The target user has an existing account with a known username.

**Step-by-step instructions:**
1. Navigate to the "Friends" page using the main navigation menu.
2. Locate the "Search" field.
3. Enter the exact username of the target user.
4. Click the "Search" button and then "Send request".
5. Verify the success message or UI update indicating the request was sent.

**Expected outcome:** The friend request is successfully sent. The UI displays a confirmation message (e.g., "Friend request sent"), and the target user's username appears in the "Pending" or "Sent Requests" section.

**Execution result (Sprint 3 UAT):** Passed — customer confirmed they could successfully search for a username and send a friend request, with appropriate success feedback displayed.

## UAT-006: Accept Friend Request

**Status:** Active

**User goal:** Navigate to the "Friends" page and accept an incoming friend request from another user.

**Preconditions:** User is logged in. Another user has already sent a friend request to this user (as tested in UAT-005).

**Step-by-step instructions:**
1. Log in to the application as the user who received the friend request.
2. Navigate to the "Friends" page using the main navigation menu.
3. Locate the "Pending requests" section.
4. Find the friend request from the target user.
5. Click the "Accept" button next to the request.
6. Verify the UI updates to show the user is now in the active "Friends" list.

**Expected outcome:** The friend request is successfully accepted. The requesting user is removed from the "Incoming Requests" list and added to the active "Friends" list. A success notification or UI update confirms the new connection.

**Execution result (Sprint 3 UAT):** Passed — customer confirmed they could view incoming requests and successfully accept a friend request, correctly moving the user to the active friends list.

## UAT-007: Upload Video Without Human (No Person Detected)

**Status:** Active

**User goal:** Upload a video that does not contain a human and receive an appropriate error message indicating no person was found in the clip.

**Preconditions:** Application is running and accessible in a browser. User is logged in. User has a supported video file (MP4, MOV, or AVI) that contains no humans (e.g., an empty basketball court, a pet, or inanimate objects).

**Step-by-step instructions:**
1. Navigate to the main upload page.
2. Upload the video file that does not contain a human.
3. Wait for the video to process (processing progress bar should display).
4. Observe the application's response after the processing completes.

**Expected outcome:** The application successfully finishes processing the video but does not generate fake analysis results. Instead, it displays a clear error message or notification stating that the video cannot be analyzed because no person was found in the clip.

**Execution result (Sprint 3 UAT):** Passed — customer confirmed the system correctly identified the absence of a human in the video and displayed the expected error message without crashing.

**Execution result (Sprint 4 UAT):** Passed — customer confirmed the person-detection validation works correctly. Videos without a human are rejected with a clear error message. Customer tested with pre-provided test videos and confirmed the behavior.

## UAT-008: Share Analysis Progress with a Friend

**Status:** Active

**User goal:** Share a completed basketball shooting analysis result with a connected friend.

**Preconditions:** User is logged in. User has at least one completed analysis result. User has at least one accepted friend in their "Friends" list.

**Step-by-step instructions:**
1. Navigate to the analysis history or results page.
2. Select a completed analysis result to view its details.
3. Locate and click the "Share with Friend" or "Share Progress" button.
4. Select a friend from the provided list of connected friends.
5. Confirm the sharing action.
6. Log in as the selected friend and verify they received the shared progress.

**Expected outcome:** The analysis result is successfully shared. The current user sees a success confirmation. The selected friend can view the shared analysis report and progress in their own dashboard or notifications.

**Execution result (Sprint 3 UAT):** Passed — customer confirmed they could successfully select a friend and share their progress, and the friend correctly received and viewed the shared analysis report.

**Execution result (Sprint 4 UAT):** Passed — customer confirmed the hover-to-share functionality works correctly via the Friends tab. The "Share" button appears when hovering over a friend's name and sharing works as expected.
