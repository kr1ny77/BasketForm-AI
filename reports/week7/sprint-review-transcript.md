June 18
Participants: Customer, ML Engineer
Length: +- 13min

**[00:00:02] — ML Engineer**
[redacted], hi. Today is our final meeting. We'll discuss what we've done this week. We'll discuss the MVP V3 we've built. What we resolved and what we didn't from week six's issues. We'll discuss how to finally launch the product. What you need to know to maintain it. And... further advice on operations, probably. Before we start, I'll ask for your permission, as usual, to record and transcribe our meeting, and we'll also have a user from Acceptance Tests recording theirs as well.

**[00:00:54] — Customer**
Yes, of course, I give my permission.

**[00:00:59] — ML Engineer**
All super. So, we'll start with what was planned for the past sprint. We didn't set any major goals for this sprint. On the contrary, we tackled a lot of small tasks, like checking all the bugs and improving the server. Specifically, making it more stable, configuring it so it doesn't crash or lag. Polishing the XML pipeline. We refactored some code here and there, tweaked the logic slightly. Basically, all minor changes to make the website work better, more stably, and more optimized. Any questions on this?

**[00:01:45] — Customer**
No, we already discussed this. Everything is good, you did a great job.

**[00:01:50] — ML Engineer**
Now I'll show you what we ended up with. That is, our MVP v3. I'll turn on the screen sharing now. Here. So, first, I'll log out of the account. The initial window, checking that everything works. I already have an account, I log in, all super. First, let's look at friends. Here you are, Arseniy, our ML Engineer. I can also check my profile; by the way, you can set an avatar, change the nickname and password. Here, in this window, all the results we've shared with friends are saved. Here's the results export. That works too. Here, the progress will be displayed if you upload several videos in a row. Our results are saved here. As you can see, one of the results came out with an error. This is exactly the test for when there's no person in the video, it gets discarded, meaning it's not analyzed, just as we wanted to do in the last sprint. Here. Right now it shows that... Yes, it's really cool. And it's not analyzed. Now let's try uploading a video of a real basketball player. Right now... And we'll see if it breaks or not. Need to wait a bit. So, while it's loading, what else can we discuss? I showed this, I showed that. The language also switches to Russian for us. What else? Now. Here, it's loaded. It calculated that the video had two shots. It calculated the average score, the best shot score, and the last shot score. We can see that the progress updated, the second video appeared, and we can see our progress. Well, it happens, for example, because the second result was lower. But that's probably all the scope of what needed to be shown regarding the website's functionality. If there are any...

**[00:04:45] — Customer**
it goes with an error, it doesn't go into the report at all...

**[00:04:51] — ML Engineer**
report, that's great. Then I'll minimize the demo. Okay, I've shown v3. Now let's talk about the bugs and issues that remained from last week, and basically from the whole product. We found the best angle, which is from the side. But we didn't explicitly write this anywhere on the website. It probably should go somewhere in the product description. We haven't figured out how to do it yet. But the point is, it's better to shoot the throw from the side, like in the video I'll send you.

**[00:05:38] — Customer**
Maybe it's better to display this before uploading the video, and while the video is processing, meaning while the user is waiting anyway, display it again. Since they have some free time anyway, they could read it.

**[00:05:52] — ML Engineer**
Then I'll do that.

**[00:05:54] — Customer**
This is just a suggestion, since you've already met the basic requirements anyway. If you have time to do it, if not, then...

**[00:06:05] — ML Engineer**
Regarding splitting into batches, meaning splitting one large video into individual shots. We rejected this, so this issue remains. I don't know if it splits or not. Most likely not, but right now nobody is splitting them. So videos need to be uploaded separately, one shot at a time. What else do we have left? Well, basically that's it, yeah. We implemented everything else we planned. Oh, right, comparing with other basketball players, we rejected that a long time ago. So only two deliverables were rejected in total. Everything else is done. At the moment, we have no pending issues. Any questions here?

**[00:06:57] — Customer**
I don't seem to have any questions. Basically, everything is great. I said during the process that you get a lot done, you're doing a good job. I'll also check it myself, I'll upload videos of basketball players and see. But I think everything is good, it basically works.

**[00:07:23] — ML Engineer**
Final Transition Status and Usefulness. We need to discuss how we'll hand over the project to you, and what you need to know about it. Basically, our product is ready for handover. The entire code is documented, all docs are there, all READMEs are there. Everything we need is in the GitHub repository. All the described ones are completed. Plus what we'll go through today. So, the entire code is ready for further development. Now, what you need to know about usage. Our ML pipeline uses an API key. It would probably be a good idea, if you want to use and maintain this product, to create your own account on OpenRouter, generate your own API key, and put it in the environment variables. But anyway, our environment variables are hidden on GitHub, so you won't find our key. We're using a free API there, so you'll only spend money on server maintenance, you know, just to keep the server and the website hosted. You won't have to spend money on the API key. That's probably all you need to know. To maintain the server and the website, you'll only need to pay for the server. That's it, we don't have any admin panels or anything, nothing like that is needed. Yeah, that's probably all you need to know to operate the diploma website. Let me think what else to say. And, probably, if you want to develop the product further, you can use our unimplemented tasks. Especially splitting into individual shots. Also, if we don't have time, you'll need to add the recommended camera angle for shot analysis. Here. Boom-boom-boom-boom. And that's it. Probably that's all. For us. Then the very last thing left is to discuss the final user acceptance test. The only thing we haven't managed to do with you yet, you can log into the website now and check how the progress bar works, meaning how it changes over time when uploading multiple videos. I sent you two different videos with athletes' shots. Try uploading the first one, then the second one, then go to the... Progress section and check if both results appeared there and if you can see the chart showing how the Progress changes. Well, how it changes. Okay.

**[00:10:57] — Customer**
The first video loaded, the second one is loading now, let's see. After loading the second video, my best score and my average result appeared in the progress section. And the last video I uploaded, the second one. Well, a chart also appeared here. Basically, everything is great.

**[00:11:43] — ML Engineer**
Super, everything is great, then this test is passed once again. I have nothing else to say. Maybe you have some advice before the final presentation, or any questions about the product, just in case.

**[00:12:00] — Customer**
I have no questions about the product. What you've done is really great. If any do appear, I'll come back to you. I hope you'll answer me even after the project ends. I wish you success at the defense. I think you should do well. But I advise you to pull up the model metrics and present them at the defense.

**[00:12:31] — ML Engineer**
Okay.

**[00:12:32] — Customer**
And in Leningrad. Here. That's probably it. Thank you so much for your work. It will be really great.

**[00:12:43] — ML Engineer**
Thank you so much. It was also a pleasure working with you as our Customer and mentor. I have absolutely no complaints about you. Everything is super. I wish you luck too, and I hope the product works well. Thank you so much.
