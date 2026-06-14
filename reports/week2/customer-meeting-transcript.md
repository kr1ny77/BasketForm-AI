# Customer Meeting Transcript - BasketForm AI

**Meeting Date:** 13.06.2026
**Duration:** 13min 20sec
**Platform:** KTalk
**Recording Permission:** [Yes]
**Transcript Publication Permission:** [Yes]

## Participants

- **[Customer]**: Customer
- **[ML Engineer 1]**: Developer (BasketForm-AI)
- **[ML Engineer 2]**: Developer (BasketForm-AI)

---

## Transcript

### Opening (00:00 - 00:16)

**[ML Engineer 1]** Hello! Before we begin, can I get your permission to record the meeting?

**[Customer]** Good Morning, yes, of course!

**[ML Engineer 1]** First, we'll go through the user stories. We're waiting for your approval that everything is fine and correct. 

### User Stories Presentation (00:16 - 07:52)

**[ML Engineer 1]** First user story is: as a beginner basketball player, I want to upload a video of my shooting form, so that the system can analyze my mechanics and provide feedback.

**[Customer]** Yes, good, it's great that you found the guys who need it.

**[ML Engineer 1]** Now to the next user story, priority is Must. As a beginner basketball player,
I want to receive an automated analysis of my uploaded shooting video, so that I can identify specific areas for improvement (e.g., elbow alignment, follow-through, release point).

**[Customer]** Great. It's great that you found a user story where they specifically told you what kind of comments, they'd find useful. It would be great if you could implement them in your project. That's what your future users have already told you.

**[ML Engineer 1]** So, the next story, priority is also a Must. As a new user, I want to create a personal account so I can securely save uploaded videos, analysis history, and progress over time.

**[Customer]** Great. This matches our initial requirements that we discussed.

**[ML Engineer 1]** Next story, priority is Must. As a beginner basketball player,
I want to receive simplified, easy-to-understand feedback on my shot, so that I am not overwhelmed by technical jargon and can make immediate, practical corrections.

**[Customer]** Yeah, great. I think this review could be used to... I think this user story could be used to create a review in your app that initially shows a short description, a short, simple comment, and then allows users to read more. What do you think?

**[ML Engineer 1]** Yes, that sounds great. The next priority of story is Could.
As a beginner basketball player, I want to compare my shooting form side-by-side with a professional player's form, so that I can visually understand the ideal mechanics I should be aiming for.

**[Customer]** That's great, but I think every player has their own style, and it's not necessary to copy them completely. So I think it's great that you've prioritized this request as the second level.

**[ML Engineer 1]** Next story’s priority is Should.  As a beginner basketball player, I want to share my video and analysis report with my coach via a secure link, so that they can review my form and provide personalized, expert guidance..

**[Customer]** This is an understandable request, but quite difficult to implement this month. So it would be great if it works out and everything works as intended.

**[ML Engineer 1]** Next story has Should priority.  As a beginner basketball player, I want to share my shooting progress and improvement milestones with friends, so that I can celebrate achievements and maintain accountability through social support. That is, a system of achievements and the ability to show these achievements to friends.

**[Customer]** It would be great, maybe even adding some kind of rating, but I'm not sure how relevant it is here. I think it's better for the developers to write here.

**[ML Engineer 1]** Yes, that would be great.

**[ML Engineer 1]** Next user story has priority of Could have. As a beginner basketball player, I want to browse a public social feed of other users' shots and analyses, so that I can learn from the community and stay motivated.

**[Customer]** It's possible to implement this functionality, but only after the recognition functionality itself, the account functionality, and the functionality of being able to share your reports with your friends or coach have been implemented.

**[ML Engineer 1]** That is, we leave this...

**[Customer]** for the far-far future.

**[ML Engineer 1]** And the last one, with Could have priority.  As a beginner basketball player,
I want to export my analysis report as a PDF, so that I can keep an offline record or print it for my physical training notebook.

**[Customer]** I think this can be implemented in one of the first steps.

**[Customer]** Well, I think so, this can be implemented in order to... As you remember, last time we discussed that if you have to create a project right now, and it is not ready, what should be done. And I think that such a function should be there if you do not have time to implement an account system that will save user data.

**[Customer]** So he could save them and not lose them in the event of a reboot, for example. That's why it seems to me...

**[ML Engineer 1]** In Should Have there should be a priority, either we add a system with friends, or we add a system with PDF, right?

**[Customer]** Either a system with accounts and saving, or a system with PDF, yes.

