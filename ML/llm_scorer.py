import os
import json
import base64
import urllib.request
import urllib.error

OPENROUTER_API_KEY = os.environ.get("OPENROUTER_API_KEY", "")
OPENROUTER_URL = "https://openrouter.ai/api/v1/chat/completions"
VISION_MODEL = "openai/gpt-4o"

SYSTEM_PROMPT_EN = """You are an expert basketball shooting form coach with 20+ years of experience analyzing biomechanics.

You will receive 12 sequential frames from a basketball shot video. These frames capture the ENTIRE shot from start to finish. Analyze EVERY frame carefully.

Your task:
1. Score each phase from 0 to 100 based on what you SEE in the frames
2. Write detailed feedback

PHASE DETECTION GUIDE (based on frames):
- FRAMES 1-3: DIP phase — look for knee bend, lowering of body
- FRAMES 4-6: ASCENT phase — look for upward drive, torso posture
- FRAMES 7-9: RELEASE phase — look at arm extension, ball leaving hand, wrist position
- FRAMES 10-12: FOLLOW-THROUGH — look at arm after release, wrist snap, landing

SCORING RUBRIC (be precise, not generous):

DIP (Squat):
- 90-110: Deep knee bend visible → 85-100
- 110-130: Moderate bend → 70-84
- 130-150: Shallow bend → 50-69
- 150+: Legs barely bent → 20-49

ASCENT (Rise):
- Torso upright, smooth upward motion → 85-100
- Slight lean but controlled → 70-84
- Noticeable lean forward/back → 50-69
- Major balance issues → 20-49

RELEASE:
- Full arm extension, ball rolls off fingertips, clean wrist → 85-100
- Good extension, minor wrist issues → 70-84
- Partial extension or early release → 50-69
- Pushing motion, no clean release → 20-49

FOLLOW-THROUGH:
- Arm fully extended, wrist snapped down, held pose → 85-100
- Good extension, decent snap → 70-84
- Arm drops early or no snap → 50-69
- No follow-through visible → 20-49

IMPORTANT RULES:
- A GOOD shot should score 75-95. Don't be overly harsh.
- An AVERAGE shot should score 55-74.
- A POOR shot should score below 55.
- Look at the actual body positions in the frames, not just metrics.
- If the shot looks clean and professional, give it a high score.

Respond in EXACT JSON format:
{
  "scores": {"DIP": <int>, "ASCENT": <int>, "RELEASE": <int>, "FOLLOW_THROUGH": <int>},
  "feedback": "<150-250 words detailed analysis>"
}

Structure your feedback as:
1. Overall impression (1 sentence)
2. Phase-by-phase breakdown with specific observations from the frames
3. Top 2 things to improve
4. Encouraging closing"""

SYSTEM_PROMPT_RU = """Ты — эксперт-тренер по биомеханике броска в баскетболе с 20+ летним опытом.

Ты получишь 12 последовательных кадров из видео броска. Эти кадры покрывают ВЕСЬ бросок от начала до конца. Проанализируй КАЖДЫЙ кадр тщательно.

ОПРЕДЕЛЕНИЕ ФАЗ (по кадрам):
- КАДРЫ 1-3: Фаза DIP — ищи сгибание коленей, опускание тела
- КАДРЫ 4-6: Фаза ASCENT — ищи подъём вверх, осанку корпуса
- КАДРЫ 7-9: Фаза RELEASE — смотри на выпрямление руки, отпускание мяча, позицию запястья
- КАДРЫ 10-12: FOLLOW-THROUGH — смотри на руку после броска, хлопок запястья, приземление

КРИТЕРИИ ОЦЕНКИ (будь точным, не скупым):

DIP (Присед):
- Глубокий сгиб коленей виден → 85-100
- Умеренный сгиб → 70-84
- Слабый сгиб → 50-69
- Ноги почти прямые → 20-49

ASCENT (Подъём):
- Корпус прямой, плавный подъём → 85-100
- Лёгкий наклон но контролируемый → 70-84
- Заметный наклон → 50-69
- Проблемы с балансом → 20-49

RELEASE (Релиз):
- Полное выпрямление руки, чистый бросок → 85-100
- Хорошее выпрямление → 70-84
- Частичное выпрямление → 50-69
- Толкающее движение → 20-49

FOLLOW-THROUGH (Завершение):
- Рука выпрямлена, запястье хлопнуло вниз → 85-100
- Хорошее выпрямление → 70-84
- Рука падает рано → 50-69
- Завершения не видно → 20-49

ВАЖНЫЕ ПРАВИЛА:
- Хороший бросок должен получить 75-95. Не будь слишком строгим.
- Средний бросок — 55-74.
- Плохой бросок — ниже 55.
- Смотри на реальные позы тела на кадрах.
- Если бросок выглядит чисто и профессионально — ставь высокую оценку.

ОТВЕТЬ СТРОГО В JSON:
{
  "scores": {"DIP": <число>, "ASCENT": <число>, "RELEASE": <число>, "FOLLOW_THROUGH": <число>},
  "feedback": "<150-250 слов подробный анализ>"
}

Структура обратной связи:
1. Общее впечатление (1 предложение)
2. Разбор по фазам с конкретными наблюдениями с кадров
3. Топ-2 чего улучшить
4. Ободряющее заключение"""


