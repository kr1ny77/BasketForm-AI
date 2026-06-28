# User Acceptance Tests

End-user-facing scenarios for BasketForm-AI. These are executed by the customer during Sprint Review.

## UAT-001: Registration and Login

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

## UAT-002: Video Upload and Analysis Result

**Status:** Active

**User goal:** Upload a basketball shooting video and view the analysis result with phases and video.

**Preconditions:** User is logged in.

**Step-by-step instructions:**
1. Navigate to `/upload`.
2. Drag and drop or select a basketball video (MP4, MOV, AVI, max 100 MB).
3. Wait for upload to complete and redirect to `/progress`.
4. Wait for processing to complete.
5. Navigate to `/results`.
6. Click on the result card.
7. Verify the modal shows: overall score, phase analysis (Stance, Arm Angle, Release, Follow-through), feedback, radar chart, and output video (if available).

**Expected outcome:** Video is uploaded, processed, and the result page shows complete analysis with phases and annotated video.

## UAT-003: Share Result with Friend

**Status:** Active

**User goal:** Share an analysis result with a friend and verify the friend can view it.

**Preconditions:** User is logged in and has at least one completed analysis result and one accepted friend.

**Step-by-step instructions:**
1. Navigate to `/results`.
2. Click on a result card.
3. In the modal, click "Share".
4. Select a friend from the dropdown.
5. Click "Send".
6. Verify success message.
7. Log in as the friend.
8. Navigate to `/shared`.
9. Verify the shared result appears in the list.
10. Click on the shared result and verify it shows the same score, phases, and feedback.

**Expected outcome:** Result is shared and the friend can view the shared result in read-only mode.

## UAT-004: Friend Request Workflow

**Status:** Active

**User goal:** Find a user by nickname, send a friend request, and have it accepted.

**Preconditions:** Two user accounts exist.

**Step-by-step instructions:**
1. Log in as User A.
2. Navigate to `/friends`.
3. Enter User B's nickname in the search field.
4. Click "Search".
5. Click "Add Friend" next to User B.
6. Verify "Request sent!" message.
7. Log in as User B.
8. Navigate to `/friends`.
9. Verify the pending request from User A appears.
10. Click "Accept".
11. Navigate to `/friends` and verify User A appears in "My Friends".

**Expected outcome:** Friend request is sent, accepted, and both users see each other in their friend list.
