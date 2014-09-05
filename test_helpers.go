package main

// https://github.com/benbjohnson/testing
import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
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
	zen := DoesKoanFulfillRule(Rule{rules}, koan)
	if zen != shouldPass {
		// TODO rephrase this error for negative case; it is confusing.
		t.Errorf(koan + " should not fulfill rule ")
	}
}
