package main

import (
	"strconv"
	"strings"
)

var validColors = [3]string{"B", "G", "R"} // blue green red

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	allRulesAreValid := true
	for _, description := range rule.ruleDescriptions {
		koanPieces := strings.Split(koan, "")
		koanPieceCount, err := koanCount(koanPieces)
		if err != nil {
			return false
		}

		rulePieces := strings.Split(description, "")
		isNegativeRule, rulePieceCount := countInRulePiece(rulePieces)

		// Get performance gains by only running rules until something comes back false
		allRulesAreValid = evaluatePiecesCountTypeRules(allRulesAreValid, koanPieceCount, rulePieceCount, isNegativeRule)
		allRulesAreValid = evaluatePiecesColorTypeRules(allRulesAreValid, rulePieces, koanPieces)
	}
	return allRulesAreValid
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

	// Nesting is bad and I should feel bad.
	ruleColor := colorOf(rulePieces)
	for _, koanPiece := range koanPieces {
		koanPieceColor := colorOf(strings.Split(koanPiece, ""))
		if koanPieceColor != ruleColor {
			return false
		}
		return allRulesAreValid
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
