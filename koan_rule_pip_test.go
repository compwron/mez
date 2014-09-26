package main

import (
	"testing"
)

func TestExactPipCountOne(t *testing.T) {
	rule := "pip(1)"
	koan := "1^SG"
	verify(rule, koan, t)
}

func TestExactPipCountOneFail(t *testing.T) {
	rule := "pip(2)"
	koan := "1^SG"
	falsify(rule, koan, t)
}

func TestExactPipCountMedium(t *testing.T) {
	rule := "pip(2)"
	koan := "1^MG"
	verify(rule, koan, t)
}

func TestExactPipCountLarge(t *testing.T) {
	rule := "pip(3)"
	koan := "1^LG"
	verify(rule, koan, t)
}

func TestSumPipCountPass(t *testing.T) {
	rule := "pip(4)"
	koan := "2^MG"
	verify(rule, koan, t)
}

func TestSumPipCountFail(t *testing.T) {
	rule := "pip(5)"
	koan := "2^MG"
	falsify(rule, koan, t)
}

func TestZeroPipCountPass(t *testing.T) {
	rule := "pip(0)"
	koan := "0^MG"
	verify(rule, koan, t)
}

func TestZeroPipCountFail(t *testing.T) {
	rule := "pip(0)"
	koan := "1^MG"
	falsify(rule, koan, t)
}
