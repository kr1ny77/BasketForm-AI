import cv2
import mediapipe as mp
import numpy as np
import argparse
import sys
import os

# Initialize MediaPipe solutions
mp_drawing = mp.solutions.drawing_utils
mp_drawing_styles = mp.solutions.drawing_styles
mp_holistic = mp.solutions.holistic


# Exponential Moving Average
class EMASmoother:
    def __init__(self, alpha=0.3):
        self.alpha = alpha
        self.value = None

    def update(self, new_value):
        if self.value is None:
            self.value = new_value
        else:
            self.value = self.alpha * new_value + (1 - self.alpha) * self.value
        return self.value


# State Machine
class ShotPhaseStateMachine:
    def __init__(self):
        self.state = "IDLE"
        self.smooth_knee = EMASmoother(alpha=0.4)
        self.smooth_elbow = EMASmoother(alpha=0.4)
        self.prev_knee = 180
        self.prev_elbow = 0

        # Colors for phases (BGR)
        self.colors = {
            "IDLE": (200, 200, 200),  # Gray
            "DIP": (0, 255, 255),  # Yellow (Crouch/Wind-up)
            "ASCENT": (0, 255, 0),  # Green (Rise/Drive)
            "RELEASE": (0, 0, 255),  # Red (Release - Peak!)
            "FOLLOW_THROUGH": (255, 0, 0)  # Blue (Follow-through/Landing)
        }

    def update(self, knee_angle, elbow_angle):
        # Smooth angles to avoid jitter
        knee = self.smooth_knee.update(knee_angle)
        elbow = self.smooth_elbow.update(elbow_angle)

        # State transition logic
        if self.state == "IDLE":
            # Player starts bending knees (crouch)
            if knee < 160:
                self.state = "DIP"

        elif self.state == "DIP":
            # Knees start extending (going up)
            if knee > self.prev_knee + 2.0:
                self.state = "ASCENT"

        elif self.state == "ASCENT":
            # Elbow almost straightened (release) or started bending back (follow-through)
            if elbow > 155 or elbow < self.prev_elbow - 3.0:
                self.state = "RELEASE"

        elif self.state == "RELEASE":
            # Arm goes down in follow-through or player finished shot
            if elbow < 140:
                self.state = "FOLLOW_THROUGH"

        elif self.state == "FOLLOW_THROUGH":
            # Reset to IDLE when player fully straightens or time passes
            # For simplicity, reset if knees bend deeply again (new shot)
            if knee < 150:
                self.state = "IDLE"

        self.prev_knee = knee
        self.prev_elbow = elbow

        return self.state, self.colors[self.state]

    def reset(self):
        self.state = "IDLE"


def calculate_angle(a, b, c):
    """Calculates angle between three points (a-b-c), where b is the vertex"""
    a = np.array(a)
    b = np.array(b)
    c = np.array(c)
    ba = a - b
    bc = c - b
    cosine_angle = np.dot(ba, bc) / (np.linalg.norm(ba) * np.linalg.norm(bc))
    angle = np.arccos(np.clip(cosine_angle, -1.0, 1.0))
    return np.degrees(angle)


def get_landmark_coords(landmarks, idx, frame_width, frame_height):
    """Extracts landmark coordinates in pixels"""
    if idx >= len(landmarks.landmark):
        return None
    lm = landmarks.landmark[idx]
    return (int(lm.x * frame_width), int(lm.y * frame_height))


def analyze_shooting_form(landmarks, frame_width, frame_height):
    """Analyzes shooting form and returns a dictionary of angles"""
    angles = {}
    if not landmarks:
        return angles

    # Right hand (for lefties swap indices to 11, 13, 15, 23, 25, 27)
    shoulder = get_landmark_coords(landmarks, 12, frame_width, frame_height)
    elbow = get_landmark_coords(landmarks, 14, frame_width, frame_height)
    wrist = get_landmark_coords(landmarks, 16, frame_width, frame_height)
    hip = get_landmark_coords(landmarks, 24, frame_width, frame_height)
    knee = get_landmark_coords(landmarks, 26, frame_width, frame_height)
    ankle = get_landmark_coords(landmarks, 28, frame_width, frame_height)

    if all([shoulder, elbow, wrist, hip, knee, ankle]):
        angles['elbow'] = {'value': calculate_angle(shoulder, elbow, wrist), 'point': elbow}
        angles['knee'] = {'value': calculate_angle(hip, knee, ankle), 'point': knee}

        vertical_point = (elbow[0], elbow[1] + 100)
        angles['forearm_vertical'] = {'value': calculate_angle(wrist, elbow, vertical_point), 'point': elbow}
        angles['torso'] = {'value': calculate_angle(shoulder, hip, knee), 'point': hip}
        angles['shoulder'] = {'value': calculate_angle(elbow, shoulder, hip), 'point': shoulder}

    return angles


