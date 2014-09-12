package main

import (
	"net/http"
	"reflect"
	"strconv"
)

func CreateGame(w http.ResponseWriter, r *http.Request) {
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
