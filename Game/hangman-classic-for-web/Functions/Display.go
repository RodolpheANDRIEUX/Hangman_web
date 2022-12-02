package Functions

import "fmt"

// UpdateWordDisplayed Display the word to find without the missing letters
func UpdateWordDisplayed() {
	data.WordDisplayed = ""
	for _, l := range data.ToFind {
		if ContainsArray(data.GoodGuess, string(l)) {
			data.WordDisplayed += string(l)
		} else {
			data.WordDisplayed += "_"
		}
	}
}

// PrintHangman Supposed to print hangman here
func PrintHangman() {
	fmt.Printf("%v tries left\n", 10-data.Tries)
}
