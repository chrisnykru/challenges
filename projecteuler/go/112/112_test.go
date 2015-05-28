package main

import (
	"testing"
	"github.com/zeroshirts/challenges/projecteuler/go/112/misc112"
)

func TestBouncyNumbersBelow1000(t *testing.T) {
	count := 0
	for i := 100; i < 1000; i++ {
		if misc112.IsBouncy(i) {
			count++
		}
	}
	if count != 525 {
		t.Errorf("count = %v, want %v", count, 525)
	}
}

func TestFindLowestNumberWithBouncyPercent(t *testing.T) {
	p := 99
  r := findLowestNumberWithBouncyPercent(p)
	if r != 1587000 {
		t.Errorf("r = %v, want %v", r, 1587000)
	}
}
