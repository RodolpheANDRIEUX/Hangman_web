package HangmanWeb

import (
	"net/http"
	"strings"
)

// DifficultyMenu : We render the DifficultyMenu template
func (ptrData *WebData) DifficultyMenu(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "difficultyMenu")
}

// Login : We render the Login template
func (ptrData *WebData) Login(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login")
}

// menuPlay : Method to handle difficulty files
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
		println("letter", guess, " received")
		guess = strings.ToUpper(guess)
		if !ptrData.ErrorInLetter(&guess) {
			ptrData.HandleLetter(&guess)
			println("state: ", ptrData.State)
			if ptrData.State == "WIN" || ptrData.State == "LOST" {
				ptrData.UpdateScore()
				http.Redirect(w, r, "gameMenu", http.StatusFound)
			}
		}
	}
	RenderTemplate(w, "game")
}

// gameMenu : We render the loose / Win menu
func (ptrData *WebData) gameMenu(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "gameMenu")
}
