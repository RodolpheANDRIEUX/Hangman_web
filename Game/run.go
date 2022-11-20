package Game

import (
	Hangman "hangman-web/Game/hangman-classic/Hangman/Game"
)

const file = "words.txt"

var data Hangman.GameData = Hangman.StartGame(file)

func HangMan(data Hangman.GameData, input string) Hangman.GameData {

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
