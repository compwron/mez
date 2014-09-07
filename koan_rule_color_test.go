package main

import (
	"testing"
)

func TestNoPiecesMayBeGreen(t *testing.T) {
	// TODO
}

func TestTwoColorsInTwoRuleChunksWithOneColorEach(t *testing.T) {
	rules := []string{"1G", "1R"}
	koan := "1^SG,1^SR"
	verifyMultiRule(rules, koan, t)
}

func TestTwoColorsInTwoRuleChunksWithOneColorRequiringTwoKoanPiecesOfItsColor(t *testing.T) {
	rules := []string{"2G", "1R"}
	koan := "2^SG,1^SR"
	verifyMultiRule(rules, koan, t)
}

func TestTwoPiecesMustBeGreenWithMultiPieceKoanBackwardsOrderFail(t *testing.T) {
	rule := "2G"
	koan := "1^SR,1^SG"
	falsify(rule, koan, t)
}

func TestTwoPiecesMustBeGreenWithMultiPieceKoanBackwardsOrderPass(t *testing.T) {
	rule := "2G"
	koan := "1^SG,1^SG"
	verify(rule, koan, t)
}

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

func TestAllPiecesMustBeGreenFail(t *testing.T) {
	rule := "G"
	koan := "1^SB"
	falsify(rule, koan, t)
}

func TestAllPiecesInMultiKoanMustBeGreenFail(t *testing.T) {
	rule := "G"
	koan := "1^SB,1^SG"
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
