package services

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
)

type feedbackTemplate struct {
	stance    string
	armAngle  string
	release   string
	follow    string
	drill     string
}

var feedbacks = []feedbackTemplate{
	{
		stance:   "Your feet are shoulder-width apart, which provides a solid base. However, your left foot is slightly ahead of your right, causing a slight imbalance. Try aligning both feet parallel to the basket for more consistent shots.",
		armAngle: "Your shooting elbow is at approximately 85 degrees, which is close to the ideal 90-degree angle. Focus on keeping your elbow tucked in closer to your body rather than flaring it outward — this will improve accuracy and reduce side-spin on the ball.",
		release:  "The release point is slightly below eye level, which can result in a flatter arc. Aim to release the ball at the highest point of your jump with your arm fully extended upward. This gives the ball a better trajectory toward the basket.",
		follow:   "Your follow-through is incomplete — your wrist snaps sideways instead of forward. A proper follow-through should have your shooting hand finishing with fingers pointing toward the basket, like reaching into a cookie jar on a high shelf. Hold the follow-through for a split second after release.",
		drill:    "Practice form shooting close to the basket (3-5 feet). Focus purely on elbow alignment, release point, and holding your follow-through. Make 20 shots before moving back.",
	},
	{
		stance:   "Your stance is well-balanced with knees slightly bent, which is good for generating power. Make sure your weight is on the balls of your feet, not your heels — this allows for a quicker, more explosive jump. Your shoulders are slightly rotated; try squaring them to the basket before you begin your shooting motion.",
		armAngle: "Your shooting arm angle is slightly open at about 100 degrees. This can cause the ball to drift left or right. Work on creating a tighter L-shape with your arm (around 90 degrees) when the ball is at forehead height. Your guide hand is positioned correctly on the side of the ball.",
		release:  "You release the ball with good timing at the peak of your jump, which is excellent. However, your fingers push the ball forward rather than flicking it upward. The last contact should be with your index and middle fingers, rolling off the tips to create natural backspin.",
		follow:   "Great follow-through extension! Your arm finishes high and straight. The only improvement needed is your wrist — it should snap forward more decisively. Think of it as 'flicking water off your fingertips' toward the rim. This snap creates the backspin that helps the ball bounce in if it hits the rim.",
		drill:    "Do the one-hand form drill: hold the ball with only your shooting hand, stand 4 feet from the basket, and shoot focusing on wrist snap and high follow-through. Make 15 in a row before stepping back.",
	},
	{
		stance:   "Your base is too narrow — your feet are almost touching, which limits your balance and power generation. Widen your stance to roughly shoulder-width. Your knees are locked straight; bend them slightly (about 15-20 degrees) to load your legs before the shot. This knee bend is where most of your shooting power comes from.",
		armAngle: "Your elbow is flaring out significantly, creating an angle of about 110 degrees. This is the most common cause of left-right misses. Imagine holding a tennis ball under your shooting elbow while you practice — this forces your elbow to stay tucked. Your guide hand is also pushing the ball, which adds unwanted spin.",
		release:  "The ball is being released too early in your jump, before you reach the apex. This results in a lower arc and less control. Wait until you feel yourself momentarily weightless at the top of your jump, then extend and release. The ball should leave your hand when your arm is at full extension above your head.",
		follow:   "Your shooting arm drops immediately after release instead of staying extended. This 'short-arming' the shot reduces arc and consistency. After releasing, your arm should stay up for a full second — long enough that a teammate could high-five your hand. Your wrist also doesn't snap; practice the wrist flick separately.",
		drill:    "Use the chair drill: sit in a chair, stand up, and immediately shoot. This forces you to use your legs and find a consistent release point. Do 3 sets of 10 shots from different spots around the paint.",
	},
	{
		stance:   "Excellent base positioning — your feet are shoulder-width apart with your shooting-side foot slightly ahead, which is textbook form. Your knees are bent at a good angle for power. One small note: you tend to lean backward slightly during the shot. Try to keep your torso upright or leaning very slightly forward for a more consistent trajectory.",
		armAngle: "Your shooting arm forms a near-perfect 90-degree angle, which is ideal. The ball rests nicely on your finger pads rather than your palm, giving you better control. Your guide hand stays to the side without interfering. This is one of the strongest aspects of your form.",
		release:  "You have a smooth, high release point which is excellent for arc. The ball comes off your fingers cleanly with good backspin. One area to refine: try to release slightly faster — there's a brief hesitation at the top of your motion that allows defenders time to contest. A quicker release while maintaining form will make your shot harder to block.",
		follow:   "Your follow-through is solid with good arm extension. Your wrist snap could be slightly more pronounced — imagine your fingers reaching down to touch the rim after release. Currently your hand stays somewhat flat. A stronger wrist snap will add more backspin and a softer touch off the rim.",
		drill:    "Try the corner 3 drill: shoot 5 shots from each corner, wing, and top of the key (total 25). Track your makes. Focus on maintaining your excellent arm angle while adding a stronger wrist snap on every shot.",
	},
	{
		stance:   "Your feet are well-positioned but you're standing quite upright with minimal knee bend. Bending your knees 15-20 degrees before the shot will add significant power from your legs, reducing the effort needed from your arms. This also helps with shot consistency since leg power is more reliable than arm strength.",
		armAngle: "Your arm angle is slightly low at about 75 degrees, causing the ball to be pushed forward rather than lifted upward. Raising your set point (where you start the shooting motion) to forehead level will naturally improve your arm angle. Your elbow is well-aligned under the ball, which is good.",
		release:  "You tend to push the ball from your chest area rather than releasing from a high set point. This creates a flat trajectory. Practice starting your shot with the ball higher — at forehead level — so you can lift and release in one fluid upward motion. The ball should travel in an arc, not a line.",
		follow:   "Your follow-through arm tends to drift to the right after release, indicating your guide hand may be pushing the ball. Focus on keeping your guide hand still after release — it should not move at all. Your shooting hand should be the only hand influencing the ball's direction. Finish with your index finger pointing at the center of the rim.",
		drill:    "Do the form shooting ladder: start 2 feet from the basket, make 5 shots, step back 2 feet, make 5 more. Continue until you reach the free-throw line. This builds muscle memory for proper form at increasing distances.",
	},
	{
		stance:   "Your stance shows good athletic positioning with feet apart and knees bent. However, your feet are angled outward (duck-footed), which can cause your hips to open up and your shots to drift right. Point your toes more toward the basket — a slight 10-degree turn is fine, but anything more affects alignment.",
		armAngle: "Your elbow is nicely tucked and your arm angle is close to ideal. The ball sits well on your shooting hand. One issue: your guide hand thumb is flicking the ball on release, which adds sidespin and causes left-right variation. Tape your guide hand thumb to your index finger during practice to break this habit.",
		release:  "Good release height and timing. The ball comes off your hand with moderate backspin, which is acceptable. To improve, focus on the last two fingers that touch the ball — your index and middle finger should be the final contact points, releasing simultaneously for a clean, straight flight.",
		follow:   "Your follow-through is inconsistent — sometimes high and straight, other times cut short. Aim for a consistent finish every time: arm fully extended, wrist snapped, fingers pointing at the target. Use a consistent pre-shot routine to make this automatic. Hold your finish until the ball hits the rim.",
		drill:    "Practice the 'goose neck' drill: after every shot, hold your follow-through and check that your hand looks like a goose neck (wrist bent, fingers down). Have a partner verify or film yourself. Do 30 free throws focusing only on follow-through consistency.",
	},
}

