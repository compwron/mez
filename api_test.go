package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateKoanValid(t *testing.T) {
	data := createKoanBody("{\"koan\":\"1>SG\"}")
	if !(data == "true") {
		t.Errorf("should return true. Actually got: " + data)
	}
}

func TestCreateKoanInvalid(t *testing.T) {
	data := createKoanBody("{\"koan\":\"0>SG\"}")
	if !(data == "false") {
		t.Errorf("should return false. Actually got: " + data)
	}
}

func TestGuessRuleGet(t *testing.T) {
	handler := GuessRule()
	server := httptest.NewServer(handler)
	defer server.Close()

	res, _ := http.Get(server.URL)
	defer res.Body.Close()

	data := body(res)

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

	data := body(res)

	if !strings.Contains(data, "Koans:") {
		t.Errorf("Should have Koans list in game summary")
	}

	if !strings.Contains(data, "current rule is original rule") {
		t.Errorf("Should have original rule in game summary")
	}
}

func TestValidGamePost(t *testing.T) {
	handler := Game()
	server := httptest.NewServer(handler)
	defer server.Close()

	newRuleJson := "{\"rule\":\"1^\", \"true\":\"1^SG\", \"false\":\"0^SG\"}"
	res, err := http.Post(server.URL, "text/json", bytes.NewBuffer([]byte(newRuleJson)))
	defer res.Body.Close()
	data := body(res)

	if err != nil {
		t.Errorf("Should not error when setting rule ")
	}

	if data != "true" {
		t.Errorf("rule should be valid", data)
	}
}

func TestNonValidGamePost(t *testing.T) {
	handler := Game()
	server := httptest.NewServer(handler)
	defer server.Close()

	badRuleJson := "{\"rule\":\"1^\", \"true\":\"0^SG\", \"false\":\"1^SG\"}"
	res, err := http.Post(server.URL, "text/json", bytes.NewBuffer([]byte(badRuleJson)))
	defer res.Body.Close()
	data := body(res)

	if err != nil {
		t.Errorf("Should not error when setting rule ")
	}

	if !strings.Contains(data, "Koans do not fulfull rule") {
		t.Errorf("rule should NOT be valid", data)
	}
}

func body(res *http.Response) string {
	dataBuf, _ := ioutil.ReadAll(res.Body)
	return string(dataBuf)
}

func createKoanBody(escapedKoanString string) string {
	handler := CreateKoan()
	server := httptest.NewServer(handler)
	defer server.Close()

	res, _ := http.Post(server.URL, "text/json", bytes.NewBuffer([]byte(escapedKoanString)))
	defer res.Body.Close()

	return body(res)
}

// test setting rule twice
