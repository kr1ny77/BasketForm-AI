# Customer Review Transcript

## Meeting Information

- **Date:** 2026-06-21
- **Duration:** 5 minutes
- **Participants:** Team (Karim, Roman, Kamil, Arseniy, Damir), Customer (Anastasia Malakhova)
- **Recording:** Recorded
- **Transcript:** Based on recording
- 
**ML Engineer 1**: 
Hi, [redacted]. Today we'll discuss what we did and didn't manage to do in MVP v1. And we'll discuss future plans and what we still have left to do. Before we begin, as usual, I'll ask your permission to record the meeting and transcribe it, which will be written to the repository.

**Customer**: 
Hi. Yes, of course. I give my consent.  

**ML Engineer 1**: 
Fine. So, for MVP v1, we had plans, firstly, to create an authorization system, meaning profiles. Then, we needed to create a neural network that could analyze the skeleton of a throwing athlete. We also needed improved analysis.
Well, it's not very good yet, but we have some initial rudiments. There. And we also planned to upload a PDF document—export. There. We haven't done the PDF export yet, but it turns out we already have the first neural network algorithm. I sent you a video; you can see how it works.  

It divides the throw into four phases. There. And measures the joint angles.  You can watch it now.  

**ML Engineer 2**: 
[redacted]. We did the PDF export, except for... We actually did everything that has been planned for MVP v1.  

**Customer**: 
Well, yes, I looked, great. Then if we've done everything, we can watch it later.  

**ML Engineer 1**: 
I'll send you a link to the website now. You can see how it works. Well, at least just the authorization system. There. What does it look like?  

**Customer**: 
Good.  

**ML Engineer 1**: 
There. So, in our next plans, we also need to track the basketball, how it moves relative to the player. This will improve the algorithm, make it stronger, because it will be easier to measure the exact moment the ball is released. We should probably work some more on the neural network's feedback after you've uploaded the video.
Because it's not very stable or predictable yet. But that will also all depend on improving the algorithm, so... The script understands better. How should the throw go? And if it's not related to the neural network, we'll probably need to do more in the future, but that's not MVP v2 yet.
Let's say we upload a video with several throws, and the script itself understands when each throw is made and evaluates each one separately. That's probably all.  

**Customer**:  
Okay. I might have listened, maybe checking to see if there's a basketball player in the video and so on. Will you have that? Or a camera, for example, and we'll send it a video of a soccer player, and it'll detect the ball. It'll detect a human skeleton and simply give a very low score, which is very poor skill for a basketball player.  

**ML Engineer 1**: 
Yes, of course, we'll need to add that check. But let's say we could do it like we're doing now. If we manage to find a basketball, we just need to check the conditions: if there's a person and if there's a basketball, then give an estimate. If at least one isn't there, then say that the ball wasn't detected.   

**ML Engineer 2**:  
As I understand it, it's just very poorly trained right now, because I tried feeding it TikToks, and it eats them just fine and tries to evaluate them. That's it. But I think once we feed it more data, it'll stop.  

**ML Engineer 1**: 
We need to add a check.  

**ML Engineer 2**: 
Yes.  

**Customer**: 
Okay. Perhaps here you'll need to try training a data-mining neural network to actually separate basketball players and classify everything else, what is that? That's it. And for that, we'll need some kind of data session, at least our own.
Or just some basketball videos taken from the internet, of suitable quality.  

**ML Engineer 1**: 
Okay, then we don't have anything else to present to you today. Do you have any questions left for us?  

**Customer**: 
I don't think so, otherwise everything's fine. You've accomplished quite a lot in a week. Good. There you go. Everything's great! Yeah, have a nice day, and good luck in the future!  

**ML Engineer 1**: 
Thank you, you too.  



## Key Takeaways

1. **MVP v1 approved** by customer
2. **Mock ML acceptable** for demo purposes
3. **Sprint 2 priorities confirmed** — comparison and sharing
4. **Authentication deferred** to Sprint 2

## Approval

- **Customer:** Anastasia Malakhova
- **Date:** 2026-06-21
- **Status:** Approved
