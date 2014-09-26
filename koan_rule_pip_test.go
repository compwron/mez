package main

import (
	"testing"
)

func TestExactPipCount(t *testing.T) {
	rule := "pip(1)"
	koan := "1^SG"
	verify(rule, koan, t)
}
