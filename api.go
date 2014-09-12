package main

import (
	"net/http"
	"strconv"
)

func StartGameWithUnknownRule() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Write([]byte(GenerateRule()))
		default:
			http.Error(w, "not supported", 405)
		}
	}
}

func Game() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write([]byte(KoanSummaries() + "\n" + RuleSummary()))
		case "POST":
			w.Write([]byte(CreateGame(w, r)))
		default:
			http.Error(w, "not supported", 405)
		}
	}
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
		w.Write([]byte("Bad input\n"))
	}

	w.Write([]byte(strconv.FormatBool(doesKoanFulfillRule)))
	w.Write([]byte("\n"))
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
	if RuleMatches(ruleGuess) {

		w.Write([]byte("Victory!\n"))

		//  reset rule and koans list
		CurrentRule = OriginalRule
		Koans = nil // move this into method in Koan.go ?
	} else {
		w.Write([]byte("incorrect guess\n"))
	}
}
