package main

import (
	"testing"
)

func TestOneOrMoreButNotThreePieces(t *testing.T) {
	rules := []string{"!3^", "1^"}
	koan := "1^SG"
	verifyMultiRule(rules, koan, t)
}

func TestThreeDoesNotFulfillNonThreeMultiRule(t *testing.T) {
	rules := []string{"1^", "!3^"}
	koan := "3^SG"
	excludeMultiRule(rules, koan, t)
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

func verify(rule string, koan string, t *testing.T) {
	zen := DoesKoanFulfillRule(Rule{[]string{rule}}, koan)
	if !zen {
		t.Errorf(koan + " should fulfill rule " + rule)
	}
}

func falsify(rule string, koan string, t *testing.T) {
	zen := DoesKoanFulfillRule(Rule{[]string{rule}}, koan)
	if zen {
		t.Errorf(koan + " should NOT fulfill rule " + rule)
	}
}

func verifyMultiRule(rules []string, koan string, t *testing.T) {
	multiRule(true, rules, koan, t)
}

func excludeMultiRule(rules []string, koan string, t *testing.T) {
	multiRule(false, rules, koan, t)
}

func multiRule(shouldPass bool, rules []string, koan string, t *testing.T) {
	zen := DoesKoanFulfillRule(Rule{rules}, koan)
	if zen != shouldPass {
		t.Errorf(koan + " should not fulfill rule ")
	}
}
