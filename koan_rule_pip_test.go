package main

import (
	"testing"
)

	// one piece koanchunk
	// multipiece koan chunk
	// one piece, two koanchunks
	// multipiece multichunk koan
	// combo with other rule tyles
func TestPipCountPassOneKoanPiecePass(t *testing.T) {
	rule := "pip(1)"
	koan := "1^SG"
	verify(rule, koan, t)
}

func TestPipCountPassOneKoanPiecePassBecauseOfMoreThanPip(t *testing.T) {
	rule := "pip(1)" // pip rules must be "more than"
	koan := "2^SG"
	verify(rule, koan, t)
}

func TestPipCountPassOneKoanPieceFail(t *testing.T) {
	rule := "pip(2)"
	koan := "1^SG"
	falsify(rule, koan, t)
}

// Gotta start by splitting rule into chunks relevant for different rule types, evaluating against koans separately... 
func TestPipCountPassMultiChunkPass(t *testing.T) {
	rule := "pip(3)"
	koan := "1^SG, 2^SG"
	verify(rule, koan, t)
}
