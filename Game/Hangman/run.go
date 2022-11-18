package Hangman

import (
	"fmt"
	"strings"
)

// StartGame : Start a new game
func StartGame() GameData {
	WordsSlice := GetFile("Hangman/Words/words.txt")
	WordToFind := TakeRandomWord(WordsSlice)
	Data := NewGame(WordToFind)
	Data.ToFind = WordToFind
	Data.Tries = InitialLetters(Data).Tries
	Data.Index = RevealInitialLetters(Data).Index
	Data = RevealInitialLetters(Data)
	Data.Tries = []string{}
	return Data
}

func HangMan(data GameData, input string) GameData {
	data.Tries = append(data.Tries, strings.ToUpper(input))
	data = FindLetter(data)
	if strings.Contains(data.ToFind, strings.ToLower(input)) {
		data.Word = RevealLetters(data)
	} else {
		data.Attempts--
		fmt.Println(data.Attempts)
	}
	return data
}
