console.log("1")

const signup_form = document.getElementById('signup_form');
const container_play_form = document.getElementById('container-play-form')
const bg_overlay = document.getElementById('bg-overlay')

function showSignupForm(open) {
    if (open) {
        signup_form.classList.add("signup_opened")
        container_play_form.classList.add("form_expend")
        bg_overlay.classList.add("showed")
    } else {
        signup_form.classList.remove("signup_opened")
        container_play_form.classList.remove("form_expend")
        bg_overlay.classList.remove("showed")
    }
}

function closeAll (){
    signup_form.classList.remove("signup_opened")
    container_play_form.classList.remove("form_expend")
    bg_overlay.classList.remove("showed")
    popup.classList.remove('opened');
    console.log("close all")
}

const popup = document.getElementById("login-popup");

function handlePopup(open) {
    if (open) {
        popup.classList.add('opened');
    } else {
        popup.classList.remove('opened');
    }
}

function refreshPage() {
    window.location.reload();
}