// API helper functions — to be expanded in Stage 2 and 3
const API = {
    async getVideos() {
        const res = await fetch('/api/videos');
        return res.json();
    },
    async getStatus(id) {
        const res = await fetch('/api/status/' + id);
        return res.json();
    },
    async getResult(id) {
        const res = await fetch('/api/result/' + id);
        return res.json();
    }
};
