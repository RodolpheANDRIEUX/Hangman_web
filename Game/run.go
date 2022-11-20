package Game

import (
	"fmt"
	Hangman "hangman-web/Game/hangman-classic/Hangman/Game"
	"strings"
)

func HangMan(data Hangman.GameData, input string) Hangman.GameData {
	input = strings.ToUpper(input)

	data = Hangman.IntputTesting(input, data)

	if data.Attempts == 9 {
		data.Error = fmt.Sprintf("Sorry, you loose!, word was: %s", data.ToFind)
		data.State = "lost"
	}

	switch data.State {
	case "goodGuess":
		data.Word = Hangman.RevealLetters(data)
		data.Error = " "

		if Hangman.WordGuessed(data) {
			data.State = "won"
			data.Error = "GG WP, YOU WON ! YOUR DICK IS ENORME"
		}

	case "badGuess":
		data.Attempts++
		data.Error = fmt.Sprintf("Sorry, the letter %s is not in the word. %d attempts lefts", input, 10-data.Attempts)

	case "badWordGuessed":
		data.Attempts += 2
		data.Error = fmt.Sprintf("Sorry, %s is not the word to find. %d attempts lefts", input, 10-data.Attempts)

	case "alreadyGuessed":
		data.Error = "This letter is already guessed"

	case "alreadyFound":
		data.Error = "This letter is already found"

	case "invalidInput":
		data.Error = "Please enter a valid input"

	}
	return data
}
