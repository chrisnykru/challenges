package main

import (
	"testing"
)

func TestSmallest_10(t *testing.T) {
	x := smallestNumberEvenlyDivisible(10)
	expected_x := int64(2520)
	if x != expected_x {
		t.Errorf("x = %v, want %v", x, expected_x)
	}
}

func TestSmallest_20(t *testing.T) {
	x := smallestNumberEvenlyDivisible(20)
	expected_x := int64(232792560)
	if x != expected_x {
		t.Errorf("x = %v, want %v", x, expected_x)
	}
}
