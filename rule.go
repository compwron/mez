package main

import (
	"reflect"
	"strings"
)

var OriginalRule = Rule{strings.Split("1^", ",")} // TODO fix syntax
var CurrentRule = OriginalRule
var ruleTypes = []string{"count", "color", "size", "orientation", "pip"} // more coming soon

type Rule struct {
	ruleDescriptions []string
}

func SyntacticallyValidRule(rule Rule) bool {
	for _, ruleChunk := range rule.ruleDescriptions {
		if lastPieceIsNegative(ruleChunk) { // cannot have ! not immediately before something else
			return false
		}
		hasValidRuleType := false
		for _, ruleType := range ruleTypes {
			if ruleContains(ruleChunk, ruleType) {
				hasValidRuleType = true
			}
		}

		if !hasValidRuleType {
			return false
		}
	}
	return true
}

func lastPieceIsNegative(chunk string) bool {
	trimmed := strings.TrimRight(chunk, "!")
	return !(trimmed == chunk)
}

func RuleSummary() string {
	same := (len(CurrentRule.ruleDescriptions) == 1) && (CurrentRule.ruleDescriptions[0] == OriginalRule.ruleDescriptions[0])
	if same {
		return "current rule is original rule\n"
	}
	return "current rule is NOT original rule\n"
}

func ParseRule(data map[string]interface{}) Rule {
	newRule := data["rule"].(string)
	return Rule{strings.Split(newRule, ",")}
}

func RuleMatches(guess string) bool {
	return guess == unparseRuleDescriptions(CurrentRule.ruleDescriptions)
}

func OkToChangeCurrentRule() bool {
	return reflect.DeepEqual(CurrentRule, OriginalRule)
}

func unparseRuleDescriptions(ruleDescriptions []string) string {
	return strings.Join(ruleDescriptions, ",")
}
