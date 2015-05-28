package main

import (
	"testing"
)

func Test(t *testing.T) {
	x := find()
	if x != 134043 {
		t.Errorf("x = %v, want %v", x, 134043)
	}
}
