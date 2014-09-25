package main

import (
	"strconv"
)

type Koan struct {
	// koanDescription  []string
	description  string
	fulfillsRule bool
}

var Koans []Koan
var validSizes = []string{"S", "M", "L"}        // small medium large
var validColors = []string{"B", "G", "R", "Y"}  // blue green red yellow
var validOrientations = []string{"^", ">", "<"} // upright, sideways right, sideways left

func AddKoanIfValid(newKoanHash map[string]interface{}) string {
	newKoan := newKoanHash["koan"].(string)
	if !SyntacticallyValidKoan(newKoan) {
		return "Invalid koan"
	}
	doesKoanFulfillRule := AddKoan(newKoan)
	return strconv.FormatBool(doesKoanFulfillRule) + "\n"
}

func isValid(list []string, c string) bool {
	for _, item := range list {
		if c == item {
			return true
		}
	}
	return false
}

func SyntacticallyValidKoan(koan string) bool {
	// No negations ("!"") in koans
	// Must have <number> <symbol>
	// Optional number
	// Requires size
	// Must have color
	// That's it

	for _, koanChunk := range chunk(koan) {
		if len(koanChunk) < 4 {
			return false
		}
		for i, koanPiece := range pieces(koanChunk) {
			if koanPiece == "!" {
				return false
			}

			if i == 0 { // count
				_, err := strconv.Atoi(koanPiece)
				if err != nil {
					return false
				}
			}

			if i == 1 { // symbol
				if !isValid(validOrientations, koanPiece) {
					return false
				}
			}

			if i == 2 {
				_, err := strconv.Atoi(koanPiece)
				if !isValid(validSizes, koanPiece) && err != nil {
					return false
				}
			}

			if i == 3 {
				if !isValid(validSizes, koanPiece) && !isValid(validColors, koanPiece) {
					return false
				}
			}

			if i == 4 {
				if !isValid(validColors, koanPiece) {
					return false
				}
			}

			if i >= 5 {
				return false
			}
		}
	}
	return true
}

func KoanSummaries() string {
	summary := "Koans:\n"
	for koanNum := range Koans {
		koan := Koans[koanNum]
		summary += koan.description + " : " + strconv.FormatBool(koan.fulfillsRule) + "\n"
	}
	return summary
}

func AddKoan(newKoan string) bool {
	doesKoanFulfillRule := DoesKoanFulfillRule(CurrentRule, newKoan)
	Koans = append(Koans, Koan{newKoan, doesKoanFulfillRule})
	return doesKoanFulfillRule
}

func AddFullKoan(koan Koan) {
	Koans = append(Koans, koan)
}

func ParseKoan(data map[string]interface{}, truthiness bool) Koan {
	if truthiness {
		return Koan{data["true"].(string), truthiness}
	} else {
		return Koan{data["false"].(string), truthiness}
	}
}
