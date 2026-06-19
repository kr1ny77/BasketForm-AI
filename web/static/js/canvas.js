// Canvas background animation placeholder — full implementation in Stage 2
(function() {
    const canvas = document.createElement('canvas');
    canvas.id = 'bgCanvas';
    document.body.prepend(canvas);
    const ctx = canvas.getContext('2d');

    function resize() {
        canvas.width = window.innerWidth;
        canvas.height = window.innerHeight;
    }
    resize();
    window.addEventListener('resize', resize);

    function draw() {
        ctx.clearRect(0, 0, canvas.width, canvas.height);
        requestAnimationFrame(draw);
    }
    draw();
})();
