package main

import (
	"testing"
)

func TestUprightKoanWithoutCountPass(t *testing.T) {
	rule := "^"
	koan := "1^SG"
	verify(rule, koan, t)
}

func TestUprightKoanPass(t *testing.T) {
	rule := "1^"
	koan := "1^SG"
	verify(rule, koan, t)
}

func TestUprightKoanFail(t *testing.T) {
	rule := "^"
	koan := "1>SG"
	falsify(rule, koan, t)

	rule = "1^"
	koan = "1>SG"
	falsify(rule, koan, t)
}
