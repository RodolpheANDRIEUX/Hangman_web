package main

import (
	"fmt"
	"github.com/toqueteos/webbrowser"
	"hangmanWeb/HangmanWeb"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const port = ":8080"

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("(http://localhost:8080) - Server started on port", port)
	err := webbrowser.Open("http://localhost:8080")
	if err != nil {
		return
	}

	http.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("Assets"))))
	http.HandleFunc("/", HangmanWeb.PathHandler)

	log.Fatal(http.ListenAndServe(port, nil))

}
