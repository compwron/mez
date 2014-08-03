package main

import (
	"strings"
)

type Rule struct {
	ruleDescriptions []string
}

func ParseRule(data map[string]interface{}) (Rule, Koan, Koan) {
	newRule := data["rule"].(string)
	return Rule{strings.Split(newRule, ",")}, Koan{}, Koan{}
}

func ruleMatches(guess string) bool {
	return true
}
