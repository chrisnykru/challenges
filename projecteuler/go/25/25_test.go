package main

import (
	"testing"
)

func Test(t *testing.T) {
	term, _ := firstFibTermOfNDigits(1000)
	if term != 4782 {
		t.Errorf("term = %v, want %v", 4782)
	}
}