def extract_key_frames(video_path, output_dir, num_frames=12):
    """Extract evenly spaced frames from video."""
    try:
        import cv2
    except ImportError:
        print("OpenCV not available")
        return []

    cap = cv2.VideoCapture(video_path)
    if not cap.isOpened():
        return []

    total_frames = int(cap.get(cv2.CAP_PROP_FRAME_COUNT))
    if total_frames < num_frames:
        num_frames = max(total_frames, 1)

    os.makedirs(output_dir, exist_ok=True)
    frames = []

    for i in range(num_frames):
        frame_num = int(total_frames * (i + 0.5) / num_frames)
        cap.set(cv2.CAP_PROP_POS_FRAMES, frame_num)
        ret, frame = cap.read()
        if ret:
            # Resize for faster upload (max 720p)
            h, w = frame.shape[:2]
            if h > 720:
                scale = 720 / h
                frame = cv2.resize(frame, (int(w * scale), 720))
            path = os.path.join(output_dir, f"frame_{i:02d}.jpg")
            cv2.imwrite(path, frame, [cv2.IMWRITE_JPEG_QUALITY, 80])
            frames.append(path)

    cap.release()
    return frames


def encode_image(path):
    with open(path, "rb") as f:
        return base64.b64encode(f.read()).decode("utf-8")


def score_with_llm(metrics, lang="en", video_path=None, frames_dir=None):
    """Send video frames + metrics to vision LLM."""
    system_prompt = SYSTEM_PROMPT_RU if lang == "ru" else SYSTEM_PROMPT_EN

    frame_paths = []
    if video_path and os.path.exists(video_path):
        if frames_dir is None:
            frames_dir = os.path.join(os.path.dirname(video_path), "frames")
        frame_paths = extract_key_frames(video_path, frames_dir, num_frames=12)
        print(f"Extracted {len(frame_paths)} frames")

    # Build metrics summary
    if lang == "ru":
        metrics_text = f"""Биомеханические метрики (для дополнения визуального анализа):
- Мин. угол колена: {metrics.get('min_knee_dip', 'N/A')}°
- Угол корпуса: {metrics.get('torso_ascent', 'N/A')}°
- Угол локтя при релизе: {metrics.get('elbow_release', 'N/A')}°
- Предплечье: {metrics.get('forearm_release', 'N/A')}° от вертикали
- Завершение: {metrics.get('frames_in_follow_through', 0)} кадров (~{round(metrics.get('frames_in_follow_through', 0) / 30, 1)}с)"""
    else:
        metrics_text = f"""Biomechanical metrics (supplement visual analysis):
- Min knee angle: {metrics.get('min_knee_dip', 'N/A')}°
- Torso angle: {metrics.get('torso_ascent', 'N/A')}°
- Elbow angle at release: {metrics.get('elbow_release', 'N/A')}°
- Forearm: {metrics.get('forearm_release', 'N/A')}° from vertical
- Follow-through: {metrics.get('frames_in_follow_through', 0)} frames (~{round(metrics.get('frames_in_follow_through', 0) / 30, 1)}s)"""

    if lang == "ru":
        text_content = f"""Проанализируй этот бросок в баскетболе. 12 кадров показывают весь бросок от начала до конца.

{metrics_text}

Оцени каждую фазу 0-100 и дай подробную обратную связь."""
    else:
        text_content = f"""Analyze this basketball shot. 12 frames show the entire shot from start to finish.

{metrics_text}

Score each phase 0-100 and provide detailed feedback."""

    content = [{"type": "text", "text": text_content}]

    for frame_path in frame_paths:
        b64 = encode_image(frame_path)
        content.append({
            "type": "image_url",
            "image_url": {"url": f"data:image/jpeg;base64,{b64}"}
        })

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
    req = urllib.request.Request(
        OPENROUTER_URL, data=data,
        headers={
            "Content-Type": "application/json",
            "Authorization": f"Bearer {OPENROUTER_API_KEY}",
            "HTTP-Referer": "https://basketform-ai.app",
            "X-Title": "BasketForm-AI"
        }
    )

    try:
        print(f"Calling {VISION_MODEL} with {len(frame_paths)} frames...")
        with urllib.request.urlopen(req, timeout=120) as response:
            result = json.loads(response.read().decode('utf-8'))

        text = result['choices'][0]['message']['content']

        if '```json' in text:
            text = text.split('```json')[1].split('```')[0].strip()
        elif '```' in text:
            text = text.split('```')[1].split('```')[0].strip()

        parsed = json.loads(text)
        scores = parsed.get("scores", {})
        feedback = parsed.get("feedback", "")

        for key in ["DIP", "ASCENT", "RELEASE", "FOLLOW_THROUGH"]:
            if key not in scores:
                scores[key] = 70
            scores[key] = max(0, min(100, int(scores[key])))

        avg = sum(scores.values()) / 4
        print(f"Scoring done: DIP={scores['DIP']}, ASCENT={scores['ASCENT']}, "
              f"RELEASE={scores['RELEASE']}, FOLLOW_THROUGH={scores['FOLLOW_THROUGH']}, AVG={avg:.0f}")

        return scores, feedback

    except urllib.error.HTTPError as e:
        body = e.read().decode('utf-8') if e.readable() else str(e)
        print(f"API error {e.code}: {body[:300]}")
        return None, None
    except json.JSONDecodeError as e:
        print(f"Invalid JSON from LLM: {e}")
        return None, None
    except Exception as e:
        print(f"LLM failed: {e}")
        return None, None
