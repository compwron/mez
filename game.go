package main

import (
	"net/http"
	"reflect"
	"strconv"
)

func CreateGame(w http.ResponseWriter, r *http.Request) string {
	parsed, err := Parse(r.Body)
	if err != nil {
		http.Error(w, "malformed JSON", 400)
	} else {
		if ruleIsSettable() {
			submittedRule := ParseRule(parsed)

			if (parsed["true"] == nil) || (parsed["false"] == nil) {
				return "need true koan and false koan\n"
			}

			trueKoan := ParseKoan(parsed, true)
			falseKoan := ParseKoan(parsed, false)
			trueKoanIsOk := DoesKoanFulfillRule(submittedRule, trueKoan.description)
			falseKoanIsOk := !DoesKoanFulfillRule(submittedRule, falseKoan.description)

			if trueKoanIsOk && falseKoanIsOk {
				AddFullKoan(trueKoan)
				AddFullKoan(falseKoan)
				CurrentRule = submittedRule
				return "Successfully set rule\n"
			} else {
				message := "Koans do not fulfull rule; game not started." +
					"\nTrue koan is ok? " + strconv.FormatBool(trueKoanIsOk) + "\n" +
					"\nFalse koan is ok? " + strconv.FormatBool(falseKoanIsOk) + "\n"
				return message
			}
		}
	}
	return "Can't create game because game is already in progress\n"
}

func ruleIsSettable() bool {
	return reflect.DeepEqual(CurrentRule.ruleDescriptions, OriginalRule.ruleDescriptions)
}
