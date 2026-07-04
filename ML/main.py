import cv2
import argparse
import sys
import os
from collections import deque
import json
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


def process_video(input_path, output_path):
    """Main processing loop."""
    if not os.path.exists(input_path):
        print(f"ERROR: Input file not found: {input_path}", file=sys.stderr)
        return False

    output_dir = os.path.dirname(output_path)
    if output_dir and not os.path.exists(output_dir):
        print(f"ERROR: Output directory does not exist: {output_dir}", file=sys.stderr)
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

    fourcc = cv2.VideoWriter_fourcc(*'mp4v')
    out = cv2.VideoWriter(output_path, fourcc, fps, (frame_width, frame_height))

    if not out.isOpened():
        print(f"ERROR: Could not create output video file: {output_path}", file=sys.stderr)
        cap.release()
        return False

    print("Initializing models...")
    # Initialize modular components
    ball_tracker = BallTracker(model_path='best.pt')
    pose_tracker = PoseTracker()
    phase_machine = ShotPhaseStateMachine()

    # For drawing ball trajectory
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

            # 1. Ball Tracking
            ball_center, ball_bbox = ball_tracker.detect(image)
            if ball_center:
                ball_trajectory.append(ball_center)

            # 2. Pose Tracking
            results = pose_tracker.process(image)
            angles, wrist_center = pose_tracker.analyze_shooting_form(
                results.pose_landmarks, frame_width, frame_height
            )

            # 3. Phase Analysis (State Machine)
            current_phase = "IDLE"
            phase_color = (200, 200, 200)

            if 'knee' in angles and 'elbow' in angles:
                torso_angle = angles.get('torso', {}).get('value', 180)
                forearm_angle = angles.get('forearm_vertical', {}).get('value', 0)

                current_phase, phase_color = phase_machine.update(
                    angles['knee']['value'],
                    angles['elbow']['value'],
                    ball_center,
                    wrist_center,
                    torso_angle,
                    forearm_angle
                )

            # --- VISUALIZATION ---
            image_bgr = image.copy()

            # Draw Ball and Trajectory
            if ball_center:
                cv2.circle(image_bgr, ball_center, 8, (0, 165, 255), -1)  # Orange dot
            if ball_bbox:
                x1, y1, x2, y2 = ball_bbox
                cv2.rectangle(image_bgr, (x1, y1), (x2, y2), (0, 165, 255), 2)

            # Draw trajectory line
            for i in range(1, len(ball_trajectory)):
                if ball_trajectory[i - 1] is None or ball_trajectory[i] is None:
                    continue
                cv2.line(image_bgr, ball_trajectory[i - 1], ball_trajectory[i], (0, 165, 255), 2)

            # Draw Skeleton
            if results.pose_landmarks:
                landmark_style = mp_drawing_styles.get_default_pose_landmarks_style()
                for idx in landmark_style:
                    landmark_style[idx].color = phase_color
                connection_style = mp_drawing.DrawingSpec(color=phase_color, thickness=2)

                mp_drawing.draw_landmarks(
                    image_bgr, results.pose_landmarks, mp_holistic.POSE_CONNECTIONS,
                    landmark_drawing_spec=landmark_style,
                    connection_drawing_spec=connection_style)

            # Draw Info Panel
            draw_angle_info_panel(image_bgr, angles, current_phase, phase_color, ball_center is not None)

            out.write(image_bgr)

    except Exception as e:
        print(f"ERROR: Exception during processing: {str(e)}", file=sys.stderr)
        cap.release()
        out.release()
        return False

    finally:
        # Принудительно завершаем анализ и считаем оценки,
        # если видео оборвалось на середине броска
        phase_machine.finalize()

        cap.release()
        out.release()
        cv2.destroyAllWindows()

    print(f"Processing complete! Processed {frame_count} frames")
    print(f"Output saved to: {output_path}")

    final_metrics = {
        "min_knee_dip": phase_machine.min_knee_dip,
        "torso_ascent": phase_machine.torso_ascent,
        "elbow_release": phase_machine.elbow_release,
        "forearm_release": phase_machine.forearm_release,
        "frames_in_follow_through": phase_machine.frames_in_follow_through,
        "elbow_snap": phase_machine.elbow_snap,
        "arm_stability_std": phase_machine.arm_stability_std
    }

    # Try LLM scoring first
    ai_feedback = ""
    final_scores = None

    lang_name = "Russian" if lang == "ru" else "English"
    generator = FeedbackGenerator()
    llm_scores, llm_feedback = generator.generate_feedback(phase_machine.scores, final_metrics, language=lang_name, video_path=input_path)
    if llm_scores is not None:
        final_scores = llm_scores
        ai_feedback = llm_feedback
        print("--- LLM COACH FEEDBACK ---")
        print(ai_feedback)
        print("--------------------------")
    else:
        print("LLM scoring failed, falling back to local scoring...")
        # Fallback: use local scoring
        phase_machine._calculate_scores()
        final_scores = phase_machine.scores
        generator = CustomFeedbackGenerator()
        ai_feedback = generator.generate(final_scores, final_metrics, lang)
        print("--- LOCAL COACH FEEDBACK ---")
        print(ai_feedback)
        print("---------------------------")

    except Exception as e:
        # Обработка всех остальных ошибок (сеть, парсинг и т.д.)
        print(f"WARNING: Could not generate AI feedback. Error: {e}", file=sys.stderr)

    return True


def main():
    """Main entry point with command-line argument parsing."""
    parser = argparse.ArgumentParser(
        description="Basketball Shot Form Analyzer - Analyzes shooting mechanics from video using Pose and Ball tracking",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
python main.py input.mp4 output.mp4
python main.py /path/to/video.mp4 /path/to/result.mp4
"""
    )
    parser.add_argument('input', type=str, help='Path to input video file')
    parser.add_argument('output', type=str, help='Path to output video file')
    args = parser.parse_args()

    valid_extensions = {'.mp4', '.avi', '.mov', '.mkv'}

    if os.path.splitext(args.input)[1].lower() not in valid_extensions:
        print("ERROR: Invalid input file extension.", file=sys.stderr)
        sys.exit(1)

    if os.path.splitext(args.output)[1].lower() not in valid_extensions:
        print("ERROR: Invalid output file extension.", file=sys.stderr)
        sys.exit(1)

    success = process_video(args.input, args.output)
    sys.exit(0 if success else 1)


# ИСПРАВЛЕНО: Добавлены пропущенные подчеркивания для корректного запуска скрипта
if __name__ == "__main__":
    main()