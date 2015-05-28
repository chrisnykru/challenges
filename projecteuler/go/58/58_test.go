package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
  r := solve()
	if r != 26241 {
		t.Errorf("r = %v, want %v", r, 26241)
	}
}
