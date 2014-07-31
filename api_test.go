package main

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	fmt.Println("Running a test")
	if 1 != 2 {
		t.Errorf("broken test")
	}
}
