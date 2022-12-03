package HangmanWeb

import (
	"net/http"
)

// LaunchGame : Game launcher. The function we call to start a new game
func LaunchGame(w http.ResponseWriter, r *http.Request) {

	//replay := r.FormValue("Replay")
	//if replay != "" {
	//	fmt.Println(data.Difficulty)
	//	data = Hangman.StartGame(data.Difficulty)
	//}
	//
	//guess := r.FormValue("Letter")
	//if guess != "" {
	//	data = DisplayMessage(&data, guess)
	//}
	//
	//RenderTemplate(w, "game")
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
