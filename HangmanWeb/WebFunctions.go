package HangmanWeb

import (
	"net/http"
	"strings"
)

// MainMenu : We render the Menu template
func (ptrData *WebData) MainMenu(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "MainMenu")
}

func (ptrData *WebData) Login(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login")
}

func (ptrData *WebData) menuPlay(w http.ResponseWriter, r *http.Request) {

	difficulty := r.FormValue("difficulty")
	if difficulty != "" {
		ptrData.Difficulty = difficulty
	}
}

// LaunchGame : HangmanWeb launcher. The method we call to start the game
func (ptrData *WebData) LaunchGame(w http.ResponseWriter, r *http.Request) {
	guess := r.FormValue("Letter")
	if guess != "" {
		guess = strings.ToUpper(guess)
		if !ptrData.ErrorInLetter(&guess) {
			ptrData.HandleLetter(&guess)
			if ptrData.State == "WIN" || ptrData.State == "LOST" {
				ptrData.UpdateScore()
				http.Redirect(w, r, "gameMenu", http.StatusFound)
			}
		}
	}
	RenderTemplate(w, "game")
}

// gameMenu : We render the Game Menu template with difficulty
func (ptrData *WebData) gameMenu(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "gameMenu")
}
