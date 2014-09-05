package main

import (
	"testing"
)

func TestTrueSimpleRuleValidity(t *testing.T) {
	rule := "1G"
	verifyValidRule(rule, t)
}

func TestTrueSimpleRuleNonValidity(t *testing.T) {
	rule := "FOOBARBAZTHING1234!!"
	falsifyValidRule(rule, t)
}

// test fancy complicated rule validity too
