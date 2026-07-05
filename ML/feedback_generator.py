import os
import json
import base64
import urllib.request
import urllib.error
import cv2

OPENROUTER_API_KEY = os.environ.get("OPENROUTER_API_KEY", "")
OPENROUTER_URL = "https://openrouter.ai/api/v1/chat/completions"
VISION_MODEL = "poolside/laguna-xs-2.1:free"


class FeedbackGenerator:
    def __init__(self):
        print(f"FeedbackGenerator initialized with {VISION_MODEL}")

    def extract_frames(self, video_path, num_frames=8):
        cap = cv2.VideoCapture(video_path)
        if not cap.isOpened():
            return []
        total_frames = int(cap.get(cv2.CAP_PROP_FRAME_COUNT))
        if total_frames < num_frames:
            num_frames = max(total_frames, 1)
        frames = []
        for i in range(num_frames):
            frame_num = int(total_frames * (i + 0.5) / num_frames)
            cap.set(cv2.CAP_PROP_POS_FRAMES, frame_num)
            ret, frame = cap.read()
            if ret:
                h, w = frame.shape[:2]
                if h > 480:
                    scale = 480 / h
                    frame = cv2.resize(frame, (int(w * scale), 480))
                _, buf = cv2.imencode('.jpg', frame, [cv2.IMWRITE_JPEG_QUALITY, 75])
                frames.append(base64.b64encode(buf).decode('utf-8'))
        cap.release()
        return frames

    def generate_feedback(self, scores, metrics, language="Russian", video_path=None):
        frames_b64 = []
        if video_path and os.path.exists(video_path):
            frames_b64 = self.extract_frames(video_path)
            print(f"Extracted {len(frames_b64)} frames for LLM analysis")
        if language == "Russian":
            system_prompt = self._system_prompt_ru()
            user_text = self._user_prompt_ru(scores, metrics, len(frames_b64))
        else:
            system_prompt = self._system_prompt_en()
            user_text = self._user_prompt_en(scores, metrics, len(frames_b64))
        content = [{"type": "text", "text": user_text}]
        for b64 in frames_b64:
            content.append({"type": "image_url", "image_url": {"url": f"data:image/jpeg;base64,{b64}"}})
        payload = {
            "model": VISION_MODEL,
            "messages": [
                {"role": "system", "content": system_prompt},
                {"role": "user", "content": content}
            ],
            "temperature": 0.2,
            "max_tokens": 1200
        }
        data = json.dumps(payload).encode('utf-8')
        req = urllib.request.Request(OPENROUTER_URL, data=data, headers={
            "Content-Type": "application/json",
            "Authorization": f"Bearer {OPENROUTER_API_KEY}",
            "HTTP-Referer": "https://basketform-ai.app",
            "X-Title": "BasketForm-AI"
        })
        try:
            print(f"Calling {VISION_MODEL} with {len(frames_b64)} frames...")
            with urllib.request.urlopen(req, timeout=120) as response:
                result = json.loads(response.read().decode('utf-8'))
            text = result['choices'][0]['message']['content']
            if '```json' in text:
                text = text.split('```json')[1].split('```')[0].strip()
            elif '```' in text:
                text = text.split('```')[1].split('```')[0].strip()
            parsed = json.loads(text)
            new_scores = parsed.get("scores", scores)
            feedback = parsed.get("feedback", "")
            for key in ["DIP", "ASCENT", "RELEASE", "FOLLOW_THROUGH"]:
                if key in new_scores:
                    new_scores[key] = max(0, min(100, int(new_scores[key])))
            avg = sum(new_scores.values()) / 4
            print(f"LLM scoring: DIP={new_scores.get('DIP')}, ASCENT={new_scores.get('ASCENT')}, RELEASE={new_scores.get('RELEASE')}, FOLLOW_THROUGH={new_scores.get('FOLLOW_THROUGH')}, AVG={avg:.0f}")
            result_text = f"OVERALL SCORE: {int(avg)}/100\n\n"
            for phase, name in [("DIP", "DIP (Squat)"), ("ASCENT", "ASCENT (Rise)"), ("RELEASE", "RELEASE"), ("FOLLOW_THROUGH", "FOLLOW-THROUGH")]:
                result_text += f"PHASE - {name}: {new_scores.get(phase, 50)}/100\n"
            result_text += f"\n{feedback}"
            return new_scores, result_text
        except urllib.error.HTTPError as e:
            body = e.read().decode('utf-8') if e.readable() else str(e)
            print(f"API error {e.code}: {body[:300]}")
            return None, None
        except Exception as e:
            print(f"LLM failed: {e}")
            return None, None

    def _system_prompt_en(self):
        return """You are an expert basketball shooting form coach. Analyze the video frames and metrics to score and give feedback. Score each phase 0-100: DIP (knee bend, ideal 90-110), ASCENT (torso posture, ideal 165-180), RELEASE (arm extension, ideal elbow>160 forearm<20), FOLLOW-THROUGH (wrist snap, arm extension). A good shot scores 75-95. Respond in JSON: {"scores": {"DIP": int, "ASCENT": int, "RELEASE": int, "FOLLOW_THROUGH": int}, "feedback": "150-250 words"}"""

    def _system_prompt_ru(self):
        return """Ты — эксперт-тренер по биомеханике броска. Проанализируй кадры и метрики. Оцени каждую фазу 0-100: DIP (сгиб коленей, идеал 90-110), ASCENT (осанка, идеал 165-180), RELEASE (выпрямление руки, идеал локоть>160), FOLLOW-THROUGH (хлопок запястья). Хороший бросок 75-95. ОТВЕТЬ В JSON: {"scores": {"DIP": число, "ASCENT": число, "RELEASE": число, "FOLLOW_THROUGH": число}, "feedback": "150-250 слов"}"""

    def _user_prompt_en(self, scores, metrics, num_frames):
        return f"""Analyze this basketball shot. {num_frames} frames show the entire shot. Metrics: knee={metrics.get('min_knee_dip','?')} deg, torso={metrics.get('torso_ascent','?')} deg, elbow={metrics.get('elbow_release','?')} deg, forearm={metrics.get('forearm_release','?')} deg, follow-through={metrics.get('frames_in_follow_through',0)} frames. Score each phase and give detailed feedback."""

    def _user_prompt_ru(self, scores, metrics, num_frames):
        return f"""Проанализируй бросок. {num_frames} кадров показывают весь бросок. Метрики: колено={metrics.get('min_knee_dip','?')} deg, корпус={metrics.get('torso_ascent','?')} deg, локоть={metrics.get('elbow_release','?')} deg, предплечье={metrics.get('forearm_release','?')} deg, завершение={metrics.get('frames_in_follow_through',0)} кадров. Оцени каждую фазу и дай обратную связь."""
