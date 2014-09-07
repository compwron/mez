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
			http.Error(w, "not supported", 405)
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
				w.Write([]byte("need true koan and false koan\n"))
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
				return
			} else {
				w.Write([]byte("Koans do not fulfull rule; game not started.\n"))
				w.Write([]byte("\nTrue koan is ok? " + strconv.FormatBool(trueKoanIsOk) + "\n"))
				w.Write([]byte("\nFalse koan is ok? " + strconv.FormatBool(falseKoanIsOk) + "\n"))
				return
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
			http.Error(w, "not supported", 405)
		}
	}
}

func acceptKoan(w http.ResponseWriter, r *http.Request) {
	newKoanHash, err := Parse(r.Body)
	newKoan := newKoanHash["koan"].(string)
	doesKoanFulfillRule := AddKoan(newKoan)
	if err != nil {
		w.Write([]byte("Bad input"))
	}

	w.Write([]byte(strconv.FormatBool(doesKoanFulfillRule)))
}

func GuessRule() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			guessRule(w, r)
		default:
			http.Error(w, "not supported", 405)
		}
	}
}

func guessRule(w http.ResponseWriter, r *http.Request) {
	// if rule matches, end game
	ruleGuessHash, err := Parse(r.Body)
	if err != nil {
		w.Write([]byte("Can't get rule from response\n"))
	}

	ruleGuess := ruleGuessHash["rule"].(string)
	if ruleMatches(ruleGuess) {

		w.Write([]byte("Victory!\n"))

		//  reset rule and koans list
		CurrentRule = OriginalRule
		Koans = nil
	} else {
		w.Write([]byte("incorrect guess\n"))
	}
}
