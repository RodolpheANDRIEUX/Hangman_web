const signup_form = document.getElementById('sinup-form');
const container_play_form = document.getElementById('container-play-form')

const showSignupForm = () => {
    signup_form.style.display = "grid"
    container_play_form.style.width = "20em"
}

const closeAll = () => {
    signup_form.style.display = "none"
    container_play_form.style.width = "10em"
}

const popup = document.getElementById("login-popup");

function handlePopup(open){
    if(open){
        popup.classList.add('opened');
    }else{
        popup.classList.remove('opened');
    }
}