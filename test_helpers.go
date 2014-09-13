package main

// https://github.com/benbjohnson/testing
import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"testing"
)

func verify(rule string, koan string, t *testing.T) {
	zen := DoesKoanFulfillRule(Rule{[]string{rule}}, koan)
	if !zen {
		t.Errorf(koan + " should fulfill rule " + rule)
	}
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
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

func falsifyPartOfMultiRule(rules []string, koan string, t *testing.T) {
	multiRule(false, rules, koan, t)
}

func multiRule(shouldPass bool, rules []string, koan string, t *testing.T) {
	rule := Rule{rules}
	zen := DoesKoanFulfillRule(rule, koan)
	if zen != shouldPass {
		t.Errorf((koan + " should be " + strconv.FormatBool(shouldPass) + " for rule " + stringRule(rule)))
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
		t.Errorf(rule + " is NOT a valid rule but should be")
	}
}

func verifyValidMultirule(rule Rule, t *testing.T) {
	zen := SyntacticallyValidRule(rule)
	if !zen {
		t.Errorf("should be a valid rule but is not")
	}
}

func verifyThatRuleIsInvalid(rule string, t *testing.T) {
	if SyntacticallyValidRule(Rule{[]string{rule}}) {
		t.Errorf(rule + " IS a valid rule but should not be")
	}
}

func verifyValidKoan(koan string, t *testing.T) {
	if !SyntacticallyValidKoan(koan) {
		t.Errorf(koan + " is not valid but should be")
	}
}

func falsifyInvalidKoan(koan string, t *testing.T) {
	if SyntacticallyValidKoan(koan) {
		t.Errorf(koan + " is valid but should NOT be")
	}
}
