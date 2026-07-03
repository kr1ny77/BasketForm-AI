import os
import json
import base64
import urllib.request
import urllib.error

OPENROUTER_API_KEY = os.environ.get("OPENROUTER_API_KEY", "")
OPENROUTER_URL = "https://openrouter.ai/api/v1/chat/completions"
VISION_MODEL = "google/gemini-2.0-flash-001"

SYSTEM_PROMPT_EN = """You are a professional basketball shooting form coach and biomechanics analyst.

You will receive key frames from a basketball shot video along with biomechanical metrics. Analyze the VISUAL evidence in the frames AND the metrics to score the shot.

Score each phase from 0 to 100:
- DIP (Squat/Knee Bend): Look at knee bend depth. Ideal: 90-110 degrees.
- ASCENT (Rise): Look at torso posture. Ideal: upright, 165-180 degrees.
- RELEASE: Look at arm extension and ball position. Ideal: full extension, clean release.
- FOLLOW_THROUGH: Look at wrist snap and arm position after release. Ideal: arm extended, wrist flexed down.

Also consider these metrics if provided:
- Knee angle, torso angle, elbow angle, forearm angle
- Follow-through duration, elbow snap, arm stability

IMPORTANT: If tracking metrics are unreliable (default values), rely more on the visual evidence from the frames.

Respond in valid JSON:
{
  "scores": {
    "DIP": <int 0-100>,
    "ASCENT": <int 0-100>,
    "RELEASE": <int 0-100>,
    "FOLLOW_THROUGH": <int 0-100>
  },
  "feedback": "<detailed feedback 100-200 words: what looks good, what needs improvement, specific drills>"
}

Be encouraging but honest. Base feedback on what you actually SEE in the frames."""

SYSTEM_PROMPT_RU = """Ты — профессиональный тренер по баскетболу и аналитик биомеханики.

Ты получишь ключевые кадры из видео броска вместе с биомеханическими метриками. Проанализируй ВИЗУАЛЬНЫЕEvidence на кадрах И метрики.

Оцени каждую фазу от 0 до 100:
- DIP (Присед): Глубина сгибания коленей. Идеал: 90-110 градусов.
- ASCENT (Подъём): Осанка корпуса. Идеал: прямо, 165-180 градусов.
- RELEASE (Релиз): Выпрямление руки и позиция мяча. Идеал: полное выпрямление, чистый бросок.
- FOLLOW-THROUGH (Завершение): Хлопок запястьем и позиция руки после броска. Идеал: рука выпрямлена, запястье согнуто вниз.

Учитывай метрики если они надёжные.

ОТВЕТЬ В JSON:
{
  "scores": {
    "DIP": <число 0-100>,
    "ASCENT": <число 0-100>,
    "RELEASE": <число 0-100>,
    "FOLLOW_THROUGH": <число 0-100>
  },
  "feedback": "<подробная обратная связь 100-200 слов>"
}

Будь ободряющим, но честным. Опирайся на то, что ВИДИШЬ на кадрах."""


def extract_key_frames(video_path, output_dir):
    """Extract 4 key frames from video: one per phase."""
    try:
        import cv2
    except ImportError:
        print("OpenCV not available for frame extraction")
        return []

    cap = cv2.VideoCapture(video_path)
    if not cap.isOpened():
        return []

    total_frames = int(cap.get(cv2.CAP_PROP_FRAME_COUNT))
    if total_frames < 4:
        cap.release()
        return []

    # Extract frames at 10%, 35%, 60%, 85% of video
    positions = [0.10, 0.35, 0.60, 0.85]
    frames = []

    os.makedirs(output_dir, exist_ok=True)

    for i, pos in enumerate(positions):
        frame_num = int(total_frames * pos)
        cap.set(cv2.CAP_PROP_POS_FRAMES, frame_num)
        ret, frame = cap.read()
        if ret:
            path = os.path.join(output_dir, f"frame_{i}.jpg")
            cv2.imwrite(path, frame, [cv2.IMWRITE_JPEG_QUALITY, 85])
            frames.append(path)

    cap.release()
    return frames


def encode_image(path):
    """Encode image to base64."""
    with open(path, "rb") as f:
        return base64.b64encode(f.read()).decode("utf-8")


