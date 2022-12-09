package HangmanWeb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	NewUser, err := json.MarshalIndent(Data.User, "", "\t")
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

func SaveUser(UsersData map[string]*UserData) {
	/*
		Save the score into the database.
	*/
	encoded, err := json.Marshal(UsersData)
	if err == nil {
		ioutil.WriteFile("Saves/users.json", encoded, 0777)
	} else {
		fmt.Println(err)
	}
}

func LoadDB(UsersData map[string]*UserData) {
	/*
		Load the scores from the database.
	*/
	data, _ := ioutil.ReadFile("Saves/users.json")
	dele := json.Unmarshal(data, &UsersData)
	if dele != nil {
		fmt.Println(dele)
	}
}

func AddUser(r *http.Request) {
	LoadDB(UsersData)
	Data.HandleUser(r)
	UsersData[Data.User.Username] = &Data.User
	SaveUser(UsersData)
}
