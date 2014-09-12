package main

import (
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
