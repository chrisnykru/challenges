package main

import (
	"testing"
)

func Test(t *testing.T) {
	d := minDiff()
	if d != 5482660 {
		t.Errorf("d = %v, want %v", d, 5482660)
	}
}
