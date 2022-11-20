package Game

import (
	Hangman "hangman-web/Game/hangman-classic/Hangman/Game"
)

var data Hangman.GameData

func HangMan(data Hangman.GameData, input, file string) Hangman.GameData {
	data = Hangman.StartGame(file)

	data.Guess = []string{}

	for {
		data = Hangman.IntputTesting(input, data)

		switch data.State {
		case "goodGuess":
			data.Word = Hangman.RevealLetters(data)

			if Hangman.WordGuessed(data) {
				data.State = "won"
			}

		case "badGuess":
			data.Attempts++

		case "badWordGuessed":
			data.Attempts += 2

		}
	}
	return data
}
