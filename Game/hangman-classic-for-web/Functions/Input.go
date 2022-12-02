package Functions

import (
	"fmt"
	"os"
	"strings"
)

// IsAlpha checks if the string is a letter
func IsAlpha(s string) bool {
	for _, l := range s {
		if l < 65 || l > 90 && l < 1040 || l > 1071 {
			return false
		}
	}
	return true
}

// AskForLetter and manage it with a switch
func AskForLetter() string {
	var guess string
	ValidLetter := false

	for !ValidLetter {
		fmt.Printf("Enter a letter: ")
		fmt.Scan(&guess)
		guess = strings.ToUpper(guess)

		switch {

		case guess == "STOP":
			SaveName := ""
			fmt.Print("Please enter the save name : ")
			fmt.Scanln(&SaveName)
			Marshal("Saves/" + SaveName + ".json")
			fmt.Printf("Game saved as : %v", SaveName)
			os.Exit(1)

		case guess == "DEVMODE":
			fmt.Println(data.ToFind)

		case !IsAlpha(guess):
			fmt.Print("\n\a\x1B[31mPlease, select a letter\x1B[0m\n") // Please, select a letter

		case ContainsArray(data.GoodGuess, guess):
			fmt.Print("\n\a\x1B[31mThis letter has already been Revealed\x1B[0m\n") // This letter has already been Revealed

		case ContainsArray(data.BadGuess, guess):
			fmt.Print("\n\a\x1B[31mThis letter has already been chosen\x1B[0m\n") // This letter has already been chosen

		default:
			ValidLetter = true
		}
	}
	return strings.ToUpper(guess)
}
