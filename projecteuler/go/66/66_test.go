package main

import "testing"

func TestFindEasy(t *testing.T) {
	maxd := int64(8)
	if d := search(maxd); d != 5 {
		t.Errorf("search(%d) = %v, want %v", maxd, d, 5)
	}
}
