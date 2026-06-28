### User Acceptance Tests (UATs)

---

### UAT-001: Upload Video and Receive Biomechanical Form Analysis
- Scenario status: Active
- User goal: A basketball player wants to upload a video of their shot to receive an automated biomechanical evaluation and personalized coaching feedback.
- Preconditions: 
  - The user has access to the live application at http://80.74.30.14.
  - The user has a supported video file (MP4, MOV, or AVI) featuring a clear, full-body view of a basketball shot.
- Step-by-step instructions:
  1. Navigate to the BasketForm-AI homepage.
  2. Locate the upload section and either drag-and-drop the video file onto the designated area or use the file picker to select the file.
  3. Wait for the system to process the video and extract biomechanical keypoints.
  4. Review the analysis results displayed on the screen once processing is complete.
- Expected outcome:
  - Step 2: The file is successfully attached. If an unsupported format is selected, the system rejects it and displays a clear error message.
  - Step 3: A real-time processing indicator is visible while the AI analyzes the video.
  - Step 4 (Overall): The results page displays the uploaded video alongside technique scores (0–100) for stance, arm angle, release point, and follow-through. It also provides exactly three actionable, AI-generated recommendations for improvement.
- Assignment-specific execution results: Expected outcome is achieved.
- Customer comments or observed issues after execution: Everything works.
- Resulting PBIs or issues after execution: No issues.

---

### UAT-002: Export Analysis Report as PDF for Offline Review
- Scenario status: Active
- User goal: A player or coach wants to download the completed shot analysis as a PDF document to review offline, share with a training staff, or keep for their personal records.
- Preconditions: 
  - The user has successfully completed a video upload (as in UAT-001).
  - The final analysis results are currently displayed on the screen.
- Step-by-step instructions:
  1. Review the analysis results and the AI-generated feedback on the screen.
  2. Locate and click the "Download PDF" or "Export Report" button.
  3. Locate the downloaded file on the local device and open it.
- Expected outcome:
  - Step 2: The browser successfully initiates a file download without errors or page reloads, utilizing the integrated jsPDF functionality.
  - Step 3 (Overall): The downloaded PDF successfully opens and accurately mirrors the web UI. It must include the player's numerical scores for all four categories (stance, arm angle, release point, follow-through) and the three personalized AI takeaways.
- Assignment-specific execution results: Expected outcome is achieved.
- Customer comments or observed issues after execution: Everything works.
- Resulting PBIs or issues after execution: No issues.

---

### UAT-003: Evaluate Interactive UI and Canvas Animation Responsiveness
- Scenario status: Active
- User goal: A user expects the application's interface to be visually engaging and responsive to user interaction as advertised in the application's design specifications.
- Preconditions: 
  - The user is using a modern web browser with hardware acceleration and JavaScript enabled.
  - The user has a functioning mouse or trackpad.
- Step-by-step instructions:
  1. Navigate to the BasketForm-AI homepage.
  2. Observe the basketball-themed background animation (e.g., floating objects, orange waves).
  3. Move the mouse cursor dynamically across the screen over the background elements.
  4. Click through the main UI elements (buttons, file pickers) to ensure they are not obscured or broken by the background animation.
- Expected outcome:
  - Step 2: The background animation loads smoothly, displaying the specified themed elements (35 objects, basketball theme).
  - Step 3: The background animation dynamically reacts to the mouse movements (mouse-reactive) without causing severe frame drops or freezing the browser.
  - Step 4 (Overall): The interactive canvas animation remains strictly in the background. All foreground UI elements (buttons, upload drop-zones, text) remain perfectly clickable, legible, and unaffected by the animation layer.
- Assignment-specific execution results: Expected outcome is achieved.
- Customer comments or observed issues after execution: Everything works. Customer is surprised about the animations.
- Resulting PBIs or issues after execution: No issues.

---

## Week 4 UAT Execution Results – Public Summary

**Sprint / Week:** Week 4
**Execution Date:** 2025-06-28
**Tester:** Customer (end user)
**Scenarios Executed:** 3

---

### UAT Scenarios Passed

All three active UAT scenarios were executed by the customer and passed:

- **UAT-001 – Upload Video and Receive Biomechanical Form Analysis:** The customer successfully uploaded a basketball shot video and received a complete biomechanical analysis, including technique scores for all four categories (stance, arm angle, release point, follow-through) and three AI-generated improvement recommendations. Scenario **passed**.
- - **UAT-002 – Export Analysis Report as PDF for Offline Review:** The customer triggered a PDF export from the results screen. The file downloaded without errors and accurately reflected all scores and AI takeaways shown on screen. Scenario **passed**.
  - - **UAT-003 – Evaluate Interactive UI and Canvas Animation Responsiveness:** The customer navigated the homepage, observed the basketball-themed background animation, and interacted with mouse movements and UI controls. The animation responded smoothly, remained in the background, and did not interfere with any foreground elements. Scenario **passed**.
   
    - ---

    ### UAT Scenarios Failed or Needing Product Changes

    None. All three scenarios met their expected outcomes. No product changes or bug fixes are required as a result of this testing cycle.

    ---

    ### Most Important Feedback Points Received

    - **Core functionality is solid:** Video processing, biomechanical analysis, PDF export, and UI interactivity all work seamlessly end-to-end.
    - - **Positive reaction to animations:** The interactive basketball-themed background animations were a standout highlight. The customer expressed genuine surprise at the quality and responsiveness of the visual effects, describing them as visually engaging and well-executed.
     
      - ---

      ### Resulting PBIs or Issues

      None. No bugs, defects, or new Product Backlog Items were generated from this Week 4 UAT cycle.
