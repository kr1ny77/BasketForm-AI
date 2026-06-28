2026-06-28
Participants: **ML Engineer**, **Customer**
Length: 00:08:40

[00:00:01.830] — **ML Engineer**
[redacted], good morning. Today we'll be conducting another sprint review. It also turns out we'll be executing user acceptance tests. Before we begin, I'd like to ask your permission to record our meeting for subsequent transcription and posting to the repository.

[00:00:28.530] — **Customer**
Good morning. Yes, I give my permission to record this meeting, transcribe it, and post it to the repository. Let's get started and run the tests later.

[00:00:39.955] — **ML Engineer**
I suggest we begin the interview now. Last week, we planned to learn how to track the ball, improve the algorithm based on that, and integrate LLM so it would provide related feedback. We've achieved all of this. Using a fine-tuned YOLO V8, we're now tracking the ball, and we downloaded the QWEN 2.5 model locally. It has 3 billion parameters, so it's not particularly heavy, but it provides a decent answer. We also have a grade for each phase, and based on the estimate for each phase, LLM produces a result. This is about the Sprint review. 

[00:01:35.230] — **Customer**
Okay, great that you got everything done. 

[00:01:40.000] — **ML Engineer**
What can we do next week? We could do it, let's say, as we wanted. One of the suggestions was to integrate videos of professional athletes so we could compare our shots to them. And a friends system. We already have an account system; we could implement a friends system. What do you think about that?

[00:02:13.930] — **Customer**
To be honest, I'm not entirely sure how much a young athlete will understand from watching a good athlete's video where they're making mistakes. That's why I think the friends feature is more useful. But if I'm missing something, please explain. Because you can simply watch an athlete's video, for example, both online and at a match. What's important to us is an explanation of the details.

[00:02:41.570] — **ML Engineer**
Well, yes, that's right. Okay, then we'll prioritize the friend system. Okay, that's all for the Sprint Review. Now for User Substance Tests. I'll be recording the steps you need to complete and telling you the expected result. Your task will be to tell me whether you got the same expected result or a different one. While I'm at it, I'll take a little time to go to the website; I'll open these scenarios. So, are you ready?

[00:03:30.510] — **Customer**
Yes. Let's get started.

[00:03:32.850] — **ML Engineer**
Then the first scenario. So, the prerequisites. You need to have access to the website and a supported MP4, MOV, or AVI video file with the video of you throwing. So, step by step. You need to go to the main page. Yeah. Find the upload section and drag the video file into the area, or upload it using the video button.

[00:04:06.272] — **Customer**
Okay, okay. Just a second.

[00:04:12.110] — **ML Engineer**
Now you'll have to wait a little while for the video to process. I think 5-7 seconds would work.

[00:04:20.790] — **Customer**
Well, yes, it's uploading. Now we'll wait for it to process. Okay, everything's ready.

[00:04:33.850] — **ML Engineer**
Now, let's review the analysis results. As expected. Your file should have uploaded successfully. While the video was loading, a processing progress bar should have been displayed. That is, it should be 20-40%. And you should now see the equipment ratings and recommendations for this idea for improvement.

[00:05:02.038] — **Customer**
Yes, I see that.

[00:05:04.029] — **ML Engineer**
Everything's great. The next User Service Test is exporting the analysis report as a PDF for offline viewing. So, you also need to go to the website and go to the "Export" heading at the top. There. You need to select one test, but you can also select the last one, for example. And export it.

[00:05:45.689] — **Customer**
Uh-huh, good.

[00:05:47.970] — **ML Engineer**
Our expected result is as follows. The browser successfully initiates the file download without errors. And the downloaded PDF can be opened and viewed. And it's structured properly.

[00:06:01.150] — **Customer**
Yes, the video is downloading. A little slowly, I think because of the internet connection. But everything is downloading fine. Good.

[00:06:10.629] — **ML Engineer**
PDF, what do you mean?

[00:06:12.517] — **Customer**
Ah, yes, yes, PDF. And it's basically hidden with the throw report.

[00:06:22.930] — **ML Engineer**
Okay, okay. Okay, then we move on to the next third one. User Assertance Test. It's related to evaluating the interactive user interface and voiceover for canvas animation. Follow the step-by-step instructions to go to our website's main page.

[00:06:44.622] — **Customer**
Got it.

[00:06:46.050] — **ML Engineer**
I was observing the basketball-themed background animation. We've already implemented all the floating objects.

[00:06:53.011] — **Customer**
Basketballs are falling.

[00:06:54.873] — **ML Engineer**
Here.

[00:06:56.971] — **Customer**
Yes, they work very well.

[00:06:58.290] — **ML Engineer**
Move the mouse cursor around the screen and click on the main elements. The user interface, its buttons, the loading area, and so on. Make sure they don't open or break due to background animation.

[00:07:13.150] — **Customer**
I didn't even know your background animation was clickable. That's really cool. Everything else, as far as I can see, works too.

[00:07:23.810] — **ML Engineer**
Let's announce the expected telescope and see if it's recessed or not. The background animation loads smoothly, displaying the specified thematic elements. Basketballs. It dynamically reacts to mouse movement during the animation without any significant FPS drops or browser freezes. After all, the active Convas animation remains strictly in the background, all foreground elements remain fully clickable, readable, and unaffected by the animation. Yep. That's great. Then, did you pass all the U-Reception tests we created? No errors? Yep. Excellent. It also asked for permission to record what I was recording, to record these user tests separately. Yes.

[00:08:21.632] — **Customer**
and give their permission. Everything is great.

[00:08:26.308] — **ML Engineer**
Aha. Then that's all for today. Do you have any questions left for us?

[00:08:33.271] — **Customer**
No, no questions left.

[00:08:35.429] — **ML Engineer**
Nice then.
