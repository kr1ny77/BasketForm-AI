import cv2
import mediapipe as mp
import numpy as np

mp_drawing = mp.solutions.drawing_utils
mp_drawing_styles = mp.solutions.drawing_styles
mp_holistic = mp.solutions.holistic

class PoseTracker:
    def __init__(self):
        """Initialize MediaPipe Holistic pipeline."""
        self.holistic = mp_holistic.Holistic(
            min_detection_confidence=0.5,
            min_tracking_confidence=0.5
        )

    def process(self, frame):
        """Processes the frame and returns MediaPipe results."""
        image_rgb = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)
        image_rgb.flags.writeable = False
        results = self.holistic.process(image_rgb)
        image_rgb.flags.writeable = True
        return results

    def get_landmark_coords(self, landmarks, idx, frame_width, frame_height):
        """Extracts landmark coordinates in pixels."""
        if not landmarks or idx >= len(landmarks.landmark):
            return None
        lm = landmarks.landmark[idx]
        return (int(lm.x * frame_width), int(lm.y * frame_height))

    def calculate_angle(self, a, b, c):
        """Calculates angle between three points (a-b-c), where b is the vertex."""
        a, b, c = np.array(a), np.array(b), np.array(c)
        ba, bc = a - b, c - b
        cosine_angle = np.dot(ba, bc) / (np.linalg.norm(ba) * np.linalg.norm(bc))
        angle = np.arccos(np.clip(cosine_angle, -1.0, 1.0))
        return np.degrees(angle)

    def analyze_shooting_form(self, landmarks, frame_width, frame_height):
        """Analyzes shooting form and returns a dictionary of angles."""
        angles = {}
        if not landmarks:
            return angles, None

        # Right hand indices (swap to 11, 13, 15, 23, 25, 27 for lefties)
        indices = {
            'shoulder': 12, 'elbow': 14, 'wrist': 16,
            'hip': 24, 'knee': 26, 'ankle': 28
        }

        coords = {}
        for joint, idx in indices.items():
            coords[joint] = self.get_landmark_coords(landmarks, idx, frame_width, frame_height)

        if all(coords.values()):
            angles['elbow'] = {
                'value': self.calculate_angle(coords['shoulder'], coords['elbow'], coords['wrist']),
                'point': coords['elbow']
            }
            angles['knee'] = {
                'value': self.calculate_angle(coords['hip'], coords['knee'], coords['ankle']),
                'point': coords['knee']
            }

            vertical_point = (coords['elbow'][0], coords['elbow'][1] + 100)
            angles['forearm_vertical'] = {
                'value': self.calculate_angle(coords['wrist'], coords['elbow'], vertical_point),
                'point': coords['elbow']
            }
            angles['torso'] = {
                'value': self.calculate_angle(coords['shoulder'], coords['hip'], coords['knee']),
                'point': coords['hip']
            }
            angles['shoulder'] = {
                'value': self.calculate_angle(coords['elbow'], coords['shoulder'], coords['hip']),
                'point': coords['shoulder']
            }

        return angles, coords.get('wrist')  # Return angles and wrist coordinates for ball logic