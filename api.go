package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var originalRule = Rule{strings.Split("1^", ",")}
var CurrentRule = originalRule

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
		if reflect.DeepEqual(CurrentRule.ruleDescriptions, originalRule.ruleDescriptions) {
			submittedRule := ParseRule(parsed)
			trueKoan := ParseKoan(parsed, true)
			falseKoan := ParseKoan(parsed, false)
			trueKoanIsOk := DoesKoanFulfillRule(submittedRule, trueKoan.description)
			falseKoanIsOk := !DoesKoanFulfillRule(submittedRule, falseKoan.description)

			if trueKoanIsOk && falseKoanIsOk {
				AddFullKoan(trueKoan)
				AddFullKoan(falseKoan)
				CurrentRule = submittedRule
			} else {
				w.Write([]byte("Koans do not fulfull rule; game not started.\n"))
				// w.Write(r.Body)
				w.Write([]byte("True koan is ok? " + strconv.FormatBool(trueKoanIsOk)))
				w.Write([]byte("False koan is ok? " + strconv.FormatBool(falseKoanIsOk)))
			}
		}
	}
}

func ViewGame(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(KoanSummaries()))
}

func CreateKoan(w http.ResponseWriter, r *http.Request) {
	newKoanHash, err := Parse(r.Body)
	newKoan := newKoanHash["koan"].(string)
	AddKoan(newKoan)
	if err != nil {
		fmt.Println("can't get koan from response")
	}
	if DoesKoanFulfillRule(CurrentRule, newKoan) == true {
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
		CurrentRule = originalRule
		w.Write([]byte("true"))
		fmt.Println("Game won! Rule reset.")
	} else {
		w.Write([]byte("false guess"))
	}
}
