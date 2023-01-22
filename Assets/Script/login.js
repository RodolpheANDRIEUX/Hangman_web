const signup_form = document.getElementById('sinup-form');
const container_play_form = document.getElementById('container-play-form')
const bg_overlay = document.getElementById('bg-overlay')

const showSignupForm = () => {
    signup_form.style.display = "grid"
    container_play_form.style.width = "20em"
    bg_overlay.style.display = "grid"
    bg_overlay.style.zIndex = "0"
}

const closeAll = () => {
    signup_form.style.display = "none"
    container_play_form.style.width = "10em"
    bg_overlay.style.display = "none"
    bg_overlay.style.zIndex = "-1"
    popup.classList.remove('opened');
}

const popup = document.getElementById("login-popup");

function handlePopup(open){
    if(open){
        popup.classList.add('opened');
    }else{
        popup.classList.remove('opened');
    }
}

