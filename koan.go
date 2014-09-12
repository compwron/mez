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

func AddKoanIfValid(newKoanHash map[string]interface{}) string {
	newKoan := newKoanHash["koan"].(string)
	if !KoanIsValid(newKoan) {
		return "Invalid koan"
	}
	doesKoanFulfillRule := AddKoan(newKoan)
	return strconv.FormatBool(doesKoanFulfillRule) + "\n"
}

func KoanIsValid(koan string) bool {
	// no !
	// must have <number> <symbol>
	// optional number
	// Must have color
	// that's it
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

func AddKoanWithOutcome(newKoan string, outcome bool) {
	Koans = append(Koans, Koan{newKoan, outcome})
}

func ParseKoan(data map[string]interface{}, truthiness bool) Koan {
	if truthiness {
		return Koan{data["true"].(string), truthiness}
	} else {
		return Koan{data["false"].(string), truthiness}
	}
}
