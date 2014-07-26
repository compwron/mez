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

var originalRule = Rule{"meaningless starter rule"}
var currentRule = originalRule

func Instructions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("How to play:" +
		"\n		POST /game {\"rule\": \"new rule\"} to start game" +
		"\n		GET /game to see current rule and current koans w/ outcomes" +
		"\n		POST /game/koan {\"koan\": \"new koan\"} to submit a koan (get boolean win/fail back)" +
		"\n		POST /game/guess {\"rule\": \"your guess for the rule\"} to possibly end game"))
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
	newKoanHash, err := json.Parse(r.Body)
	newKoan := newKoanHash["koan"].(string)
	if err != nil {
		fmt.Println("can't get koan from response")
	}
	if doesKoanFulfillRule(newKoan) == true {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func doesKoanFulfillRule(koan string) bool {

	// rule = "1^MG, 1>S, !2>"

	// koan = "1>SG, 1^MG"

	// rule = "1^"

	// koan = "1^MG"

	return false
}

func GuessRule(w http.ResponseWriter, r *http.Request) {
	// if rule matches, end game
	ruleGuessHash, err := json.Parse(r.Body)
	if err != nil {
		fmt.Println("Can't get rule from response")
	}
	ruleGuess := ruleGuessHash["rule"].(string)
	if ruleMatches(ruleGuess) {
		currentRule = Rule{"meaningless starter rule"}
		w.Write([]byte("true"))
		fmt.Println("Game won! Rule reset.")
	} else {
		w.Write([]byte("false"))
	}
}

func ruleMatches(guess string) bool {
	return true
}
