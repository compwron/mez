package main

import (
	"net/http"
)

func Instructions() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("How to play:" +
			"\n		POST /game {\"rule\":\"!3^\", \"true\" : \"1^SG\", \"false\" : \"3^SG\"} to start game" +
			"\n		GET /game to see current rule status and current koans w/ outcomes" +
			"\n		POST /game/koan {\"koan\": \"new koan\"} to submit a koan (get boolean win/fail back)" +
			"\n 	POST /game/generate to create a new game with a randomly generated rule (only works when no game is in progress)" +
			"\n		POST /game/guess {\"rule\": \"your guess for the rule\"} to possibly end game\n"))
	}
}
