package main

import (
	"strings"
	"testing"
)

func TestChecksForTrueAndFalseKoan(t *testing.T) {
	emptyJSONReader := strings.NewReader("{\"rule\":\"foo\"}")
	Parsed, _ := Parse(emptyJSONReader)

	result := CreateGame(Parsed)
	if result != "need true koan and false koan" {
		t.Error("Should detect missing true and false koans")
	}
}

func TestFailsToStartGameIfGameIsInProgress(t *testing.T) {
	CurrentRule = Rule{[]string{"some rule"}}
	emptyJSONReader := strings.NewReader("{\"rule\":\"foo\"}")
	Parsed, _ := Parse(emptyJSONReader)

	result := CreateGame(Parsed)
	if result != "Can't create game because game is already in progress" {
		t.Error("Should not try to create game if game exists")
	}
	CurrentRule = OriginalRule // test cleanup
}
