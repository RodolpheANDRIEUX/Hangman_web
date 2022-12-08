package HangmanWeb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func CheckUsernameAvailability(r *http.Request) bool {
	Data.Message = ""
	UserFile, err := os.ReadFile("Saves/users.json")
	if err != nil {
		fmt.Println(err)
	}
	if strings.Contains(string(UserFile), r.FormValue("username")) {
		Data.Message = "This username is already taken"
		return false
	}
	return true
}

func (ptrData *WebData) HandleUser(r *http.Request) {
	ptrData.User.Username = r.FormValue("username")
	ptrData.User.Password = r.FormValue("password")
	ptrData.User.Language = r.Header.Get("Accept-Language")[:2]
}

func MarshallUser() {
	NewUser, err := json.Marshal(Data.User)

	UserFile, err := os.OpenFile("Saves/users.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer UserFile.Close()

	n, err := UserFile.Write(NewUser)
	if err != nil {
		fmt.Println(n, err)
	}

	if n, err = UserFile.WriteString("\n"); err != nil {
		fmt.Println(n, err)
	}
}

func AddUser(r *http.Request) {
	Data.HandleUser(r)
	MarshallUser()
}
