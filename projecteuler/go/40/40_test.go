package main

import (
	"testing"
)

func Test(t *testing.T) {
	x := expr()
	if x != 210 {
		t.Errorf("x = %v, want %v", x, 210)
	}
}
