package main

import (
	"reflect"
	"testing"
)

func TestGenerateResetsExistingKoans(t *testing.T) {
	AddKoan("1^G")
	if len(Koans) != 1 {
		t.Errorf("Test setup should contain 1 koan")
	}

	GenerateRule()

	if len(Koans) != 0 {
		t.Errorf("Generate should wipe existing koans.")
	}
}

func TestGenerateDoesNotChangeNonDefaultRule(t *testing.T) {
	CurrentRule = OriginalRule // Test data setup

	GenerateRule()
	if reflect.DeepEqual(OriginalRule, CurrentRule) {
		t.Errorf("Generate should change rule to not be original")
	}
	firstGeneratedRule := CurrentRule

	GenerateRule()
	if !reflect.DeepEqual(firstGeneratedRule, CurrentRule) {
		t.Errorf("Generate should not be able to change change rule from non-original rule")
	}
}