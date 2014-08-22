package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGuessRuleGet(t *testing.T) {
	handler := GuessRule()
	server := httptest.NewServer(handler)
	defer server.Close()

	res, _ := http.Get(server.URL)
	defer res.Body.Close()

	dataBuf, _ := ioutil.ReadAll(res.Body)
	data := string(dataBuf)

	if !strings.Contains(data, "not supported") {
		t.Errorf("Should not support Get for Guess")
	}
}

func TestGuessRulePost(t *testing.T) {
	handler := GuessRule()
	server := httptest.NewServer(handler)
	defer server.Close()

	// Set up some data
	CurrentRule = Rule{strings.Split("1rule,2rule", ",")}
	AddKoan("1 new koan")
	// Confirm that data is present

	// End game by guessing rule
	res, _ := http.Post(server.URL, "text/json", bytes.NewBuffer([]byte("{\"rule\":\"rule1,rule3\"}")))
	defer res.Body.Close()

	// Check that data from previous game is not in new game
	if CurrentRule.ruleDescriptions[0] != OriginalRule.ruleDescriptions[0] {
		t.Errorf("Current rule should be OriginalRule")
	}
	if len(Koans) != 0 {
		t.Errorf("Length of koans should be 0")
	}
}

func TestGameGet(t *testing.T) {
	handler := Game()
	server := httptest.NewServer(handler)
	defer server.Close()

	res, _ := http.Get(server.URL)
	defer res.Body.Close()

	dataBuf, _ := ioutil.ReadAll(res.Body)
	data := string(dataBuf)

	if !strings.Contains(data, "Koans:") {
		t.Errorf("Should have Koans list in game summary")
	}

	if !strings.Contains(data, "current rule is original rule") {
		t.Errorf("Should have original rule in game summary")
	}
}
