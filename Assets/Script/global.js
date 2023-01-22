// custom cursor
const Light = document.querySelector('#Light');
const moveCursor = (e)=> {
    const mouseY = e.clientY;
    const mouseX = e.clientX;
    Light.style.transform = `translate3d(${mouseX}px, ${mouseY}px, 0)`;
}
window.addEventListener('mousemove', moveCursor);

// loading screen
setTimeout(() => {
    document.querySelector("#loader").style.display = "none"
}, 5000)


function UpdateFirstPlan(route, name, id) {
    let req = new XMLHttpRequest();
    req.open("POST", route);
    req.onload = function (event) {
        const parser = new DOMParser();
        const doc = parser.parseFromString(event.target.response, "text/html");
        const FirstPlan = doc.getElementById("FirstPlan");
        document.getElementById("FirstPlan").innerHTML = FirstPlan.innerHTML;
    };
    const Data = new FormData();
    const input_form = document.getElementById(id)
    Data.append(name, input_form.value);
    req.send(Data);
}

function FocusLetter(){
    setTimeout(() => {
        document.getElementById("letter-input").focus();
    }, 100);
}

function SendDifficulty(route, name, id){
    let difficulty = document.querySelector('input[name="difficulty"]:checked').value;
    UpdateFirstPlan(route, name, difficulty)
}