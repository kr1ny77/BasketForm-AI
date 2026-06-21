(function () {
    var translations = {
        en: {
            nav_upload: 'Upload',
            nav_results: 'Results',
            nav_progress: 'Progress',
            nav_export: 'Export',
            nav_profile: 'Profile',
            upload_title: 'Upload Video',
            upload_desc: 'Upload a basketball shooting video for analysis',
            upload_drag: 'Drag & drop your video here',
            upload_hint: 'or click to browse \u00b7 MP4, MOV, AVI \u00b7 max 100 MB',
            upload_msg: 'Uploading...',
            results_title: 'Analysis Results',
            results_desc: 'View your shooting form analyses',
            results_filter: 'Filter by filename...',
            results_empty: 'No analyses yet. Upload a video to get started.',
            results_no_match: 'No results found.',
            results_sort_new: 'Newest first',
            results_sort_old: 'Oldest first',
            results_sort_high: 'Highest score',
            results_sort_low: 'Lowest score',
            results_view: 'View',
            results_detail_score: 'Score',
            results_detail_status: 'Status',
            results_detail_id: 'Video ID',
            results_detail_date: 'Date',
            results_feedback: 'Analysis Report',
            results_radar: 'Score Breakdown',
            progress_title: 'Progress',
            progress_desc: 'Score history for your last shots',
            progress_proc: 'Processing video...',
            progress_done: 'Analysis complete!',
            progress_view: 'View Results',
            progress_total: 'Total Shots',
            progress_avg: 'Average Score',
            progress_best: 'Best Score',
            progress_last: 'Last Score',
            progress_table_title: 'Last 10 Shots',
            progress_th_num: '#',
            progress_th_video: 'Video',
            progress_th_score: 'Score',
            progress_th_stance: 'Stance',
            progress_th_arm: 'Arm Angle',
            progress_th_release: 'Release',
            progress_th_follow: 'Follow-through',
            progress_th_date: 'Date',
            progress_empty: 'No shots yet. Upload a video to begin.',
            progress_trend: 'Score Trend',
            export_title: 'Export Report',
            export_desc: 'Download analysis reports as PDF',
            export_select: 'Select a completed analysis',
            export_loading: 'Loading...',
            export_none: 'No completed analyses',
            export_choose: 'Choose...',
            export_btn: 'Download PDF',
            profile_title: 'Profile',
            profile_desc: 'Manage your account settings',
            profile_name: 'Name',
            profile_email: 'Email',
            profile_pw: 'New Password',
            profile_pw2: 'Confirm New Password',
            profile_save: 'Save Changes',
            profile_saved: 'Profile saved (demo).',
            profile_mismatch: 'Passwords do not match.',
            profile_pw_hint: 'Leave blank to keep current',
            profile_pw_hint2: 'Repeat new password',
            feedback_show: '\u25bc Show detailed analysis',
            feedback_hide: '\u25b2 Hide detailed analysis',
        },
        ru: {
            nav_upload: '\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430',
            nav_results: '\u0420\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u044b',
            nav_progress: '\u041f\u0440\u043e\u0433\u0440\u0435\u0441\u0441',
            nav_export: '\u042d\u043a\u0441\u043f\u043e\u0440\u0442',
            nav_profile: '\u041f\u0440\u043e\u0444\u0438\u043b\u044c',
            upload_title: '\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430 \u0432\u0438\u0434\u0435\u043e',
            upload_desc: '\u0417\u0430\u0433\u0440\u0443\u0437\u0438\u0442\u0435 \u0432\u0438\u0434\u0435\u043e \u0431\u0430\u0441\u043a\u0435\u0442\u0431\u043e\u043b\u044c\u043d\u043e\u0433\u043e \u0431\u0440\u043e\u0441\u043a\u0430 \u0434\u043b\u044f \u0430\u043d\u0430\u043b\u0438\u0437\u0430',
            upload_drag: '\u041f\u0435\u0440\u0435\u0442\u0430\u0449\u0438\u0442\u0435 \u0432\u0438\u0434\u0435\u043e \u0441\u044e\u0434\u0430',
            upload_hint: '\u0438\u043b\u0438 \u043d\u0430\u0436\u043c\u0438\u0442\u0435 \u0434\u043b\u044f \u0432\u044b\u0431\u043e\u0440\u0430 \u00b7 MP4, MOV, AVI \u00b7 \u0434\u043e 100 \u041c\u0411',
            upload_msg: '\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430...',
            results_title: '\u0420\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u044b \u0430\u043d\u0430\u043b\u0438\u0437\u0430',
            results_desc: '\u041f\u0440\u043e\u0441\u043c\u043e\u0442\u0440 \u0430\u043d\u0430\u043b\u0438\u0437\u0430 \u0432\u0430\u0448\u0435\u0433\u043e \u0431\u0440\u043e\u0441\u043a\u0430',
            results_filter: '\u0424\u0438\u043b\u044c\u0442\u0440 \u043f\u043e \u0438\u043c\u0435\u043d\u0438...',
            results_empty: '\u041f\u043e\u043a\u0430 \u043d\u0435\u0442 \u0440\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u043e\u0432. \u0417\u0430\u0433\u0440\u0443\u0437\u0438\u0442\u0435 \u0432\u0438\u0434\u0435\u043e.',
            results_no_match: '\u041d\u0438\u0447\u0435\u0433\u043e \u043d\u0435 \u043d\u0430\u0439\u0434\u0435\u043d\u043e.',
            results_sort_new: '\u0421\u043d\u0430\u0447\u0430\u043b\u0430 \u043d\u043e\u0432\u044b\u0435',
            results_sort_old: '\u0421\u043d\u0430\u0447\u0430\u043b\u0430 \u0441\u0442\u0430\u0440\u044b\u0435',
            results_sort_high: '\u0412\u044b\u0441\u043e\u043a\u0438\u0439 \u0431\u0430\u043b\u043b',
            results_sort_low: '\u041d\u0438\u0437\u043a\u0438\u0439 \u0431\u0430\u043b\u043b',
            results_view: '\u041f\u043e\u0434\u0440\u043e\u0431\u043d\u0435\u0435',
            results_detail_score: '\u0411\u0430\u043b\u043b',
            results_detail_status: '\u0421\u0442\u0430\u0442\u0443\u0441',
            results_detail_id: 'ID \u0432\u0438\u0434\u0435\u043e',
            results_detail_date: '\u0414\u0430\u0442\u0430',
            results_feedback: '\u041e\u0442\u0447\u0435\u0442 \u0430\u043d\u0430\u043b\u0438\u0437\u0430',
            results_radar: '\u0420\u0430\u0437\u0431\u0438\u0435\u043d\u0438\u0435 \u0431\u0430\u043b\u043b\u0430',
            progress_title: '\u041f\u0440\u043e\u0433\u0440\u0435\u0441\u0441',
            progress_desc: '\u0418\u0441\u0442\u043e\u0440\u0438\u044f \u0431\u0430\u043b\u043b\u043e\u0432 \u0437\u0430 \u043f\u043e\u0441\u043b\u0435\u0434\u043d\u0438\u0435 \u0431\u0440\u043e\u0441\u043a\u0438',
            progress_proc: '\u041e\u0431\u0440\u0430\u0431\u043e\u0442\u043a\u0430 \u0432\u0438\u0434\u0435\u043e...',
            progress_done: '\u0410\u043d\u0430\u043b\u0438\u0437 \u0437\u0430\u0432\u0435\u0440\u0448\u0435\u043d!',
            progress_view: '\u0420\u0435\u0437\u0443\u043b\u044c\u0442\u0430\u0442\u044b',
            progress_total: '\u0412\u0441\u0435\u0433\u043e \u0431\u0440\u043e\u0441\u043a\u043e\u0432',
            progress_avg: '\u0421\u0440\u0435\u0434\u043d\u0438\u0439 \u0431\u0430\u043b\u043b',
            progress_best: '\u041b\u0443\u0447\u0448\u0438\u0439 \u0431\u0430\u043b\u043b',
            progress_last: '\u041f\u043e\u0441\u043b\u0435\u0434\u043d\u0438\u0439 \u0431\u0430\u043b\u043b',
            progress_table_title: '\u041f\u043e\u0441\u043b\u0435\u0434\u043d\u0438\u0435 10 \u0431\u0440\u043e\u0441\u043a\u043e\u0432',
            progress_th_num: '#',
            progress_th_video: '\u0412\u0438\u0434\u0435\u043e',
            progress_th_score: '\u0411\u0430\u043b\u043b',
            progress_th_stance: '\u0421\u0442\u043e\u0439\u043a\u0430',
            progress_th_arm: '\u0423\u0433\u043e\u043b \u0440\u0443\u043a\u0438',
            progress_th_release: '\u0412\u044b\u043f\u0443\u0441\u043a',
            progress_th_follow: '\u0417\u0430\u0432\u0435\u0440\u0448\u0435\u043d\u0438\u0435',
            progress_th_date: '\u0414\u0430\u0442\u0430',
            progress_empty: '\u0411\u0440\u043e\u0441\u043a\u043e\u0432 \u043f\u043e\u043a\u0430 \u043d\u0435\u0442. \u0417\u0430\u0433\u0440\u0443\u0437\u0438\u0442\u0435 \u0432\u0438\u0434\u0435\u043e.',
            progress_trend: '\u0414\u0438\u043d\u0430\u043c\u0438\u043a\u0430 \u0431\u0430\u043b\u043b\u043e\u0432',
            export_title: '\u042d\u043a\u0441\u043f\u043e\u0440\u0442 \u043e\u0442\u0447\u0435\u0442\u0430',
            export_desc: '\u0421\u043a\u0430\u0447\u0430\u0442\u044c \u043e\u0442\u0447\u0435\u0442 \u0430\u043d\u0430\u043b\u0438\u0437\u0430 \u0432 PDF',
            export_select: '\u0412\u044b\u0431\u0435\u0440\u0438\u0442\u0435 \u0437\u0430\u0432\u0435\u0440\u0448\u0451\u043d\u043d\u044b\u0439 \u0430\u043d\u0430\u043b\u0438\u0437',
            export_loading: '\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430...',
            export_none: '\u041d\u0435\u0442 \u0437\u0430\u0432\u0435\u0440\u0448\u0451\u043d\u043d\u044b\u0445 \u0430\u043d\u0430\u043b\u0438\u0437\u043e\u0432',
            export_choose: '\u0412\u044b\u0431\u0435\u0440\u0438\u0442\u0435...',
            export_btn: '\u0421\u043a\u0430\u0447\u0430\u0442\u044c PDF',
            profile_title: '\u041f\u0440\u043e\u0444\u0438\u043b\u044c',
            profile_desc: '\u0423\u043f\u0440\u0430\u0432\u043b\u0435\u043d\u0438\u0435 \u043d\u0430\u0441\u0442\u0440\u043e\u0439\u043a\u0430\u043c\u0438',
            profile_name: '\u0418\u043c\u044f',
            profile_email: '\u042d\u043b. \u043f\u043e\u0447\u0442\u0430',
            profile_pw: '\u041d\u043e\u0432\u044b\u0439 \u043f\u0430\u0440\u043e\u043b\u044c',
            profile_pw2: '\u041f\u043e\u0432\u0442\u043e\u0440\u0438\u0442\u0435 \u043f\u0430\u0440\u043e\u043b\u044c',
            profile_save: '\u0421\u043e\u0445\u0440\u0430\u043d\u0438\u0442\u044c',
            profile_saved: '\u041f\u0440\u043e\u0444\u0438\u043b\u044c \u0441\u043e\u0445\u0440\u0430\u043d\u0451\u043d (demo).',
            profile_mismatch: '\u041f\u0430\u0440\u043e\u043b\u0438 \u043d\u0435 \u0441\u043e\u0432\u043f\u0430\u0434\u0430\u044e\u0442.',
            profile_pw_hint: '\u041e\u0441\u0442\u0430\u0432\u044c\u0442\u0435 \u043f\u0443\u0441\u0442\u044b\u043c \u0434\u043b\u044f \u0441\u043e\u0445\u0440\u0430\u043d\u0435\u043d\u0438\u044f',
            profile_pw_hint2: '\u041f\u043e\u0432\u0442\u043e\u0440\u0438\u0442\u0435 \u043f\u0430\u0440\u043e\u043b\u044c',
            feedback_show: '\u25bc \u041f\u043e\u043a\u0430\u0437\u0430\u0442\u044c \u043f\u043e\u0434\u0440\u043e\u0431\u043d\u044b\u0439 \u0430\u043d\u0430\u043b\u0438\u0437',
            feedback_hide: '\u25b2 \u0421\u043a\u0440\u044b\u0442\u044c \u043f\u043e\u0434\u0440\u043e\u0431\u043d\u044b\u0439 \u0430\u043d\u0430\u043b\u0438\u0437',
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
