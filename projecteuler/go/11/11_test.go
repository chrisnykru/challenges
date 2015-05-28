package main

import (
	"testing"
)

func Test(t *testing.T) {
	gp := greatestProduct(grid)
	expected := 70600674
	if gp != expected {
		t.Errorf("gp = %v, want %v", gp, expected)
	}
}
