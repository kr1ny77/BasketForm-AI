#!/usr/bin/env python3
"""
Mock video processing script for BasketForm-AI.
Called via exec.Command from Go backend.
Usage: python3 process_video.py <video_path>
Output: JSON to stdout with score, feedback, pose_data.
"""

import sys
import json
import random
import time

FEEDBACKS = [
    "Good release point. Work on keeping your elbow aligned with the basket.",
    "Strong stance. Try to extend your follow-through higher for better arc.",
    "Nice arm angle. Focus on squaring your shoulders to the basket before release.",
    "Consistent base. Keep your shooting hand under the ball for a cleaner release.",
    "Good rhythm. Work on a quicker release without sacrificing form.",
    "Solid footwork. Try to jump straight up rather than drifting forward.",
]


def generate_pose():
    skeleton = [
        (50, 8), (50, 22), (50, 40),
        (32, 52), (68, 52), (20, 62), (80, 62),
        (50, 40), (38, 68), (62, 68), (36, 88), (64, 88),
    ]
    return [{"x": x + random.uniform(-2, 2), "y": y + random.uniform(-1.5, 1.5)} for x, y in skeleton]


def main():
    if len(sys.argv) < 2:
        print(json.dumps({"error": "usage: process_video.py <video_path>"}))
        sys.exit(1)

    video_path = sys.argv[1]
    time.sleep(3)

    score = random.randint(40, 90)
    scores = [
        max(0, min(100, score + random.randint(-10, 10))),
        max(0, min(100, score + random.randint(-15, 15))),
        max(0, min(100, score + random.randint(-5, 5))),
        max(0, min(100, score + random.randint(-12, 12))),
    ]

    result = {
        "score": score,
        "feedback": random.choice(FEEDBACKS),
        "pose_data": generate_pose(),
        "scores": scores,
    }

    print(json.dumps(result))


if __name__ == "__main__":
    main()
