package main

import (
	"strconv"
	"strings"
)

func remove(data []string, item string) []string {
	newData := make([]string, len(data)-1)
	for i := 0; i < len(data); i++ {
		if data[i] != item {
			newData = append(newData, data[i])
		}
	}
	return newData
}

func analyzeSingleRule(ruleDescription string) (bool, int) {
	ruleNot := false
	ruleCharacters := strings.Split(ruleDescription, "")
	if ruleCharacters[0] == "!" {
		ruleNot = true
		ruleCharacters = ruleCharacters[1:]
	}

	rulePieceCount := intOf(ruleCharacters[0])
	return ruleNot, rulePieceCount
}

func splitKoan(koan string) (bool, int) {
	invalidKoan := false
	koanCharacters := strings.Split(koan, "")
	if koanCharacters[0] == "!" {
		invalidKoan = true
	}

	koanPieceCount := intOf(koanCharacters[0])

	return invalidKoan, koanPieceCount
}

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	allRulesAreValid := true
	for i := 0; i < len(rule.ruleDescriptions); i++ {
		invalidKoan, koanPieceCount := splitKoan(koan)
		if invalidKoan {
			return false
		}
		ruleNot, rulePieceCount := analyzeSingleRule(rule.ruleDescriptions[i])

		// if rule is a not, check that koanCount is anything other than ruleCount
		if ruleNot {
			if koanPieceCount == rulePieceCount {
				allRulesAreValid = false
			}
		} else {
			if !(koanPieceCount >= rulePieceCount) {
				allRulesAreValid = false
			}
		}

	}
	return allRulesAreValid
}

func intOf(char string) int {
	i, err := strconv.Atoi(char)
	if err != nil {
		return 0
	}
	return i
}
