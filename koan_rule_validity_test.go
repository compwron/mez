package main

import (
	"testing"
)

func RuleWithTwoColorsInTwoChunksIsValid(t *testing.T) {
	multiColorRule := Rule{[]string{"G", "R"}}
	verifyValidMultirule(multiColorRule, t)
}

func TestRuleWithTwoColorsInOneChunkIsInvalid(t *testing.T) {
	rule := "GR"
	verifyThatRuleIsInvalid(rule, t)
}

func TestTrueSimpleRuleValidity(t *testing.T) {
	rule := "1G"
	verifyValidRule(rule, t)
}

func TestEmptyRuleNonValidity(t *testing.T) {
	rule := ""
	verifyThatRuleIsInvalid(rule, t)
}

func TestNonsenseRuleNonValidity(t *testing.T) {
	rule := "FOO"
	verifyThatRuleIsInvalid(rule, t)
}

func TestLeadingReverseCharactersInRulesOk(t *testing.T) {
	rule := "!1^"
	verifyValidRule(rule, t)
}

func TestNoNonLeadingReverseCharactersInRules(t *testing.T) {
	rule := "1^!"
	verifyThatRuleIsInvalid(rule, t)
}
