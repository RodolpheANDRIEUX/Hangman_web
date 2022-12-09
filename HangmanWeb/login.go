package HangmanWeb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var UsersData = make(map[string]*UserData)

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
	ptrData.User.Score = 0
}

func MarshalUser(UsersData map[string]*UserData) {
	/*
		Save the score into the database.
	*/
	encoded, err := json.MarshalIndent(UsersData, "", "\t")
	if err == nil {
		ioutil.WriteFile("Saves/users.json", encoded, 0777)
	} else {
		fmt.Println(err)
	}
}

func UnmarshalDataBase(UsersData map[string]*UserData) {
	/*
		Load the scores from the database.
	*/
	data, _ := ioutil.ReadFile("Saves/users.json")
	dele := json.Unmarshal(data, &UsersData)
	if dele != nil {
		fmt.Println(dele)
	}
}

func SucessfulLogin(r *http.Request) bool {
	UserFile, err := os.ReadFile("Saves/users.json")
	if err != nil {
		fmt.Println(err)
	}
	if (strings.Contains(string(UserFile), r.FormValue("username"))) && (strings.Contains(string(UserFile), r.FormValue("password"))) {
		return true
	}
	Data.Message = "Wrong username of password"
	return false
}

func (ptrData *WebData) LoadUser(user string) {
	UnmarshalDataBase(UsersData)
	ptrData.User = *UsersData[user]
}

func UpdateDatabase(r *http.Request) {
	UnmarshalDataBase(UsersData)
	if r != nil {
		Data.HandleUser(r)
	}
	UsersData[Data.User.Username] = &Data.User
	MarshalUser(UsersData)
}
