# Meeting Transcript, 12 Jule
# Length: 5 min 39 sec

**Participants:** ML Engineer, Customer

[00:00:01] — ML Engineer
[redacted], hi! Today we'll look again at what we managed to get done over the last sprint. We'll discuss what can be fixed, what can be added, and how to prepare for the presentation in the final week. Also, we'll go through the new user acceptance tests with you. Before we begin, as usual, I'll ask for your permission to record, transcribe, and post this transcript in our repository.

[00:00:33] — Customer
Good morning. Yes, of course, I give my consent. Let's begin.

[00:00:37] — ML Engineer
Super. So, for this sprint, our plan was to implement automatic splitting of a large video into separate clips upon user upload. However, during development, we ran into various pitfalls, so we realized we wouldn't be able to implement it in just one week. To avoid breaking anything, we shifted our focus to fixing another bug. Previously, it was possible to upload a random video that didn't even contain a basketball player or a person in general, and the site would still try to perform some analysis and output results. Now, if you upload a video without a basketball player throwing a ball—meaning no person is identified in the video—the scoring and analysis will not be performed. It will simply display a message stating that it couldn't be analyzed.

[00:01:38] — Customer
Great. Of course, it's a pity that your initial idea didn't work out as easily as we planned, but you can implement it after the course is over. And it's great that you came up with another similar alternative to improve human recognition in the video, so you still made progress toward this feature's functionality. You guys did a great job in this regard.

[00:02:06] — ML Engineer
Okay. Next week, we probably won't set any serious goals since it's the final week. We'll just polish things up, look around, try to find some bugs, and attempt to optimize the site to ensure it runs smoothly. Maybe we'll refactor the code or do something similar.

[00:02:30] — Customer
Yes, I think that's great. Leaving the last week for testing. Actually, a week or two ago, we discussed that you had fully completed the Definition of Done that I gave you, and you've essentially finished all your main tasks. After that, you were even working beyond the project scope. So now you have time to double-check everything, make sure it all works, and finish the course with flying colors.

[00:02:58] — ML Engineer
Okay, thanks, [redacted]. Let's move on to the user acceptance tests now. They will focus on our first new features. The first one is about the inability to upload a video without a person. I sent you some videos in advance that don't contain a person. You can try uploading one for analysis now and see what feedback the site gives you.

[00:03:30] — Customer
Yes, I'm on it now. So, I see the video is uploading, and now I see it's returning an inability to analyze due to the absence of a person. That's really great.

[00:04:01] — ML Engineer
That's exactly how it should be. Excellent. And the second one. Continuing from last week, when we added the friends feature, try sharing your result with your friends. To do this, you need to go to the Friends tab. Then, at the bottom, in the list of friends you already have, hover over each of them, and a "Share" button will appear. You can use it to share your result. Try doing that now.

[00:04:48] — Customer
Yes, it works. That's really great.

[00:04:51] — ML Engineer
Everything is super, everything works. Then that's all for today. Do you have any questions, instructions, or wishes? No? Well, basically everything is great.

[00:05:06] — Customer
I like how the site works. I hope you guys get some good rest. I'll review your site again at the beginning of the week. Maybe I'll find some bugs, and then I'll let you know so you can definitely fix them by next week. But overall, I hope you'll do some thorough testing, test your site well, and ensure it works properly. Basically, all the functionality needs to be completed by now. That's great.

[00:05:33] — ML Engineer
Super, thanks. See you at the next meeting. Bye-bye.

[00:05:37] — Customer
See you at the next meeting.
