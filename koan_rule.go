package main

import (
	"strconv"
	"strings"
)

var NONE = "none"
var ALL = 100

func DoesKoanFulfillRule(rule Rule, koan string) bool {
	if !SyntacticallyValidRule(rule) {
		return false
	}

	if !SyntacticallyValidKoan(koan) {
		return false
	}
	koanChunks := chunk(koan)
	if koanFailsPipRule(rule, koanChunks) {
		return false
	}

	if !allColorRulesAreValid(rule, koan) {
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
			if koanFailsSizeRule(koanChunk, ruleChunk) {
				return false
			}
			if koanFailsOrientationRule(koanChunk, ruleChunk) {
				return false
			}
		}
	}
	return true
}

func koanFailsPipRule(rule Rule, koanChunks []string) bool {
	for _, ruleChunk := range rule.ruleDescriptions {
		pipRuleExists, rulePips := rulePips(ruleChunk)

		if pipRuleExists {
			koanPips := 0
			for _, koanChunk := range koanChunks {
				koanPips += pipsFor(sizeOf(koanChunk)) * koanCount(koanChunk)
			}
			if !isNegativeRule(ruleChunk) {
				return rulePips > koanPips || (rulePips == 0 && koanPips != 0)
			}
		}
	}
	return false
}

func rulePips(ruleChunk string) (bool, int) {
	i := strings.Index(ruleChunk, "pip")
	if i > -1 {
		return true, intOf(pieces(ruleChunk)[i+4]) // pip(<digit we want>) <- 4
	}
	return false, 0
}

func pipsFor(size string) int {
	switch size {
	case "S":
		return 1
	case "M":
		return 2
	case "L":
		return 3
	default:
		println("Getting to default pips should be impossible because of the syntax check for rule")
		return 0
	}
}

func koanFailsOrientationRule(koanChunk string, ruleChunk string) bool {
	ruleOrientation := orientation(ruleChunk) // duplication - can use lambdas?
	koanOrientation := orientation(koanChunk)
	if !isNegativeRule(ruleChunk) {
		return ruleOrientation != NONE && ruleOrientation != koanOrientation
	}
	//  Strange bug here, negative needs different aggregation
	return false
}

func orientation(chunk string) string {
	pieces := pieces(chunk)
	for _, piece := range pieces {
		for _, orientation := range validOrientations {
			if piece == orientation {
				return orientation
			}
		}
	}
	return NONE
}

func koanFailsSizeRule(koanChunk string, ruleChunk string) bool {
	ruleSize := sizeOf(ruleChunk)
	koanSize := sizeOf(koanChunk)
	if !isNegativeRule(ruleChunk) {
		return ruleSize != NONE && ruleSize != koanSize
	}
	return ruleSize == koanSize
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
	count := 0
	for i, piece := range pieces(chunk) {
		if piece == ruleColor {
			if i == 0 { // There is no count for the color
				return i
			}
			count = koanCount(chunk)
		}
	}
	return count
}

func handleAllColorRule(koanChunks []string, colorRuleCount int) int {
	if colorRuleCount == 0 {
		return len(koanChunks) + ALL
	}
	return colorRuleCount
}

func allColorRulesAreValid(rule Rule, koan string) bool {
	colorRules := colorRules(rule)
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
		colorsOfKoanChunks[colorOf(koanChunk)]++
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
	case "orientation":
		return orientation(ruleChunk) != NONE
	case "pip":
		return strings.Index(ruleChunk, "pip") != -1
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
			colorCount++
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
		println("Something went wrong with finding count of first digit in ", koanChunk, " ...Returning count:", count)
	}
	return count
}

func rulePieceCount(ruleChunk string) int {
	rulePieces := pieces(ruleChunk)
	if isNegativeRule(ruleChunk) {
		return intOf(rulePieces[1])
	}
	return intOf(rulePieces[0])
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

func intOf(char string) int {
	i, err := strconv.Atoi(char)
	if err != nil {
		// println("Help, tried to get int from", char)
		return 0
	}
	return i
}
