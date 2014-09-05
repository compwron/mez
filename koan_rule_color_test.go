package main

import (
	"testing"
)

// A nice failing test, ready for me to come back from vacation
func TestAllPiecesMustBeSameColorFailWithMultiColorKoan(t *testing.T) {
	rule := "G"
	multiColorKoan := "1^SB,1^SG"
	// A rule that there must be at least one G passes if there is one G and one non-G in the koan
	verify(rule, multiColorKoan, t)
}

func TestMultipleColorPiecesPass(t *testing.T) {
	rule := "2G"
	multiColorKoan := "1^SG,1^SG"
	verify(rule, multiColorKoan, t)
}

func TestMultipleColorPiecesFail(t *testing.T) {
	rule := "2G"
	multiColorKoan := "1^SB,1^SG"
	falsify(rule, multiColorKoan, t)
}

func TestNegativeAllPiecesMustBeSameColorPass(t *testing.T) {}
func TestNegativeAllPiecesMustBeSameColorFail(t *testing.T) {}

func TestAllPiecesMustBeSameColorPassSimple(t *testing.T) {
	rule := "G" // All pieces must be green
	koan := "1^SG"
	verify(rule, koan, t)
}

func TestAllPiecesMustBeSameColorFail(t *testing.T) {
	rule := "G"
	koan := "1^SB"
	falsify(rule, koan, t)
}

func TestAllPiecesMustBeSameColorPassWithMultiColorKoan(t *testing.T) {
	rule := "G"
	multiColorKoan := "1^SG,1^SG"
	verify(rule, multiColorKoan, t)
}
