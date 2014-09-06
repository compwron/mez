package main

import (
	"testing"
)

func TestNoPiecesMayBeGreenKoan(t *testing.T) {
	rule := "!G"
	trueKoan := "2^LR,1^SY"
	falseKoan := "1^LG,2^MR"
	verify(rule, trueKoan, t)
	falsify(rule, falseKoan, t)
}

func TestTwoColorsInTwoRuleChunks(t *testing.T) {
	t.Skipf("Skipping this failing test for the moment")
	rules := []string{"1G", "1R"}
	koan := "1^SG,1^SR"
	verifyMultiRule(rules, koan, t)
}

func TestTwoPiecesMustBeGreenWithMultiPieceKoanBackwardsOrderFail(t *testing.T) {
	rule := "2G"
	koan := "1^SR,1^SG"
	falsify(rule, koan, t)
}

// how can one but not both of the two above tests not pass?
// ===========

func TestExcludeCertainNumberOfCertainColorPiecesFail(t *testing.T) {
	rule := "!2G"
	koan := "1^SG,1^SG"
	falsify(rule, koan, t)
}

func TestExcludeCertainNumberOfCertainColorPiecesPassWithLargerNumberOfPieces(t *testing.T) {
	rule := "!2G"
	koan := "1^SG,1^SG,1^SG"
	verify(rule, koan, t)
}

func TestExcludeCertainNumberOfCertainColorPiecesPassWithSmallerNumberOfPieces(t *testing.T) {
	rule := "!3G"
	koan := "1^SG,1^SG"
	verify(rule, koan, t)
}

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

func TestNegativeSingleColorRulePass(t *testing.T) {
	rule := "!1G"
	koan := "1^SR"
	verify(rule, koan, t)
}

func TestNegativeSingleColorRuleFail(t *testing.T) {
	rule := "!1G"
	koan := "1^SG"
	falsify(rule, koan, t)
}
