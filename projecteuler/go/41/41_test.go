package main

import (
	"testing"
)

func Test(t *testing.T) {
	x := find()
	if x != 7652413 {
		t.Errorf("x = %v, want %v", x, 7652413)
	}
}
