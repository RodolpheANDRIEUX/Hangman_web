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

func GetWordDifficulty(user *UserData) string {
	switch user.Difficulty {
	case "easy":
		return "Assets/words/words.txt"
	case "medium":
		return "Assets/words/words2.txt"
	case "hard":
		return "Assets/words/words3.txt"
	case "impossible":
		return "Assets/words/words3.txt"
	default:
		return "Assets/words/words.txt"
	}
}
