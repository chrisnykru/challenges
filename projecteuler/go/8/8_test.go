package main

import (
	"testing"
)

func Test(t *testing.T) {
	_, gp := greatestProductOfFiveConsecutiveDigits(s1000)
	if gp != 40824 {
		t.Errorf("gp = %v, want %v", gp, 40824)
	}
}
