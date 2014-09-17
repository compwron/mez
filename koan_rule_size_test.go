package main

import (
	"testing"
)

func TestSizeMustMatchSizeRuleValid(t *testing.T) {
	rule := "S"
	koan := "1^SG"
	verify(rule, koan, t)
}
