package main

import (
	"testing"
)

func TestCombinedColorCountInvalid(t *testing.T) {
	rule := "4^2R"
	koan := "1^R"
	falsify(rule, koan, t)
}

// set a rule (generated), then call false koan
// koanCount tries to count part of the rule ... why?
func TestFalseKoanOnNonDefaultRuleShouldNotCallKoanCount(t *testing.T) {
	t.Errorf("FIX ALL THE COUNTING WRONG THING THINGS")
}
