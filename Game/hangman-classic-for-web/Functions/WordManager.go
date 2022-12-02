package Functions

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
)

// GetWord returns a random word from the words.txt file
func GetWord(file string) string {
	Content, err := os.ReadFile("Files/" + file)
	if err != nil {
		log.Fatal(fmt.Println("Error: ", err))
	}
	sep := "\n"
	if runtime.GOOS == "windows" {
		sep = string([]byte{13, 10})
	}
	Words := strings.Split(string(Content), string(sep))
	Word := strings.ToUpper(Words[rand.Intn(len(Words))])
	return Word

}

// ContainsArray returns true if the given list contains the given letter
func ContainsArray(list []string, n string) bool {
	for _, l := range list {
		if l == n {
			return true
		}
	}
	return false
}

// StartingLettersReveal Reveals (len(word) /2 - 1) letter(s) to the word at the beginning
func StartingLettersReveal() {
	data.GoodGuess = nil
	NumberOfLetterRevealed := (len(data.ToFind) / 2) - 1

	if len(data.ToFind) >= 3 {
		for len(data.GoodGuess) < NumberOfLetterRevealed {
			RandomLetter := string(data.ToFind[rand.Intn(len(data.ToFind))])
			if !ContainsArray(data.GoodGuess, RandomLetter) {
				data.GoodGuess = append(data.GoodGuess, RandomLetter)
			}
		}
	}
}
