package Hangman

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

// GameData : The structure of our game
type GameData struct {
	Word     string   // Word with _ ( ex A _ _ _ )
	ToFind   string   // Final word chosen by the program at the beginning. It is the word to find
	Index    []int    // Indexes of the good letter guessed
	Attempts int      // Number of attempts left
	Tries    []string // Letters tried
}

// NewGame : Start a new game
func NewGame(WordToFind string) GameData {
	return GameData{
		"",
		WordToFind,
		[]int{},
		10,
		[]string{},
	}
}

// GetFile : Take the words.txt file convert it into []string
func GetFile(filePath string) []string {
	WordFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	WordStr := string(WordFile)                // Transform our file into string
	WordsSlice := strings.Split(WordStr, "\n") // strings.Split function slice a string when it found a parameter, here is "/n"
	return WordsSlice
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
	var initialLetters []string
	var n int
	if len(data.ToFind) < 4 { // error when len word < 3 ( 3/2-1 == 0 ) We cannot take a random number between 0 & 0
		n = 1

	} else {
		n = GetRandomNumber((len(data.ToFind) / 2) - 1)
	}
	if n == 0 {
		n = (len(data.ToFind) / 2) - 1
	}
	for i := 0; i < n; i++ {
		index := GetRandomNumber(len(data.ToFind))
		initialLetters = append(initialLetters, string(data.ToFind[index]))
		data.Tries = initialLetters
	}
	return data
}

// RevealInitialLetters : Take InitialLetters indexes and return the initial word to show
func RevealInitialLetters(data GameData) GameData {
	LetterInWord := strings.Split(data.ToFind, "")
	for _, letter := range data.Tries {
		for i, v := range LetterInWord {
			if letter == v {
				data.Index = append(data.Index, i)
			}
		}
	}
	InitialWord := make([]string, len(data.ToFind))
	for i := range data.ToFind {
		InitialWord[i] = "_"
	}
	for _, index := range data.Index {
		InitialWord[index] = string(data.ToFind[index])
	}
	reveal := strings.Join(InitialWord, "")
	data.Word = strings.ToUpper(reveal)
	return data
}

// RevealLetters : return the word with a letter revealed
func RevealLetters(data GameData) string {
	revealRune := []rune(data.Word)
	wordRune := []rune(data.ToFind)
	for _, v := range data.Index {
		revealRune[v] = wordRune[v]
	}
	data.Word = strings.ToUpper(string(revealRune))
	return data.Word
}

// FindLetter : Store the index of the letter guessed on the word to print it.
func FindLetter(data GameData) GameData {
	letter := strings.ToLower(data.Tries[len(data.Tries)-1])
	letters := strings.Split(data.ToFind, "")
	for i, v := range letters {
		if v == letter {
			data.Index = append(data.Index, i)
		}
	}
	return data
}
