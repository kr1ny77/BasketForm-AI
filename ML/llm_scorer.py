import os
import json
import urllib.request
import urllib.error

OPENROUTER_API_KEY = os.environ.get("OPENROUTER_API_KEY", "")
OPENROUTER_URL = "https://openrouter.ai/api/v1/chat/completions"

SYSTEM_PROMPT_EN = """You are a professional basketball shooting form coach and biomechanics analyst.

You will receive raw biomechanical metrics from a basketball shot video analysis. Your job is to:
1. Score each phase (DIP, ASCENT, RELEASE, FOLLOW_THROUGH) from 0 to 100
2. Calculate an overall average score
3. Write a detailed feedback report explaining WHY each score was given and WHAT the player should improve

Scoring guidelines:
- DIP (Squat/Knee Bend): Ideal minimum knee angle is 90-110 degrees. Deeper = more power.
- ASCENT (Rise): Ideal torso angle is 165-180 degrees (upright). Leaning = balance issues.
- RELEASE: Ideal elbow > 160 degrees, forearm < 20 degrees from vertical. Clean wrist snap.
- FOLLOW_THROUGH: Based on duration, elbow snap intensity, and arm extension stability.

You MUST respond in valid JSON with this exact structure:
{
  "scores": {
    "DIP": <int 0-100>,
    "ASCENT": <int 0-100>,
    "RELEASE": <int 0-100>,
    "FOLLOW_THROUGH": <int 0-100>
  },
  "feedback": "<detailed feedback text, 100-200 words>"
}

The feedback should be encouraging but honest. Start with the overall assessment, then break down each phase with specific actionable advice."""

SYSTEM_PROMPT_RU = """Ты — профессиональный тренер по баскетболу и аналитик биомеханики броска.

Ты получишь сырые биомеханические метрики из видеоанализа броска. Твоя задача:
1. Оценить каждую фазу (DIP, ASCENT, RELEASE, FOLLOW_THROUGH) от 0 до 100
2. Посчитать средний балл
3. Написать подробный отчёт с объяснением ПОЧЕМУ поставлена такая оценка и ЧТО нужно исправить

Критерии оценки:
- DIP (Присед): Идеальный минимальный угол колена 90-110 градусов. Глубже = больше мощи.
- ASCENT (Подъём): Идеальный угол корпуса 165-180 градусов (прямо). Наклон = проблемы с балансом.
- RELEASE (Релиз): Идеальный локоть > 160 градусов, предплечье < 20 градусов от вертикали.
- FOLLOW-THROUGH (Завершение): По длительности, интенсивности сгибания локтя и стабильности руки.

Ты ОБЯЗАН ответить в формате валидного JSON:
{
  "scores": {
    "DIP": <число 0-100>,
    "ASCENT": <число 0-100>,
    "RELEASE": <число 0-100>,
    "FOLLOW_THROUGH": <число 0-100>
  },
  "feedback": "<подробный текст обратной связи, 100-200 слов>"
}

Отчёт должен быть ободряющим, но честным. Начни с общей оценки, затем разбери каждую фазу с конкретными советами."""


def score_with_llm(metrics, lang="en"):
    """Send metrics to LLM via OpenRouter and get scores + feedback."""
    system_prompt = SYSTEM_PROMPT_RU if lang == "ru" else SYSTEM_PROMPT_EN

    user_message = f"""Analyze this basketball shot and provide scores + feedback.

Biomechanical Metrics:
- Minimum knee angle during squat (DIP phase): {metrics.get('min_knee_dip', 'N/A')} degrees
- Torso angle during rise (ASCENT phase): {metrics.get('torso_ascent', 'N/A')} degrees
- Maximum elbow angle at release (RELEASE phase): {metrics.get('elbow_release', 'N/A')} degrees
- Forearm angle from vertical at release: {metrics.get('forearm_release', 'N/A')} degrees
- Follow-through duration: {metrics.get('frames_in_follow_through', 'N/A')} frames (~{round(metrics.get('frames_in_follow_through', 0) / 30, 1)} seconds)
- Elbow snap during follow-through: {metrics.get('elbow_snap', 'N/A')} degrees of bend
- Arm extension stability (lower = more stable): {metrics.get('arm_stability_std', 'N/A')} degrees std deviation"""

    if lang == "ru":
        user_message = f"""Проанализируй этот бросок и дай оценки + обратную связь.

Биомеханические метрики:
- Минимальный угол колена при приседе (фаза DIP): {metrics.get('min_knee_dip', 'N/A')} градусов
- Угол корпуса при подъёме (фаза ASCENT): {metrics.get('torso_ascent', 'N/A')} градусов
- Максимальный угол локтя при релизе (фаза RELEASE): {metrics.get('elbow_release', 'N/A')} градусов
- Угол предплечья от вертикали при релизе: {metrics.get('forearm_release', 'N/A')} градусов
- Длительность завершения: {metrics.get('frames_in_follow_through', 'N/A')} кадров (~{round(metrics.get('frames_in_follow_through', 0) / 30, 1)} секунд)
- Сгибание локтя при завершении: {metrics.get('elbow_snap', 'N/A')} градусов
- Стабильность предплечья (меньше = стабильнее): {metrics.get('arm_stability_std', 'N/A')} отклонение"""

    payload = {
        "model": "openai/gpt-4o-mini",
        "messages": [
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": user_message}
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
        print("Calling LLM for scoring...")
        with urllib.request.urlopen(req, timeout=60) as response:
            result = json.loads(response.read().decode('utf-8'))

        content = result['choices'][0]['message']['content']

        # Try to extract JSON from the response
        # The LLM might wrap it in markdown code blocks
        if '```json' in content:
            content = content.split('```json')[1].split('```')[0].strip()
        elif '```' in content:
            content = content.split('```')[1].split('```')[0].strip()

        parsed = json.loads(content)

        scores = parsed.get("scores", {})
        feedback = parsed.get("feedback", "")

        # Validate scores
        for key in ["DIP", "ASCENT", "RELEASE", "FOLLOW_THROUGH"]:
            if key not in scores:
                scores[key] = 50
            scores[key] = max(0, min(100, int(scores[key])))

        print(f"LLM scoring complete: DIP={scores['DIP']}, ASCENT={scores['ASCENT']}, "
              f"RELEASE={scores['RELEASE']}, FOLLOW_THROUGH={scores['FOLLOW_THROUGH']}")

        return scores, feedback

    except urllib.error.HTTPError as e:
        error_body = e.read().decode('utf-8') if e.readable() else str(e)
        print(f"LLM API HTTP error {e.code}: {error_body}", flush=True)
        return None, None
    except json.JSONDecodeError as e:
        print(f"LLM returned invalid JSON: {e}", flush=True)
        print(f"Raw content: {content[:500] if 'content' in dir() else 'N/A'}", flush=True)
        return None, None
    except Exception as e:
        print(f"LLM scoring failed: {e}", flush=True)
        return None, None
