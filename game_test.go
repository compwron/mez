package main

import (
	"strings"
	"testing"
)

func TestChecksForTrueAndFalseKoan(t *testing.T) {
	emptyJsonReader := strings.NewReader("{\"rule\":\"foo\"}")
	parsed, _ := Parse(emptyJsonReader)

	result := CreateGame(parsed)
	if result != "need true koan and false koan" {
		println("RESULT", result)
		t.Errorf("Should detect missing true and false koans")
	}
}
