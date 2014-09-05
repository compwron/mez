package main

import (
	"strconv"
	"strings"
)

var validColors = [3]string{"B", "G", "R"}

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	allRulesAreValid := true
	for i := 0; i < len(rule.ruleDescriptions); i++ {
		koanPieces := strings.Split(koan, "")
		invalidKoan, koanPieceCount := koanCountValidity(koanPieces)
		if invalidKoan {
			return false
		}

		rulePieces := strings.Split(rule.ruleDescriptions[i], "")
		ruleNot, rulePieceCount := analyzeSingleRule(rulePieces)

		// We could get performance gains by only running rules until something comes back false, but wait until optimization is needed.
		allRulesAreValid = evaluatePiecesCountTypeRules(allRulesAreValid, koanPieceCount, rulePieceCount, ruleNot)
		allRulesAreValid = evaluatePiecesColorTypeRules(allRulesAreValid, rulePieces, koanPieces)

	}
	return allRulesAreValid
}

func analyzeSingleRule(ruleCharacters []string) (bool, int) {
	ruleNot := false
	if ruleCharacters[0] == "!" {
		ruleNot = true
		ruleCharacters = ruleCharacters[1:]
	}

	rulePieceCount := intOf(ruleCharacters[0])
	return ruleNot, rulePieceCount
}

func koanCountValidity(koanCharacters []string) (bool, int) {
	invalidKoan := false
	if koanCharacters[0] == "!" {
		invalidKoan = true
	}

	koanPieceCount := intOf(koanCharacters[0])

	return invalidKoan, koanPieceCount
}

func evaluatePiecesCountTypeRules(allRulesAreValid bool, koanPieceCount int, rulePieceCount int, ruleNot bool) bool {
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
