package HangmanWeb

import (
	HangmanClassic "hangmanWeb/hangman-classic-for-web/Functions"
	"html/template"
	"net/http"
)

type UserData struct {
	Username string
	Password string
	Language string
	Record   int
}

type WebData struct {
	Game       HangmanClassic.Game
	User       UserData
	State      string
	Error      bool
	Message    string
	Difficulty string
}

var Data WebData

// PathHandler : handle every path in a switch
func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {

	case "/hangman":
		Data.LaunchGame(w, r)

	case "/login": // TODO: make a login page here
		Data.Login(w, r)
	case "/login-Submit": // TODO: make a submit log page here
		if CheckUsernameAvailability(r) {
			AddUser(r)
			http.Redirect(w, r, "mainMenu", http.StatusFound)
		} else {
			http.Redirect(w, r, "login", http.StatusFound)
		}

	case "/mainMenu": // TODO: make the main menu page here
		Data.MainMenu(w, r)

	case "/mainMenu-Play":
		Data.menuPlay(w, r)
		Data.Reset()
		HangmanClassic.NewGame(Data.GetWordDifficulty(), &Data.Game)
		http.Redirect(w, r, "hangman", http.StatusFound)

	case "/gameMenu": // TODO: win and lose screen
		Data.gameMenu(w, r)

	default: // redirect to the login page instead of error (currently mainMenu tho)
		http.Redirect(w, r, "login", http.StatusFound)
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
