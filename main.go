package main

import (
	"fmt"
	"hangman-web/Game"
	"log"
	"net/http"
)

const port = ":8080"

func main() {

	fmt.Println("(http://localhost:8080) - Server started on port", port)

	http.HandleFunc("/", Game.PathHandler)

	log.Fatal(http.ListenAndServe(port, nil))

}
