package Game

import (
	"fmt"
	Hangman "hangman-web/Game/hangman-classic/Hangman/Game"
	"html/template"
	"net/http"
	"strings"
)

const file = "words.txt"

func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hangman":
		LaunchGame(w, r)
	case "/login": // make a login page here
	case "/loginSubmit": // make a submit log page here
	case "/mainMenu": // make the main menu page here
	default: // redirect to the login page instead of error
		http.Redirect(w, r, "login", http.StatusFound)
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *Hangman.GameData) {

	t, err := template.ParseFiles("./Templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, *data)
}

func LaunchGame(w http.ResponseWriter, r *http.Request) {

	data := Hangman.StartGame(file)

	replay := r.FormValue("Replay")
	if replay != "" {
		data = Hangman.StartGame(file)
	}

	guess := r.FormValue("Letter")
	if guess != "" {
		data = DisplayMessage(&data, guess)
	}

	RenderTemplate(w, "game", &data)
}

func DisplayMessage(data *Hangman.GameData, input string) Hangman.GameData {
	input = strings.ToUpper(input)

	*data = Hangman.IntputTesting(input, *data)

	if data.Attempts == 9 {
		data.Error = fmt.Sprintf("Sorry, you loose!, word was: %s", data.ToFind)
		data.State = "lost"
	}

	switch data.State {

	case "won":
		data.State = "won"
		data.Error = "GG WP, YOU WON ! YOUR DICK IS ENORME"

	case "goodGuess":
		data.Word = Hangman.RevealLetters(*data)
		data.Error = "GG"

	case "badGuess":
		data.Attempts++
		data.Error = fmt.Sprintf("Sorry, the letter %s is not in the word. %d attempts lefts", input, 10-data.Attempts)

	case "badWordGuessed":
		data.Attempts += 2
		data.Error = fmt.Sprintf("Sorry, %s is not the word to find. %d attempts lefts", input, 10-data.Attempts)

	case "alreadyGuessed":
		data.Error = "This letter is already guessed"

	case "alreadyFound":
		data.Error = "This letter is already found"

	case "invalidInput":
		data.Error = "Please enter a valid input"

	}
	return *data
}
