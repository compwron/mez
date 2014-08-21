package main

import (
	"fmt"
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

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	// if first character is !, set "not" and remove it
	ruleNot := false
	ruleCharacters := strings.Split(rule.ruleDescriptions[0], "")
	if ruleCharacters[0] == "!" {
		ruleNot = true
		ruleCharacters = ruleCharacters[1:]
	}

	rulePieceCount := 0

	if ruleCharacters[0] == "" {
		rulePieceCount = intOf(ruleCharacters[1])
	} else {
		rulePieceCount = intOf(ruleCharacters[0])
	}

	koanCharacters := strings.Split(koan, "")
	// koan is not allowed to contain !
	if koanCharacters[0] == "!" {
		fmt.Println("Koans are not allowed to have !, returning false")
		return false
	}

	koanPieceCount := intOf(koanCharacters[0])

	// if rule is a not, check that koanCount is anything other than ruleCount
	if ruleNot {
		return koanPieceCount != rulePieceCount
	} else {
		return koanPieceCount >= rulePieceCount
	}

	return false
}

func intOf(char string) int {
	i, err := strconv.Atoi(char)
	if err != nil {
		return 0
	}
	return i
}
