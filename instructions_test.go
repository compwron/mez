package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInstructions(t *testing.T) {
	handler := Instructions()
	server := httptest.NewServer(handler)
	defer server.Close()

	res, _ := http.Get(server.URL)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("Expected 200")
	}
	dataBuf, _ := ioutil.ReadAll(res.Body)
	data := string(dataBuf)

	if !strings.Contains(data, "game") {
		t.Errorf("Expected response to include the word game")
	}
}
