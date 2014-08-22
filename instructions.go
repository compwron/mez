package main

import (
	"net/http"
)

func Instructions() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("How to play:" +
			"\n		POST /game {\"rule\": \"new rule\"} to start game" +
			"\n		GET /game to see current rule and current koans w/ outcomes" +
			"\n		POST /game/koan {\"koan\": \"new koan\"} to submit a koan (get boolean win/fail back)" +
			"\n		POST /game/guess {\"rule\": \"your guess for the rule\"} to possibly end game"))
	}
}
