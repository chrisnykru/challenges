package main

import (
	primegen "github.com/jbarham/primegen.go"
	"testing"
)

func Test6thPrime(t *testing.T) {
	gen := primegen.New()
	var prime uint64
	for i := 0; i < 6; i++ {
		prime = gen.Next()
	}

	if prime != 13 {
		t.Errorf("prime = %v, want %v", prime, 13)
	}
}

func Test10001stPrime(t *testing.T) {
	gen := primegen.New()
	var prime uint64
	for i := 0; i < 10001; i++ {
		prime = gen.Next()
	}

	if prime != 104743 {
		t.Errorf("prime = %v, want %v", prime, 104743)
	}
}
