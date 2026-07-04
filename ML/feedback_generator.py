import requests
import json
import os
from dotenv import load_dotenv

load_dotenv()

class FeedbackGenerator:
    def __init__(self, api_key=None):
        # Сначала смотрим, передали ли ключ явно, потом ищем в переменных окружения (из .env)
        self.api_key = api_key or os.environ.get("OPENROUTER_API_KEY")

        if not self.api_key:
            raise ValueError(
                "API key for OpenRouter is missing. Please create a .env file with OPENROUTER_API_KEY=your_key")

        self.api_url = "https://openrouter.ai/api/v1/chat/completions"
        self.model_name = "poolside/laguna-m.1:free"  # Или ваша рабочая модель
        print(f"Initialized OpenRouter API client with model: {self.model_name}")

    def generate_feedback(self, scores, metrics, language="English"):
        # 1. Промпт полностью на английском
        prompt_data = f"""Scores by phase (out of 100):
DIP (Squat): {scores.get('DIP ', scores.get('DIP', 0))}
ASCENT (Rise): {scores.get('ASCENT ', scores.get('ASCENT', 0))}
RELEASE: {scores.get('RELEASE ', scores.get('RELEASE', 0))}

Biomechanical metrics:
Minimum knee angle (DIP): {metrics.get('min_knee_dip ', metrics.get('min_knee_dip', 0)):.1f}°
Torso lean angle (ASCENT): {metrics.get('torso_ascent ', metrics.get('torso_ascent', 0)):.1f}°
Maximum elbow angle (RELEASE): {metrics.get('elbow_release ', metrics.get('elbow_release', 0)):.1f}°
Forearm angle relative to vertical (RELEASE): {metrics.get('forearm_release ', metrics.get('forearm_release', 0)):.1f}°

Write a short, encouraging summary (2-3 sentences) and 3 specific bullet points on what the player needs to fix in their technique. Length: up to 150 words. Write entirely in English."""

        # 2. Системное сообщение тоже на английском
        messages = [
            {
                "role": "system",
                "content": "You are a professional basketball coach. Your task is to analyze shooting biomechanical metrics and provide brief, constructive feedback in English."
            },
            {
                "role": "user",
                "content": prompt_data
            }
        ]

        headers = {
            "Authorization": f"Bearer {self.api_key}",
            "Content-Type": "application/json",
        }

        payload = {
            "model": self.model_name,
            "messages": messages,
        }

        try:
            response = requests.post(
                url=self.api_url,
                headers=headers,
                data=json.dumps(payload),
                timeout=60
            )

            response.raise_for_status()
            response_data = response.json()

            assistant_message = response_data['choices'][0]['message']
            feedback = assistant_message.get('content', '').strip()

            return feedback

        except requests.exceptions.RequestException as e:
            print(f"ERROR: API request failed: {e}")
            if hasattr(e, 'response') and e.response is not None:
                print(f"Response body: {e.response.text}")
            return "Failed to get feedback from AI coach due to a network error."
        except (KeyError, IndexError) as e:
            print(f"ERROR: Failed to parse API response: {e}")
            return "Failed to parse AI coach response."