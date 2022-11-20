package Hangman

import (
	"fmt"
	"strings"
)

// GameData : The structure of our game
type GameData struct {
	Word             string   // Word with _ ( ex A _ _ _ )
	ToFind           string   // Final word chosen by the program at the beginning. It is the word to find
	LetterIsCorrect  bool     // True if we found the letter, false if not
	State            string   // The state of pur game
	Attempts         int      // Number of attempts left
	Guess            []string // Letters tried
	HangmanPositions int      // State of Jose
	Error            string   // Error to display ( for web version )
}

// RevealLetters : return the word with a letter revealed
func RevealLetters(data GameData) string {
	WordByte := []byte(data.Word)
	WordToFind := []byte(data.ToFind)
	for index, letter := range WordToFind {
		if string(letter) == data.Guess[len(data.Guess)-1] {
			WordByte[index] = WordToFind[index]
		}
	}
	return strings.ToUpper(string(WordByte))
}

// IntputTesting : Game input testing and logic
func IntputTesting(guess string, data GameData) GameData {
	if guess == "" || guess < "A" || guess > "Z" {
		data.State = "invalidInput"
		return data
	}
	if len(guess) < 2 {
		if strings.Contains(data.Word, guess) {
			data.State = "alreadyFound"
			return data
		}
		if strings.Contains(strings.Join(data.Guess, " "), guess) {
			data.State = "alreadyGuessed"
		}
		if strings.Contains(data.ToFind, guess) && !(strings.Contains(data.Word, guess)) {
			data.State = "goodGuess"
			data.Guess = append(data.Guess, guess)
		}
		if !(strings.Contains(data.ToFind, guess)) && !(strings.Contains(strings.Join(data.Guess, " "), guess)) {
			data.State = "badGuess"
			data.Guess = append(data.Guess, guess)
		}

	}
	if len(guess) >= 2 {
		data.State = "badWordGuessed"
		if guess == data.ToFind {
			data.State = "won"
		}
	}
	return data
}

// FindLetter : check if the letter is on the word or not
func FindLetter(input string, data GameData) GameData {
	data.State = "badGuess"
	if strings.Contains(data.ToFind, input) {
		data.State = "goodGuess"
	}
	return data
}

// WordGuessed : Check if we found the word or not
func WordGuessed(data GameData) bool {
	count := 0
	for _, v := range data.Word {
		if v == 95 {
			count++
		}
	}
	if count == 0 {
		return true
	}
	return false
}

// PrintWord : Simply prints spaces between letters
func PrintWord(word string) {
	for _, v := range word {
		fmt.Print(string(v) + " ")
	}
	fmt.Println()
}
