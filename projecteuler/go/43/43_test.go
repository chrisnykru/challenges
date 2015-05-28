package main

import (
	"testing"
)

func Test(t *testing.T) {
	sum := sumSpecialPandigitals()
	expectedSum := uint64(16695334890)
	if sum != expectedSum {
		t.Errorf("sum = %v, want %v", sum, expectedSum)
	}
}
