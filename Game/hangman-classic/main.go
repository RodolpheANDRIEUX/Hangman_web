package main

import (
	"fmt"
	Hangman "hangman-web/Game/hangman-classic/Hangman/Game"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please enter words.txt as argument to run the game")
		os.Exit(1)
	}
	arg := args[1]
	Hangman.Run(arg)
}