def draw_angle_info_panel(image, angles, current_phase, phase_color, position=(10, 30)):
    """Draws an info panel with angles in the corner of the video"""
    x, y = position
    # Draw panel background
    cv2.rectangle(image, (x - 5, y - 25), (x + 320, y + 180), (0, 0, 0), -1)
    cv2.rectangle(image, (x - 5, y - 25), (x + 320, y + 180), (255, 255, 255), 2)

    # Header and CURRENT PHASE
    cv2.putText(image, "Shooting Form Analysis", (x, y),
                cv2.FONT_HERSHEY_SIMPLEX, 0.6, (255, 255, 255), 2)

    # Display phase in large font
    cv2.putText(image, f"PHASE: {current_phase}", (x, y + 30),
                cv2.FONT_HERSHEY_SIMPLEX, 0.8, phase_color, 3)

    y_offset = y + 60
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
            # Highlight in red if angle is out of normal range
            color = (0, 255, 255) if min_val <= value <= max_val else (0, 0, 255)
            text = f"{desc}: {value:.1f}"
            cv2.putText(image, text, (x, y_offset), cv2.FONT_HERSHEY_SIMPLEX, 0.5, color, 2)
            y_offset += 25


def process_video(input_path, output_path):
    """
    Process basketball shooting video and generate analysis output.

    Args:
        input_path (str): Path to input video file
        output_path (str): Path to save output video file

    Returns:
        bool: True if processing succeeded, False otherwise
    """
    # Validate input file
    if not os.path.exists(input_path):
        print(f"ERROR: Input file not found: {input_path}", file=sys.stderr)
        return False

    # Validate output directory
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

    print(f"Processing video...")

    # Initialize state machine
    phase_machine = ShotPhaseStateMachine()
    frame_count = 0

    try:
        with mp_holistic.Holistic(min_detection_confidence=0.5, min_tracking_confidence=0.5) as holistic:
            while cap.isOpened():
                success, image = cap.read()
                if not success:
                    break

                frame_count += 1
                if frame_count % 30 == 0:  # Progress update every 30 frames
                    progress = (frame_count / total_frames) * 100
                    print(f"Progress: {progress:.1f}% ({frame_count}/{total_frames} frames)")

                image_rgb = cv2.cvtColor(image, cv2.COLOR_BGR2RGB)
                image_rgb.flags.writeable = False
                results = holistic.process(image_rgb)
                image_rgb.flags.writeable = True
                image_bgr = cv2.cvtColor(image_rgb, cv2.COLOR_RGB2BGR)

                angles = analyze_shooting_form(results.pose_landmarks, frame_width, frame_height)

                # --- PHASE LOGIC ---
                current_phase = "IDLE"
                phase_color = (200, 200, 200)

                if 'knee' in angles and 'elbow' in angles:
                    current_phase, phase_color = phase_machine.update(
                        angles['knee']['value'],
                        angles['elbow']['value']
                    )

                # Draw skeleton, changing color based on phase!
                if results.pose_landmarks:
                    # 1. Styles for points (landmarks) - this is a dict, change color for each point
                    landmark_style = mp_drawing_styles.get_default_pose_landmarks_style()
                    for idx in landmark_style:
                        landmark_style[idx].color = phase_color

                    # 2. Styles for lines (connections) - this is a single DrawingSpec object
                    connection_style = mp_drawing.DrawingSpec(color=phase_color, thickness=2)

                    # 3. Draw
                    mp_drawing.draw_landmarks(
                        image_bgr,
                        results.pose_landmarks,
                        mp_holistic.POSE_CONNECTIONS,
                        landmark_drawing_spec=landmark_style,
                        connection_drawing_spec=connection_style)

                # Draw panel with angles and phase
                draw_angle_info_panel(image_bgr, angles, current_phase, phase_color)

                out.write(image_bgr)

    except Exception as e:
        print(f"ERROR: Exception during processing: {str(e)}", file=sys.stderr)
        cap.release()
        out.release()
        return False

    finally:
        cap.release()
        out.release()
        cv2.destroyAllWindows()

    print(f"Processing complete! Processed {frame_count} frames")
    print(f"Output saved to: {output_path}")
    return True


def main():
    """Main entry point with command-line argument parsing"""
    parser = argparse.ArgumentParser(
        description="Basketball Shot Form Analyzer - Analyzes shooting mechanics from video",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  python main.py input.mp4 output.mp4
  python main.py /path/to/video.mp4 /path/to/result.mp4
        """
    )

    parser.add_argument(
        'input',
        type=str,
        help='Path to input video file (e.g., footage.mp4)'
    )

    parser.add_argument(
        'output',
        type=str,
        help='Path to output video file (e.g., output.mp4)'
    )

    args = parser.parse_args()

    # Validate file extensions
    valid_extensions = {'.mp4', '.avi', '.mov', '.mkv'}

    input_ext = os.path.splitext(args.input)[1].lower()
    if input_ext not in valid_extensions:
        print(f"ERROR: Invalid input file extension: {input_ext}", file=sys.stderr)
        print(f"Supported formats: {', '.join(valid_extensions)}", file=sys.stderr)
        sys.exit(1)

    output_ext = os.path.splitext(args.output)[1].lower()
    if output_ext not in valid_extensions:
        print(f"ERROR: Invalid output file extension: {output_ext}", file=sys.stderr)
        print(f"Supported formats: {', '.join(valid_extensions)}", file=sys.stderr)
        sys.exit(1)

    # Process video
    success = process_video(args.input, args.output)

    # Exit with appropriate code for backend
    sys.exit(0 if success else 1)


if __name__ == "__main__":
    main()