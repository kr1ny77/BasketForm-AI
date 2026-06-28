import random


class CustomFeedbackGenerator:
    def __init__(self):
        pass

    def generate(self, scores, metrics, lang="en"):
        dip = scores.get('DIP', 50)
        ascent = scores.get('ASCENT', 50)
        release = scores.get('RELEASE', 50)
        avg = (dip + ascent + release) / 3

        min_knee = metrics.get('min_knee_dip', 180)
        torso = metrics.get('torso_ascent', 0)
        elbow = metrics.get('elbow_release', 90)
        forearm = metrics.get('forearm_release', 20)

        if lang == "ru":
            return self._ru(dip, ascent, release, avg, min_knee, torso, elbow, forearm)
        return self._en(dip, ascent, release, avg, min_knee, torso, elbow, forearm)

    def _rating(self, s):
        if s >= 80:
            return ("Excellent", "Отлично")
        if s >= 60:
            return ("Good", "Хорошо")
        if s >= 40:
            return ("Needs Work", "Требует работы")
        return ("Poor", "Слабо")

    def _en(self, dip, ascent, release, avg, min_knee, torso, elbow, forearm):
        r = self._rating
        lines = []
        lines.append(f"OVERALL SCORE: {int(avg)}/100")
        lines.append("")

        lines.append(f"PHASE 1 — DIP (Squat): {int(dip)}/100")
        if dip >= 80:
            lines.append(f"Excellent knee bend! Your knees are bent to {min_knee:.0f}\u00b0, generating solid power from your legs. This is the foundation of a strong shot.")
        elif dip >= 60:
            lines.append(f"Good squat. Knee angle is {min_knee:.0f}\u00b0. Try bending a bit deeper before the shot for more explosive power.")
        elif dip >= 40:
            lines.append(f"Average squat. Knee angle {min_knee:.0f}\u00b0 is too straight. Bend your knees 15-20 degrees before the shot to generate power from your legs.")
        else:
            lines.append(f"Weak squat. Knees are nearly straight at {min_knee:.0f}\u00b0. Bend your knees like sitting in a chair before shooting \u2014 this is critical for technique.")

        lines.append("")
        lines.append(f"PHASE 2 — ASCENT (Rise): {int(ascent)}/100")
        if ascent >= 80:
            lines.append(f"Excellent rise! Torso is stable at {torso:.0f}\u00b0. Smooth upward motion transfers energy efficiently from legs to upper body.")
        elif ascent >= 60:
            lines.append(f"Good rise. Torso lean is {torso:.0f}\u00b0. Try keeping your torso more vertical for better balance and accuracy.")
        elif ascent >= 40:
            lines.append(f"Rise needs work. Torso lean of {torso:.0f}\u00b0 causes balance issues. Practice smooth, vertical rises \u2014 imagine a string pulling you up from the crown of your head.")
        else:
            lines.append(f"Significant rise issues. Torso at {torso:.0f}\u00b0. This reduces accuracy and power. Start with close-range form shooting.")

        lines.append("")
        lines.append(f"PHASE 3 — RELEASE: {int(release)}/100")
        if release >= 80:
            lines.append(f"Great release! Elbow at {elbow:.0f}\u00b0, forearm at {forearm:.0f}\u00b0 from vertical. Wrist and fingers work in sync for a clean release.")
        elif release >= 60:
            lines.append(f"Good release. Elbow angle {elbow:.0f}\u00b0, forearm {forearm:.0f}\u00b0 from vertical. For a perfect release, keep the wrist flexed and let the ball roll off your fingertips.")
        elif release >= 40:
            lines.append(f"Release needs correction. Elbow {elbow:.0f}\u00b0, forearm {forearm:.0f}\u00b0. Keep your elbow under your wrist and direct the ball to the rim with the highest arc. Practice wall shooting.")
        else:
            lines.append(f"Critical release issues. Elbow {elbow:.0f}\u00b0 is too open/closed, forearm {forearm:.0f}\u00b0. This is the main cause of inaccurate shots. Work on basics: elbow under wrist, ball above head, wrist flexed.")

        lines.append("")
        lines.append("RECOMMENDATIONS:")
        recs = []
        if dip < 70:
            recs.append("\u2022 Bend your knees deeper before shooting \u2014 it's your main power source")
        if ascent < 70:
            recs.append("\u2022 Keep your torso vertical during the rise, strengthen your core")
        if release < 70:
            recs.append("\u2022 Control your elbow and wrist position at release")
        if avg >= 80:
            recs.append("\u2022 Excellent technique! Focus on consistency and repeatability")
        elif avg >= 60:
            recs.append("\u2022 Good foundation, work on your weakest phases one at a time")
        else:
            recs.append("\u2022 Start with close-range form shooting, gradually increase distance")
        if not recs:
            recs.append("\u2022 Keep practicing, pay attention to each phase")
        lines.extend(recs)

        return "\n".join(lines)

    def _ru(self, dip, ascent, release, avg, min_knee, torso, elbow, forearm):
        r = self._rating
        lines = []
        lines.append(f"\u041e\u0411\u0429\u0410\u042f \u041e\u0426\u0415\u041d\u041a\u0410: {int(avg)}/100")
        lines.append("")

        lines.append(f"\u0424\u0410\u0417\u0410 1 \u2014 DIP (\u041f\u043e\u0434\u0441\u0435\u0434): {int(dip)}/100")
        if dip >= 80:
            lines.append(f"\u041e\u0442\u043b\u0438\u0447\u043d\u044b\u0439 \u043f\u043e\u0434\u0441\u0435\u0434! \u041a\u043e\u043b\u0435\u043d\u0438 \u0441\u043e\u0433\u043d\u0443\u0442\u044b \u043d\u0430 {min_knee:.0f}\u00b0, \u043e\u0431\u0435\u0441\u043f\u0435\u0447\u0438\u0432\u0430\u044f \u043c\u043e\u0449\u043d\u0443\u044e \u0433\u0435\u043d\u0435\u0440\u0430\u0446\u0438\u044e \u0438\u0437 \u043d\u043e\u0433.")
        elif dip >= 60:
            lines.append(f"\u0425\u043e\u0440\u043e\u0448\u0438\u0439 \u043f\u043e\u0434\u0441\u0435\u0434. \u0423\u0433\u043e\u043b \u043a\u043e\u043b\u0435\u043d\u0430 {min_knee:.0f}\u00b0. \u041f\u043e\u043f\u0440\u043e\u0431\u0443\u0439\u0442\u0435 \u0441\u043e\u0433\u043d\u0443\u0442\u044c \u0433\u043b\u0443\u0431\u0436\u0435 \u043f\u0435\u0440\u0435\u0434 \u0431\u0440\u043e\u0441\u043a\u043e\u043c \u0434\u043b\u044f \u0431\u043e\u043b\u0435\u0435 \u043c\u043e\u0449\u043d\u043e\u0433\u043e \u0432\u044b\u043f\u0443\u0441\u043a\u0430.")
        elif dip >= 40:
            lines.append(f"\u0421\u0440\u0435\u0434\u043d\u0438\u0439 \u043f\u043e\u0434\u0441\u0435\u0434. \u0423\u0433\u043e\u043b \u043a\u043e\u043b\u0435\u043d\u0430 {min_knee:.0f}\u00b0 \u0441\u043b\u0438\u0448\u043a\u043e\u043c \u043f\u0440\u044f\u043c\u043e\u0439. \u0421\u043e\u0433\u043d\u0438\u0442\u0435 \u043a\u043e\u043b\u0435\u043d\u0438 \u043d\u0430 15-20 \u0433\u0440\u0430\u0434\u0443\u0441\u043e\u0432 \u043f\u0435\u0440\u0435\u0434 \u0431\u0440\u043e\u0441\u043a\u043e\u043c \u0434\u043b\u044f \u0433\u0435\u043d\u0435\u0440\u0430\u0446\u0438\u0438 \u043c\u043e\u0449\u043d\u043e\u0441\u0442\u0438 \u0438\u0437 \u043d\u043e\u0433.")
        else:
            lines.append(f"\u0421\u043b\u0430\u0431\u044b\u0439 \u043f\u043e\u0434\u0441\u0435\u0434. \u041d\u043e\u0433\u0438 \u043f\u0447\u0435\u043c \u043f\u0440\u044f\u043c\u044b\u0435 \u2014 {min_knee:.0f}\u00b0. \u0421\u043e\u0433\u043d\u0438\u0442\u0435 \u043a\u043e\u043b\u0435\u043d\u0438, \u043a\u0430\u043a \u0441\u0430\u0434\u0438\u0442\u0435\u0441\u044c \u043d\u0430 \u0441\u0442\u0443\u043b \u2014 \u044d\u0442\u043e \u043a\u0440\u0438\u0442\u0438\u0447\u043d\u043e \u0434\u043b\u044f \u0442\u0435\u0445\u043d\u0438\u043a\u0438.")

        lines.append("")
        lines.append(f"\u0424\u0410\u0417\u0410 2 \u2014 ASCENT (\u041f\u043e\u0434\u044a\u0451\u043c): {int(ascent)}/100")
        if ascent >= 80:
            lines.append(f"\u041e\u0442\u043b\u0438\u0447\u043d\u044b\u0439 \u043f\u043e\u0434\u044a\u0451\u043c! \u041a\u043e\u0440\u043f\u0443\u0441 \u0441\u0442\u0430\u0431\u0438\u043b\u0435\u043d ({torso:.0f}\u00b0), \u043f\u043b\u0430\u0432\u043d\u043e\u0435 \u0434\u0432\u0438\u0436\u0435\u043d\u0438\u0435 \u043f\u0435\u0440\u0435\u0434\u0430\u0451\u0442 \u044d\u043d\u0435\u0440\u0433\u0438\u044e \u0438\u0437 \u043d\u043e\u0433 \u0432 \u0432\u0435\u0440\u0445\u043d\u044e\u044e \u0447\u0430\u0441\u0442\u044c \u0442\u0435\u043b\u0430.")
        elif ascent >= 60:
            lines.append(f"\u0425\u043e\u0440\u043e\u0448\u0438\u0439 \u043f\u043e\u0434\u044a\u0451\u043c. \u041d\u0430\u043a\u043b\u043e\u043d \u043a\u043e\u0440\u043f\u0443\u0441\u0430 {torso:.0f}\u00b0. \u0421\u0442\u0430\u0440\u0430\u0439\u0442\u0435\u0441\u044c \u0434\u0435\u0440\u0436\u0430\u0442\u044c \u043a\u043e\u0440\u043f\u0443\u0441 \u0431\u043e\u043b\u0435\u0435 \u0432\u0435\u0440\u0442\u0438\u043a\u0430\u043b\u044c\u043d\u043e.")
        elif ascent >= 40:
            lines.append(f"\u041f\u043e\u0434\u044a\u0451\u043c \u0442\u0440\u0435\u0431\u0443\u0435\u0442 \u0434\u043e\u0440\u0430\u0431\u043e\u0442\u043a\u0438. \u041d\u0430\u043a\u043b\u043e\u043d {torso:.0f}\u00b0 \u0432\u044b\u0437\u044b\u0432\u0430\u0435\u0442 \u043f\u0440\u043e\u0431\u043b\u0435\u043c\u044b \u0441 \u0431\u0430\u043b\u0430\u043d\u0441\u043e\u043c. \u0422\u0440\u0435\u043d\u0438\u0440\u0443\u0439\u0442\u0435\u0441\u044c \u043d\u0430 \u043f\u043b\u0430\u0432\u043d\u043e\u043c \u043f\u043e\u0434\u044a\u0451\u043c\u0435 \u0431\u0435\u0437 \u043d\u0430\u043a\u043b\u043e\u043d\u0430.")
        else:
            lines.append(f"\u041f\u0440\u043e\u0431\u043b\u0435\u043c\u044b \u0441 \u043f\u043e\u0434\u044a\u0451\u043c\u043e\u043c. \u041a\u043e\u0440\u043f\u0443\u0441 \u0441\u0438\u043b\u044c\u043d\u043e \u043d\u0430\u043a\u043b\u043e\u043d\u0435\u043d ({torso:.0f}\u00b0). \u041d\u0430\u0447\u043d\u0438\u0442\u0435 \u0441 \u0431\u0440\u043e\u0441\u043a\u043e\u0432 \u0443 \u043a\u043e\u043b\u044c\u0446\u0430.")

        lines.append("")
        lines.append(f"\u0424\u0410\u0417\u0410 3 \u2014 RELEASE (\u0420\u0435\u043b\u0438\u0437): {int(release)}/100")
        if release >= 80:
            lines.append(f"\u041e\u0442\u043b\u0438\u0447\u043d\u044b\u0439 \u0440\u0435\u043b\u0438\u0437! \u041b\u043e\u043a\u043e\u0442\u044c {elbow:.0f}\u00b0, \u043f\u0440\u0435\u0434\u043f\u043b\u0435\u0447\u044c\u0435 {forearm:.0f}\u00b0 \u043e\u0442 \u0432\u0435\u0440\u0442\u0438\u043a\u0430\u043b\u0438. \u0417\u0430\u043f\u044f\u0441\u0442\u044c\u0435 \u0438 \u043f\u0430\u043b\u044c\u0446\u044b \u0440\u0430\u0431\u043e\u0442\u0430\u044e\u0442 \u0441\u0438\u043d\u0445\u0440\u043e\u043d\u043d\u043e.")
        elif release >= 60:
            lines.append(f"\u0425\u043e\u0440\u043e\u0448\u0438\u0439 \u0440\u0435\u043b\u0438\u0437. \u0423\u0433\u043e\u043b \u043b\u043e\u043a\u0442\u044f {elbow:.0f}\u00b0, \u043f\u0440\u0435\u0434\u043f\u043b\u0435\u0447\u044c\u0435 {forearm:.0f}\u00b0. \u0414\u043b\u044f \u0438\u0434\u0435\u0430\u043b\u044c\u043d\u043e\u0433\u043e \u0440\u0435\u043b\u0438\u0437\u0430 \u0441\u043b\u0435\u0434\u0438\u0442\u0435 \u0437\u0430 \u043f\u043e\u043b\u043e\u0436\u0435\u043d\u0438\u0435\u043c \u0437\u0430\u043f\u044f\u0441\u0442\u044c\u044f.")
        elif release >= 40:
            lines.append(f"\u0420\u0435\u043b\u0438\u0437 \u0442\u0440\u0435\u0431\u0443\u0435\u0442 \u043a\u043e\u0440\u0440\u0435\u043a\u0442\u0438\u0440\u043e\u0432\u043a\u0438. \u041b\u043e\u043a\u043e\u0442\u044c {elbow:.0f}\u00b0, \u043f\u0440\u0435\u0434\u043f\u043b\u0435\u0447\u044c\u0435 {forearm:.0f}\u00b0. \u0423\u0434\u0435\u0440\u0436\u0438\u0432\u0430\u0439\u0442\u0435 \u043b\u043e\u043a\u043e\u0442\u044c \u043f\u043e\u0434 \u0437\u0430\u043f\u044f\u0441\u0442\u044c\u0435\u043c.")
        else:
            lines.append(f"\u041a\u0440\u0438\u0442\u0438\u0447\u0435\u0441\u043a\u0438\u0435 \u043f\u0440\u043e\u0431\u043b\u0435\u043c\u044b \u0441 \u0440\u0435\u043b\u0438\u0437\u043e\u043c. \u041b\u043e\u043a\u043e\u0442\u044c {elbow:.0f}\u00b0, \u043f\u0440\u0435\u0434\u043f\u043b\u0435\u0447\u044c\u0435 {forearm:.0f}\u00b0. \u042d\u0442\u043e \u0433\u043b\u0430\u0432\u043d\u0430\u044f \u043f\u0440\u0438\u0447\u0438\u043d\u0430 \u043d\u0435\u0442\u043e\u0447\u043d\u044b\u0445 \u0431\u0440\u043e\u0441\u043a\u043e\u0432.")

        lines.append("")
        lines.append("\u0420\u0415\u041a\u041e\u041c\u0415\u041d\u0414\u0410\u0426\u0418\u0418:")
        recs = []
        if dip < 70:
            recs.append("\u2022 \u0421\u043e\u0433\u043d\u0438\u0442\u0435 \u043a\u043e\u043b\u0435\u043d\u0438 \u0433\u043b\u0443\u0431\u0436\u0435 \u043f\u0435\u0440\u0435\u0434 \u0431\u0440\u043e\u0441\u043a\u043e\u043c \u2014 \u044d\u0442\u043e \u043e\u0441\u043d\u043e\u0432\u0430 \u043c\u043e\u0449\u043d\u043e\u0441\u0442\u0438")
        if ascent < 70:
            recs.append("\u2022 \u0414\u0435\u0440\u0436\u0438\u0442\u0435 \u043a\u043e\u0440\u043f\u0443\u0441 \u0432\u0435\u0440\u0442\u0438\u043a\u0430\u043b\u044c\u043d\u043e \u0432\u043e \u0432\u0440\u0435\u043c\u044f \u043f\u043e\u0434\u044a\u0451\u043c\u0430")
        if release < 70:
            recs.append("\u2022 \u041a\u043e\u043d\u0442\u0440\u043e\u043b\u0438\u0440\u0443\u0439\u0442\u0435 \u043f\u043e\u0437\u0438\u0446\u0438\u044e \u043b\u043e\u043a\u0442\u044f \u0438 \u0437\u0430\u043f\u044f\u0441\u0442\u044c\u044f \u043f\u0440\u0438 \u0440\u0435\u043b\u0438\u0437\u0435")
        if avg >= 80:
            recs.append("\u2022 \u041e\u0442\u043b\u0438\u0447\u043d\u0430\u044f \u0442\u0435\u0445\u043d\u0438\u043a\u0430! \u0424\u043e\u043a\u0443\u0441\u0438\u0440\u0443\u0439\u0442\u0435\u0441\u044c \u043d\u0430 \u043f\u043e\u0432\u0442\u043e\u0440\u044f\u0435\u043c\u043e\u0441\u0442\u0438")
        elif avg >= 60:
            recs.append("\u2022 \u0425\u043e\u0440\u043e\u0448\u0430\u044f \u0431\u0430\u0437\u0430, \u0440\u0430\u0431\u043e\u0442\u0430\u0439\u0442\u0435 \u043d\u0430\u0434 \u0441\u043b\u0430\u0431\u044b\u043c\u0438 \u0444\u0430\u0437\u0430\u043c\u0438")
        else:
            recs.append("\u2022 \u041d\u0430\u0447\u043d\u0438\u0442\u0435 \u0441 \u0431\u0440\u043e\u0441\u043a\u043e\u0432 \u0443 \u043a\u043e\u043b\u044c\u0446\u0430, \u043f\u043e\u0441\u0442\u0435\u043f\u0435\u043d\u043d\u043e \u0443\u0432\u0435\u043b\u0438\u0447\u0438\u0432\u0430\u044f \u0434\u0438\u0441\u0442\u0430\u043d\u0446\u0438\u044e")
        if not recs:
            recs.append("\u2022 \u041f\u0440\u043e\u0434\u043e\u043b\u0436\u0430\u0439\u0442\u0435 \u0442\u0440\u0435\u043d\u0438\u0440\u043e\u0432\u0430\u0442\u044c\u0441\u044f, \u0443\u0434\u0435\u043b\u044f\u0439\u0442\u0435 \u0432\u043d\u0438\u043c\u0430\u043d\u0438\u0435 \u043a\u0430\u0436\u0434\u043e\u0439 \u0444\u0430\u0437\u0435")
        lines.extend(recs)

        return "\n".join(lines)
