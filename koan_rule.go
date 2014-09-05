package main

import (
	"strconv"
	"strings"
)

var validColors = [3]string{"B", "G", "R"}

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

		// We could get performance gains by only running rules until something comes back false, but wait to optimize until optimization is needed.
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
			allRulesAreValid = false
		}
	} else {
		if !(koanPieceCount >= rulePieceCount) {
			allRulesAreValid = false
		}
	}
	return allRulesAreValid
}

func ruleColor(rulePieces []string) string {
	for _, piece := range rulePieces {
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
	ruleColor := ruleColor(rulePieces)
	if ruleColor != "none" {
		for _, koanPiece := range koanPieces {
			for _, color := range validColors {
				if koanPiece == color {
					if koanPiece != ruleColor {
						return false
					}
					return allRulesAreValid
				}
			}
		}
		return false // koan must have color to be valid
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
