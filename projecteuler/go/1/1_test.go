package main

import (
	"testing"
)

func Test10(t *testing.T) {
	sum10 := sumOfMultiples(10)
	if sum10 != 23 {
		t.Errorf("sum(3,5 multiples < 10) = %v, want %v", sum10, 23)
	}
	sum10 = fasterSumOfMultiples(10)
	if sum10 != 23 {
		t.Errorf("sum(3,5 multiples < 10) = %v, want %v", sum10, 23)
	}
}

func Test1000(t *testing.T) {
	sum1000 := sumOfMultiples(1000)
	if sum1000 != 233168 {
		t.Errorf("sum(3,5 multiples < 1000) = %v, want %v", sum1000, 233168)
	}
	sum1000 = fasterSumOfMultiples(1000)
	if sum1000 != 233168 {
		t.Errorf("sum(3,5 multiples < 1000) = %v, want %v", sum1000, 233168)
	}
}

func BenchmarkNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sumOfMultiples(1000)
	}
}

func BenchmarkFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fasterSumOfMultiples(1000)
	}
}
