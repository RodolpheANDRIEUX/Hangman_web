package HangmanWeb

import (
	HangmanClassic "hangmanWeb/hangman-classic-for-web/Functions"
	"net/http"
	"strings"
)

// MainMenu : We render the Menu template
func (ptrData *WebData) MainMenu(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "MainMenu")
}

func (ptrData *WebData) menuPlay(w http.ResponseWriter, r *http.Request) {

	difficulty := r.FormValue("difficulty")
	if difficulty != "" {
		ptrData.User.Difficulty = difficulty
	}
}

// LaunchGame : HangmanWeb launcher. The method we call to start the game
func (ptrData *WebData) LaunchGame(w http.ResponseWriter, r *http.Request) {
	replay := r.FormValue("Replay")
	if replay != "" {
		ptrData.Reset()
		HangmanClassic.NewGame(ptrData.User.GetWordDifficulty(), &ptrData.Game)
	}

	guess := r.FormValue("Letter")
	if guess != "" {
		guess = strings.ToUpper(guess)
		if !ptrData.ErrorInLetter(&guess) {
			ptrData.HandleLetter(&guess)
			//if ptrData.State == "won" || ptrData.State == "lost" {
			//	http.Redirect(w, r, "gameMenu.gohtml", http.StatusFound)
			//}
		}
	}
	RenderTemplate(w, "game")
}

func (ptrData *WebData) gameMenu(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "gameMenu.gohtml")
}
