package HangmanWeb

import (
	HangmanClassic "hangmanWeb/hangman-classic-for-web/Functions"
	"strings"
)

// Reset : Simply reset the game if we click on replay
func (ptrData *WebData) Reset() {
	ptrData.State = "InGame"
	ptrData.Error = false
	ptrData.Message = ""
}

// CheckTries : Check if the game is lost
func (ptrData *WebData) CheckTries() {
	if ptrData.Game.Tries >= 10 {
		ptrData.State = "LOST"
	}
}

// HandleLetter : Our game loop. We deal with our guess
func (ptrData *WebData) HandleLetter(guess *string) {
	if len(*guess) > 1 {
		if *guess == ptrData.Game.ToFind {
			ptrData.Game.WordDisplayed = ptrData.Game.ToFind
			ptrData.State = "WIN"
		} else {
			ptrData.Game.Tries += 2
			ptrData.CheckTries()
		}
		return
	}
	if !strings.Contains(ptrData.Game.ToFind, *guess) {
		ptrData.Game.BadGuess = append(ptrData.Game.BadGuess, *guess)
		ptrData.Game.Tries++
		ptrData.CheckTries()
		return
	}
	ptrData.Game.GoodGuess = append(ptrData.Game.GoodGuess, *guess)
	HangmanClassic.UpdateWordDisplayed(&ptrData.Game)
	if ptrData.Game.WordDisplayed == ptrData.Game.ToFind {
		ptrData.State = "WIN"
	}
}

// ErrorInLetter : Check if the input is valid or not
func (ptrData *WebData) ErrorInLetter(guess *string) bool {
	switch {

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

// GetWordDifficulty : Assign a file for each difficulty chosen
func (ptrData *WebData) GetWordDifficulty() string {
	switch ptrData.Difficulty {
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

// UpdateScore : Update the player score for each difficulty chosen
func (ptrData *WebData) UpdateScore() {
	if ptrData.State == "WIN" {
		switch ptrData.Difficulty {
		case "easy":
			ptrData.User.Score++
		case "medium":
			ptrData.User.Score += 2
		case "hard":
			ptrData.User.Score += 3
		}
	}
	UpdateDatabase(ptrData.User.Username)
}
