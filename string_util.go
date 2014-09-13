package main

import (
	"strings"
)

func chunk(thingWithComma string) []string {
	return strings.Split(thingWithComma, ",")
}

func pieces(chunk string) []string {
	return strings.Split(chunk, "")
}
