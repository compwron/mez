package main

import (
	"testing"
)

// negative size test

func TestSizeMustMatchSizeRuleValid(t *testing.T) {
	rule := "S"
	koan := "1^SG"
	verify(rule, koan, t)
}
