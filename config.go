package main

import (
	"flag"
)

var Configuration struct {
	pipS int
	pipM int
	pipL int
}

func Config() {
	configuration := &Configuration

	pipS := flag.Int("pipS", 1, "Pips on a small piece")
	pipM := flag.Int("pipM", 2, "Pips on a medium piece")
	pipL := flag.Int("pipL", 3, "Pips on a large piece")

	flag.Parse()
	configuration.pipS = *pipS
	configuration.pipM = *pipM
	configuration.pipL = *pipL
}
