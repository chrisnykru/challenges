package main

import (
	"testing"
)

func Test(t *testing.T) {
	start, maxLen := startWithLongestSequence(1000000)
	if start != 837799 {
		t.Errorf("start = %v, want %v", start, 837799)
	}
	if maxLen != 525 {
		t.Errorf("maxLen = %v, want %v", maxLen, 525)
	}
}
