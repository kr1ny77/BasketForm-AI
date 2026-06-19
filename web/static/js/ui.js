var UI = {
    notify: function (message, type) {
        var el = document.createElement('div');
        var bg = type === 'error' ? '#e74c3c' : type === 'success' ? '#27ae60' : '#ff6a00';
        el.style.cssText = 'position:fixed;top:80px;right:20px;padding:0.8rem 1.4rem;border-radius:8px;z-index:300;font-size:0.88rem;color:#fff;background:' + bg + ';box-shadow:0 4px 20px rgba(0,0,0,0.4);transition:opacity 0.3s;';
        el.textContent = message;
        document.body.appendChild(el);
        setTimeout(function () { el.style.opacity = '0'; setTimeout(function () { el.remove(); }, 300); }, 2500);
    },
    formatDate: function (iso) {
        return new Date(iso).toLocaleDateString(undefined, { year: 'numeric', month: 'short', day: 'numeric' });
    },
    formatDateTime: function (iso) {
        return new Date(iso).toLocaleString();
    },
    escHtml: function (s) {
        var d = document.createElement('div');
        d.textContent = s || '';
        return d.innerHTML;
    }
};
