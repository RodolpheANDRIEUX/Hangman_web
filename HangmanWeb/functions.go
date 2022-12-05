package HangmanWeb

import (
	"net/http"
)

//const EASY = "words.txt"
//const MEDIUM = "words2.txt"
//const HARD = "words3.txt"
//const IMPOSSIBLE = "words3.txt"

// MainMenu : We render the Menu template 1
func MainMenu(w http.ResponseWriter, r *http.Request, data *WebData) {
	RenderTemplate(w, "MainMenu")
}

func GetWordDifficulty(data *WebData) {
	//switch r.FormValue("difficulty") {
	//case "easy":
	//
	//case "medium":
	//	data.User.Difficulty = MEDIUM
	//case "hard":
	//	data.User.Difficulty = HARD
	//case "impossible":
	//	data.User.Difficulty = IMPOSSIBLE
	//}
}
