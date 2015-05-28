package main

import (
	"testing"
)

func Test_10(t *testing.T) {
	sum := sumOfPrimesBelow(10).Int64()
	expectedSum := int64(17)
	if sum != expectedSum {
		t.Errorf("sum = %v, want %v", sum, expectedSum)
	}
}

func Test_2e6(t *testing.T) {
	sum := sumOfPrimesBelow(2000000).Int64()
	expectedSum := int64(142913828922)
	if sum != expectedSum {
		t.Errorf("sum = %v, want %v", sum, expectedSum)
	}
}
