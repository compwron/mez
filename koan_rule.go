package main

import (
	"strconv"
)

var ruleTypes = [3]string{"count", "color", "size"} // more coming soon
var NONE = "none"
var ALL = 100

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	if !SyntacticallyValidRule(rule) {
		return false
	}

	koanChunks := chunk(koan)

	if !allColorRulesAreValid(colorRules(rule), koan) {
		return false
	}

	for _, ruleChunk := range rule.ruleDescriptions {
		if koanHasNumberOfColorDisallowedByNegativeColorRule(ruleChunk, koan) {
			return false
		}
		for _, koanChunk := range koanChunks {
			if ruleContains(ruleChunk, "count") && !koanPassesCountRule(koanChunk, ruleChunk) {
				return false
			}
			if ruleContains(ruleChunk, "color") && isNegativeRule(ruleChunk) && !koanPassesColorRule(ruleChunk, koanChunk) {
				return false
			}
			if koanPassesSizeRule(koanChunk, ruleChunk) {
				return true
			}
		}
	}
	return true
}

func koanPassesSizeRule(koanChunk string, ruleChunk string) bool {
	size := sizeOf(ruleChunk)
	return size != NONE && size == sizeOf(koanChunk)
}

func colorRules(rule Rule) []string {
	var colorRules []string
	for _, ruleChunk := range rule.ruleDescriptions {
		if colorOf(ruleChunk) != NONE && !isNegativeRule(ruleChunk) {
			colorRules = append(colorRules, ruleChunk)
		}
	}
	return colorRules
}

func koanHasNumberOfColorDisallowedByNegativeColorRule(ruleChunk string, koan string) bool {
	if ruleContains(ruleChunk, "color") && isNegativeRule(ruleChunk) {
		koanHasDisallowedNumberOfColor := koanHasDisallowedNumberOf(colorOf(ruleChunk), diallowedColorCount(ruleChunk), koan)
		if koanHasDisallowedNumberOfColor {
			return true
		}
	}
	return false
}

func countOfColor(chunk string, ruleColor string) int {
	for _, piece := range pieces(chunk) {
		if piece == ruleColor {
			return koanCount(chunk)
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

func allColorRulesAreValid(colorRules []string, koan string) bool {
	koanChunks := chunk(koan)

	for _, ruleChunk := range colorRules {
		ruleColor := colorOf(ruleChunk)
		ruleColorCountInKoanChunks := 0

		for _, koanChunk := range koanChunks {
			if colorOf(koanChunk) == ruleColor {
				ruleColorCountInKoanChunks += countOfColor(koanChunk, ruleColor)
			}
		}

		if ruleColorCountInKoanChunks == len(koanChunks) { // all are the color of the rule
			return true
		}
		if !(ruleColorCountInKoanChunks >= handleAllColorRule(koanChunks, countOfColor(ruleChunk, ruleColor))) {
			return false
		}

	}
	return true
}

func SyntacticallyValidRule(rule Rule) bool {
	for _, ruleChunk := range rule.ruleDescriptions {
		hasValidRuleType := false
		for _, ruleType := range ruleTypes {
			if ruleContains(ruleChunk, ruleType) {
				hasValidRuleType = true
			}
		}

		if !hasValidRuleType {
			return false
		}
	}
	return true
}

func isNegativeRule(ruleChunk string) bool {
	return pieces(ruleChunk)[0] == "!"
}

func initializeEmptyColorCount() map[string]int {
	colorsOfKoanChunks := map[string]int{}
	for _, validColor := range validColors {
		colorsOfKoanChunks[validColor] = 0
	}

	return colorsOfKoanChunks
}

func koanHasDisallowedNumberOf(ruleColor string, diallowedColorCount int, koan string) bool {
	colorsOfKoanChunks := initializeEmptyColorCount()

	for _, koanChunk := range chunk(koan) {
		colorsOfKoanChunks[colorOf(koanChunk)] += 1
	}
	return colorsOfKoanChunks[ruleColor] == diallowedColorCount
}

func diallowedColorCount(ruleChunk string) int {
	rulePieces := pieces(ruleChunk)
	for i, rulePiece := range rulePieces {
		if isValid(validColors, rulePiece) {
			numberOfColor := intOf(rulePieces[i-1])
			if numberOfColor > 0 {
				return numberOfColor
			}
		}
	}
	return 1 // "!G" is the same of "!1G"
}

func ruleContains(ruleChunk string, ruleType string) bool {
	rulePieces := pieces(ruleChunk)
	switch ruleType {
	case "count":
		for i, rulePiece := range rulePieces {
			if rulePiece == "!" && i != 0 { // no non-leading negatives
				return false
			}
		}

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
		if colorOf(ruleChunk) != NONE && !multipleColors(ruleChunk) {
			// for rule "G" instead of "1G"
			return true
		}
		return false
	case "size":
		return sizeOf(ruleChunk) != NONE
	default:
		return false
	}
}

func sizeOf(chunk string) string {
	pieces := pieces(chunk)
	for _, piece := range pieces {
		for _, size := range validSizes {
			if piece == size {
				return size
			}
		}
	}
	return NONE
}

func multipleColors(ruleChunk string) bool {
	colorCount := 0
	for _, piece := range pieces(ruleChunk) {
		if colorOf(piece) != NONE {
			colorCount += 1
		}
	}
	return colorCount > 1
}

func nextPieceIsAColor(ruleChunk string, currentIndex int) bool {
	rulePieces := pieces(ruleChunk)
	colorOfNextPiece := colorOf(rulePieces[currentIndex+1])
	if colorOfNextPiece != NONE {
		return true
	}
	return false
}

func koanCount(koanChunk string) int {
	count, err := strconv.Atoi(pieces(koanChunk)[0])
	if err != nil {
		println("koanCount has been wrongly called with a non-int first character of ", koanChunk, " ...Returning count:", count)
	}
	return count
}

func rulePieceCount(ruleChunk string) int {
	rulePieces := pieces(ruleChunk)
	if isNegativeRule(ruleChunk) {
		return intOf(rulePieces[1])
	} else {
		return intOf(rulePieces[0])
	}
}

func koanPassesCountRule(koanChunk string, ruleChunk string) bool {
	koanCount := koanCount(koanChunk)
	rulePieceCount := rulePieceCount(ruleChunk)

	if isNegativeRule(ruleChunk) {
		return !(koanCount == rulePieceCount)
	}

	return koanCount >= rulePieceCount
}

func colorOf(chunk string) string {
	pieces := pieces(chunk)
	for _, piece := range pieces {
		for _, color := range validColors {
			if piece == color {
				return color
			}
		}
	}
	return NONE
}

func koanPassesColorRule(ruleChunk string, koanChunk string) bool {
	// TODO Why is this?
	// colorsMatch := colorOf(koanChunk) == colorOf(ruleChunk)
	// return colorsMatch || isNegativeRule(ruleChunk)
	return isNegativeRule(ruleChunk)
}

func intOf(char string) int {
	i, err := strconv.Atoi(char)
	if err != nil {
		return 0
	}
	return i
}
