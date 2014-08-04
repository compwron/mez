package main

import (
	"fmt"
	"strconv"
	"strings"
)

func remove(data []string, item string) []string {
	newData := make([]string, len(data))
	for i := 0; i < len(data); i++ {
		if data[i] != item {
			newData = append(newData, data[i])
		}
	}
	fmt.Println("new data without removed element:")
	return newData
}

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	// if first character is !, set "not" and remove it
	ruleNot := false
	ruleCharacters := strings.Split(rule.ruleDescriptions[0], "")
	if ruleCharacters[0] == "!" {
		fmt.Println("Setting ruleNot to true")
		ruleNot = true

		remove(ruleCharacters, "!")
		fmt.Println("Removed leading ! from ruleCharacters now that ruleNot is set")
		fmt.Println(ruleCharacters)
	}
	rulePieceCount, ruleErr := intOf(ruleCharacters[0])

	koanCharacters := strings.Split(koan, "")
	// koan is not allowed to contain !
	if koanCharacters[0] == "!" {
		fmt.Println("Koans are not allowed to have !, returning false")
		return false
	}

	koanPieceCount, koanErr := intOf(koanCharacters[0])

	if ruleErr != nil || koanErr != nil {
		return false
	}

	// if rule is a not, check that koanCount is anything other than ruleCount
	if ruleNot {
		return koanPieceCount != rulePieceCount
	} else {
		return koanPieceCount >= rulePieceCount
	}

	return false
}

func intOf(char string) (int64, error) {
	return strconv.ParseInt(char, 10, 8)
}
