package main

import (
	"testing"
)

// A nice failing test, ready for me to come back from vacation
func TestAllPiecesMustBeSameColorFailWithMultiColorKoan(t *testing.T) {
	rule := "G"
	multiColorKoan := "1^SB,1^SG"
	verify(rule, multiColorKoan, t)
}

func TestNegativeAllPiecesMustBeSameColorPass(t *testing.T) {}
func TestNegativeAllPiecesMustBeSameColorFail(t *testing.T) {}
