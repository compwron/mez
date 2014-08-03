package main

import (
	"fmt"
	"strconv"
)

var Koans []Koan

func KoanSummaries() string {
	summary := "Koans:\n"
	fmt.Println("koans:", Koans)
	for koanNum := range Koans {
		koan := Koans[koanNum]
		summary += koan.description + " : " + strconv.FormatBool(koan.fulfillsRule) + "\n"
	}
	return summary
}
