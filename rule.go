package main

import (
	"strings"
	"reflect"
	"fmt"
)

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

func UnparseRuleDescriptions(ruleDescriptions []string) string {
	return strings.Join(ruleDescriptions, ",")
}

func ruleMatches(guess string) bool {
	println("ZMD DEBUG A")
	println(guess)
	fmt.Println(reflect.TypeOf(guess))
	println(UnparseRuleDescriptions(CurrentRule.ruleDescriptions))
	return guess == UnparseRuleDescriptions(CurrentRule.ruleDescriptions)
}
