package main

import (
	"testing"
)

func Test(t *testing.T) {
	sum := find()
	if sum != 972 {
		t.Errorf("sum = %v, want %v", sum, 972)
	}
}
