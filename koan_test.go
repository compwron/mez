package main

import (
	"testing"
)

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
