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
		t.Errorf("Should detect missing true and false koans")
	}
}

func TestFailsToStartGameIfGameIsInProgress(t *testing.T) {
	CurrentRule = Rule{[]string{"some rule"}}
	emptyJsonReader := strings.NewReader("{\"rule\":\"foo\"}")
	parsed, _ := Parse(emptyJsonReader)

	result := CreateGame(parsed)
	if result != "Can't create game because game is already in progress" {
		t.Errorf("Should not try to create game if game exists")
	}
	CurrentRule = OriginalRule // test cleanup
}