def score_with_llm(metrics, lang="en", video_path=None, frames_dir=None):
    """Send video frames + metrics to vision LLM and get scores + feedback."""
    system_prompt = SYSTEM_PROMPT_RU if lang == "ru" else SYSTEM_PROMPT_EN

    # Try to extract frames if video_path is provided
    frame_paths = []
    if video_path and os.path.exists(video_path):
        if frames_dir is None:
            frames_dir = os.path.join(os.path.dirname(video_path), "frames")
        frame_paths = extract_key_frames(video_path, frames_dir)
        print(f"Extracted {len(frame_paths)} key frames from video")

    # Build user message
    phase_names = ["DIP (Присед/Сquat)", "ASCENT (Подъём/Rise)", "RELEASE (Релиз)", "FOLLOW-THROUGH (Завершение)"]

    metrics_text = ""
    if lang == "ru":
        unreliable_count = 0
        if metrics.get('min_knee_dip', 180) >= 175: unreliable_count += 1
        if metrics.get('elbow_release', 0) <= 5: unreliable_count += 1
        if metrics.get('forearm_release', 90) >= 85: unreliable_count += 1
        if metrics.get('frames_in_follow_through', 0) <= 1: unreliable_count += 1

        tracking = "ХОРОШЕЕ" if unreliable_count < 2 else "ЧАСТИЧНОЕ" if unreliable_count < 4 else "ПЛОХОЕ"

        metrics_text = f"""
Метрики трекинга (качество: {tracking}):
- Угол колена (DIP): {metrics.get('min_knee_dip', 'N/A')}°
- Угол корпуса (ASCENT): {metrics.get('torso_ascent', 'N/A')}°
- Угол локтя (RELEASE): {metrics.get('elbow_release', 'N/A')}°
- Предплечье от вертикали: {metrics.get('forearm_release', 'N/A')}°
- Завершение: {metrics.get('frames_in_follow_through', 0)} кадров (~{round(metrics.get('frames_in_follow_through', 0) / 30, 1)}с)
- Сгибание локтя: {metrics.get('elbow_snap', 0)}°
- Стабильность предплечья: {metrics.get('arm_stability_std', 0)}° откл."""
    else:
        unreliable_count = 0
        if metrics.get('min_knee_dip', 180) >= 175: unreliable_count += 1
        if metrics.get('elbow_release', 0) <= 5: unreliable_count += 1
        if metrics.get('forearm_release', 90) >= 85: unreliable_count += 1
        if metrics.get('frames_in_follow_through', 0) <= 1: unreliable_count += 1

        tracking = "GOOD" if unreliable_count < 2 else "PARTIAL" if unreliable_count < 4 else "POOR"

        metrics_text = f"""
Tracking metrics (quality: {tracking}):
- Knee angle (DIP): {metrics.get('min_knee_dip', 'N/A')}°
- Torso angle (ASCENT): {metrics.get('torso_ascent', 'N/A')}°
- Elbow angle (RELEASE): {metrics.get('elbow_release', 'N/A')}°
- Forearm from vertical: {metrics.get('forearm_release', 'N/A')}°
- Follow-through: {metrics.get('frames_in_follow_through', 0)} frames (~{round(metrics.get('frames_in_follow_through', 0) / 30, 1)}s)
- Elbow snap: {metrics.get('elbow_snap', 0)}°
- Arm stability: {metrics.get('arm_stability_std', 0)}° std"""

    text_content = f"""Analyze this basketball shot. Key frames show the shot at different phases.

Frame order: {', '.join(phase_names[:len(frame_paths)])}
{metrics_text}

Score each phase 0-100 and provide detailed feedback."""

    if lang == "ru":
        text_content = f"""Проанализируй этот бросок. Ключевые кадры показывают бросок на разных фазах.

Порядок кадров: {', '.join(phase_names[:len(frame_paths)])}
{metrics_text}

Оцени каждую фазу 0-100 и дай подробную обратную связь."""

    # Build message content with images
    content = []

    # Add text
    content.append({"type": "text", "text": text_content})

    # Add images if available
    for frame_path in frame_paths:
        b64 = encode_image(frame_path)
        content.append({
            "type": "image_url",
            "image_url": {
                "url": f"data:image/jpeg;base64,{b64}"
            }
        })

    payload = {
        "model": VISION_MODEL,
        "messages": [
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": content}
        ],
        "temperature": 0.3,
        "max_tokens": 1000
    }

    data = json.dumps(payload).encode('utf-8')
    req = urllib.request.Request(
        OPENROUTER_URL,
        data=data,
        headers={
            "Content-Type": "application/json",
            "Authorization": f"Bearer {OPENROUTER_API_KEY}",
            "HTTP-Referer": "https://basketform-ai.app",
            "X-Title": "BasketForm-AI"
        }
    )

    try:
        print(f"Calling vision LLM ({VISION_MODEL}) with {len(frame_paths)} frames...")
        with urllib.request.urlopen(req, timeout=90) as response:
            result = json.loads(response.read().decode('utf-8'))

        content_text = result['choices'][0]['message']['content']

        # Extract JSON from response
        if '```json' in content_text:
            content_text = content_text.split('```json')[1].split('```')[0].strip()
        elif '```' in content_text:
            content_text = content_text.split('```')[1].split('```')[0].strip()

        parsed = json.loads(content_text)

        scores = parsed.get("scores", {})
        feedback = parsed.get("feedback", "")

        for key in ["DIP", "ASCENT", "RELEASE", "FOLLOW_THROUGH"]:
            if key not in scores:
                scores[key] = 50
            scores[key] = max(0, min(100, int(scores[key])))

        print(f"Vision LLM scoring: DIP={scores['DIP']}, ASCENT={scores['ASCENT']}, "
              f"RELEASE={scores['RELEASE']}, FOLLOW_THROUGH={scores['FOLLOW_THROUGH']}")

        return scores, feedback

    except urllib.error.HTTPError as e:
        error_body = e.read().decode('utf-8') if e.readable() else str(e)
        print(f"Vision LLM API error {e.code}: {error_body}", flush=True)
        return None, None
    except json.JSONDecodeError as e:
        print(f"Vision LLM returned invalid JSON: {e}", flush=True)
        return None, None
    except Exception as e:
        print(f"Vision LLM scoring failed: {e}", flush=True)
        return None, None
