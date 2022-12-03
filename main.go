package main

import (
	"fmt"
	"hangmanWeb/Game"
	"log"
	"net/http"
)

const port = ":8080"

func main() {

	fmt.Println("(http://localhost:8080) - Server started on port", port)

	http.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("Assets"))))
	http.HandleFunc("/", Game.PathHandler)

	log.Fatal(http.ListenAndServe(port, nil))

}
