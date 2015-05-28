package main

import (
	"testing"
)

func TestSumFib(t *testing.T) {
	sum := sumFibUpTo(4000000)
	if sum != 4613732 {
		t.Errorf("sum = %v, want %v", sum, 4613732)
	}
}

func BenchmarkSumFibUpTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumFibUpTo(4000000)
	}
}
