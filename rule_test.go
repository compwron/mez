package main

import (
	"testing"
)

func TestRuleSummaryIsNotOriginalRule(t *testing.T) {
	CurrentRule = Rule{[]string{"some rule"}}
	s := RuleSummary()
	if s != "current rule is NOT original rule\n" {
		t.Error("Undesired rule summary output: " + s)
	}
}
