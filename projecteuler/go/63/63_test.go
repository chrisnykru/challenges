package main

import (
	"testing"
)

func Test(t *testing.T) {
	r := solve()
	if r != 49 {
		t.Errorf("r = %v, want %v\n", r, 49)
	}
}
