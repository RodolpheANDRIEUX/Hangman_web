package main

import (
	"fmt"
	"hangman-web/Game"
	"log"
	"net/http"
)

// bonjour

const port = ":8080"

func main() {

	fmt.Println("(http://localhost:8080) - Server started on port", port)

	http.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("Assets"))))
	http.HandleFunc("/", Game.PathHandler)

	log.Fatal(http.ListenAndServe(port, nil))

}
