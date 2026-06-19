(function () {
    const canvas = document.createElement('canvas');
    canvas.id = 'bgCanvas';
    document.body.prepend(canvas);
    const ctx = canvas.getContext('2d');

    let W, H;
    let mouseX = -1000, mouseY = -1000;
    const balls = [];
    const BALL_COUNT = 35;
    const TRAIL_LEN = 8;

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

    function Ball() {
        this.reset(true);
    }

    Ball.prototype.reset = function (init) {
        this.r = 4 + Math.random() * 10;
        this.x = Math.random() * W;
        this.y = init ? Math.random() * H : -this.r * 2;
        this.vx = (Math.random() - 0.5) * 1.2;
        this.vy = 0.3 + Math.random() * 0.8;
        this.trail = [];
        this.hue = Math.random() > 0.6 ? 0 : 25;
        this.alpha = 0.4 + Math.random() * 0.5;
    };

    Ball.prototype.update = function () {
        var dx = mouseX - this.x;
        var dy = mouseY - this.y;
        var dist = Math.sqrt(dx * dx + dy * dy);
        var repel = 120;

        if (dist < repel) {
            var force = (repel - dist) / repel * 0.08;
            this.vx -= dx * force * 0.03;
            this.vy -= dy * force * 0.03;
        }

        this.vx *= 0.995;
        this.vy *= 0.995;
        this.vy += 0.012;

        this.x += this.vx;
        this.y += this.vy;

        this.trail.push({ x: this.x, y: this.y });
        if (this.trail.length > TRAIL_LEN) this.trail.shift();

        if (this.y > H + this.r * 2 || this.x < -this.r * 4 || this.x > W + this.r * 4) {
            this.reset(false);
        }
    };

    Ball.prototype.draw = function () {
        if (this.trail.length > 1) {
            ctx.beginPath();
            ctx.moveTo(this.trail[0].x, this.trail[0].y);
            for (var i = 1; i < this.trail.length; i++) {
                ctx.lineTo(this.trail[i].x, this.trail[i].y);
            }
            ctx.strokeStyle = 'hsla(' + this.hue + ',100%,55%,' + (this.alpha * 0.25) + ')';
            ctx.lineWidth = 1.5;
            ctx.stroke();
        }

        ctx.beginPath();
        ctx.arc(this.x, this.y, this.r, 0, Math.PI * 2);
        ctx.fillStyle = 'hsla(' + this.hue + ',100%,55%,' + this.alpha + ')';
        ctx.fill();

        if (this.r > 7) {
            ctx.beginPath();
            ctx.arc(this.x - this.r * 0.2, this.y - this.r * 0.2, this.r * 0.3, 0, Math.PI * 2);
            ctx.fillStyle = 'hsla(' + this.hue + ',100%,80%,' + (this.alpha * 0.4) + ')';
            ctx.fill();
        }
    };

    for (var i = 0; i < BALL_COUNT; i++) {
        balls.push(new Ball());
    }

    var waveOffset = 0;

    function drawWaves() {
        waveOffset += 0.004;
        for (var w = 0; w < 3; w++) {
            ctx.beginPath();
            var baseY = H * 0.65 + w * 40;
            ctx.moveTo(0, baseY);
            for (var x = 0; x <= W; x += 8) {
                var y = baseY + Math.sin(x * 0.003 + waveOffset + w * 1.5) * 25
                    + Math.sin(x * 0.007 - waveOffset * 0.6) * 15;
                ctx.lineTo(x, y);
            }
            ctx.lineTo(W, H);
            ctx.lineTo(0, H);
            ctx.closePath();
            ctx.fillStyle = 'rgba(255,106,0,' + (0.02 - w * 0.005) + ')';
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
