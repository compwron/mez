package main

import (
	// "fmt"
	"testing"
)

func verify(rule string, koan string, t *testing.T) {
	zen := doesKoanFulfillRule(Rule{[]string{rule}}, koan)
	if !zen {
		t.Errorf(koan + " should fulfill rule " + rule)
	}
}

func TestSimplestKoanWithSimplestRule(t *testing.T) {
	rule := "1^" 
	koan := "1^SG"
	verify(rule, koan, t)
}

func TestTwoPiecesShouldMatchAtLeastOnePieceRule(t *testing.T) {
	rule := "1^" 
	koan := "2^SG"
	verify(rule, koan, t)
}

func TestOnePieceShouldNotMatchTwoPieceRule(t *testing.T) {
 	rule := "2^" 
	koan := "1^SG"
	verify(rule, koan, t)
}