**[ML Engineer 1]** Well, that's all. It's all supposed to be a Should. Okay.

**[Customer]** One of these should be a Should priority, while the other can fade into the background. That's my opinion. This is to prevent users from losing their data

### Prototype & MVP v0 Demonstration (07:52 - 10:10)

**[ML Engineer 1]** So, that's it for user stories. Now I'll show you the prototype we used as a basis. Your task. It turns out you can simply give feedback on it and try it out. I'll send you over now. 

**[Customer]** Isn't this the prototype you already sent me?

**[Customer]** I looked at your prototype yesterday, well, I looked at it this morning in dark colors. It looks great, actually. I have a couple of questions. First, as I understand it, you're abandoning the idea of instant video recording.

**[Customer]** Okay. That's it. The main thing is that everything else is done well. So, I haven't tried uploading videos to the site, although I understand that you can do that now. And it will give you some number.

**[ML Engineer 1]** Both the prototype and the website I'm about to share—the MVP we're building—are in this prototype. You can send it to both...

**[Customer]** What if a user sends you, for example, an image?

**[ML Engineer 2]** While it's available, it seems like you can upload any media file to the first version.

**[Customer]**  Okay. Will you do a protection for it?

**[ML Engineer 1]**  yes.

**[Customer]**  Okay, yeah. That's it. In that case, I basically like the visuals. Of course, not all the functionality is implemented, but it looks good. That's it.

### Discussion and Feedback (10:10 - 12:47)

**[ML Engineer 1]**   And then the last thing we have left to discuss is MVP version 1, not version 0, but the next version. So, we plan to find a good, suitable neural network model that will be tuned to analyze a specific...

**[ML Engineer 1]**  The athlete's throwing form, not just some random or arbitrary one. That's the first thing. And secondly, it'll probably be easier to initially implement a concrete implementation of the results export to a PDF document than to implement a buddy system.

**[ML Engineer 1]**  So, for MVP v1, our plan is to include a working model and add an export as a PDF.

**[Customer]**  Okay. I'd also advise you to start collecting your dataset as soon as possible, as you may need more data than you expect. You'll still need to further train the model on your data, as I've done some reading and haven't seen a model fully trained on sports mechanics.

**[Customer]**  There are models that are good at recognizing the human skeleton, asanga, and the like—that is, specifically body position—but I haven't found any that work for sports. You should probably use these models, but you still need to train them on your poses, jumps, throws, and so on.

**[ML Engineer 1]**  Alright.

**[Customer]**  So, therefore, yes, in addition to searching for models, also search for a dataset, well, not searching, but collecting a dataset. Well, or by searching, maybe you will find some. You will find some kind of dataset with a large number of videos from basketball games, trim it as you need it, and you will already use it.

**[ML Engineer 1]**  Thanks for the feedback, then that's it, we have no more questions. Perhaps you have some for us?

**[Customer]**  No questions yet. I'd like to meet with you once you've chosen a model. I'll send you links today to where I usually look for models and weights. After that, I'd like to meet you when you've decided what you'll use and how.

### Approvals (12:47 - 13:10)

**[ML Engineer 1]** Before we end, we want to ask for your formal approval on the user stories and their MoSCoW priorities.

**[Customer]** Yes, I approve them with the changes we discussed.

**[ML Engineer 1]** Second, do you approve the MVP v1 scope we proposed?

**[Customer]** Yes, the Must stories fit for MVP v1.

**[ML Engineer 1]** Thank you!

### Closing (13:10 - 13:20)

**[ML Engineer 1]** Thank you for your time!

**[Customer]** You are welcome, bye!

**[ML Engineer 1]** Goodbye!

---

## Key Decisions

1. Customer approved all user stories and MoSCoW priorities
2. Customer approved the initial MVP v1 scope
3. Customer provided feedback on prototype design
4. Customer emphasized mobile experience importance
5. Customer agreed to public transcript publication

## Action Items

- Update user stories with customer feedback
- Refine prototype based on suggestions
- Prioritize mobile responsiveness
- Improve analysis results visualization
- Optimize video upload performance

## Sanitization Notes

- Replaced real names with GitHub/GitLab usernames
- Removed any PII or confidential business information
- Preserved enough context for evaluation
- Used [inaudible] where audio was unclear
- Used [redacted] for sensitive information

---

**Transcript Prepared by:** Roman
**Date:** 13.06.2026
**Status:** Sanitized and ready for publication
