package main

import (
	"strconv"
	"strings"
)

var ValidColors = [3]string{"B", "G", "R"}
// var ValidSize []string = ["S", "M", "L"]

func remove(data []string, item string) []string {
	newData := make([]string, len(data)-1)
	for i := 0; i < len(data); i++ {
		if data[i] != item {
			newData = append(newData, data[i])
		}
	}
	return newData
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

func splitKoan(koan string) (bool, int) {
	invalidKoan := false
	koanCharacters := strings.Split(koan, "")
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

func containsColor(rulePieces []string) bool {
	for piece := range rulePieces {
	    for color := range ValidColors {
	        if piece == color {
	            return true
	        }
	    }
	}
	return false
}

func evaluatePiecesColorTypeRules(allRulesAreValid bool, rulePieces []string) bool {
	if containsColor(rulePieces) {
		// eval rule and return
	}
	return allRulesAreValid
}

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	allRulesAreValid := true
	for i := 0; i < len(rule.ruleDescriptions); i++ {
		invalidKoan, koanPieceCount := splitKoan(koan)
		if invalidKoan {
			return false
		}
		rulePieces := strings.Split(rule.ruleDescriptions[i], "")
		ruleNot, rulePieceCount := analyzeSingleRule(rulePieces)

		// We could get performance gains by only running rules until something comes back false, but wait until optimization is needed.
		allRulesAreValid = evaluatePiecesCountTypeRules(allRulesAreValid, koanPieceCount, rulePieceCount, ruleNot)
		allRulesAreValid = evaluatePiecesColorTypeRules(allRulesAreValid, rulePieces)

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
