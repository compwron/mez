package main

import (
	// "fmt"
	"testing"
)

func TestSimplestKoanWithSimplestRule(t *testing.T) {
	oneUprightRule := Rule{[]string{"1^"}}
	oneUprightSmallGreen := "1^SG"

	zen := doesKoanFulfillRule(oneUprightRule, oneUprightSmallGreen)
	if !zen {
		t.Errorf("Koan should fulfill rule")
		// TODO get better to_s for Rule and Koan
	}
}
