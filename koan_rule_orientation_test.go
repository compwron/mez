package main

import (
	"testing"
)

// - DONE it contains two or more upright pieces. [Example: "2^"] # wait does this actually work
func TestTwoOrMoreUprightPieces(t *testing.T) {
	t.Skipf("skipping koan chunk addition for now")

	rule := "2^"
	koan := "2^SG"
	verify(rule, koan, t)

	koan = "1^SG,1^SG"
	verify(rule, koan, t)

	koan = "2^SG,1^SG"
	verify(rule, koan, t)

	koan = "2>SG"
	falsify(rule, koan, t)

	koan = "2>SG,1^SG"
	falsify(rule, koan, t)

	koan = "1^SG"
	falsify(rule, koan, t)
}

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
