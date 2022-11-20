package Hangman

import (
	"io/ioutil"
	"log"
	"strings"
)

// OpenJose : Open hangman.txt file and return a string sliced fore each positions of the hangman
func OpenJose() []string {
	JoseFile, err := ioutil.ReadFile("./Hangman/Assets/jose.txt")
	if err != nil {
		log.Fatal(err)
	}
	JoseStr := string(JoseFile)              // Transform our file into string
	JoseSlice := strings.Split(JoseStr, "1") // strings.Split function slice a string when it found a parameter, here is "\n\n"
	return JoseSlice
}
