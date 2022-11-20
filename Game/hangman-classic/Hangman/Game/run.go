package Hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Data GameData

// StartGame : Start a new game
func StartGame(arg string) GameData {
	Data.Word = ""
	WordsSlice := GetFile("Game/hangman-classic/Hangman/Assets/" + arg)
	Data.ToFind = strings.ToUpper(TakeRandomWord(WordsSlice))
	Data = InitialLetters(Data)
	return Data
}

// InputRead : Simply read an input form our terminal
func InputRead() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose:")
	guess, _ := reader.ReadString('\n')
	guess = strings.TrimSuffix(strings.ToUpper(guess), "\n")
	return guess
}

func Run(arg string) {
	Data = StartGame(arg)
	fmt.Println("Hangman\n" + "You have 10 attempts to find the word")

	PrintWord(Data.Word)

	Data.Guess = []string{}

	for Data.Attempts < 10 {
		if Data.State == "won" {
			fmt.Println(Lime("Congrats ! You find the word."))
			os.Exit(0)
		}

		guess := InputRead()
		Data = IntputTesting(guess, Data)

		switch Data.State {
		case "alreadyFound":
			fmt.Println(Yellow("Letter " + guess + " was already found..."))
		case "invalidInput":
			fmt.Println(Red("Please enter a valid guess"))
		case "goodGuess":
			fmt.Println(Green("Good guess !"))
			Data.Word = RevealLetters(Data)
			PrintWord(Data.Word)
			if WordGuessed(Data) {
				Data.State = "won"
			}
		case "alreadyGuessed":
			fmt.Println(Yellow("Letter " + guess + " was already used..."))
		case "badGuess":
			Data.Attempts++
			fmt.Print(Yellow("Sorry, " + guess + " is not in the word... "))
			fmt.Print(Purple(10 - Data.Attempts))
			fmt.Println(Yellow(" attempts remaining"))
			fmt.Println(OpenJose()[Data.Attempts-1])
			PrintWord(Data.Word)
		case "badWordGuessed":
			Data.Attempts += 2
			fmt.Print(Yellow("Sorry, " + guess + " is not the word... "))
			fmt.Print(Purple(10 - Data.Attempts))
			fmt.Println(Yellow(" attempts remaining"))
			fmt.Println(OpenJose()[Data.Attempts-2])
			PrintWord(Data.Word)
		}
	}
	Data.State = "lose"
	fmt.Println(Red("Sorry, you loose!, word was:", Data.ToFind))
}
