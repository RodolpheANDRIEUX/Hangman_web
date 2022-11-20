package main

import (
	"hangman-web/Game"
	Hangman "hangman-web/Game/hangman-classic/Hangman/Game"
	"html/template"
	"net/http"
)

const file = "words.txt"

var data = Hangman.StartGame(file)

func RenderTemplate(w http.ResponseWriter, tmpl string, data Hangman.GameData) {
	t, err := template.ParseFiles("./Templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func game(w http.ResponseWriter, r *http.Request) {

	replay := r.FormValue("Replay")
	if replay != "" {
		data = Hangman.StartGame(file)
	}

	guess := r.FormValue("Letter")
	if guess != "" {
		data = Game.HangMan(data, guess)
	}
	RenderTemplate(w, "game", data)
}
