package main

import (
	"testing"
)

func Test(t *testing.T) {
	x := largest()
	if x != 6857 {
		t.Errorf("x = %v, want %v", x, 6857)
	}
}

func BenchmarkLargest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = largest()
	}
}
