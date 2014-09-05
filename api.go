package main

import (
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
	parsed, err := Parse(r.Body)
	if err != nil {
		http.Error(w, "malformed JSON", 400)
	} else {
		if ruleIsSettable() {
			submittedRule := ParseRule(parsed)

			if (parsed["true"] == nil) || (parsed["false"] == nil) {
				w.Write([]byte("need true koan and false koan"))
				return
			}

			trueKoan := ParseKoan(parsed, true)
			falseKoan := ParseKoan(parsed, false)
			trueKoanIsOk := DoesKoanFulfillRule(submittedRule, trueKoan.description)
			falseKoanIsOk := !DoesKoanFulfillRule(submittedRule, falseKoan.description)

			if trueKoanIsOk && falseKoanIsOk {
				AddFullKoan(trueKoan)
				AddFullKoan(falseKoan)
				CurrentRule = submittedRule
				w.Write([]byte("true"))
			} else {
				w.Write([]byte("Koans do not fulfull rule; game not started.\n"))
				w.Write([]byte("\nTrue koan is ok? " + strconv.FormatBool(trueKoanIsOk)))
				w.Write([]byte("\nFalse koan is ok? " + strconv.FormatBool(falseKoanIsOk)))
			}
		}
	}
}

func ruleIsSettable() bool {
	return reflect.DeepEqual(CurrentRule.ruleDescriptions, OriginalRule.ruleDescriptions)
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
	if DoesKoanFulfillRule(CurrentRule, newKoan) {
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

		//  reset rule and koans list
		CurrentRule = OriginalRule
		Koans = nil
	} else {
		w.Write([]byte("false guess"))
	}
}
