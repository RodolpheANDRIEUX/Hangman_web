/*----- POPUP -----*/

const bg_overlay = document.getElementById('bg-overlay');
const popup = document.getElementById("popup");


console.log("file charged")

function handlePopup(open) {
    if (open) {
        popup.classList.add('popup_open');
        bg_overlay.classList.add("showed")
    } else {
        popup.classList.remove('popup_open');
        bg_overlay.classList.remove("showed");
    }
}

function closePopup (){
    bg_overlay.classList.remove("showed")
    popup.classList.remove('popup_open');
}

/*----- END -----*/