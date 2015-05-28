package main

import (
	"testing"
)

func Test(t *testing.T) {
	sum := sumOfAmicableNumbersUnder(10000)
	if sum != 31626 {
		t.Errorf("sum = %v, want %v", sum, 31626)
	}
}
