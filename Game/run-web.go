package Game

import (
	"hangmanWeb/Game/HangmanWeb"
	hangmanClassic "hangmanWeb/Game/hangman-classic-for-web/Functions"
	"html/template"
	"net/http"
)

type UserData struct {
	Name   string
	Record int
}

type WebData struct {
	Game         hangmanClassic.Game
	User         UserData
	Status       string
	JoseFilePath string
	Message      string
}

const EASY = "words.txt"
const MEDIUM = "words2.txt"
const HARD = "words3.txt"
const IMPOSSIBLE = "words3.txt"

var data WebData

// PathHandler : handle every path in a switch
func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hangman":
		HangmanWeb.LaunchGame(w, r)
	case "/login": // make a login page here
	case "/loginSubmit": // make a submit log page here
	case "/mainMenu":
		// make the main menu page here
		HangmanWeb.MainMenu(w, r)
	case "/mainMenu-Play":
		// Functions.NewGame("Assets/words/words.txt")
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
