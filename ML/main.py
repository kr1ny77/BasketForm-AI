import cv2
import argparse
import sys
import os
from collections import deque
import json
from feedback_generator import FeedbackGenerator
from ball_tracker import BallTracker
from pose_tracker import PoseTracker
from shot_analyzer import ShotPhaseStateMachine
from custom_feedback import CustomFeedbackGenerator
import mediapipe as mp

mp_drawing = mp.solutions.drawing_utils
mp_drawing_styles = mp.solutions.drawing_styles
mp_holistic = mp.solutions.holistic


def draw_angle_info_panel(image, angles, current_phase, phase_color, ball_detected, position=(10, 30)):
    """Draws an info panel with angles and phase in the corner of the video."""
    x, y = position
    cv2.rectangle(image, (x - 5, y - 25), (x + 320, y + 200), (0, 0, 0), -1)
    cv2.rectangle(image, (x - 5, y - 25), (x + 320, y + 200), (255, 255, 255), 2)
    cv2.putText(image, "Shooting Form Analysis", (x, y),
                cv2.FONT_HERSHEY_SIMPLEX, 0.6, (255, 255, 255), 2)

    cv2.putText(image, f"PHASE: {current_phase}", (x, y + 30),
                cv2.FONT_HERSHEY_SIMPLEX, 0.8, phase_color, 3)

    # Ball status
    ball_status = "TRACKED" if ball_detected else "LOST"
    ball_color = (0, 255, 0) if ball_detected else (0, 0, 255)
    cv2.putText(image, f"Ball: {ball_status}", (x, y + 60),
                cv2.FONT_HERSHEY_SIMPLEX, 0.5, ball_color, 2)

    y_offset = y + 85
    angle_descriptions = {
        'elbow': ('Elbow Angle', 70, 110),
        'knee': ('Knee Angle', 90, 130),
        'forearm_vertical': ('Forearm vs Vert', 0, 30),
        'torso': ('Torso Lean', 150, 180),
        'shoulder': ('Shoulder Angle', 20, 60)
    }

    for angle_name, (desc, min_val, max_val) in angle_descriptions.items():
        if angle_name in angles:
            value = angles[angle_name]['value']
            color = (0, 255, 255) if min_val <= value <= max_val else (0, 0, 255)
            text = f"{desc}: {value:.1f}"
            cv2.putText(image, text, (x, y_offset), cv2.FONT_HERSHEY_SIMPLEX, 0.5, color, 2)
            y_offset += 25


def process_video(input_path, output_path, lang="en"):
    """Main processing loop."""
    if not os.path.exists(input_path):
        print(f"ERROR: Input file not found: {input_path}", file=sys.stderr)
        return False

    cap = cv2.VideoCapture(input_path)
    if not cap.isOpened():
        print(f"ERROR: Could not open video file: {input_path}", file=sys.stderr)
        return False

    frame_width = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
    frame_height = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))
    fps = cap.get(cv2.CAP_PROP_FPS)
    total_frames = int(cap.get(cv2.CAP_PROP_FRAME_COUNT))

    print(f"Input video: {frame_width}x{frame_height} @ {fps} FPS, {total_frames} frames")

    print("Initializing models...")
    ball_tracker = BallTracker(model_path='best.pt')
    pose_tracker = PoseTracker()
    phase_machine = ShotPhaseStateMachine()

    ball_trajectory = deque(maxlen=30)
    frame_count = 0

    try:
        print("Processing video...")
        while cap.isOpened():
            success, image = cap.read()
            if not success:
                break

            frame_count += 1
            if frame_count % 30 == 0:
                progress = (frame_count / total_frames) * 100
                print(f"Progress: {progress:.1f}% ({frame_count}/{total_frames} frames)")

            ball_center, ball_bbox = ball_tracker.detect(image)
            if ball_center:
                ball_trajectory.append(ball_center)

            results = pose_tracker.process(image)
            angles, wrist_center = pose_tracker.analyze_shooting_form(
                results.pose_landmarks, frame_width, frame_height
            )

            if 'knee' in angles and 'elbow' in angles:
                torso_angle = angles.get('torso', {}).get('value', 180)
                forearm_angle = angles.get('forearm_vertical', {}).get('value', 0)

                phase_machine.update(
                    angles['knee']['value'],
                    angles['elbow']['value'],
                    ball_center,
                    wrist_center,
                    torso_angle,
                    forearm_angle
                )

    except Exception as e:
        print(f"ERROR: Exception during processing: {str(e)}", file=sys.stderr)
        cap.release()
        return False

    finally:
        phase_machine.finalize()
        cap.release()
        try:
            cv2.destroyAllWindows()
        except Exception:
            pass

    print(f"Processing complete! Processed {frame_count} frames")

    final_scores = phase_machine.scores
    final_metrics = {
        "min_knee_dip": phase_machine.min_knee_dip,
        "torso_ascent": phase_machine.torso_ascent,
        "elbow_release": phase_machine.elbow_release,
        "forearm_release": phase_machine.forearm_release
    }

    ai_feedback = ""
    try:
        generator = CustomFeedbackGenerator()
        ai_feedback = generator.generate(final_scores, final_metrics, lang)
        print("--- AI COACH FEEDBACK ---")
        print(ai_feedback)
        print("--------------------------")
    except Exception as e:
        print(f"WARNING: Could not generate feedback. Error: {e}", file=sys.stderr)
        ai_feedback = "Scores received. Work on your shooting technique."

    report = {
        "scores": final_scores,
        "metrics": {k: float(v) for k, v in final_metrics.items()},
        "ai_feedback": ai_feedback
    }

    report_path = os.path.splitext(output_path)[0] + ".json" if output_path else os.path.splitext(input_path)[0] + "_report.json"
    with open(report_path, 'w', encoding='utf-8') as f:
        json.dump(report, f, ensure_ascii=False, indent=2)

    print(f"Report saved to {report_path}")
    return True


def main():
    """Main entry point with command-line argument parsing."""
    parser = argparse.ArgumentParser(
        description="Basketball Shot Form Analyzer - Analyzes shooting mechanics from video using Pose and Ball tracking",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
python main.py input.mp4 output.mp4
python main.py input.mp4
python main.py /path/to/video.mp4 /path/to/result.mp4
"""
    )
    parser.add_argument('input', type=str, help='Path to input video file')
    parser.add_argument('output', type=str, nargs='?', default=None, help='Path to output report JSON (optional)')
    parser.add_argument('--lang', type=str, default='en', choices=['en', 'ru'], help='Feedback language')

    args = parser.parse_args()

    valid_extensions = {'.mp4', '.avi', '.mov', '.mkv'}

    if os.path.splitext(args.input)[1].lower() not in valid_extensions:
        print(f"ERROR: Invalid input file extension.", file=sys.stderr)
        sys.exit(1)

    if args.output and os.path.splitext(args.output)[1].lower() not in valid_extensions:
        pass

    success = process_video(args.input, args.output, args.lang)
    sys.exit(0 if success else 1)


if __name__ == "__main__":
    main()