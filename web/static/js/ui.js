// UI helper functions — notifications, formatting, etc.
const UI = {
    showNotification(message, type) {
        const el = document.createElement('div');
        el.className = 'notification notification-' + (type || 'info');
        el.textContent = message;
        el.style.cssText = 'position:fixed;top:80px;right:20px;padding:1rem 1.5rem;border-radius:8px;z-index:200;font-size:0.9rem;' +
            (type === 'error' ? 'background:#e74c3c;color:#fff;' : 'background:#27ae60;color:#fff;');
        document.body.appendChild(el);
        setTimeout(() => el.remove(), 3000);
    }
};
