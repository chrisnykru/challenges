package main

import (
	"testing"
)

func Test(t *testing.T) {
	solCount, p := maxSol()
	if solCount != 8 || p != 840 {
		t.Errorf("solCount=%v, p=%v, want %v, %v", solCount, p, 8, 840)
	}
}
