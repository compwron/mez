package main

import (
	"testing"
)

// - DONE Number of pieces excluded in rule is avoided [Example: "!3"] // There can be 1, 2, 4, etc pieces
func TestCombinedPiecesCountIsForbiddenPass(t *testing.T) {}
func TestCombinedPiecesCountIsForbiddenFail(t *testing.T) {}

func TestOneOrMoreButNotThreePieces(t *testing.T) {
	rules := []string{"!3^", "1^"}
	koan := "1^SG"
	verifyMultiRule(rules, koan, t)
}

func TestNegativeOfThreePiecesPassesWithOnePiece(t *testing.T) {
	rules := []string{"!3^"}
	koan := "1^SG"
	verifyMultiRule(rules, koan, t)
}

func TestNegativeOfThreePiecesFailsWithThreePieces(t *testing.T) {
	rule := "!3^"
	koan := "3^SG"
	falsify(rule, koan, t)
}

func TestNegativeKoanFailsBecauseKoansCannotBeNegative(t *testing.T) {
	rule := "1^"
	koan := "!1^SG"
	falsify(rule, koan, t)
}

func TestMultiKoanFulfillsRuleForMustHaveSingle(t *testing.T) {
	rule := "1^"
	koan := "3^SG"
	verify(rule, koan, t)
}

func TestThreeDoesNotFulfillNonThreeMultiRule(t *testing.T) {
	rules := []string{"1^", "!3^"}
	koan := "3^SG"
	falsifyPartOfMultiRule(rules, koan, t)
}

func TestOnePieceShouldNotMatchTwoPieceRule(t *testing.T) {
	rule := "2^"
	koan := "1^SG"
	falsify(rule, koan, t)
}

func TestTwoPiecesShouldMatchAtLeastOnePieceRule(t *testing.T) {
	rule := "1^"
	koan := "2^SG"
	verify(rule, koan, t)
}

func TestSimplestKoanWithSimplestRule(t *testing.T) {
	rule := "1^"
	koan := "1^SG"
	verify(rule, koan, t)
}
