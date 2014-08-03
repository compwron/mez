package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Koan struct {
	description  string
	fulfillsRule bool
}

var originalRule = Rule{strings.Split("1^", ",")}
var currentRule = originalRule

func Instructions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("How to play:" +
		"\n		POST /game {\"rule\": \"new rule\"} to start game" +
		"\n		GET /game to see current rule and current koans w/ outcomes" +
		"\n		POST /game/koan {\"koan\": \"new koan\"} to submit a koan (get boolean win/fail back)" +
		"\n		POST /game/guess {\"rule\": \"your guess for the rule\"} to possibly end game"))
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Trying to create game")
	parsed, err := Parse(r.Body)
	if err != nil {
		http.Error(w, "malformed JSON", 400)
	} else {
		if reflect.DeepEqual(currentRule.ruleDescriptions, originalRule.ruleDescriptions) {
			currentRule, _, _ = ParseRule(parsed)
			fmt.Println("Set new Rule to", currentRule)
		}
	}
}

func ViewGame(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(KoanSummaries()))
}

func CreateKoan(w http.ResponseWriter, r *http.Request) {
	newKoanHash, err := Parse(r.Body)
	newKoan := newKoanHash["koan"].(string)
	Koans = append(Koans, Koan{newKoan, DoesKoanFulfillRule(currentRule, newKoan)})
	if err != nil {
		fmt.Println("can't get koan from response")
	}
	if DoesKoanFulfillRule(currentRule, newKoan) == true {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func GuessRule(w http.ResponseWriter, r *http.Request) {
	// if rule matches, end game
	ruleGuessHash, err := Parse(r.Body)
	if err != nil {
		fmt.Println("Can't get rule from response")
	}
	ruleGuess := ruleGuessHash["rule"].(string)
	if ruleMatches(ruleGuess) {
		currentRule = originalRule
		w.Write([]byte("true"))
		fmt.Println("Game won! Rule reset.")
	} else {
		w.Write([]byte("false guess"))
	}
}
