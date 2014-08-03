package main

import (
	"strconv"
	"strings"
)

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	ruleCharacters := strings.Split(rule.ruleDescriptions[0], "")
	koanCharacters := strings.Split(koan, "")

	rulePieceCount, ruleErr := intOf(ruleCharacters[0])
	koanPieceCount, koanErr := intOf(koanCharacters[0])

	if ruleErr != nil || koanErr != nil {
		return false
	}

	if koanPieceCount > rulePieceCount {
		return true
	}

	if strings.Contains(koan, currentRule.ruleDescriptions[0]) {
		return true
	}
	return false
}

func intOf(char string) (int64, error) {
	return strconv.ParseInt(char, 10, 8)
}
