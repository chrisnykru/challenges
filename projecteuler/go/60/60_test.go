package main

import (
	"testing"
)

func TestIsRemarkable(t *testing.T) {
	// "The primes 3, 7, 109, and 673, are quite remarkable"
	stim := []struct {
		x, y       uint64
		remarkable bool
	}{
		{3, 5, false},
		{3, 7, true},
		{3, 109, true},
		{3, 673, true},
		{7, 109, true},
		{7, 673, true},
		{109, 673, true},
	}
	primeLut := newPrimeLut()
	for _, s := range stim {
		r := isRemarkable(primeLut, s.x, s.y)
		if r != s.remarkable {
			t.Errorf("(%d,%d) = %v, want %v", s.x, s.y, r, s.remarkable)
		}
	}
}

func TestSetOf4(t *testing.T) {
	sum := findLowestSetSum(4)
	if sum != 792 {
		t.Errorf("sum = %v, want %v", sum, 792)
	}
}

func TestSetOfFive(t *testing.T) {
	sum := findLowestSetSum(5)
	if sum != 26033 {
		t.Errorf("sum = %v, want %v", sum, 26033)
	}
}
