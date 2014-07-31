package main

import (
	"encoding/json"
	"io"
)

func Parse(data io.Reader) (map[string]interface{}, error) {
	var jsonHolder interface{}
	var parsedMap map[string]interface{}
	err := json.NewDecoder(data).Decode(&jsonHolder)
	if err == nil {
		parsedMap = jsonHolder.(map[string]interface{})
	}
	return parsedMap, err
}
