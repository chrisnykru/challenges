package main

import (
	"testing"
)

func Test(t *testing.T) {
	n, d := solve(findCuriousFractions())
	if n != 1 || d != 100 {
		t.Errorf("n/d = %v/%v, want %v/%v", n, d, 1, 100)
	}
}
