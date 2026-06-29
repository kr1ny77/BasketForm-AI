(function () {
    var translations = {
        en: {
            nav_upload: 'Upload', nav_results: 'Results', nav_progress: 'Progress',
            nav_export: 'Export', nav_shared: 'Shared', nav_friends: 'Friends', nav_profile: 'Profile',
            upload_title: 'Upload Video', upload_desc: 'Upload a basketball shooting video for analysis',
            upload_drag: 'Drag & drop your video here', upload_hint: 'or click to browse \u00b7 MP4, MOV, AVI \u00b7 max 100 MB', upload_msg: 'Uploading...',
            results_title: 'Analysis Results', results_desc: 'View your shooting form analyses',
            results_filter: 'Filter by filename...',             results_empty: 'No analyses yet. Upload a video to get started.',
            results_selected: 'selected', results_select_all: 'Select All', results_delete: 'Delete',
            results_confirm_delete: 'Delete selected results? This cannot be undone.',
            results_no_match: 'No results found.', results_sort_new: 'Newest first', results_sort_old: 'Oldest first',
            results_sort_high: 'Highest score', results_sort_low: 'Lowest score', results_detail_score: 'Score',
            results_detail_status: 'Status', results_detail_id: 'Video ID', results_detail_date: 'Date',
            results_feedback: 'Analysis Report', results_radar: 'Score Breakdown',
            results_phase_analysis: 'Phase Analysis', results_loading: 'Loading...', results_no_friends: 'No friends yet',
            processing_failed: 'Processing failed.',
            profile_loading: 'Loading...',
            progress_title: 'Progress', progress_desc: 'Score history for your last shots',
            progress_proc: 'Processing video...', progress_done: 'Analysis complete!', progress_view: 'View Results',
            progress_total: 'Total Shots', progress_avg: 'Average Score', progress_best: 'Best Score', progress_last: 'Last Score',
            progress_table_title: 'Last 10 Shots', progress_th_num: '#', progress_th_video: 'Video',
            progress_th_score: 'Score', progress_th_stance: 'Stance', progress_th_arm: 'Arm Angle',
            progress_th_release: 'Release', progress_th_follow: 'Follow-through', progress_th_date: 'Date',
            progress_empty: 'No shots yet. Upload a video to begin.', progress_trend: 'Score Trend',
            export_title: 'Export Report', export_desc: 'Download analysis reports as PDF',
            shared_title: 'Shared Results', shared_desc: 'Results your friends shared with you', shared_empty: 'No shared results yet.',
            friends_title: 'Friends', friends_desc: 'Find players, connect, and share your results',
            friends_find: 'Find Players', friends_search_ph: 'Search players by nickname...', friends_search_btn: 'Search',
            friends_requests: 'Pending Requests', friends_requests_empty: 'No pending requests.',
            friends_my: 'My Friends', friends_my_empty: 'No friends yet. Search above to connect!',
            friends_share_title: 'Share a Result with', friends_share_loading: 'Loading your results...',
            friends_share_empty: 'No completed results to share. Upload a video first!', friends_share_btn: 'Share',
            friends_accept: 'Accept', friends_decline: 'Decline', friends_add: 'Add Friend', friends_hint: 'Select to share',
            profile_title: 'Profile', profile_desc: 'Manage your account settings',
            profile_nickname: 'Nickname', profile_email: 'Email', profile_change_pw: 'Change Password',
            profile_current_pw: 'Current Password', profile_new_pw: 'New Password', profile_pw2: 'Confirm New Password',
            profile_pw_hint: 'Enter current password', profile_pw_hint2: 'Repeat new password', profile_pw_hint3: 'At least 6 characters',
            profile_update_pw: 'Update Password', profile_logout: 'Logout',
            profile_mismatch: 'Passwords do not match.', profile_pw_short: 'Password must be at least 6 characters.',
            profile_updated: 'Password updated successfully!',
            feedback_show: '\u25bc Show detailed analysis', feedback_hide: '\u25b2 Hide detailed analysis',
            share_btn: 'Share', export_pdf_btn: 'Export PDF',
            share_title: 'Share Result', share_select: 'Select a friend', share_loading: 'Loading friends...', share_send: 'Send',
            friends_search_empty: 'No users found.',
            login_subtitle: 'Sign in to your account', login_email: 'Email', login_password: 'Password',
            login_email_ph: 'Enter your email', login_password_ph: 'Enter your password',
            login_btn: 'Sign In', login_no_account: "Don't have an account?", login_signup_link: 'Sign Up',
            login_error: 'Connection error. Please try again.',
            signup_subtitle: 'Create your account', signup_nickname: 'Nickname', signup_email: 'Email',
            signup_password: 'Password', signup_confirm: 'Confirm Password',
            signup_nickname_ph: 'Choose a unique nickname', signup_email_ph: 'Enter your email',
            signup_password_ph: 'At least 6 characters', signup_confirm_ph: 'Repeat your password',
            signup_btn: 'Sign Up', signup_has_account: 'Already have an account?', signup_login_link: 'Sign In',
            signup_mismatch: 'Passwords do not match.', signup_short: 'Password must be at least 6 characters.',
            signup_nick_short: 'Nickname must be at least 3 characters.',
        },
        ru: {
            nav_upload: '\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430', nav_results: '\u0420\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u044b', nav_progress: '\u041f\u0440\u043e\u0433\u0440\u0435\u0441\u0441',
            nav_export: '\u042d\u043a\u0441\u043f\u043e\u0440\u0442', nav_shared: '\u041e\u0431\u0449\u0438\u0435', nav_friends: '\u0414\u0440\u0443\u0437\u044c\u044f', nav_profile: '\u041f\u0440\u043e\u0444\u0438\u043b\u044c',
            upload_title: '\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430 \u0432\u0438\u0434\u0435\u043e', upload_desc: '\u0417\u0430\u0433\u0440\u0443\u0437\u0438\u0442\u0435 \u0432\u0438\u0434\u0435\u043e \u0431\u0430\u0441\u043a\u0435\u0442\u0431\u043e\u043b\u044c\u043d\u043e\u0433\u043e \u0431\u0440\u043e\u0441\u043a\u0430 \u0434\u043b\u044f \u0430\u043d\u0430\u043b\u0438\u0437\u0430',
            upload_drag: '\u041f\u0435\u0440\u0435\u0442\u0430\u0449\u0438\u0442\u0435 \u0432\u0438\u0434\u0435\u043e \u0441\u044e\u0434\u0430', upload_hint: '\u0438\u043b\u0438 \u043d\u0430\u0436\u043c\u0438\u0442\u0435 \u0434\u043b\u044f \u0432\u044b\u0431\u043e\u0440\u0430 \u00b7 MP4, MOV, AVI \u00b7 \u0434\u043e 100 \u041c\u0411', upload_msg: '\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430...',
            results_title: '\u0420\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u044b \u0430\u043d\u0430\u043b\u0438\u0437\u0430', results_desc: '\u041f\u0440\u043e\u0441\u043c\u043e\u0442\u0440 \u0430\u043d\u0430\u043b\u0438\u0437\u0430 \u0432\u0430\u0448\u0435\u0433\u043e \u0431\u0440\u043e\u0441\u043a\u0430',
            results_filter: '\u0424\u0438\u043b\u044c\u0442\u0440 \u043f\u043e \u0438\u043c\u0435\u043d\u0438...',             results_empty: '\u041f\u043e\u043a\u0430 \u043d\u0435\u0442 \u0440\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u043e\u0432. \u0417\u0430\u0433\u0440\u0443\u0437\u0438\u0442\u0435 \u0432\u0438\u0434\u0435\u043e.',
            results_selected: '\u0432\u044b\u0431\u0440\u0430\u043d\u043e', results_select_all: '\u0412\u044b\u0431\u0440\u0430\u0442\u044c \u0432\u0441\u0435', results_delete: '\u0423\u0434\u0430\u043b\u0438\u0442\u044c',
            results_delete_selected: '\u0423\u0434\u0430\u043b\u0438\u0442\u044c', results_cancel: '\u041e\u0442\u043c\u0435\u043d\u0430',
            results_confirm_delete: '\u0423\u0434\u0430\u043b\u0438\u0442\u044c \u0432\u044b\u0431\u0440\u0430\u043d\u043d\u044b\u0435? \u0414\u0435\u0439\u0441\u0442\u0432\u0438\u0435 \u043d\u0435\u043e\u0431\u0440\u0430\u0442\u0438\u043c\u043e.',
            results_no_match: '\u041d\u0438\u0447\u0435\u0433\u043e \u043d\u0435 \u043d\u0430\u0439\u0434\u0435\u043d\u043e.', results_sort_new: '\u0421\u043d\u0430\u0447\u0430\u043b\u0430 \u043d\u043e\u0432\u044b\u0435', results_sort_old: '\u0421\u043d\u0430\u0447\u0430\u043b\u0430 \u0441\u0442\u0430\u0440\u044b\u0435',
            results_sort_high: '\u0412\u044b\u0441\u043e\u043a\u0438\u0439 \u0431\u0430\u043b\u043b', results_sort_low: '\u041d\u0438\u0437\u043a\u0438\u0439 \u0431\u0430\u043b\u043b', results_detail_score: '\u0411\u0430\u043b\u043b',
            results_detail_status: '\u0421\u0442\u0430\u0442\u0443\u0441', results_detail_id: 'ID \u0432\u0438\u0434\u0435\u043e', results_detail_date: '\u0414\u0430\u0442\u0430',
            results_feedback: '\u041e\u0442\u0447\u0435\u0442 \u0430\u043d\u0430\u043b\u0438\u0437\u0430', results_radar: '\u0420\u0430\u0437\u0431\u0438\u0435\u043d\u0438\u0435 \u0431\u0430\u043b\u043b\u0430',
            results_phase_analysis: '\u0410\u043d\u0430\u043b\u0438\u0437 \u0444\u0430\u0437', results_loading: '\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430...', results_no_friends: '\u0414\u0440\u0443\u0437\u0435\u0439 \u043f\u043e\u043a\u0430 \u043d\u0435\u0442',
            processing_failed: '\u041e\u0448\u0438\u0431\u043a\u0430 \u043e\u0431\u0440\u0430\u0431\u043e\u0442\u043a\u0438.',
            profile_loading: '\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430...',
            progress_title: '\u041f\u0440\u043e\u0433\u0440\u0435\u0441\u0441', progress_desc: '\u0418\u0441\u0442\u043e\u0440\u0438\u044f \u0431\u0430\u043b\u043b\u043e\u0432 \u0437\u0430 \u043f\u043e\u0441\u043b\u0435\u0434\u043d\u0438\u0435 \u0431\u0440\u043e\u0441\u043a\u0438',
            progress_proc: '\u041e\u0431\u0440\u0430\u0431\u043e\u0442\u043a\u0430 \u0432\u0438\u0434\u0435\u043e...', progress_done: '\u0410\u043d\u0430\u043b\u0438\u0437 \u0437\u0430\u0432\u0435\u0440\u0448\u0435\u043d!', progress_view: '\u0420\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u044b',
            progress_total: '\u0412\u0441\u0435\u0433\u043e \u0431\u0440\u043e\u0441\u043a\u043e\u0432', progress_avg: '\u0421\u0440\u0435\u0434\u043d\u0438\u0439 \u0431\u0430\u043b\u043b', progress_best: '\u041b\u0443\u0447\u0448\u0438\u0439 \u0431\u0430\u043b\u043b', progress_last: '\u041f\u043e\u0441\u043b\u0435\u0434\u043d\u0438\u0439 \u0431\u0430\u043b\u043b',
            progress_table_title: '\u041f\u043e\u0441\u043b\u0435\u0434\u043d\u0438\u0435 10 \u0431\u0440\u043e\u0441\u043a\u043e\u0432', progress_th_num: '#', progress_th_video: '\u0412\u0438\u0434\u0435\u043e',
            progress_th_score: '\u0411\u0430\u043b\u043b', progress_th_stance: '\u0421\u0442\u043e\u0439\u043a\u0430', progress_th_arm: '\u0423\u0433\u043e\u043b \u0440\u0443\u043a\u0438',
            progress_th_release: '\u0412\u044b\u043f\u0443\u0441\u043a', progress_th_follow: '\u0417\u0430\u0432\u0435\u0440\u0448\u0435\u043d\u0438\u0435', progress_th_date: '\u0414\u0430\u0442\u0430',
            progress_empty: '\u0411\u0440\u043e\u0441\u043a\u043e\u0432 \u043f\u043e\u043a\u0430 \u043d\u0435\u0442. \u0417\u0430\u0433\u0440\u0443\u0437\u0438\u0442\u0435 \u0432\u0438\u0434\u0435\u043e.', progress_trend: '\u0414\u0438\u043d\u0430\u043c\u0438\u043a\u0430 \u0431\u0430\u043b\u043b\u043e\u0432',
            export_title: '\u042d\u043a\u0441\u043f\u043e\u0440\u0442 \u043e\u0442\u0447\u0435\u0442\u0430', export_desc: '\u0421\u043a\u0430\u0447\u0430\u0442\u044c \u043e\u0442\u0447\u0435\u0442 \u0430\u043d\u0430\u043b\u0438\u0437\u0430 \u0432 PDF',
            shared_title: '\u041e\u0431\u0449\u0438\u0435 \u0440\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u044b', shared_desc: '\u0420\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u044b, \u043a\u043e\u0442\u043e\u0440\u044b\u0435 \u043f\u043e\u0434\u0435\u043b\u0438\u043b\u0438\u0441\u044c \u0434\u0440\u0443\u0437\u044c\u044f', shared_empty: '\u041f\u043e\u043a\u0430 \u043d\u0435\u0442 \u043e\u0431\u0449\u0438\u0445 \u0440\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u043e\u0432.',
            friends_title: '\u0414\u0440\u0443\u0437\u044c\u044f', friends_desc: '\u041d\u0430\u0439\u0434\u0438\u0442\u0435 \u0438\u0433\u0440\u043e\u043a\u043e\u0432, \u0441\u0432\u044f\u0436\u0438\u0442\u0435\u0441\u044c \u0438 \u043f\u043e\u0434\u0435\u043b\u0438\u0442\u0435\u0441\u044c \u0440\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u0430\u043c\u0438',
            friends_find: '\u041d\u0430\u0439\u0442\u0438 \u0438\u0433\u0440\u043e\u043a\u043e\u0432', friends_search_ph: '\u041f\u043e\u0438\u0441\u043a \u043f\u043e \u043f\u0441\u0435\u0432\u0434\u043e\u043d\u0438\u043c\u0443...', friends_search_btn: '\u041f\u043e\u0438\u0441\u043a',
            friends_requests: '\u0417\u0430\u044f\u0432\u043a\u0438', friends_requests_empty: '\u041d\u0435\u0442 \u0437\u0430\u044f\u0432\u043e\u043a.',
            friends_my: '\u041c\u043e\u0438 \u0434\u0440\u0443\u0437\u044c\u044f', friends_my_empty: '\u041f\u043e\u043a\u0430 \u043d\u0435\u0442 \u0434\u0440\u0443\u0437\u0435\u0439. \u041f\u043e\u0438\u0449\u0438\u0442\u0435 \u0438\u0433\u0440\u043e\u043a\u043e\u0432 \u0432\u044b\u0448\u0435!',
            friends_share_title: '\u041f\u043e\u0434\u0435\u043b\u0438\u0442\u044c\u0441\u044f \u0440\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u043e\u043c \u0441', friends_share_loading: '\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430 \u0432\u0430\u0448\u0438\u0445 \u0440\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u043e\u0432...',
            friends_share_empty: '\u041d\u0435\u0442 \u0437\u0430\u0432\u0435\u0440\u0448\u0451\u043d\u043d\u044b\u0445 \u0440\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u043e\u0432. \u0417\u0430\u0433\u0440\u0443\u0437\u0438\u0442\u0435 \u0432\u0438\u0434\u0435\u043e!', friends_share_btn: '\u041f\u043e\u0434\u0435\u043b\u0438\u0442\u044c\u0441\u044f',
            friends_accept: '\u041f\u0440\u0438\u043d\u044f\u0442\u044c', friends_decline: '\u041e\u0442\u043a\u043b\u043e\u043d\u0438\u0442\u044c', friends_add: '\u0414\u043e\u0431\u0430\u0432\u0438\u0442\u044c', friends_hint: '\u0412\u044b\u0431\u0435\u0440\u0438\u0442\u0435 \u0434\u043b\u044f \u0448\u0435\u0440\u0438\u043d\u0433\u0430',
            profile_title: '\u041f\u0440\u043e\u0444\u0438\u043b\u044c', profile_desc: '\u0423\u043f\u0440\u0430\u0432\u043b\u0435\u043d\u0438\u0435 \u043d\u0430\u0441\u0442\u0440\u043e\u0439\u043a\u0430\u043c\u0438',
            profile_nickname: '\u041f\u0441\u0435\u0432\u0434\u043e\u043d\u0438\u043c', profile_email: '\u042d\u043b. \u043f\u043e\u0447\u0442\u0430', profile_change_pw: '\u0421\u043c\u0435\u043d\u0438\u0442\u044c \u043f\u0430\u0440\u043e\u043b\u044c',
            profile_current_pw: '\u0422\u0435\u043a\u0443\u0449\u0438\u0439 \u043f\u0430\u0440\u043e\u043b\u044c', profile_new_pw: '\u041d\u043e\u0432\u044b\u0439 \u043f\u0430\u0440\u043e\u043b\u044c', profile_pw2: '\u041f\u043e\u0432\u0442\u043e\u0440\u0438\u0442\u0435 \u043f\u0430\u0440\u043e\u043b\u044c',
            profile_pw_hint: '\u0412\u0432\u0435\u0434\u0438\u0442\u0435 \u0442\u0435\u043a\u0443\u0449\u0438\u0439 \u043f\u0430\u0440\u043e\u043b\u044c', profile_pw_hint2: '\u041f\u043e\u0432\u0442\u043e\u0440\u0438\u0442\u0435 \u043f\u0430\u0440\u043e\u043b\u044c', profile_pw_hint3: '\u041c\u0438\u043d\u0438\u043c\u0443\u043c 6 \u0441\u0438\u043c\u0432\u043e\u043b\u043e\u0432',
            profile_update_pw: '\u041e\u0431\u043d\u043e\u0432\u0438\u0442\u044c \u043f\u0430\u0440\u043e\u043b\u044c', profile_logout: '\u0412\u044b\u0439\u0442\u0438',
            profile_mismatch: '\u041f\u0430\u0440\u043e\u043b\u0438 \u043d\u0435 \u0441\u043e\u0432\u043f\u0430\u0434\u0430\u044e\u0442.', profile_pw_short: '\u041f\u0430\u0440\u043e\u043b\u044c \u0434\u043e\u043b\u0436\u0435\u043d \u0431\u044b\u0442\u044c \u043d\u0435 \u043c\u0435\u043d\u0435\u0435 6 \u0441\u0438\u043c\u0432\u043e\u043b\u043e\u0432.',
            profile_updated: '\u041f\u0430\u0440\u043e\u043b\u044c \u0443\u0441\u043f\u0435\u0448\u043d\u043e \u043e\u0431\u043d\u043e\u0432\u043b\u0451\u043d!',
            feedback_show: '\u25bc \u041f\u043e\u043a\u0430\u0437\u0430\u0442\u044c \u043f\u043e\u0434\u0440\u043e\u0431\u043d\u044b\u0439 \u0430\u043d\u0430\u043b\u0438\u0437', feedback_hide: '\u25b2 \u0421\u043a\u0440\u044b\u0442\u044c \u043f\u043e\u0434\u0440\u043e\u0431\u043d\u044b\u0439 \u0430\u043d\u0430\u043b\u0438\u0437',
            share_btn: '\u041f\u043e\u0434\u0435\u043b\u0438\u0442\u044c\u0441\u044f', export_pdf_btn: '\u042d\u043a\u0441\u043f\u043e\u0440\u0442 PDF',
            share_title: '\u041f\u043e\u0434\u0435\u043b\u0438\u0442\u044c\u0441\u044f \u0440\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u043e\u043c', share_select: '\u0412\u044b\u0431\u0435\u0440\u0438\u0442\u0435 \u0434\u0440\u0443\u0433\u0430', share_loading: '\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430...', share_send: '\u041e\u0442\u043f\u0440\u0430\u0432\u0438\u0442\u044c',
            friends_search_empty: '\u041f\u043e\u043b\u044c\u0437\u043e\u0432\u0430\u0442\u0435\u043b\u0435\u0439 \u043d\u0435 \u043d\u0430\u0439\u0434\u0435\u043d\u043e.',
            login_subtitle: '\u0412\u043e\u0439\u0434\u0438\u0442\u0435 \u0432 \u0430\u043a\u043a\u0430\u0443\u043d\u0442', login_email: '\u042d\u043b. \u043f\u043e\u0447\u0442\u0430', login_password: '\u041f\u0430\u0440\u043e\u043b\u044c',
            login_email_ph: '\u0412\u0432\u0435\u0434\u0438\u0442\u0435 \u0432\u0430\u0448\u0443 \u043f\u043e\u0447\u0442\u0443', login_password_ph: '\u0412\u0432\u0435\u0434\u0438\u0442\u0435 \u043f\u0430\u0440\u043e\u043b\u044c',
            login_btn: '\u0412\u043e\u0439\u0442\u0438', login_no_account: '\u041d\u0435\u0442 \u0430\u043a\u043a\u0430\u0443\u043d\u0442\u0430?', login_signup_link: '\u0417\u0430\u0440\u0435\u0433\u0438\u0441\u0442\u0440\u0438\u0440\u043e\u0432\u0430\u0442\u044c\u0441\u044f',
            login_error: '\u041e\u0448\u0438\u0431\u043a\u0430 \u0441\u043e\u0435\u0434\u0438\u043d\u0435\u043d\u0438\u044f. \u041f\u043e\u043f\u0440\u043e\u0431\u0443\u0439\u0442\u0435 \u0441\u043d\u043e\u0432\u0430.',
            signup_subtitle: '\u0421\u043e\u0437\u0434\u0430\u0439\u0442\u0435 \u0430\u043a\u043a\u0430\u0443\u043d\u0442', signup_nickname: '\u041f\u0441\u0435\u0432\u0434\u043e\u043d\u0438\u043c', signup_email: '\u042d\u043b. \u043f\u043e\u0447\u0442\u0430',
            signup_password: '\u041f\u0430\u0440\u043e\u043b\u044c', signup_confirm: '\u041f\u043e\u0432\u0442\u043e\u0440\u0438\u0442\u0435 \u043f\u0430\u0440\u043e\u043b\u044c',
            signup_nickname_ph: '\u0412\u044b\u0431\u0435\u0440\u0438\u0442\u0435 \u0443\u043d\u0438\u043a\u0430\u043b\u044c\u043d\u044b\u0439 \u043f\u0441\u0435\u0432\u0434\u043e\u043d\u0438\u043c', signup_email_ph: '\u0412\u0432\u0435\u0434\u0438\u0442\u0435 \u0432\u0430\u0448\u0443 \u043f\u043e\u0447\u0442\u0443',
            signup_password_ph: '\u041c\u0438\u043d\u0438\u043c\u0443\u043c 6 \u0441\u0438\u043c\u0432\u043e\u043b\u043e\u0432', signup_confirm_ph: '\u041f\u043e\u0432\u0442\u043e\u0440\u0438\u0442\u0435 \u043f\u0430\u0440\u043e\u043b\u044c',
            signup_btn: '\u0417\u0430\u0440\u0435\u0433\u0438\u0441\u0442\u0440\u0438\u0440\u043e\u0432\u0430\u0442\u044c\u0441\u044f', signup_has_account: '\u0423\u0436\u0435 \u0435\u0441\u0442\u044c \u0430\u043a\u043a\u0430\u0443\u043d\u0442?', signup_login_link: '\u0412\u043e\u0439\u0442\u0438',
            signup_mismatch: '\u041f\u0430\u0440\u043e\u043b\u0438 \u043d\u0435 \u0441\u043e\u0432\u043f\u0430\u0434\u0430\u044e\u0442.', signup_short: '\u041f\u0430\u0440\u043e\u043b\u044c \u0434\u043e\u043b\u0436\u0435\u043d \u0431\u044b\u0442\u044c \u043d\u0435 \u043c\u0435\u043d\u0435\u0435 6 \u0441\u0438\u043c\u0432\u043e\u043b\u043e\u0432.',
            signup_nick_short: '\u041f\u0441\u0435\u0432\u0434\u043e\u043d\u0438\u043c \u0434\u043e\u043b\u0436\u0435\u043d \u0431\u044b\u0442\u044c \u043d\u0435 \u043c\u0435\u043d\u0435\u0435 3 \u0441\u0438\u043c\u0432\u043e\u043b\u043e\u0432.',
        }
    };

    var currentLang = localStorage.getItem('lang') || 'en';

    function applyLang(lang) {
        currentLang = lang;
        localStorage.setItem('lang', lang);
        var t = translations[lang];
        if (!t) return;

        document.querySelectorAll('[data-i18n]').forEach(function (el) {
            var key = el.getAttribute('data-i18n');
            if (t[key]) el.textContent = t[key];
        });
        document.querySelectorAll('[data-i18n-placeholder]').forEach(function (el) {
            var key = el.getAttribute('data-i18n-placeholder');
            if (t[key]) el.placeholder = t[key];
        });

        var selector = document.getElementById('langSelect');
        if (selector) selector.value = lang;
    }

    window.getLang = function () { return currentLang; };
    window.applyLang = applyLang;
    window.translations = translations;

    document.addEventListener('DOMContentLoaded', function () {
        applyLang(currentLang);
    });
})();
