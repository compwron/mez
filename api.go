package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var OriginalRule = Rule{strings.Split("1^", ",")}
var CurrentRule = OriginalRule

func Game() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write([]byte(KoanSummaries() + "\n" + RuleSummary()))
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
		if reflect.DeepEqual(CurrentRule.ruleDescriptions, OriginalRule.ruleDescriptions) {
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

func CreateKoan() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			acceptKoan(w, r)
		default:
			w.Write([]byte("not supported"))
		}
	}
}

func acceptKoan(w http.ResponseWriter, r *http.Request) {
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

func GuessRule() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			guessRule(w, r)
		default:
			w.Write([]byte("not supported"))
		}
	}
}

func guessRule(w http.ResponseWriter, r *http.Request) {
	// if rule matches, end game
	ruleGuessHash, err := Parse(r.Body)
	if err != nil {
		w.Write([]byte("Can't get rule from response"))
	}

	ruleGuess := ruleGuessHash["rule"].(string)
	if ruleMatches(ruleGuess) {

		w.Write([]byte("true"))
		fmt.Println("Game won! Rule reset.")

		//  reset rule and koans list
		CurrentRule = OriginalRule
		Koans = nil
	} else {
		w.Write([]byte("false guess"))
	}
}
