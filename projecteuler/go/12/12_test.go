package main

import (
	"testing"
)

func Test(t *testing.T) {
	tri := findTriNumWithOver500Divisors()
	expected := int64(76576500)
	if tri != expected {
		t.Errorf("tri = %v, want %v", tri, expected)
	}
}
