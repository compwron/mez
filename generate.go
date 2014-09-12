package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func GenerateRule() {
	cmd := exec.Command("ruby", "rulegen.rb")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		println("Failed to generate new rule.")
	}
	println("Generated rule", out.String())
	CurrentRule = Rule{strings.Split(out.String(), ",")} // Put the parsing in rule.go?
}
