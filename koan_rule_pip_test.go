package main

import (
	"testing"
)

func TestPipCountPassOneKoanPiecePassWithOnlyOnePip(t *testing.T) {
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

func TestPipCountPassMultiChunkPass(t *testing.T) {
	rule := "pip(3)"
	koan := "1^SG,2^SG"
	verify(rule, koan, t)
}
