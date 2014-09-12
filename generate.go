package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func GenerateRule() string {
	if OkToChangeCurrentRule() {
		cmd := exec.Command("ruby", "rulegen.rb")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			message := "Failed to generate new rule.\n"
			println(message)
			return message
		}
		println("Generated rule", out.String())
		CurrentRule = Rule{strings.Split(out.String(), ",")} // Put the parsing in rule.go?
		Koans = nil
		return "Generated and set rule\n"
	}
	return "It is not ok to set a rule right now because a non-default rule currently in play.\n"
}
