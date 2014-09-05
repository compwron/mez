package main

import (
	"testing"
)

func TestTrueSimpleRuleValidity(t *testing.T) {
	rule := "1G"
	verifyValidRule(rule, t)
}

func TestEmptyRuleNonValidity(t *testing.T) {
	rule := ""
	falsifyValidRule(rule, t)
}

func TestNonsenseRuleNonValidity(t *testing.T) {
	rule := "FOO"
	falsifyValidRule(rule, t)
}

// test fancy complicated rule validity also
