const Light = document.querySelector('#Light');
const moveCursor = (e)=> {
    const mouseY = e.clientY;
    const mouseX = e.clientX;
    Light.style.transform = `translate3d(${mouseX}px, ${mouseY}px, 0)`;
}
window.addEventListener('mousemove', moveCursor);
setTimeout(() => {
    document.querySelector("#loader").style.display = "none"
}, 3000)
