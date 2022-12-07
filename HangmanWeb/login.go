package HangmanWeb

import (
	"net/http"
)

func (ptrData *WebData) HandleUser(r *http.Request) {
	ptrData.User.Username = r.FormValue("Username")
	ptrData.User.Password = r.FormValue("Password")
}

func (ptrData *WebData) GetLanguage(r *http.Request) {
	ptrData.User.Language = r.Header.Get("Accept-Language")[:2]
}
