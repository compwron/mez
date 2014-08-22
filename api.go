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

func Game() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write([]byte(KoanSummaries()))
		case "POST":
			createGame(w, r)
		default:
			w.Write([]byte("not supported"))
		}
	}
}

func createGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Trying to create game")
	parsed, err := Parse(r.Body)
	if err != nil {
		http.Error(w, "malformed JSON", 400)
	} else {
		if reflect.DeepEqual(CurrentRule.ruleDescriptions, originalRule.ruleDescriptions) {
			fmt.Println("Current rule is default rule")
			submittedRule := ParseRule(parsed)
			trueKoan := ParseKoan(parsed, true)
			falseKoan := ParseKoan(parsed, false)
			trueKoanIsOk := DoesKoanFulfillRule(submittedRule, trueKoan.description)
			falseKoanIsOk := !DoesKoanFulfillRule(submittedRule, falseKoan.description)

			if trueKoanIsOk && falseKoanIsOk {
				fmt.Println("Valid new rule because its true and false koans are true and false")
				AddFullKoan(trueKoan)
				AddFullKoan(falseKoan)
				CurrentRule = submittedRule
			} else {
				w.Write([]byte("Koans do not fulfull rule; game not started.\n"))
				// w.Write(r.Body)
				w.Write([]byte("\nTrue koan is ok? " + strconv.FormatBool(trueKoanIsOk)))
				w.Write([]byte("\nFalse koan is ok? " + strconv.FormatBool(falseKoanIsOk)))
			}
		}
	}
}

func CreateKoan(w http.ResponseWriter, r *http.Request) {
	newKoanHash, err := Parse(r.Body)
	newKoan := newKoanHash["koan"].(string)
	AddKoan(newKoan)
	if err != nil {
		w.Write([]byte("Bad input"))
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
