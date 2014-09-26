package main

import (
	"reflect"
	"strconv"
)

func CreateGame(parsed map[string]interface{}) string {
	if ruleIsSettable() {
		submittedRule := ParseRule(parsed)

		if (parsed["true"] == nil) || (parsed["false"] == nil) {
			return "need true koan and false koan"
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
		}
		message := "Koans do not fulfull rule; game not started." +
			"\nTrue koan is ok? " + strconv.FormatBool(trueKoanIsOk) + "\n" +
			"\nFalse koan is ok? " + strconv.FormatBool(falseKoanIsOk)
		return message

	}
	return "Can't create game because game is already in progress"
}

func ruleIsSettable() bool {
	return reflect.DeepEqual(CurrentRule.ruleDescriptions, OriginalRule.ruleDescriptions)
}
