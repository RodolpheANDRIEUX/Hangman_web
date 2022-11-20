package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {

	fmt.Println("(http://localhost:8080) - Server started on port", port)

	http.HandleFunc("/", game)

	http.ListenAndServe(port, nil)

}
