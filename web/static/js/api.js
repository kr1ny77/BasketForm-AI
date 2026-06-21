var API = {
    getVideos: function () {
        return fetch('/api/videos').then(function (r) { return r.json(); });
    },
    getStatus: function (id) {
        return fetch('/api/status/' + id).then(function (r) { return r.json(); });
    },
    getResult: function (id) {
        return fetch('/api/result/' + id).then(function (r) { return r.json(); });
    },
    upload: function (file, onProgress) {
        return new Promise(function (resolve, reject) {
            var formData = new FormData();
            formData.append('video', file);
            var xhr = new XMLHttpRequest();
            xhr.upload.addEventListener('progress', function (e) {
                if (e.lengthComputable && onProgress) {
                    onProgress(Math.round((e.loaded / e.total) * 100));
                }
            });
            xhr.addEventListener('load', function () {
                if (xhr.status === 201) resolve(JSON.parse(xhr.responseText));
                else reject(new Error(xhr.statusText));
            });
            xhr.addEventListener('error', function () { reject(new Error('Network error')); });
            xhr.open('POST', '/api/upload');
            xhr.send(formData);
        });
    }
};
