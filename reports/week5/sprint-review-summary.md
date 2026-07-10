# Meeting Notes: Sprint Review, Planning, and UAT

## Meeting Details
- **Participants**: ML Engineer, Customer, Full-stack Developer
- **Agenda**: Sprint Review, Next Sprint Planning, User Acceptance Testing (UAT)

---

## 1. Sprint Review (Accomplishments)
The team reviewed the completed tasks from the current sprint. All planned items were successfully implemented.

- **Friend System**: Successfully implemented and fully functional. 
- **ML System Redesign**: 
  - Transitioned from a locally hosted `[redacted]` model (which caused server overload and slow startup times) to an API-based `[redacted]` model.
  - **Result**: Significantly faster feedback generation and no server overload. The core algorithm remains unchanged.
- **Localization**: Fixed identified errors in the Russian translation.

## 2. Sprint Planning (Next Sprint)
*Note: This is the second-to-last sprint before the final defense (approx. 2 weeks remaining).*

### Feature Discussions & Decisions
- **Camera Positioning Recommendations**: Discussed adding tips for users on how to better position their cameras when recording throws.
- **Multi-Throw Video Upload & Auto-Segmentation (APPROVED)**: 
  - **Concept**: Allow users to upload a single continuous video of multiple throws. The app will automatically detect and cut out the throwing segments for individual analysis.
  - **Technical Approach**: The Customer suggested using the moment the ball leaves the person's hand as a "trigger" to capture and process the preceding frames from a queue.
  - **Decision**: Approved for the next sprint. The Customer advised managing time carefully to ensure this doesn't compromise the stability of the currently finished product.
- **Direct Video-to-Video Comparison (REJECTED)**:
  - **Concept**: Comparing two specific videos (e.g., today's throw vs. yesterday's throw) to see exact feature changes.
  - **Decision**: Dropped. The Full-stack Developer clarified that the existing progress system already tracks and displays specific criteria (throw, foot position, arm position, etc.) using a points-based system, making this feature redundant.

## 3. User Acceptance Testing (UAT)
The team conducted UAT for the newly implemented **Friends Window**. Both tests were executed and passed successfully.

- **Test 1: Sending a Friend Request**
  - **Action**: Customer navigated to the Friends page, entered the provided nickname in the search field, and clicked "Search".
  - **Result**: **Passed**. Notification confirmed the request was successfully sent.
- **Test 2: Accepting a Friend Request**
  - **Action**: ML Engineer sent a friend request. Customer navigated to the "Pending requests" section, refreshed the page, and clicked "Accept".
  - **Result**: **Passed**. The ML Engineer's profile successfully appeared in the Customer's "My Friends" section.

## 4. Action Items & Next Steps
- **Development Team**: Implement the multi-throw video upload with automatic segmentation (using the ball-release trigger method) for the upcoming sprint.
- **Development Team**: Ensure thorough testing of the current build to prevent regressions while adding the new feature.
- **All Participants**: Reconvene at the end of the next sprint for the final review and project defense.
