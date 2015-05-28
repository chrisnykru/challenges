package main

import (
	"testing"
)

func Test7(t *testing.T) {
	family := npvf(7)
	if family[0] != 56003 {
		t.Errorf("family[0] = %v, want %v", family[0], 56003)
	}
}

func Test8(t *testing.T) {
	family := npvf(8)
	if family[0] != 121313 {
		t.Errorf("family[0] = %v, want %v", family[0], 121313)
	}
}
