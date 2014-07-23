package api

import (
	"fmt"
	"mez/json"
	"net/http"
)

type Rule struct{}

type Koan struct{}

func Instructions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("How to play:\n		POST /game {rule: [<TODO>]} to start game"))
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	parsed, err := json.Parse(r.Body)
	if err != nil {
		http.Error(w, "malformed JSON", 400)
	} else {
		fmt.Println(parsed)
		rule, correct, incorrect := parseRule(parsed)
		fmt.Println(rule, correct, incorrect)
	}
}

func parseRule(data map[string]interface{}) (Rule, Koan, Koan) {
	return Rule{}, Koan{}, Koan{}
}

func ViewGame(w http.ResponseWriter, r *http.Request)   {}
func CreateKoan(w http.ResponseWriter, r *http.Request) {}
func GuessRule(w http.ResponseWriter, r *http.Request)  {}
