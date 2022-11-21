package Hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

// GetFile : Take the words.txt file convert it into []string
func GetFile(filePath string) []string {
	WordFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(WordFile)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	WordFile.Close()
	return txtlines
}

// GetRandomNumber : Return a random int value
func GetRandomNumber(i int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomIndex := r1.Intn(i)
	return randomIndex
}

// TakeRandomWord : Takes a Random word in a slice of words
func TakeRandomWord(txt []string) string {
	RandomIndex := GetRandomNumber(len(txt))
	ToFind := txt[RandomIndex]
	return ToFind
}

// InitialLetters : chose 0<n<len(word) / 2 - 1 letters in the word to print and return the index of these letters in th word
func InitialLetters(data GameData) GameData {
	var nbLetterToReveal int
	for range data.ToFind {
		data.Word += "_"
	}
	if len(data.ToFind) < 4 {
		nbLetterToReveal = 1
	} else {
		nbLetterToReveal = GetRandomNumber((len(data.ToFind) / 2) - 1)

	}
	if nbLetterToReveal == 0 {
		nbLetterToReveal = (len(data.ToFind) / 2) - 1
	}

	for i := 0; i < nbLetterToReveal; i++ {
		index := GetRandomNumber(len(data.ToFind))
		data.Guess = append(data.Guess, string(data.ToFind[index]))
		data.Word = RevealLetters(data)
	}
	return data
}
