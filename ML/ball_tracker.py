import cv2
import math
import torch
from ultralytics import YOLO


class BallTracker:
    def __init__(self, model_path='best.pt', conf_threshold=0.3):
        print(f"Loading custom YOLO model from {model_path}...")
        self.model = YOLO(model_path)
        self.conf_threshold = conf_threshold
        self.device = '0' if torch.cuda.is_available() else 'cpu'

        # History for cleaning erratic jumps
        self.ball_pos_history = []
        self.frame_count = 0

    def _clean_ball_pos(self):
        """
        Removes inaccurate ball positions to prevent jumping to wrong objects.
        """
        if len(self.ball_pos_history) > 1:
            w1, h1 = self.ball_pos_history[-2][2], self.ball_pos_history[-2][3]
            w2, h2 = self.ball_pos_history[-1][2], self.ball_pos_history[-1][3]
            x1, y1 = self.ball_pos_history[-2][0]
            x2, y2 = self.ball_pos_history[-1][0]
            f1, f2 = self.ball_pos_history[-2][1], self.ball_pos_history[-1][1]
            f_dif = f2 - f1

            dist = math.sqrt((x2 - x1) ** 2 + (y2 - y1) ** 2)
            max_dist = 4 * math.sqrt(w1 ** 2 + h1 ** 2)

            # Ball should not move 4x its diameter within 5 frames
            if (dist > max_dist) and (f_dif < 5):
                self.ball_pos_history.pop()
            # Ball should be relatively square
            elif (w2 * 1.4 < h2) or (h2 * 1.4 < w2):
                self.ball_pos_history.pop()

        # Remove points older than 30 frames to keep the list short
        if len(self.ball_pos_history) > 0:
            if self.frame_count - self.ball_pos_history[0][1] > 30:
                self.ball_pos_history.pop(0)

    def detect(self, frame):
        self.frame_count += 1
        results = self.model(frame, stream=True, device=self.device, verbose=False)

        best_ball = None
        max_conf = self.conf_threshold

        for r in results:
            boxes = r.boxes
            for box in boxes:
                cls = int(box.cls[0])
                # Class 0 is 'Basketball' in your custom model
                if cls == 0:
                    conf = float(box.conf[0])
                    if conf > max_conf:
                        max_conf = conf
                        x1, y1, x2, y2 = box.xyxy[0]
                        x1, y1, x2, y2 = int(x1), int(y1), int(x2), int(y2)
                        w, h = x2 - x1, y2 - y1
                        center = (int(x1 + w / 2), int(y1 + h / 2))
                        best_ball = (center, self.frame_count, w, h, conf)

        if best_ball:
            self.ball_pos_history.append(best_ball)

        # Clean history to remove erratic jumps
        self._clean_ball_pos()

        # Return current valid ball state
        if len(self.ball_pos_history) > 0 and self.ball_pos_history[-1][1] == self.frame_count:
            # Current frame detection is valid
            cx, cy = self.ball_pos_history[-1][0]
            w, h = self.ball_pos_history[-1][2], self.ball_pos_history[-1][3]
            return (cx, cy), (cx - w // 2, cy - h // 2, cx + w // 2, cy + h // 2)
        elif len(self.ball_pos_history) > 0:
            # Current frame was erratic, fallback to previous valid position
            cx, cy = self.ball_pos_history[-1][0]
            w, h = self.ball_pos_history[-1][2], self.ball_pos_history[-1][3]
            return (cx, cy), (cx - w // 2, cy - h // 2, cx + w // 2, cy + h // 2)
        else:
            return None, None