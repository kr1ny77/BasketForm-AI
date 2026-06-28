#!/usr/bin/env python3
"""
Basketball shooting form analyzer for BasketForm-AI backend.

Reads a video file, extracts pose keypoints with MediaPipe,
scores the player's shooting form, generates an annotated output video,
and writes a JSON result to stdout.

Usage:
  python3 analyze_video.py <video_path> --user_id <uid> --video_id <vid>
  python3 analyze_video.py <video_path>  # fallback, no output video
"""

import argparse
import json
import math
import sys
import os

import cv2
import numpy as np
import mediapipe as mp

mp_holistic = mp.solutions.holistic
mp_drawing = mp.solutions.drawing_utils

L_SHOULDER, L_ELBOW, L_WRIST = 11, 13, 15
R_SHOULDER, R_ELBOW, R_WRIST = 12, 14, 16
L_HIP, L_KNEE, L_ANKLE = 23, 25, 27
R_HIP, R_KNEE, R_ANKLE = 24, 26, 28

POSE_INDICES = [
    0, 11, 12, 13, 14, 15, 16, 23, 24, 25, 26, 27,
]

PHASE_NAMES = ["Stance", "Arm Angle", "Release", "Follow-through"]

COLORS = {
    "keypoint": (255, 106, 0),
    "skeleton": (0, 200, 255),
    "text": (255, 255, 255),
    "bg": (0, 0, 0),
    "good": (46, 213, 115),
    "warn": (255, 165, 2),
    "bad": (255, 71, 87),
}


def angle(a, b, c):
    a, b, c = np.array(a, dtype=float), np.array(b, dtype=float), np.array(c, dtype=float)
    ba, bc = a - b, c - b
    cos = np.dot(ba, bc) / (np.linalg.norm(ba) * np.linalg.norm(bc) + 1e-9)
    return float(np.degrees(np.arccos(np.clip(cos, -1.0, 1.0))))


def clamp(v, lo, hi):
    return max(lo, min(hi, v))


def score_angle(value, ideal, tolerance):
    dist = abs(value - ideal)
    return clamp(int(100 * max(0, 1 - dist / tolerance)), 0, 100)


def score_color(s):
    if s >= 80:
        return COLORS["good"]
    if s >= 50:
        return COLORS["warn"]
    return COLORS["bad"]


def rating(s):
    if s >= 80:
        return "Excellent"
    if s >= 60:
        return "Good"
    if s >= 40:
        return "Needs Work"
    return "Poor"


def analyze_frame(landmarks, w, h):
    if not landmarks:
        return None, None
    lm = landmarks.landmark
    pts = {}
    for idx in POSE_INDICES:
        if idx < len(lm):
            pts[idx] = (lm[idx].x * w, lm[idx].y * h)
    required = [R_SHOULDER, R_ELBOW, R_WRIST, R_HIP, R_KNEE, R_ANKLE]
    if not all(idx in pts for idx in required):
        return None, {i: pts.get(i, (0, 0)) for i in POSE_INDICES if i in pts}
    angles = {
        "elbow": angle(pts[R_SHOULDER], pts[R_ELBOW], pts[R_WRIST]),
        "knee": angle(pts[R_HIP], pts[R_KNEE], pts[R_ANKLE]),
        "shoulder": angle(pts[R_ELBOW], pts[R_SHOULDER], pts[R_HIP]),
        "torso": angle(pts[R_SHOULDER], pts[R_HIP], pts[R_KNEE]),
    }
    elbow = pts[R_ELBOW]
    vertical_up = (elbow[0], elbow[1] - 0.2)
    angles["forearm_vertical"] = angle(pts[R_WRIST], elbow, vertical_up)
    pose_px = {i: pts.get(i, (0, 0)) for i in POSE_INDICES if i in pts}
    return angles, pose_px


def draw_landmarks_on_frame(frame, pose_landmarks, w, h):
    overlay = frame.copy()
    if pose_landmarks and pose_landmarks.landmark:
        lm = pose_landmarks.landmark
        connections = [
            (R_SHOULDER, R_ELBOW), (R_ELBOW, R_WRIST),
            (L_SHOULDER, L_ELBOW), (L_ELBOW, L_WRIST),
            (R_SHOULDER, L_SHOULDER),
            (R_SHOULDER, R_HIP), (L_SHOULDER, L_HIP),
            (R_HIP, L_HIP),
            (R_HIP, R_KNEE), (R_KNEE, R_ANKLE),
            (L_HIP, L_KNEE), (L_KNEE, L_ANKLE),
        ]
        for a, b in connections:
            if a < len(lm) and b < len(lm):
                pt1 = (int(lm[a].x * w), int(lm[a].y * h))
                pt2 = (int(lm[b].x * w), int(lm[b].y * h))
                cv2.line(overlay, pt1, pt2, COLORS["skeleton"], 2, cv2.LINE_AA)
        for idx in POSE_INDICES:
            if idx < len(lm):
                cx, cy = int(lm[idx].x * w), int(lm[idx].y * h)
                cv2.circle(overlay, (cx, cy), 5, COLORS["keypoint"], -1, cv2.LINE_AA)
                cv2.circle(overlay, (cx, cy), 5, COLORS["text"], 1, cv2.LINE_AA)
    alpha = 0.7
    cv2.addWeighted(overlay, alpha, frame, 1 - alpha, 0, frame)


