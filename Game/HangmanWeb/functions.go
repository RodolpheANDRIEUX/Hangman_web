package HangmanWeb

import "net/http"

// MainMenu : We render the Menu template 1
func MainMenu(w http.ResponseWriter, r *http.Request) {
	//	RenderTemplate(w, "MainMenu")
	//}
	//
	//// Difficulty : We chose a difficulty on the menu. Here we start the game with the difficulty chosen
	//func Difficulty(r *http.Request) {
	//	switch r.FormValue("difficulty") {
	//	case "easy":
	//		data = Hangman.StartGame(EASY)
	//		data.Difficulty = EASY
	//	case "medium":
	//		data = Hangman.StartGame(MEDIUM)
	//		data.Difficulty = MEDIUM
	//	case "hard":
	//		data = Hangman.StartGame(HARD)
	//		data.Difficulty = HARD
	//	case "impossible":
	//		data = Hangman.StartGame(IMPOSSIBLE)
	//		data.Difficulty = IMPOSSIBLE
	//	}
}
