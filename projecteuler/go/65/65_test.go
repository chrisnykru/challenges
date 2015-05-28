package main

import "testing"

func Test10thConvergent(t *testing.T) {
	r := sumOfNumeratorDigits(10)
	if r != 17 {
		t.Errorf("r = %v, want %v", r, 17)
	}
}

func Test100thConvergent(t *testing.T) {
	r := sumOfNumeratorDigits(100)
	if r != 272 {
		t.Errorf("r = %v, want %v", r, 272)
	}
}
