package main

import (
	"strings"
)

type Rule struct {
	ruleDescriptions []string
}

func ParseRule(data map[string]interface{}) Rule {
	newRule := data["rule"].(string)
	return Rule{strings.Split(newRule, ",")}
}

func ruleMatches(guess string) bool {
	return true
}
