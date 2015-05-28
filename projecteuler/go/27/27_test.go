package main

import (
	"testing"
)

func Test(t *testing.T) {
	a, b, primes := optimalAB()
	if a != -61 {
		t.Errorf("a = %v, want %v", a, -61)
	}
	if b != 971 {
		t.Errorf("b = %v, want %v", b, 971)
	}
	if len(primes) != 71 {
		t.Errorf("len(primes) = %v, want %v", len(primes), 71)
	}
}
