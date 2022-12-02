package Game

import (
	"fmt"
	Hangman "hangman-web/Game/hangman-classic/Hangman/Game"
	"html/template"
	"net/http"
	"strings"
)

const EASY = "words.txt"
const MEDIUM = "words2.txt"
const HARD = "words3.txt"
const IMPOSSIBLE = "words3.txt"

var data Hangman.GameData

// MainMenu : We render the Menu template
func MainMenu(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "MainMenu")
}

// Difficulty : We chose a difficulty on the menu. Here we start the game with the difficulty chosen
func Difficulty(r *http.Request) {
	switch r.FormValue("difficulty") {
	case "easy":
		data = Hangman.StartGame(EASY)
		data.Difficulty = EASY
	case "?difficulty=medium":
		data = Hangman.StartGame(MEDIUM)
		data.Difficulty = MEDIUM
	case "?difficulty=hard":
		data = Hangman.StartGame(HARD)
		data.Difficulty = HARD
	case "?difficulty=impossible":
		data = Hangman.StartGame(IMPOSSIBLE)
		data.Difficulty = IMPOSSIBLE
	}
}

// PathHandler : We render a specific template for every path
func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hangman":
		LaunchGame(w, r)
	case "/login": // make a login page here
	case "/loginSubmit": // make a submit log page here
	case "/mainMenu":
		// make the main menu page here
		MainMenu(w, r)

	case "/mainMenu-Play":
		Difficulty(r)
		http.Redirect(w, r, "hangman", http.StatusFound)

	default: // redirect to the login page instead of error (currently mainmenu tho)
		http.Redirect(w, r, "mainMenu", http.StatusFound)
	}
}

// RenderTemplate : helps to render the html templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	t, err := template.ParseFiles("./Templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
}

// LaunchGame : Game launcher. The function we call to start a new game
func LaunchGame(w http.ResponseWriter, r *http.Request) {

	replay := r.FormValue("Replay")
	if replay != "" {
		fmt.Println(data.Difficulty)
		data = Hangman.StartGame(data.Difficulty)
	}

	guess := r.FormValue("Letter")
	if guess != "" {
		data = DisplayMessage(&data, guess)
	}

	RenderTemplate(w, "game")
}

// DisplayMessage : Basically the game loop with the message displayed to the user
func DisplayMessage(data *Hangman.GameData, input string) Hangman.GameData {
	input = strings.ToUpper(input)

	*data = Hangman.IntputTesting(input, *data)

	if data.Attempts >= 9 {
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
