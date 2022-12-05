package HangmanWeb

import (
	"fmt"
	HangmanClassic "hangmanWeb/hangman-classic-for-web/Functions"
	"net/http"
	"strings"
)

func HandleLetter(guess *string) {

}

func ErrorInLetter(w http.ResponseWriter, guess *string, ptrData *WebData) bool {
	switch {

	case *guess == "DEVMODE":
		fmt.Fprintf(w, ptrData.Game.ToFind)

	case !HangmanClassic.IsAlpha(*guess):
		ptrData.Error = true
		ptrData.Message = "Please, select a letter"
		return true

	case HangmanClassic.ContainsArray(ptrData.Game.GoodGuess, *guess):
		ptrData.Error = true
		ptrData.Message = "This letter has already been Revealed"
		return true

	case HangmanClassic.ContainsArray(ptrData.Game.BadGuess, *guess):
		ptrData.Error = true
		ptrData.Message = "This letter has already been chosen"
		return true

	}
	ptrData.Error = false
	return false
}

// LaunchGame : HangmanWeb launcher. The function we call to start a new game
func LaunchGame(w http.ResponseWriter, r *http.Request, data *WebData) {
	replay := r.FormValue("Replay")
	if replay != "" {
		HangmanClassic.NewGame(GetWordDifficulty(&Data.User), &Data.Game)
	}

	guess := r.FormValue("Letter")
	if guess != "" {
		guess = strings.ToUpper(guess)
		if !ErrorInLetter(w, &guess, &Data) {
			HandleLetter(&guess)
		}
	}

	RenderTemplate(w, "game")
}

//// DisplayMessage : Basically the game loop with the message displayed to the user
//func DisplayMessage(data *Hangman.GameData, input string) Hangman.GameData {
//	input = strings.ToUpper(input)
//
//	*data = Hangman.IntputTesting(input, *data)
//
//	switch data.State {
//
//	case "won":
//		data.State = "won"
//		data.Error = "GG WP, YOU WON ! YOUR DICK IS ENORME"
//
//	case "goodGuess":
//		data.Error = "GG"
//
//	case "badGuess":
//		data.Attempts++
//		data.Error = fmt.Sprintf("Sorry, the letter %s is not in the word. %d attempts lefts", input, 10-data.Attempts)
//		if data.Attempts >= 10 {
//			data.Error = fmt.Sprintf("Sorry, you loose!, word was: %s", data.ToFind)
//			data.State = "lost"
//		}
//
//	case "badWordGuessed":
//		data.Attempts += 2
//		data.Error = fmt.Sprintf("Sorry, %s is not the word to find. %d attempts lefts", input, 10-data.Attempts)
//
//	case "alreadyGuessed":
//		data.Error = "This letter is already guessed"
//
//	case "alreadyFound":
//		data.Error = "This letter is already found"
//
//	case "invalidInput":
//		data.Error = "Please enter a valid input"
//
//	}
//	return *data
//}
