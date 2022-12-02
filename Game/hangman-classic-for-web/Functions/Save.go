package Functions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Marshal is used to save the current game in a file called in argument
func Marshal(file string) {
	GameData, err := json.Marshal(data)
	err = os.WriteFile(file, GameData, 0644)
	if err != nil {
		log.Fatal(fmt.Println("Error: ", err))
	}
}

// ReadJSON assigns the game saved to the current game
func ReadJSON(file string) {
	jsonFile, err := os.Open(file)
	err = json.NewDecoder(jsonFile).Decode(&data)
	err = jsonFile.Close()
	err = os.Remove(file)
	if err != nil {
		log.Fatal(fmt.Println("Error: ", err))
	}
}
