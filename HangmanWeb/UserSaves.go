package HangmanWeb

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var UsersData = make(map[string]*UserData)

// CheckUsernameAvailability : All username must be unique or non nil. This function make sure of this
func CheckUsernameAvailability(r *http.Request) bool {
	Data.Message = ""
	UserFile, err := os.ReadFile("Saves/users.json")
	if err != nil {
		fmt.Println(err)
	}
	if strings.Contains(string(UserFile), r.FormValue("username")) {
		Data.Message = "This username is already taken"
		return false
	} else if (r.FormValue("username")) == "" {
		Data.Message = "Username cannot be blank"
		return false
	}
	return true
}

// GetPasswordHash : Wash the password using bcrypt package
func GetPasswordHash(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return []byte(nil), err
	}
	return hash, nil
}

// HandleUser : We update the Structure with the values of the signup form
func (ptrData *WebData) HandleUser(r *http.Request) {
	ptrData.User.Username = r.FormValue("username")
	ptrData.User.Password, _ = GetPasswordHash(r.FormValue("password"))
	ptrData.User.Language = r.Header.Get("Accept-Language")[:2]
	ptrData.User.Score = 0
}

// MarshalUser : We save the User information into the database.
func MarshalUser(UsersData map[string]*UserData) {
	encoded, err := json.MarshalIndent(UsersData, "", "\t")
	if err == nil {
		ioutil.WriteFile("Saves/users.json", encoded, 0777)
	} else {
		fmt.Println(err)
	}
}

// UnmarshalDataBase : We load the database.
func UnmarshalDataBase(UsersData map[string]*UserData) {
	data, _ := ioutil.ReadFile("Saves/users.json")
	dele := json.Unmarshal(data, &UsersData)
	if dele != nil {
		fmt.Println(dele)
	}
}

// SuccessfulLogin : Check if the combination login-password is good
func SuccessfulLogin(r *http.Request) bool {
	UnmarshalDataBase(UsersData)
	for key, _ := range UsersData {
		if r.FormValue("username") == key {
			if bcrypt.CompareHashAndPassword(UsersData[key].Password, []byte(r.FormValue("password"))) == nil {
				return true
			}
		}
	}
	Data.Message = "Wrong username or password"
	return false
}

// LoadUser : We update the user infos after log in
func (ptrData *WebData) LoadUser(user string) {
	UnmarshalDataBase(UsersData)
	ptrData.User = *UsersData[user]
}

// UpdateDatabase : We simply update the database with de Data we have
func UpdateDatabase(r *http.Request) {
	UnmarshalDataBase(UsersData)
	if r != nil {
		Data.HandleUser(r)
	}
	UsersData[Data.User.Username] = &Data.User
	MarshalUser(UsersData)
}
