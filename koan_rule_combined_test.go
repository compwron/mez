package main

import (
	"testing"
)

func TestCombinedColorCountInvalid(t *testing.T) {
	rule := "4^2R"
	koan := "1^R"
	falsify(rule, koan, t)
}
