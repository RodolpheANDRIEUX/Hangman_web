package HangmanWeb

import (
	HangmanClassic "hangmanWeb/hangman-classic-for-web/Functions"
	"strings"
)

func (ptrData *WebData) Reset() {
	ptrData.State = "InGame"
	ptrData.Error = false
	ptrData.Message = ""
}

func (ptrData *WebData) CheckTries() {
	if ptrData.Game.Tries >= 10 {
		ptrData.State = "lost"
	}
}

func (ptrData *WebData) HandleLetter(guess *string) {
	if len(*guess) > 1 {
		if *guess == ptrData.Game.ToFind {
			ptrData.Game.WordDisplayed = ptrData.Game.ToFind
			ptrData.State = "won"
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
		ptrData.State = "won"
	}
}

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

func (user *UserData) GetWordDifficulty() string {
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
