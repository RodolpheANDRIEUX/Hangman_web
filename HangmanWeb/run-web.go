package HangmanWeb

import (
	HangmanClassic "hangmanWeb/hangman-classic-for-web/Functions"
	"html/template"
	"net/http"
)

type UserData struct {
	Name       string
	Language   string
	Difficulty string
	Record     int
}

type WebData struct {
	Game         HangmanClassic.Game
	User         UserData
	State        string
	JoseFilePath string
	Error        bool
	Message      string
}

var Data WebData

// PathHandler : handle every path in a switch
func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {

	case "/hangman":
		Data.LaunchGame(w, r)

	case "/login": // TODO: make a login page here

	case "/loginSubmit": // TODO: make a submit log page here

	case "/mainMenu": // TODO: make the main menu page here
		Data.MainMenu(w, r)

	case "/mainMenu-Play":
		HangmanClassic.NewGame(Data.User.GetWordDifficulty(), &Data.Game)
		http.Redirect(w, r, "hangman", http.StatusFound)

	default: // redirect to the login page instead of error (currently mainMenu tho)
		http.Redirect(w, r, "mainMenu", http.StatusFound)
	}
}

// RenderTemplate : helps to render the html templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	t, err := template.ParseFiles("./Templates/" + tmpl + ".gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, Data)
}
