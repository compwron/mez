package main

import (
	"testing"
)

func RuleWithTwoColorsInTwoChunksIsValid(t *testing.T) {
	multiColorRule := Rule{[]string{"G", "R"}}
	verifyValidMultirule(multiColorRule, t)
}

func RuleWithTwoColorsInOneChunkIsInvalid(t *testing.T) {
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

func LeadingReverseCharactersInRulesOk(t *testing.T) {
	rule := "!1^"
	verifyValidRule(rule, t)
}

func NoNonLeadingReverseCharactersInRules(t *testing.T) {
	rule := "1^!"
	verifyThatRuleIsInvalid(rule, t)
}
