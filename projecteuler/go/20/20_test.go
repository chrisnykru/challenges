package main

import (
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"testing"
)

func Test(t *testing.T) {
	f := misc.Factorial(100)
	sum := sumDigits(f)
	if sum != 648 {
		t.Errorf("sum = %v, want %v", sum, 648)
	}
}
