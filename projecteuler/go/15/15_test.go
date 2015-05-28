package main

import (
	"testing"
)

func Test2(t *testing.T) {
	x := numRoutes(2)
	expected_x := int64(6)
	if x != expected_x {
		t.Errorf("x = %v, want %v", x, expected_x)
	}
}

func Test20(t *testing.T) {
	x := numRoutes(20)
	expected_x := int64(137846528820)
	if x != expected_x {
		t.Errorf("x = %v, want %v", x, expected_x)
	}
}
