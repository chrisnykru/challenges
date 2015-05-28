package main

import (
	"testing"
)

func Test_a_2to5_b_2to5(t *testing.T) {
	m := termMap(2, 5, 2, 5)
	if len(m) != 15 {
		t.Errorf("len(m) = %v, want %v", len(m), 15)
	}
}

func Test_a_2to100_b_2to100(t *testing.T) {
	m := termMap(2, 100, 2, 100)
	if len(m) != 9183 {
		t.Errorf("len(m) = %v, want %v", len(m), 9183)
	}
}