def draw_hud(frame, phase_scores, overall, frame_num, total_frames):
    h, w = frame.shape[:2]
    bar_w = 120
    bar_h = 8
    margin = 15
    x0 = w - bar_w - margin
    y0 = margin

    progress = frame_num / max(total_frames, 1)
    cv2.rectangle(frame, (margin, h - 25), (margin + int((w - 2 * margin) * progress), h - 18),
                  COLORS["keypoint"], -1)
    cv2.rectangle(frame, (margin, h - 25), (w - margin, h - 18), (60, 60, 60), 1)

    for i, name in enumerate(PHASE_NAMES):
        y = y0 + i * 28
        sc = phase_scores[i] if i < len(phase_scores) else 0
        cv2.putText(frame, name, (x0, y + 12), cv2.FONT_HERSHEY_SIMPLEX, 0.38,
                    COLORS["text"], 1, cv2.LINE_AA)
        cv2.rectangle(frame, (x0 + 70, y + 4), (x0 + 70 + bar_w, y + 4 + bar_h),
                      (40, 40, 40), -1)
        cv2.rectangle(frame, (x0 + 70, y + 4), (x0 + 70 + int(bar_w * sc / 100), y + 4 + bar_h),
                      score_color(sc), -1)

    label = f"Overall: {overall}"
    cv2.putText(frame, label, (x0, y0 + len(PHASE_NAMES) * 28 + 22),
                cv2.FONT_HERSHEY_SIMPLEX, 0.5, COLORS["keypoint"], 2, cv2.LINE_AA)


def build_feedback(stance_s, arm_s, release_s, follow_s, details):
    parts = [
        f"STANCE ({stance_s}/100 \u2014 {rating(stance_s)})",
        details.get("stance", ""),
        "",
        f"ARM ANGLE ({arm_s}/100 \u2014 {rating(arm_s)})",
        details.get("arm_angle", ""),
        "",
        f"RELEASE POINT ({release_s}/100 \u2014 {rating(release_s)})",
        details.get("release", ""),
        "",
        f"FOLLOW-THROUGH ({follow_s}/100 \u2014 {rating(follow_s)})",
        details.get("follow_through", ""),
    ]
    return "\n".join(parts)


def process_video(video_path, output_video_path=None):
    cap = cv2.VideoCapture(video_path)
    if not cap.isOpened():
        return None

    w = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
    h = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))
    fps = cap.get(cv2.CAP_PROP_FPS) or 30.0
    total = int(cap.get(cv2.CAP_PROP_FRAME_COUNT))
    if total == 0:
        cap.release()
        return None

    all_angles = []
    all_pose = []
    all_landmarks = []
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
            all_landmarks.append(results.pose_landmarks)

    cap.release()

    if not all_angles:
        return None

    avg = {k: np.mean([a[k] for a in all_angles]) for k in all_angles[0]}

    stance = score_angle(avg.get("knee", 180), 150, 40)
    arm = score_angle(avg.get("elbow", 90), 90, 40)
    release = score_angle(avg.get("forearm_vertical", 20), 0, 45)
    follow = score_angle(avg.get("shoulder", 45), 45, 40)
    overall = clamp(int(0.30 * stance + 0.30 * arm + 0.20 * release + 0.20 * follow), 0, 100)

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

    phase_scores = [stance, arm, release, follow]
    phase_feedbacks = [
        details.get("stance", ""),
        details.get("arm_angle", ""),
        details.get("release", ""),
        details.get("follow_through", ""),
    ]

    if output_video_path:
        os.makedirs(os.path.dirname(output_video_path), exist_ok=True)
        fourcc = cv2.VideoWriter_fourcc(*"mp4v")
        writer = cv2.VideoWriter(output_video_path, fourcc, fps, (w, h))
        cap = cv2.VideoCapture(video_path)
        idx = 0
        while cap.isOpened():
            ok, frame = cap.read()
            if not ok:
                break
            lm = all_landmarks[idx] if idx < len(all_landmarks) else None
            draw_landmarks_on_frame(frame, lm, w, h)
            draw_hud(frame, phase_scores, overall, idx, total)
            writer.write(frame)
            idx += 1
        cap.release()
        writer.release()

    last_pose = all_pose[-1] if all_pose else {}
    pose_data = []
    for pi in POSE_INDICES:
        if pi in last_pose:
            x, y = last_pose[pi]
            pose_data.append({"x": round(x * 100, 2), "y": round(y * 100, 2)})
        else:
            pose_data.append({"x": 0, "y": 0})

    feedback = build_feedback(stance, arm, release, follow, details)

    return {
        "score": overall,
        "scores": phase_scores,
        "feedback": feedback,
        "pose_data": pose_data,
        "phases": [
            {"phase_name": PHASE_NAMES[i], "score": phase_scores[i], "feedback": phase_feedbacks[i]}
            for i in range(4)
        ],
        "output_video": output_video_path or "",
    }


def main():
    parser = argparse.ArgumentParser(description="Basketball shooting form analyzer")
    parser.add_argument("video_path", help="Path to input video")
    parser.add_argument("--user_id", default="", help="User ID for output path")
    parser.add_argument("--video_id", default="", help="Video ID for output path")
    args = parser.parse_args()

    if not os.path.isfile(args.video_path):
        print(json.dumps({"error": f"file not found: {args.video_path}"}))
        sys.exit(1)

    output_path = ""
    if args.user_id and args.video_id:
        output_path = os.path.join("results", args.user_id, args.video_id, "output.mp4")

    result = process_video(args.video_path, output_path or None)
    if result is None:
        print(json.dumps({"error": "could not analyze video (no pose detected)"}))
        sys.exit(1)

    print(json.dumps(result))


if __name__ == "__main__":
    main()
