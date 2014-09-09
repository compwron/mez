package main

import (
	"strconv"
	"strings"
)

var validColors = [3]string{"B", "G", "R"} // blue green red

var ruleTypes = [2]string{"count", "color"} // more coming soon

var NONE = "none"

var ALL = 100

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	if !validRule(rule) {
		return false
	}

	allRulesAreValid := true
	koanChunks := chunk(koan)

	hasColorRules, colorRules := multipleColorRules(rule)
	if !allColorRulesAreValid(colorRules, koan) {
		return false
	}

	for _, ruleChunk := range rule.ruleDescriptions {

		isCountRule := ruleContains(ruleChunk, "count")
		isColorRule := ruleContains(ruleChunk, "color")

		// if rule is a negative COUNT of color rule, must evaluate against all koans
		rulePieces := strings.Split(ruleChunk, "")
		if isColorRule && isNegativeRule(rulePieces) {
			koanHasDisallowedNumberOfColor := koanHasDisallowedNumberOf(colorOf(ruleChunk), diallowedColorCount(rulePieces), koan)
			if koanHasDisallowedNumberOfColor {
				return false
			}
			// else continue to checking other things
		}

		for _, koanChunk := range koanChunks {
			koanPieces := strings.Split(koanChunk, "")
			koanPieceCount, err := koanCount(koanPieces)
			if err != nil {
				return false
			}

			if isCountRule {
				isNegativeRule, rulePieceCount := countInRulePiece(rulePieces)
				allRulesAreValid = evaluatePiecesCountTypeRules(allRulesAreValid, koanPieceCount, rulePieceCount, isNegativeRule)
			}

			if isColorRule && !hasColorRules {
				allRulesAreValid = evaluatePiecesColorTypeRules(allRulesAreValid, ruleChunk, koanChunk)
			}
		}
	}
	return allRulesAreValid
}

func multipleColorRules(rule Rule) (bool, []string) {
	var colorRules []string
	for _, ruleChunk := range rule.ruleDescriptions {
		if colorOf(ruleChunk) != NONE && !isNegativeRule(strings.Split(ruleChunk, "")) {
			colorRules = append(colorRules, ruleChunk)
		}
	}
	return len(colorRules) != 0, colorRules
}

func countOfColor(chunk string, ruleColor string) int {
	pieces := strings.Split(chunk, "")
	for _, piece := range pieces {
		if piece == ruleColor {
			koanCount, _ := koanCount(pieces)
			return koanCount
		}
	}
	return 0
}

func handleAllColorRule(koanChunks []string, colorRuleCount int) int {
	if colorRuleCount == 0 {
		return len(koanChunks) + ALL
	}
	return colorRuleCount
}

func chunk(thingWithComma string) []string {
	return strings.Split(thingWithComma, ",")
}

func allColorRulesAreValid(colorRules []string, koan string) bool {
	allColorRulesFulfilled := true
	koanChunks := chunk(koan)

	for _, ruleChunk := range colorRules {
		ruleColor := colorOf(ruleChunk)
		ruleColorCountInKoanChunks := 0

		for _, koanChunk := range koanChunks {
			if colorOf(koanChunk) == ruleColor {
				ruleColorCountInKoanChunks += countOfColor(koanChunk, ruleColor)
			}
		}

		ruleColorCount := handleAllColorRule(koanChunks, countOfColor(ruleChunk, ruleColor))

		if ruleColorCountInKoanChunks == len(koanChunks) { // all are the color of the rule
			return true
		}
		if !(ruleColorCountInKoanChunks >= ruleColorCount) {
			return false
		}

	}
	return allColorRulesFulfilled
}

func validRule(rule Rule) bool {
	for _, ruleChunk := range rule.ruleDescriptions {
		hasValidRuleType := false
		for _, ruleType := range ruleTypes {
			if ruleContains(ruleChunk, ruleType) {
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

func isNegativeRule(rulePieces []string) bool {
	return rulePieces[0] == "!"
}

func initializeEmptyColorCount() map[string]int {
	colorsOfKoanChunks := map[string]int{}
	for _, validColor := range validColors {
		colorsOfKoanChunks[validColor] = 0
	}

	return colorsOfKoanChunks
}

func koanHasDisallowedNumberOf(ruleColor string, diallowedColorCount int, koan string) bool {
	koanChunks := chunk(koan)
	colorsOfKoanChunks := initializeEmptyColorCount()

	for _, koanChunk := range koanChunks {
		koanColor := colorOf(koanChunk)
		colorsOfKoanChunks[koanColor] += 1
	}
	return colorsOfKoanChunks[ruleColor] == diallowedColorCount
}

func diallowedColorCount(rulePieces []string) int {
	for i, rulePiece := range rulePieces {
		if isValidColor(rulePiece) {
			numberOfColor := intOf(rulePieces[i-1])
			if numberOfColor != 0 {
				return numberOfColor
			}
		}
	}
	return 1 // "!G" is the same of "!1G"
}

func isValidColor(c string) bool {
	for _, color := range validColors {
		if color == c {
			return true
		}
	}
	return false
}

func ruleContains(ruleChunk string, ruleType string) bool {
	rulePieces := strings.Split(ruleChunk, "")
	switch ruleType {
	case "count":
		for i, rulePiece := range rulePieces {
			if intOf(rulePiece) != 0 && !nextPieceIsAColor(ruleChunk, i) {
				return true
			}
		}
		return false
	case "color":
		for i, rulePiece := range rulePieces {
			if intOf(rulePiece) != 0 && nextPieceIsAColor(ruleChunk, i) {
				return true
			}
		}
		if colorOf(ruleChunk) != NONE {
			// for rule "G" instead of "1G"
			return true
		}
		return false
	default:
		return false
	}
}

func nextPieceIsAColor(ruleChunk string, currentIndex int) bool {
	rulePieces := strings.Split(ruleChunk, "")
	colorOfNextPiece := colorOf(rulePieces[currentIndex+1])
	if colorOfNextPiece != NONE {
		return true
	}
	return false
}

func countInRulePiece(rulePieces []string) (bool, int) {
	if rulePieces[0] == "!" {
		return true, intOf(rulePieces[1])
	}
	return false, intOf(rulePieces[0])
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

func colorOf(chunk string) string {
	pieces := strings.Split(chunk, "")
	for _, piece := range pieces {
		for _, color := range validColors {
			if piece == color {
				return color
			}
		}
	}
	return NONE
}

func evaluatePiecesColorTypeRules(allRulesAreValid bool, ruleChunk string, koanChunk string) bool {
	ruleColor := colorOf(ruleChunk)
	if ruleColor == NONE {
		return allRulesAreValid
	}

	koanPieceColor := colorOf(koanChunk)

	colorsMatch := koanPieceColor == ruleColor

	rulePieces := strings.Split(ruleChunk, "")
	if !colorsMatch && !isNegativeRule(rulePieces) { // because negative rules have been previously handled
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
