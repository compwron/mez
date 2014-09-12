package main

import (
	"strings"
)

var OriginalRule = Rule{strings.Split("1^", ",")} // TODO fix syntax
var CurrentRule = OriginalRule

type Rule struct {
	ruleDescriptions []string
}

func RuleSummary() string {
	same := (len(CurrentRule.ruleDescriptions) == 1) && (CurrentRule.ruleDescriptions[0] == OriginalRule.ruleDescriptions[0])
	if same {
		return "current rule is original rule\n"
	} else {
		return "current rule is not original rule\n"
	}
}

func ParseRule(data map[string]interface{}) Rule {
	newRule := data["rule"].(string)
	return Rule{strings.Split(newRule, ",")}
}

func RuleMatches(guess string) bool {
	return guess == unparseRuleDescriptions(CurrentRule.ruleDescriptions)
}

func unparseRuleDescriptions(ruleDescriptions []string) string {
	return strings.Join(ruleDescriptions, ",")
}
