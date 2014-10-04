package main

import (
	"strconv"
	"testing"
)

func verify(rule string, koan string, t *testing.T) {
	zen := DoesKoanFulfillRule(Rule{[]string{rule}}, koan)
	if !zen {
		t.Error(koan + " should fulfill rule " + rule)
	}
}

func falsify(rule string, koan string, t *testing.T) {
	zen := DoesKoanFulfillRule(Rule{[]string{rule}}, koan)
	if zen {
		t.Error(koan + " should NOT fulfill rule " + rule)
	}
}

func verifyMultiRule(rules []string, koan string, t *testing.T) {
	multiRule(true, rules, koan, t)
}

func falsifyPartOfMultiRule(rules []string, koan string, t *testing.T) {
	multiRule(false, rules, koan, t)
}

func multiRule(shouldPass bool, rules []string, koan string, t *testing.T) {
	rule := Rule{rules}
	zen := DoesKoanFulfillRule(rule, koan)
	if zen != shouldPass {
		t.Error((koan + " should be " + strconv.FormatBool(shouldPass) + " for rule " + stringRule(rule)))
	}
}

func stringRule(rule Rule) string {
	all := ""
	for _, d := range rule.ruleDescriptions {
		all += "," + d
	}
	return all
}

func verifyValidRule(rule string, t *testing.T) {
	if !SyntacticallyValidRule(Rule{[]string{rule}}) {
		t.Error(rule + " is NOT a valid rule but should be")
	}
}

func verifyValidMultirule(rule Rule, t *testing.T) {
	zen := SyntacticallyValidRule(rule)
	if !zen {
		t.Error("should be a valid rule but is not")
	}
}

func verifyThatRuleIsInvalid(rule string, t *testing.T) {
	if SyntacticallyValidRule(Rule{[]string{rule}}) {
		t.Error(rule + " IS a valid rule but should not be")
	}
}

func verifyValidKoan(koan string, t *testing.T) {
	if !SyntacticallyValidKoan(koan) {
		t.Error(koan + " is not valid but should be")
	}
}

func falsifyInvalidKoan(koan string, t *testing.T) {
	if SyntacticallyValidKoan(koan) {
		t.Error(koan + " is valid but should NOT be")
	}
}

// for testing
func RuleToString(rule Rule) string {
	return unparseRuleDescriptions(rule.ruleDescriptions)
}
