package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestCreateKoanWithWrongHttpMethodFails(t *testing.T) {
	handler := CreateKoan()
	server := httptest.NewServer(handler)
	defer server.Close()

	res, _ := http.Get(server.URL)
	defer res.Body.Close()

	if res.Status != "405 Method Not Allowed" || body(res) != "not supported" {
		println(res.Status, body(res))
		t.Errorf("Should not support Get")
	}
}

func TestParseFailInCreateKoan(t *testing.T) {
	handler := CreateKoan()
	server := httptest.NewServer(handler)
	defer server.Close()

	res, _ := http.Post(server.URL, "text/json", bytes.NewBuffer([]byte("{")))
	defer res.Body.Close()

	if res.Status != "405 Method Not Allowed" || body(res) != "Impossible to parse formatting" {
		println(res.Status, body(res))
		t.Errorf("Should not succeed when parsing bad json")
	}
}

func TestCreateGameSadPath(t *testing.T) {
	handler := Game()
	server := httptest.NewServer(handler)
	defer server.Close()

	res, _ := http.Head(server.URL)
	defer res.Body.Close()

	if res.Status != "405 Method Not Allowed" {
		println(res.Status)
		t.Errorf("Should not support Get")
	}
}

func TestGenerateRuleSadPath(t *testing.T) {
	handler := StartGameWithUnknownRule()
	server := httptest.NewServer(handler)
	defer server.Close()

	res, _ := http.Get(server.URL)
	defer res.Body.Close()

	if res.Status != "405 Method Not Allowed" {
		t.Errorf("Should not support Get")
	}
}

func TestGenerateRuleHappyPath(t *testing.T) {
	handler := StartGameWithUnknownRule()
	server := httptest.NewServer(handler)
	defer server.Close()

	res, _ := http.Post(server.URL, "text/json", bytes.NewBuffer([]byte("{}")))
	defer res.Body.Close()

	if reflect.DeepEqual(OriginalRule, CurrentRule) || RuleMatches("") {
		t.Errorf("Should have set original rule to a new non-blank rule")
	}
}

func TestCreateKoanValid(t *testing.T) {
	CurrentRule = OriginalRule // Setup state for test
	data := createKoanBody("{\"koan\":\"1^SG\"}")
	if !(data == "true") {
		t.Errorf("should return true. Actually got: " + data + " and current rule is: " + RuleToString(CurrentRule))
	}
}

func TestCreateKoanInvalid(t *testing.T) {
	CurrentRule = OriginalRule // Setup state for test
	data := createKoanBody("{\"koan\":\"0>SG\"}")
	if !(data == "false") {
		t.Errorf("should return false. Actually got: " + data + " and current rule is: " + RuleToString(CurrentRule))
	}
}

func TestGuessRuleGet(t *testing.T) {
	handler := GuessRule()
	server := httptest.NewServer(handler)
	defer server.Close()

	res, _ := http.Get(server.URL)
	defer res.Body.Close()

	if res.Status != "405 Method Not Allowed" {
		t.Errorf(fmt.Sprintf("Should not support GET for Guess; expected status \"405 Method Not Allowed\", but got \"%s\"", res.Status))
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
	res, _ := http.Post(server.URL, "text/json", bytes.NewBuffer([]byte("{\"rule\":\"1rule,2rule\"}")))
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

	if data != "Successfully set rule" {
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

func TestCorrectGuess(t *testing.T) {
	handler := GuessRule()
	server := httptest.NewServer(handler)
	defer server.Close()

	correctGuess := "{\"rule\": \"" + CurrentRule.ruleDescriptions[0] + "\"}"
	res, err := http.Post(server.URL, "text/json", bytes.NewBuffer([]byte(correctGuess)))
	defer res.Body.Close()
	data := body(res)

	if err != nil {
		t.Errorf("Should not get error when submitting correct guess")
	}

	if !strings.Contains(data, "Victory") {
		t.Errorf("correct guess should prompt victory message", data)
	}

	if !strings.Contains(data, "Victory") {
		t.Errorf("correct guess should prompt victory message", data)
	}
}

func TestIncorrectGuess(t *testing.T) {
	handler := GuessRule()
	server := httptest.NewServer(handler)
	defer server.Close()

	incorrectGuess := "{\"rule\": \"What is Toronto????\"}"

	res, _ := http.Post(server.URL, "text/json", bytes.NewBuffer([]byte(incorrectGuess)))
	defer res.Body.Close()
	data := body(res)

	if !strings.Contains(data, "incorrect guess") {
		t.Errorf("incorrect guess should prompt corresponding message", data)
	}
}

func body(res *http.Response) string {
	dataBuf, _ := ioutil.ReadAll(res.Body)
	return strings.Trim(string(dataBuf), "\n")
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
