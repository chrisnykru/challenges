package main

import (
	"testing"
)

func Test(t *testing.T) {
	x := find()
	expectedX := uint(1533776805)
	if x != expectedX {
		t.Errorf("x = %v, want %v", x, expectedX)
	}
}
