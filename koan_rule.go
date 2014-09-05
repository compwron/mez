package main

import (
	// "fmt"
	"strconv"
	"strings"
)

var validColors = [3]string{"B", "G", "R"} // blue green red
var ruleTypes = [2]string{"count", "color"}

func validRule(rule Rule) bool {
	// rule must be a color or count (or both) rule (more rules coming soon)
		for _, ruleChunk := range rule.ruleDescriptions {
			hasValidRuleType := false
			for _, ruleType := range ruleTypes {
				if ruleContains(strings.Split(ruleChunk, ""), ruleType){
					hasValidRuleType = true
				}
			}

			// All rule chunks must have a valid rule type
			if hasValidRuleType != true {
				return false
			}
		}
	return true
}

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	if !validRule(rule) {
		return false
	}

	allRulesAreValid := true
	for _, description := range rule.ruleDescriptions {
		koanChunks := strings.Split(koan, ",")
		for _, koanChunk := range koanChunks {
			koanPieces := strings.Split(koanChunk, "")
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
	}
	return allRulesAreValid
}

func ruleContains(rulePieces []string, ruleType string) bool {
	// TODO fix duplication
	switch ruleType {
	case "count":
		for i, rulePiece := range rulePieces {
			if intOf(rulePiece) != 0 && !nextPieceIsAColor(rulePieces, i) {
				return true
			}
		}
		return false
	case "color":
		for i, rulePiece := range rulePieces {
			if intOf(rulePiece) != 0 && nextPieceIsAColor(rulePieces, i) {
				return true
			}
		}
		if colorOf(rulePieces) != "none" {
			// for rule "G" instead of "1G"
			return true
		}
		return false
	default:
		return false
	}
}

func nextPieceIsAColor(rulePieces []string, currentIndex int) bool {
	// TODO make this shorter but still clear

	colorOfNextPiece := colorOf(strings.Split(rulePieces[currentIndex+1], ""))
	if colorOfNextPiece != "none" {
		return true
	}
	return false
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
	ruleColor := colorOf(rulePieces)
	if ruleColor == "none" {
		return allRulesAreValid
	}

	koanPieceColor := colorOf(koanPieces)

	colorsMatch := koanPieceColor == ruleColor

	if isNegativeColorRule(rulePieces) {
		return !colorsMatch
	} else if !colorsMatch {
		return false
	}
	return allRulesAreValid
}

func isNegativeColorRule(rulePieces []string) bool {
	return rulePieces[0] == "!" 
}

func intOf(char string) int {
	i, err := strconv.Atoi(char)
	if err != nil {
		return 0
	}
	return i
}
