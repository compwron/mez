package main

import (
	// "fmt"
	"strconv"
	"strings"
)

var validColors = [3]string{"B", "G", "R"} // blue green red

func validRule(rule Rule) bool {
	return true
}

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	if !validRule(rule) {
		return false
	}

	allRulesAreValid := true
	for _, description := range rule.ruleDescriptions {
		koanPieces := strings.Split(koan, "")
		koanPieceCount, err := koanCount(koanPieces)
		if err != nil {
			return false
		}

		rulePieces := strings.Split(description, "")
		isNegativeRule, rulePieceCount := countInRulePiece(rulePieces)

		if ruleContains(rulePieces, "count") {
			allRulesAreValid = evaluatePiecesCountTypeRules(allRulesAreValid, koanPieceCount, rulePieceCount, isNegativeRule)
		}
		if ruleContains(rulePieces, "color") {
			allRulesAreValid = evaluatePiecesColorTypeRules(allRulesAreValid, rulePieces, koanPieces) // need loop per koans
		}
		if allRulesAreValid == false {
			// for performance
			return allRulesAreValid
		}
	}
	return allRulesAreValid
}

func ruleContains(rulePieces []string, ruleType string) bool {
	switch ruleType {
	case "count":
		return true
	case "color":
		return true
	default:
		return false
	}
}

func countInRulePiece(ruleCharacters []string) (bool, int) {
	if ruleCharacters[0] == "!" {
		return true, intOf(ruleCharacters[1])
	}
	return false, intOf(ruleCharacters[0])
}

func koanCount(koanCharacters []string) (int, error) {
	return strconv.Atoi(koanCharacters[0])
}

func evaluatePiecesCountTypeRules(allRulesAreValid bool, koanPieceCount int, rulePieceCount int, isNegativeRule bool) bool {
	// if rule is a not, check that koanCount is anything other than ruleCount
	if isNegativeRule {
		// how do you handle a negative rule without color?
		if koanPieceCount == rulePieceCount {
			return false
		}
	} else {
		if !(koanPieceCount >= rulePieceCount) {
			return false
		}
	}
	return allRulesAreValid
}

func colorOf(pieces []string) string {
	for _, piece := range pieces {
		for _, color := range validColors {
			if piece == color {
				return color
			}
		}
	}
	return "none"
}

func evaluatePiecesColorTypeRules(allRulesAreValid bool, rulePieces []string, koanPieces []string) bool {
	// check for multiple colors in rules&koans? Or do rule validation elsewhere?
	ruleColor := colorOf(rulePieces)
	if ruleColor == "none" {
		return allRulesAreValid
	}

	koanPieceColor := colorOf(koanPieces)

	if koanPieceColor != ruleColor {
		return false
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
