from transformers import AutoModelForCausalLM, AutoTokenizer
import torch


class FeedbackGenerator:
    def __init__(self, model_name="Qwen/Qwen2.5-3B-Instruct"):
        print(f"Loading local LLM ({model_name})... This may take a minute on first run.")
        self.tokenizer = AutoTokenizer.from_pretrained(model_name)

        # Use float16 for GPU, float32 for CPU
        dtype = torch.float16 if torch.cuda.is_available() else torch.float32
        self.model = AutoModelForCausalLM.from_pretrained(
            model_name,
            torch_dtype=dtype,
            device_map="auto" if torch.cuda.is_available() else None
        )
        if not torch.cuda.is_available():
            self.model.to('cpu')

    def generate_feedback(self, scores, metrics, language="Russian"):
        # 1. Формируем промпт только с данными
        prompt_data = f"""Оценки по фазам (из 100):
DIP (Подсед): {scores.get('DIP', 0)}
ASCENT (Подъем): {scores.get('ASCENT', 0)}
RELEASE (Релиз): {scores.get('RELEASE', 0)}
Биомеханические метрики:
Минимальный угол в коленях (DIP): {metrics.get('min_knee_dip', 0):.1f}°
Угол наклона корпуса (ASCENT): {metrics.get('torso_ascent', 0):.1f}°
Максимальный угол в локте (RELEASE): {metrics.get('elbow_release', 0):.1f}°
Угол предплечья относительно вертикали (RELEASE): {metrics.get('forearm_release', 0):.1f}°

Напиши короткое, ободряющее резюме (2-3 предложения) и 3 конкретных пункта в виде списка, что игроку нужно исправить в технике. Объем: до 150 слов."""

        # 2. Используем правильный формат чата (Chat Template)
        messages = [
            {"role": "system",
             "content": "Ты — профессиональный тренер по баскетболу. Твоя задача — анализировать биомеханические метрики броска и давать краткую, конструктивную обратную связь на русском языке."},
            {"role": "user", "content": prompt_data}
        ]

        # Применяем шаблон чата, чтобы модель поняла структуру диалога
        text = self.tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
        inputs = self.tokenizer(text, return_tensors="pt").to(self.model.device)

        # 3. Генерация с правильными параметрами (без галлюцинаций)
        with torch.no_grad():
            outputs = self.model.generate(
                **inputs,
                max_new_tokens=350,
                temperature=0.1,  # Почти детерминированный вывод
                do_sample=False,  # Жадный декодинг (greedy)
                pad_token_id=self.tokenizer.eos_token_id
            )

        response = self.tokenizer.decode(outputs[0], skip_special_tokens=True)

        # Извлекаем только сгенерированную часть (отсекаем сам промпт)
        feedback = response[len(text):].strip()
        return feedback