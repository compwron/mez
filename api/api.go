package api

import (
	"net/http"
)

func Instructions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("How to play:\n		POST /game {rule: [<TODO>]} to start game"))
}

func CreateGame(w http.ResponseWriter, r *http.Request) {

}

func ViewGame(w http.ResponseWriter, r *http.Request)   {}
func CreateKoan(w http.ResponseWriter, r *http.Request) {}
func GuessRule(w http.ResponseWriter, r *http.Request)  {}
