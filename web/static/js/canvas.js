(function () {
    var canvas = document.createElement('canvas');
    canvas.id = 'bgCanvas';
    document.body.prepend(canvas);
    var ctx = canvas.getContext('2d');

    var W, H;
    var mouseX = -9999, mouseY = -9999;
    var balls = [];
    var BALL_COUNT = 20;
    var emojiImg = null;
    var emojiLoaded = false;

    function resize() {
        W = canvas.width = window.innerWidth;
        H = canvas.height = window.innerHeight;
    }
    resize();
    window.addEventListener('resize', resize);

    document.addEventListener('mousemove', function (e) {
        mouseX = e.clientX;
        mouseY = e.clientY;
    });

    // Pre-render basketball emoji to offscreen canvas
    var emojiCanvas = document.createElement('canvas');
    var emojiCtx = emojiCanvas.getContext('2d');
    emojiCanvas.width = 64;
    emojiCanvas.height = 64;
    emojiCtx.font = '56px serif';
    emojiCtx.textAlign = 'center';
    emojiCtx.textBaseline = 'middle';
    emojiCtx.fillText('\uD83C\uDFC0', 32, 34);
    emojiLoaded = true;

    function Ball() { this.reset(true); }

    Ball.prototype.reset = function (init) {
        this.r = 14 + Math.random() * 16;
        this.x = Math.random() * W;
        this.y = init ? Math.random() * H : -this.r * 3 - Math.random() * 80;
        this.vx = (Math.random() - 0.5) * 0.2;
        this.vy = 0.3 + Math.random() * 0.6;
        this.rotation = (Math.random() - 0.5) * 0.4;
        this.rotSpeed = (Math.random() - 0.5) * 0.008;
        this.alpha = 0.06 + Math.random() * 0.1;
    };

    Ball.prototype.update = function () {
        var dx = mouseX - this.x;
        var dy = mouseY - this.y;
        var dist = Math.sqrt(dx * dx + dy * dy);
        var repel = 150;

        if (dist < repel && dist > 0) {
            var force = (repel - dist) / repel * 0.06;
            this.vx -= (dx / dist) * force * 2;
            this.vy -= (dy / dist) * force * 2;
        }

        this.vy += 0.06;
        this.vx *= 0.992;
        this.vy *= 0.996;
        this.rotation += this.rotSpeed;

        this.x += this.vx;
        this.y += this.vy;

        if (this.y > H + this.r * 3) this.reset(false);
        if (this.x < -this.r * 4) this.x = W + this.r * 2;
        if (this.x > W + this.r * 4) this.x = -this.r * 2;
    };

    Ball.prototype.draw = function () {
        if (!emojiLoaded) return;
        ctx.save();
        ctx.translate(this.x, this.y);
        ctx.rotate(this.rotation);
        ctx.globalAlpha = this.alpha;
        var size = this.r * 2.2;
        ctx.drawImage(emojiCanvas, -size / 2, -size / 2, size, size);
        ctx.globalAlpha = 1;
        ctx.restore();
    };

    for (var i = 0; i < BALL_COUNT; i++) {
        balls.push(new Ball());
    }

    var waveOffset = 0;

    function drawWaves() {
        waveOffset += 0.002;
        for (var w = 0; w < 3; w++) {
            ctx.beginPath();
            var baseY = H * 0.75 + w * 30;
            ctx.moveTo(0, baseY);
            for (var x = 0; x <= W; x += 6) {
                var y = baseY
                    + Math.sin(x * 0.002 + waveOffset + w * 1.2) * 25
                    + Math.sin(x * 0.004 - waveOffset * 0.3) * 12;
                ctx.lineTo(x, y);
            }
            ctx.lineTo(W, H);
            ctx.lineTo(0, H);
            ctx.closePath();
            ctx.fillStyle = 'rgba(255,106,0,' + (0.012 - w * 0.003) + ')';
            ctx.fill();
        }
    }

    function animate() {
        ctx.clearRect(0, 0, W, H);
        drawWaves();
        for (var i = 0; i < balls.length; i++) {
            balls[i].update();
            balls[i].draw();
        }
        requestAnimationFrame(animate);
    }
    animate();
})();
