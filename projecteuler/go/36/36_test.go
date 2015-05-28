package main

import (
	"testing"
)

func Test(t *testing.T) {
	sum := base2_and_10_palindromes()
	if sum != 872187 {
		t.Errorf("sum = %v, want %v", sum, 872187)
	}
}
