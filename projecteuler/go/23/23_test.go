package main

import (
	"testing"
)

func Test(t *testing.T) {
	sum := Sum()
	if sum != 4179871 {
		t.Errorf("sum = %v, want %v", sum, 4179871)
	}
}
