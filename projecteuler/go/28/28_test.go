package main

import (
	"testing"
)

func Test_5x5(t *testing.T) {
	sum := sumAllDiags(5)
	if sum != 101 {
		t.Errorf("sum = %v, want %v", sum, 101)
	}
}

func Test_1001x1001(t *testing.T) {
	sum := sumAllDiags(1001)
	if sum != 669171001 {
		t.Errorf("sum = %v, want %v", sum, 669171001)
	}
}
