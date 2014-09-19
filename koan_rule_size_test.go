package main

import (
	"testing"
)

func TestNegativeSizeShouldPassWithDifferentSizes(t *testing.T) {
	rule := "!S"
	koan := "1^MG"
	verify(rule, koan, t)
}

func TestNegativeSizeShouldFailWithSameSizes(t *testing.T) {
	rule := "!S"
	koan := "1^SG"
	falsify(rule, koan, t)
}

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
