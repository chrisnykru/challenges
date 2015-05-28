package main

import (
	"testing"
)

func Test(t *testing.T) {
	sum := tri67()
	if sum != 7273 {
		t.Errorf("sum = %v, want %v", sum, 7273)
	}
}
