# Meeting Transcript

**Participants:** Full-stack Developer, Customer, ML Engineer
**Length:** 12:27

---

### [00:00:07] — ML Engineer
[redacted], hello! Today is our next meeting. We'll discuss what we accomplished during the last sprint and try running some user asseptance tests. Before we begin, I'd like to ask your permission to record this meeting, as well as to record user acceptance tests and transcribe the entire meeting.

### [00:00:38] — Customer
Good morning. Yes, of course, I agree to everything.

### [00:00:42] — ML Engineer
Great. So, during the last sprint, we discussed what we could plan. We were deciding between adding a friend system and comparing the mechanics of someone who quits with the mechanics of professional players. You and I came to the conclusion that adding a friend system was a higher priority. We've implemented it, and it works. We'll check that out today. That's it. Regarding the ML part. We... It turns out we redesigned the system using feedback from LLM. Previously, we had a downloaded [redacted] model, which was really loading up the server and taking a long time to start up. Now we use an API key we got from [redacted]. And we use it to get feedback. It's much faster, and the server isn't overloaded. There you go. Yeah.

### [00:01:46] — Customer
That's great.

### [00:01:48] — ML Engineer
Yes, I think I've mentioned everything. We didn't change the algorithm at all. Okay, those are all the updates we've come up with. What can we do for the next sprint? Oh, and the translation. We've also fixed the Russian translation. Remember, you mentioned something was wrong somewhere, something was incorrect somewhere.

### [00:02:11] — Customer
Yes, I remember. That's also very cool. You've accomplished a lot this week.

### [00:02:20] — ML Engineer
What can we do in the next sprint? It's our second-to-last one, right? The next one, it seems.

### [00:02:30] — Customer
Yes, I think it was. Your defense was 20... 20 something? 21...

### [00:02:40] — ML Engineer
Then what can we do? Probably come up with some recommendations for users on how to position the camera better when recording their throws. Perhaps, if we have time, we could, for example, create a system that would allow users to upload not just one throw, but videos of several throws. And the app will automatically cut out the parts where you throw and analyze them separately, so you don't have to trim them yourself. What do you think of that?

### [00:03:29] — Customer
Now I have a question: how useful is this for the user? After all, to send them a video of several throws in a row, they need to edit it, as I understand it. Well, because everything...

### [00:03:48] — ML Engineer
So, let's say you set up a camera, filmed a video, that's it, turned off the recording, and immediately uploaded it. You don't need to trim anything, no. And the app will automatically select the sections where you threw and analyze them with your help.

### [00:04:06] — Customer
That actually sounds great. It's probably even great from the perspective that you can then upload full games and throwing moments. A neural network will be able to determine this. That's how I see the future and speculate. But how difficult will it be to separate this, then... I have a suggestion to separate this using a ball. That is, make the moment when the ball leaves the person a trigger and store the frames that came before it in a queue. So, when the ball leaves the person, that's the trigger for you, that's it, we need to take X previous frames and process them. But this is my suggestion, just how I would do it. Probably, yes, it can be implemented in the remaining two weeks, maybe even faster, if possible. But still, allow yourself time so that if it suddenly turns out to be difficult, you simply don't do it because you've already reached the definition of done, just so you don't end up with a broken product in the worst case scenario, and leave yourself time to test the currently finished product. Okay?

### [00:05:30] — ML Engineer
Okay. What other suggestions do you have, what can be done, how can it be improved?

### [00:05:37] — Customer
What needs to be improved? If you don't mind, I'll open our spreadsheet for a second and see what we discussed in detail. So, you ended up making the apps into websites, which is cool and convenient. But we definitely won't have time to make an app. Yes, I understand. I'm just skimming through it. In fact, it seems like you've more or less fulfilled all the requirements I originally wrote down. Yes, some of them have been modified, but firstly, I approved that, and secondly, they've been improved. There you go. Okay, and now we're discussing what we can do besides adding video uploads with several throws in a row. What if... as a suggestion, we could compare the two videos. Specifically, what's changed in the feature. That is, compare, for example, today's throw with yesterday's throw.

### [00:08:15] — ML Engineer
But we already have a progress system that shows how much...

### [00:08:19] — Customer
by the way, it shows And it shows it in descriptions and points, or does it just describe it specifically? Yesterday it was like this, today it was like that. No, it works in bulk. I'm talking about creating a system that would directly report that today your brine was better, but your foot placement was incorrect. But that's just a guess. I don't know how useful that would be. We...

### [00:08:55] — Full-stack Developer
The system works by displaying points, and the points are assigned using four criteria, and the criteria are also displayed. So, in my opinion, it's throw, foot position, arm position, and so on. And they're also displayed there, and I can see how it's divided up.

### [00:09:20] — Customer
Okay, then it probably won't be a very useful feature.

### [00:09:24] — ML Engineer
Well, it's hard to classify them as such. Well, I guess we'll stick with uploading the entire video but dividing it into segments. Okay.

### [00:09:35] — Customer
Yeah, I guess that's it...

### [00:09:37] — ML Engineer
Okay, then let's move on to the user testы. We've already discussed the filming decision. While I'm giving you a little time to log into the site, we'll be testing our new friends window. Tell me, how will you be ready? You already have an account, right?

### [00:10:07] — Customer
Yes, I did.

### [00:10:09] — ML Engineer
Then for the first test recipe, we'll be testing the ability to send a friend request. So, you need to go to the friends page. Then, in the field where you need to enter your nickname, enter the nickname I'll send you now on Telegram. After that, you can click "Search." A notification should appear indicating that the request was successfully sent.

### [00:10:56] — Customer
Yes, yes, good, it worked.

### [00:11:00] — ML Engineer
Okay, then the next test is to ensure you can accept a friend request. I'll send you a request now. You'll also need to go to the "Friends" window. And mine will appear in the "Pending requests" section. Are you ready? I'll send you a request then. You can refresh the page, and my request should appear. After that, I'll try clicking "accept." The record should be accepted, and I should appear in the "My Friends" section.

### [00:11:49] — Customer
Yes, it worked.

### [00:11:51] — ML Engineer
Everything's great. Then both User Acceptance tests have passed. Any comments on them?

### [00:11:58] — Customer
No, everything's great.

### [00:12:00] — ML Engineer
Okay, fine. That's all for today. We discussed how we managed to complete this sprint, planned the next sprint, and passed the User Acceptance Test. Are there any final questions or suggestions?

### [00:12:18] — Customer
No, basically, everything's fine. See you at the end of the next sprint.

### [00:12:24] — ML Engineer
Okay. See you there, too.
