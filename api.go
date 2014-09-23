package main

import (
	"net/http"
)

var NOT_SUPPORTED = "not supported"

func StartGameWithUnknownRule() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Write([]byte(GenerateRule()))
		default:
			http.Error(w, NOT_SUPPORTED, 405)
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
			http.Error(w, NOT_SUPPORTED, 405)
		}
	}
}

func CreateKoan() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			newKoanHash, err := Parse(r.Body)
			if err != nil {
				http.Error(w, "Impossible to parse formatting\n", 405)
				return
			}
			w.Write([]byte(AddKoanIfValid(newKoanHash)))
			w.Write([]byte("\n"))
		default:
			http.Error(w, NOT_SUPPORTED, 405)
		}
	}
}

func GuessRule() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			ruleGuessHash, err := Parse(r.Body)
			if err != nil || (ruleGuessHash["rule"] == nil) {
				http.Error(w, "Can't get rule from response\n", 400)
				return
			}

			w.Write([]byte(guessRule(ruleGuessHash) + "\n"))
		default:
			http.Error(w, NOT_SUPPORTED, 405)
		}
	}
}

func guessRule(ruleGuessHash map[string]interface{}) string {
	// if rule matches, end game
	ruleGuess := ruleGuessHash["rule"].(string)
	if RuleMatches(ruleGuess) {

		//  reset rule and koans list
		CurrentRule = OriginalRule
		Koans = nil // move this into method in Koan.go ?
		return "Victory!"
	}
	return "incorrect guess"
}
