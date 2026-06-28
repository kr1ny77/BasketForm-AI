import numpy as np


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


class ShotPhaseStateMachine:
    def __init__(self):
        self.state = "IDLE"
        self.smooth_knee = EMASmoother(alpha=0.4)
        self.smooth_elbow = EMASmoother(alpha=0.4)
        self.smooth_ball_dist = EMASmoother(alpha=0.5)

        self.prev_knee = 180
        self.prev_elbow = 0
        self.prev_ball_y = None

        self.frames_since_release = 0
        self.frames_in_follow_through = 0

        # Colors for phases (BGR)
        self.colors = {
            "IDLE": (200, 200, 200), "DIP": (0, 255, 255), "ASCENT": (0, 255, 0),
            "RELEASE": (0, 0, 255), "FOLLOW_THROUGH": (255, 0, 0)
        }

        # --- SCORING METRICS ---
        self.min_knee_dip = 180
        self.torso_ascent = 180
        self.elbow_release = 0
        self.forearm_release = 90
        self.scores = {"DIP": 0, "ASCENT": 0, "RELEASE": 0}

    def update(self, knee_angle, elbow_angle, ball_center, wrist_center, torso_angle, forearm_angle):
        knee = self.smooth_knee.update(knee_angle)
        elbow = self.smooth_elbow.update(elbow_angle)

        ball_wrist_dist = None
        if ball_center and wrist_center:
            dist = np.linalg.norm(np.array(ball_center) - np.array(wrist_center))
            ball_wrist_dist = self.smooth_ball_dist.update(dist)

        # Collect metrics based on current phase
        if self.state == "DIP":
            self.min_knee_dip = min(self.min_knee_dip, knee)
        elif self.state == "ASCENT":
            self.torso_ascent = torso_angle
        elif self.state == "RELEASE":
            self.elbow_release = max(self.elbow_release, elbow)
            self.forearm_release = forearm_angle

        # State transition logic
        if self.state == "IDLE":
            if knee < 160:
                self.state = "DIP"
                self.frames_since_release = 0
                self.min_knee_dip = 180  # Reset metrics for new shot

        elif self.state == "DIP":
            if knee > self.prev_knee + 2.0:
                self.state = "ASCENT"

        elif self.state == "ASCENT":
            release_by_pose = (elbow > 155 or elbow < self.prev_elbow - 3.0)
            release_by_ball = False

            if ball_wrist_dist is not None and self.prev_ball_y is not None:
                ball_moving_up = ball_center[1] < self.prev_ball_y
                ball_separated = ball_wrist_dist > 60.0
                if ball_moving_up and ball_separated:
                    release_by_ball = True

            if release_by_pose or release_by_ball:
                self.state = "RELEASE"
                self.frames_since_release = 0
                self.elbow_release = 0

        elif self.state == "RELEASE":
            self.frames_since_release += 1
            if self.frames_since_release > 10 and elbow < 140:
                self.state = "FOLLOW_THROUGH"
                self.frames_in_follow_through = 0

        elif self.state == "FOLLOW_THROUGH":
            self.frames_in_follow_through += 1
            # Force follow-through to last at least 30 frames (~1 second)
            # AND wait until the player bends their knees again (new shot or landing)
            if self.frames_in_follow_through > 30 and knee < 140:
                self.state = "IDLE"
                self._calculate_scores()  # Calculate scores at the end of the shot

        self.prev_knee = knee
        self.prev_elbow = elbow
        if ball_center:
            self.prev_ball_y = ball_center[1]

        return self.state, self.colors[self.state]

    def _calculate_scores(self):
        """Evaluates the shot based on collected biomechanical metrics."""
        import random

        # DIP: Ideal knee bend is 90-110 degrees
        knee = self.min_knee_dip
        if 90 <= knee <= 110:
            base = 95
        elif 85 <= knee < 90 or 110 < knee <= 115:
            base = 85
        elif 115 < knee <= 130:
            base = 70
        elif 130 < knee <= 145:
            base = 55
        elif 145 < knee <= 160:
            base = 40
        else:
            base = 25
        self.scores["DIP"] = max(0, min(100, base + random.randint(-5, 5)))

        # ASCENT: Ideal torso is upright (160-180 degrees)
        torso = self.torso_ascent
        if 165 <= torso <= 180:
            base = 95
        elif 155 <= torso < 165 or 180 < torso <= 185:
            base = 80
        elif 140 <= torso < 155:
            base = 60
        elif 120 <= torso < 140:
            base = 40
        else:
            base = 25
        self.scores["ASCENT"] = max(0, min(100, base + random.randint(-5, 5)))

        # RELEASE: Ideal elbow > 160, forearm < 20
        elbow = self.elbow_release
        forearm = self.forearm_release
        score_rel = 0
        if elbow > 170:
            score_rel += 50
        elif elbow > 155:
            score_rel += 40
        elif elbow > 140:
            score_rel += 25
        elif elbow > 120:
            score_rel += 15

        if forearm < 15:
            score_rel += 50
        elif forearm < 25:
            score_rel += 40
        elif forearm < 35:
            score_rel += 25
        elif forearm < 50:
            score_rel += 15

        self.scores["RELEASE"] = max(0, min(100, score_rel + random.randint(-3, 3)))

    def finalize(self):
        """
        Принудительно считает оценки, если видео закончилось
        до естественного завершения цикла броска.
        """
        if self.state != "IDLE":
            self._calculate_scores()

    def reset(self):
        self.state = "IDLE"
        self.frames_since_release = 0
        self.frames_in_follow_through = 0