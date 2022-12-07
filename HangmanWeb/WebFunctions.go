package HangmanWeb

import (
	HangmanClassic "hangmanWeb/hangman-classic-for-web/Functions"
	"net/http"
	"strings"
)

// MainMenu : We render the Menu template 1
func (ptrData *WebData) MainMenu(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "MainMenu")
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
			//	http.Redirect(w, r, "gameMenu", http.StatusFound)
			//}
		}
	}
	RenderTemplate(w, "game")
}

func (ptrData *WebData) gameMenu(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "gameMenu")
}
