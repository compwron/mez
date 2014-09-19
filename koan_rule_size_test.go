package main

import (
	"testing"
)

func TestSizeMustMatchSizeRuleInvalid(t *testing.T) {
	rule := "S"
	koan := "1^MG"
	falsify(rule, koan, t)
}

func TestSizeMustMatchSizeRuleValid(t *testing.T) {
	rule := "S"
	koan := "1^SG"
	verify(rule, koan, t)
}
