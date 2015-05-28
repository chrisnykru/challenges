package main

import (
	"testing"
)

func Test(t *testing.T) {
	sum := sumAll(factorialOfDigits())
	if sum != 40730 {
		t.Errorf("sum = %v, want %v", sum, 40730)
	}
}
