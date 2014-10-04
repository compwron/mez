package main

import (
	"encoding/json"
	"io"
)

func Parse(data io.Reader) (map[string]interface{}, error) {
	var jsonHolder interface{}
	var ParsedMap map[string]interface{}
	err := json.NewDecoder(data).Decode(&jsonHolder)
	if err == nil {
		ParsedMap = jsonHolder.(map[string]interface{})
	}
	return ParsedMap, err
}
