package Functions

import (
	"fmt"
	"os"
	"strings"
)

type Game struct {
	ToFind        string
	WordDisplayed string
	Tries         int
	GoodGuess     []string
	BadGuess      []string
}

var data Game

// NewGame set all to zero
func NewGame(WordFile string) {
	data.ToFind = GetWord(WordFile)
	data.Tries = 0
	data.GoodGuess = nil
	data.BadGuess = nil
	StartingLettersReveal()
	UpdateWordDisplayed()
}

// ProcessLetter manages the guess to increment or not the tries
func ProcessLetter(guess string) {
	if len(guess) > 1 {
		if guess == data.ToFind {
			data.WordDisplayed = data.ToFind
		} else {
			data.Tries += 2
			PrintHangman()
		}
		return
	}
	if !strings.Contains(data.ToFind, guess) {
		data.BadGuess = append(data.BadGuess, guess)
		data.Tries++
		PrintHangman()
		return
	}
	data.GoodGuess = append(data.GoodGuess, guess)
}

// Run is a call function that runs all the game
func Run(Save string) {
	if Save != "" {
		ReadJSON("Saves/" + Save + ".json")
	} else {
		NewGame(os.Args[1])
	}

	for data.Tries < 10 {
		fmt.Println(data.WordDisplayed)
		guess := AskForLetter()
		ProcessLetter(guess)
		UpdateWordDisplayed()

		if data.WordDisplayed == data.ToFind {
			fmt.Printf("GG!")
			return
		}
	}
	fmt.Println("GAME OVER")
}
