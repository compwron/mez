package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Rule struct {
	ruleDescriptions []string
}

type Koan struct {
	description  string
	fulfillsRule bool
}

var originalRule = Rule{strings.Split("1^", ",")}
var currentRule = originalRule
var koans []Koan

func Instructions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("How to play:" +
		"\n		POST /game {\"rule\": \"new rule\"} to start game" +
		"\n		GET /game to see current rule and current koans w/ outcomes" +
		"\n		POST /game/koan {\"koan\": \"new koan\"} to submit a koan (get boolean win/fail back)" +
		"\n		POST /game/guess {\"rule\": \"your guess for the rule\"} to possibly end game"))
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	parsed, err := Parse(r.Body)
	if err != nil {
		http.Error(w, "malformed JSON", 400)
	} else {
		fmt.Println(parsed)
		// need golang comparator

		if reflect.DeepEqual(currentRule.ruleDescriptions, originalRule.ruleDescriptions) {
			currentRule, _, _ = parseRule(parsed)
		}
		fmt.Println("new current rule is", currentRule)
		// currentRule = Rule{"2^MG"} // must have two upright medium size green pyramids
	}
}

func parseRule(data map[string]interface{}) (Rule, Koan, Koan) {
	newRule := data["rule"].(string)
	fmt.Println(newRule, "new rule!!")
	return Rule{strings.Split(newRule, ",")}, Koan{}, Koan{}
}

func ViewGame(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(koanSummaries()))
}

func koanSummaries() string {
	summary := "Koans:\n"
	fmt.Println("koans:", koans)
	for koanNum := range koans {
		koan := koans[koanNum]
		summary += koan.description + " : " + strconv.FormatBool(koan.fulfillsRule) + "\n"
	}
	return summary
}

func CreateKoan(w http.ResponseWriter, r *http.Request) {
	newKoanHash, err := Parse(r.Body)
	newKoan := newKoanHash["koan"].(string)
	koans = append(koans, Koan{newKoan, doesKoanFulfillRule(currentRule, newKoan)})
	if err != nil {
		fmt.Println("can't get koan from response")
	}
	if doesKoanFulfillRule(currentRule, newKoan) == true {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func doesKoanFulfillRule(rule Rule, koan string) bool {
	if strings.Contains(koan, currentRule.ruleDescriptions[0]) {
		return true
	}
	return false
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

func ruleMatches(guess string) bool {
	return true
}
