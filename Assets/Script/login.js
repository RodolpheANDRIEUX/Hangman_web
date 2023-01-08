const login_form = document.getElementById('login-form');
const signup_form = document.getElementById('sinup-form');
const container_login_form = document.getElementById('container-login-form')
const container_play_form = document.getElementById('container-play-form')

const showLoginForm = () => {
    login_form.style.display = "grid"
    container_login_form.style.width = "20em"
    signup_form.style.display = "none"
    container_play_form.style.width = "10em"
}

const showSignupForm = () => {
    signup_form.style.display = "grid"
    container_play_form.style.width = "20em"
    login_form.style.display = "none"
    container_login_form.style.width = "10em"
}

const closeAll = () => {
    login_form.style.display = "none"
    signup_form.style.display = "none"
    container_login_form.style.width = "10em"
    container_play_form.style.width = "10em"
}

const popup = document.getElementById("popup");

function handlePopup(open){
    if(open){
        popup.classList.add('opened');
    }else{
        popup.classList.remove('opened');
    }
}