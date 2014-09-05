package main

import (
	"testing"
)

func TestOnePieceMustBeGreenPass(t *testing.T) {
	rule := "G" // At least one piece must be green
	koan := "1^SG"
	verify(rule, koan, t)
}

func TestOnePieceMustBeGreenFail(t *testing.T) {
	rule := "G"
	koan := "1^SB"
	falsify(rule, koan, t)
}

func TestTwoPiecesMustBeGreenWithMultiPieceKoanPass(t *testing.T) {
	rule := "2G"
	koan := "1^SG,1^SG"
	verify(rule, koan, t)
}

func TestTwoPiecesMustBeGreenWithMultiPieceKoanFail(t *testing.T) {
	rule := "2G"
	koan := "1^SG,1^SR"
	falsify(rule, koan, t)
}

func TestNegativeColorRulePass(t *testing.T) {
	rule := "!1G"
	koan := "1^SR"
	verify(rule, koan, t)
}

func TestNegativeColorRuleFail(t *testing.T) {
	rule := "!1G"
	koan := "1^SG"
	falsify(rule, koan, t)
}
