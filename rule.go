package main

import (
	"strings"
)

type Rule struct {
	ruleDescriptions []string
}

func RuleSummary() string {
	same := (len(CurrentRule.ruleDescriptions) == 1) && (CurrentRule.ruleDescriptions[0] == OriginalRule.ruleDescriptions[0])
	if same {
		return "current rule is original rule"
	} else {
		return "current rule is not original rule"
	}
}

func ParseRule(data map[string]interface{}) Rule {
	newRule := data["rule"].(string)
	return Rule{strings.Split(newRule, ",")}
}

func ruleMatches(guess string) bool {
	return true
}
