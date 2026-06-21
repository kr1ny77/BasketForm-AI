#!/usr/bin/env python3
"""
Basketball shooting form analyzer for BasketForm-AI backend.

Reads a video file, extracts pose keypoints with MediaPipe,
scores the player's shooting form, and writes a JSON result to stdout.

Usage:  python3 analyze_video.py <video_path>
Output: JSON to stdout with score, feedback, pose_data, scores.
"""

import json
import math
import sys
import os

import cv2
import numpy as np
import mediapipe as mp

mp_holistic = mp.solutions.holistic

# Landmark indices (MediaPipe Pose)
L_SHOULDER, L_ELBOW, L_WRIST = 11, 13, 15
R_SHOULDER, R_ELBOW, R_WRIST = 12, 14, 16
L_HIP, L_KNEE, L_ANKLE = 23, 25, 27
R_HIP, R_KNEE, R_ANKLE = 24, 26, 28

# Key landmarks to return as pose data (12 points)
POSE_INDICES = [
    0,   # nose
    11,  # left shoulder
    12,  # right shoulder
    13,  # left elbow
    14,  # right elbow
    15,  # left wrist
    16,  # right wrist
    23,  # left hip
    24,  # right hip
    25,  # left knee
    26,  # right knee
    27,  # left ankle
]


def angle(a, b, c):
    """Angle at vertex b formed by points a-b-c, in degrees."""
    a, b, c = np.array(a, dtype=float), np.array(b, dtype=float), np.array(c, dtype=float)
    ba, bc = a - b, c - b
    cos = np.dot(ba, bc) / (np.linalg.norm(ba) * np.linalg.norm(bc) + 1e-9)
    return float(np.degrees(np.arccos(np.clip(cos, -1.0, 1.0))))


def lm_px(landmarks, idx, w, h):
    """Return (x, y) pixel coords for a landmark, or None."""
    if idx >= len(landmarks):
        return None
    l = landmarks[idx]
    return (l.x * w, l.y * h)


def clamp(v, lo, hi):
    return max(lo, min(hi, v))


def score_angle(value, ideal, tolerance):
    """Score 0-100: 100 when value == ideal, decays with distance."""
    dist = abs(value - ideal)
    return clamp(int(100 * max(0, 1 - dist / tolerance)), 0, 100)


def build_feedback(stance_s, arm_s, release_s, follow_s, details):
    """Build a human-readable feedback string."""
    def rating(s):
        if s >= 80: return "Excellent"
        if s >= 60: return "Good"
        if s >= 40: return "Needs Work"
        return "Poor"

    parts = [
        f"STANCE ({stance_s}/100 — {rating(stance_s)})",
        details.get("stance", ""),
        "",
        f"ARM ANGLE ({arm_s}/100 — {rating(arm_s)})",
        details.get("arm_angle", ""),
        "",
        f"RELEASE POINT ({release_s}/100 — {rating(release_s)})",
        details.get("release", ""),
        "",
        f"FOLLOW-THROUGH ({follow_s}/100 — {rating(follow_s)})",
        details.get("follow_through", ""),
    ]
    return "\n".join(parts)


def analyze_frame(landmarks, w, h):
    """Analyze a single frame's pose landmarks.

    Returns (angles_dict, pose_pixels) or (None, None) if insufficient data.
    """
    if not landmarks:
        return None, None

    lm = landmarks.landmark
    pts = {}
    for idx in POSE_INDICES:
        if idx < len(lm):
            pts[idx] = (lm[idx].x * w, lm[idx].y * h)

    # Need at least right-side landmarks for a right-handed shooter
    required = [R_SHOULDER, R_ELBOW, R_WRIST, R_HIP, R_KNEE, R_ANKLE]
    if not all(idx in pts for idx in required):
        return None, {i: pts.get(i, (0, 0)) for i in POSE_INDICES if i in pts}

    angles = {
        "elbow": angle(pts[R_SHOULDER], pts[R_ELBOW], pts[R_WRIST]),
        "knee": angle(pts[R_HIP], pts[R_KNEE], pts[R_ANKLE]),
        "shoulder": angle(pts[R_ELBOW], pts[R_SHOULDER], pts[R_HIP]),
        "torso": angle(pts[R_SHOULDER], pts[R_HIP], pts[R_KNEE]),
    }

    # Forearm vs vertical
    elbow = pts[R_ELBOW]
    vertical_up = (elbow[0], elbow[1] - 0.2)
    angles["forearm_vertical"] = angle(pts[R_WRIST], elbow, vertical_up)

    pose_px = {i: pts.get(i, (0, 0)) for i in POSE_INDICES if i in pts}
    return angles, pose_px


