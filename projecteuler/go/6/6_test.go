package main

import (
	"testing"
)

func TestNaive(t *testing.T) {
	x := naive(100)
	expectedX := -25164150
	if x != expectedX {
		t.Errorf("naive = %v, want %v", x, expectedX)
	}
}

func TestSmarter(t *testing.T) {
	x := smarter(100)
	expectedX := -25164150
	if x != expectedX {
		t.Errorf("smarter = %v, want %v", x, expectedX)
	}
}

func BenchmarkNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = naive(100)
	}
}

func BenchmarkSmarter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = smarter(100)
	}
}
