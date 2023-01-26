package HangmanWeb

import (
	"bytes"
	"fmt"
	HangmanClassic "hangmanWeb/hangman-classic-for-web/Functions"
	"html/template"
	"net/http"
	"path/filepath"
)

type UserData struct {
	Username  string `json:"user_id"`
	Password  []byte `json:"password"`
	Language  string `json:"locale"`
	BestScore int    `json:"best_score"`
	Score     int    `json:"score"`
}

type WebData struct {
	Game       HangmanClassic.Game
	User       UserData
	State      string
	Error      bool
	Message    string
	Difficulty string
	Scores     [][]string
}

var Data WebData

// PathHandler : handle every path in a switch
func PathHandler(w http.ResponseWriter, r *http.Request) {

	println(r.URL.Path)

	switch r.URL.Path {

	case "/hangman":
		Data.LaunchGame(w, r)

	case "/hangman-replay":
		Data.Reset()
		HangmanClassic.NewGame(Data.GetWordDifficulty(), &Data.Game)
		http.Redirect(w, r, "hangman", http.StatusFound)

	case "/hangman-stop":
		Data.State = "LOST"
		Data.UpdateScore()
		http.Redirect(w, r, "gameMenu", http.StatusFound)

	case "/hangman-game":
		Data.Login(w, r)

	case "/sign-up":
		if CheckUsernameAvailability(r) {
			UpdateDatabase(r)
			http.Redirect(w, r, "difficulty", http.StatusFound)
		} else {
			http.Redirect(w, r, "hangman-game", http.StatusFound)
		}

	case "/log-in":
		if SuccessfullLogin(r) {
			Data.LoadUser(r.FormValue("log_username"))
			http.Redirect(w, r, "difficulty", http.StatusFound)
		} else {
			http.Redirect(w, r, "hangman-game", http.StatusFound)
		}

	case "/logout":
		Data.Logout(r)
		fmt.Println(string("\033[31m"), "user logout", string("\033[0m"))
		http.Redirect(w, r, "hangman-game", http.StatusFound)

	case "/difficulty":
		Data.DifficultyMenu(w, r)

	case "/difficulty-play":
		Data.menuPlay(w, r)
		Data.Reset()
		HangmanClassic.NewGame(Data.GetWordDifficulty(), &Data.Game)
		http.Redirect(w, r, "hangman", http.StatusFound)

	case "/gameMenu":
		Data.gameMenu(w, r)

	default: // redirect to the login page instead of error (currently mainMenu tho)
		http.Redirect(w, r, "hangman-game", http.StatusFound)
	}
}

// RenderTemplate : helps to render the html templates
func RenderTemplate(w http.ResponseWriter, tmplName string) {
	templateCache, err := CreateTemplateCache()

	if err != nil {
		panic(err)
	}

	tmpl, ok := templateCache[tmplName+".gohtml"]

	if !ok {
		http.Error(w, "The template do not exist", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, Data)
	buffer.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./Templates/*.gohtml")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layout, err := filepath.Glob("./Templates/layouts/*.layout.gohtml")

		if err != nil {
			return cache, err
		}
		if len(layout) > 0 {
			tmpl.ParseGlob("./Templates/layouts/*.layout.gohtml")
		}
		cache[name] = tmpl
	}
	return cache, nil
}