type Processor struct {
	storage *Storage
}

func NewProcessor(storage *Storage) *Processor {
	return &Processor{storage: storage}
}

func (p *Processor) ProcessVideo(id string) {
	p.storage.UpdateStatus(id, "processing", 0)
	log.Printf("Processing video %s", id)

	for progress := 10; progress <= 90; progress += 20 {
		time.Sleep(time.Second)
		p.storage.UpdateStatus(id, "processing", progress)
	}

	time.Sleep(2 * time.Second)

	score := 40 + rand.Intn(51)
	fb := feedbacks[rand.Intn(len(feedbacks))]

	scores := []int{
		clamp(score-5+rand.Intn(11), 0, 100),
		clamp(score-10+rand.Intn(21), 0, 100),
		clamp(score-3+rand.Intn(7), 0, 100),
		clamp(score-8+rand.Intn(17), 0, 100),
	}

	feedback := buildFeedback(scores, fb)

	pose := generateMockPose()

	video, ok := p.storage.GetVideo(id)
	if !ok {
		log.Printf("Video %s not found during processing", id)
		return
	}

	result := &models.Result{
		ID:        id,
		VideoID:   id,
		Filename:  video.Filename,
		Score:     score,
		Feedback:  feedback,
		PoseData:  pose,
		Scores:    scores,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if err := p.storage.SaveResult(result); err != nil {
		log.Printf("Failed to save result for %s: %v", id, err)
		p.storage.UpdateStatus(id, "error", 0)
		return
	}

	p.storage.SetScore(id, score)
	p.storage.UpdateStatus(id, "done", 100)
	log.Printf("Video %s processed: score=%d", id, score)
}

func buildFeedback(scores []int, fb feedbackTemplate) string {
	var sections []string

	rating := func(s int) string {
		switch {
		case s >= 80:
			return "Excellent"
		case s >= 60:
			return "Good"
		case s >= 40:
			return "Needs Work"
		default:
			return "Poor"
		}
	}

	sections = append(sections, fmt.Sprintf("STANCE (%d/100 — %s)\n%s", scores[0], rating(scores[0]), fb.stance))
	sections = append(sections, fmt.Sprintf("\nARM ANGLE (%d/100 — %s)\n%s", scores[1], rating(scores[1]), fb.armAngle))
	sections = append(sections, fmt.Sprintf("\nRELEASE POINT (%d/100 — %s)\n%s", scores[2], rating(scores[2]), fb.release))
	sections = append(sections, fmt.Sprintf("\nFOLLOW-THROUGH (%d/100 — %s)\n%s", scores[3], rating(scores[3]), fb.follow))
	sections = append(sections, fmt.Sprintf("\nRECOMMENDED DRILL\n%s", fb.drill))

	return strings.Join(sections, "\n")
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func generateMockPose() []models.Point {
	skeleton := [][2]float64{
		{50, 8}, {50, 22}, {50, 40},
		{32, 52}, {68, 52}, {20, 62}, {80, 62},
		{50, 40}, {38, 68}, {62, 68}, {36, 88}, {64, 88},
	}
	var pts []models.Point
	for _, s := range skeleton {
		pts = append(pts, models.Point{
			X: s[0] + (rand.Float64()-0.5)*4,
			Y: s[1] + (rand.Float64()-0.5)*3,
		})
	}
	return pts
}
