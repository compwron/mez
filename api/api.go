package api

import (
	"fmt"
	"mez/json"
	"net/http"
)

type Rule struct {
	ruleDescription string
}

type Koan struct{}

var currentRule = Rule{"meaningless starter rule"}

func Instructions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("How to play:\n		POST /game {rule: [<TODO>]} to start game"))
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	parsed, err := json.Parse(r.Body)
	if err != nil {
		http.Error(w, "malformed JSON", 400)
	} else {
		fmt.Println(parsed)
		if currentRule.ruleDescription == "meaningless starter rule" { // should not be able to override game rule until game is won
			currentRule, _, _ = parseRule(parsed)
		}
		fmt.Println("new current rule is", currentRule)
		// currentRule = Rule{"2^MG"} // must have two upright medium size green pyramids
	}
}

func parseRule(data map[string]interface{}) (Rule, Koan, Koan) {
	newRule := data["rule"].(string)
	fmt.Println(newRule, "new rule!!")
	return Rule{newRule}, Koan{}, Koan{}
}

func ViewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println("viewing new rule", currentRule)
	w.Write([]byte(currentRule.ruleDescription))
}

func CreateKoan(w http.ResponseWriter, r *http.Request) {
	doesKoanFulfillRule := false
	if doesKoanFulfillRule == true {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func GuessRule(w http.ResponseWriter, r *http.Request) {}
