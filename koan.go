package main

import (
	"strconv"
)

type Koan struct {
	description  string
	fulfillsRule bool
}

var Koans []Koan

func KoanSummaries() string {
	summary := "Koans:\n"
	for koanNum := range Koans {
		koan := Koans[koanNum]
		summary += koan.description + " : " + strconv.FormatBool(koan.fulfillsRule) + "\n"
	}
	return summary
}

func AddKoan(newKoan string) {
	Koans = append(Koans, Koan{newKoan, DoesKoanFulfillRule(CurrentRule, newKoan)})
}

func AddFullKoan(koan Koan) {

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
