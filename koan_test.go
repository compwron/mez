package main

import (
	"strings"
	"testing"
)

func TestKoanIsInvalidWhenThirdCharacterIsAnInvalidSize(t *testing.T) {
	falsifyInvalidKoan("1^ZG", t)
}

func TestKoanIsInvalidWhenFourthCharacterIsNotAValidSizeOrColor(t *testing.T) {
	falsifyInvalidKoan("1^SZ", t)
}

func TestKoanIsInvalidWhenFifthCharacterIsNotAValidColor(t *testing.T) {
	falsifyInvalidKoan("1^2SZ", t)
}

func TestKoanIsNotLongerThanFiveCharacters(t *testing.T) {
	falsifyInvalidKoan("1^2SGZ", t)
}

func TestAddKoanIfValidCatchesBadKoanSyntax(t *testing.T) {
	invalidKoan := strings.NewReader("{\"koan\":\"ABCD\"}")
	parsed, _ := Parse(invalidKoan)
	result := AddKoanIfValid(parsed)
	if result != "Invalid koan" {
		t.Error("Koan should be detected as invalid")
	}
}

func TestValidKoanIsValid(t *testing.T) {
	verifyValidKoan("1^SG", t)
}

func TestKoanWithMultiplePiecesAndColorsIsValid(t *testing.T) {
	verifyValidKoan("3^3SG", t)
}

func TestInvalidKoanIsInvalid(t *testing.T) {
	falsifyInvalidKoan("foo", t)
}

func TestNegativeKoanIsInvalid(t *testing.T) {
	falsifyInvalidKoan("!1^SG", t)
}

func TestKoanWithoutCountIsInvalid(t *testing.T) {
	falsifyInvalidKoan("^SG", t)
}

func TestKoanWithoutPieceIsInvalid(t *testing.T) {
	falsifyInvalidKoan("1SG", t)
}

func TestKoanWithoutColorIsInvalid(t *testing.T) {
	falsifyInvalidKoan("1^S", t)
}

func TestKoanWithoutSizeIsInvalid(t *testing.T) {
	falsifyInvalidKoan("1^G", t)
}