def process_video(video_path):
    """Process a video and return analysis results dict."""
    cap = cv2.VideoCapture(video_path)
    if not cap.isOpened():
        return None

    w = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
    h = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))
    total = int(cap.get(cv2.CAP_PROP_FRAME_COUNT))
    if total == 0:
        cap.release()
        return None

    all_angles = []
    all_pose = []
    frame_count = 0

    with mp_holistic.Holistic(min_detection_confidence=0.5, min_tracking_confidence=0.5) as holistic:
        while cap.isOpened():
            ok, frame = cap.read()
            if not ok:
                break
            frame_count += 1

            rgb = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)
            rgb.flags.writeable = False
            results = holistic.process(rgb)

            angles, pose_px = analyze_frame(results.pose_landmarks, w, h)
            if angles:
                all_angles.append(angles)
            if pose_px:
                all_pose.append(pose_px)

    cap.release()

    if not all_angles:
        return None

    # Average angles across all detected frames
    avg = {k: np.mean([a[k] for a in all_angles]) for k in all_angles[0]}

    # Use the last frame's pose for display
    last_pose = all_pose[-1] if all_pose else {}

    # --- Score each dimension ---
    # Stance: knee bend ideal ~150 (slightly bent), tolerance 40
    stance = score_angle(avg.get("knee", 180), 150, 40)

    # Arm angle: elbow ~90 is ideal, tolerance 40
    arm = score_angle(avg.get("elbow", 90), 90, 40)

    # Release: forearm near vertical is ideal (0° = straight up), tolerance 45
    release = score_angle(avg.get("forearm_vertical", 20), 0, 45)

    # Follow-through: shoulder openness ~45°, tolerance 40
    follow = score_angle(avg.get("shoulder", 45), 45, 40)

    overall = clamp(int(0.30 * stance + 0.30 * arm + 0.20 * release + 0.20 * follow), 0, 100)

    # --- Build feedback details ---
    details = {}

    knee_v = avg.get("knee", 180)
    if knee_v > 165:
        details["stance"] = "Your knees are nearly straight. Bend them about 15-20 degrees before the shot to generate power from your legs."
    elif knee_v < 120:
        details["stance"] = "You are bending your knees too deeply. A moderate bend (15-20 degrees) provides better balance and a quicker release."
    else:
        details["stance"] = "Good knee bend. Your stance provides a solid base for power generation."

    elbow_v = avg.get("elbow", 90)
    if elbow_v > 110:
        details["arm_angle"] = "Your elbow is flaring out (open angle). Tuck your elbow in to create a tighter L-shape (~90 degrees) for better accuracy."
    elif elbow_v < 70:
        details["arm_angle"] = "Your elbow angle is too closed. Open it slightly toward 90 degrees to get a more natural shooting motion."
    else:
        details["arm_angle"] = "Your arm angle is close to the ideal 90-degree L-shape. This gives you good control over the ball's trajectory."

    fv = avg.get("forearm_vertical", 20)
    if fv > 30:
        details["release"] = "Your forearm is angled significantly away from vertical at release. Try to release the ball with your arm extended more directly upward for a higher arc."
    else:
        details["release"] = "Good forearm alignment at release. Your arm is near vertical, which helps produce a clean arc toward the basket."

    shoulder_v = avg.get("shoulder", 45)
    if shoulder_v > 65:
        details["follow_through"] = "Your shoulder is very open at follow-through. Keep your shooting arm finishing high with your wrist snapped forward toward the rim."
    elif shoulder_v < 25:
        details["follow_through"] = "Your follow-through is cut short. Extend your arm fully upward and snap your wrist as if reaching into a cookie jar on a high shelf."
    else:
        details["follow_through"] = "Solid follow-through. Your arm extends well toward the basket. Focus on holding the finish for a moment after release."

    # --- Build pose data (12 keypoints normalized to 0-100 range) ---
    pose_data = []
    for idx in POSE_INDICES:
        if idx in last_pose:
            x, y = last_pose[idx]
            pose_data.append({"x": round(x * 100, 2), "y": round(y * 100, 2)})
        else:
            pose_data.append({"x": 0, "y": 0})

    feedback = build_feedback(stance, arm, release, follow, details)

    return {
        "score": overall,
        "scores": [stance, arm, release, follow],
        "feedback": feedback,
        "pose_data": pose_data,
    }


def main():
    if len(sys.argv) < 2:
        print(json.dumps({"error": "usage: analyze_video.py <video_path>"}))
        sys.exit(1)

    video_path = sys.argv[1]
    if not os.path.isfile(video_path):
        print(json.dumps({"error": f"file not found: {video_path}"}))
        sys.exit(1)

    result = process_video(video_path)
    if result is None:
        print(json.dumps({"error": "could not analyze video (no pose detected)"}))
        sys.exit(1)

    print(json.dumps(result))


if __name__ == "__main__":
    main()
