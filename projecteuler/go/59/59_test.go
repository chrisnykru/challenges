package main

import (
	"testing"
)

func Test(t *testing.T) {
	sum := sum(bruteForce(loadCipher()))
	if sum != 107359 {
		t.Errorf("sum = %v, want %v", sum, 107359)
	}
}
