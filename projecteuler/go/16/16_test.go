package main

import (
	"math/big"
	"testing"
)

func Test_2tothe15th(t *testing.T) {
	s := sumOfDigits(big.NewInt(2), big.NewInt(15))
	if s.Int64() != 26 {
		t.Errorf("sum = %v, want %v", s.Int64(), 26)
	}
}

func Test_2tothe1000th(t *testing.T) {
	s := sumOfDigits(big.NewInt(2), big.NewInt(1000))
	if s.Int64() != 1366 {
		t.Errorf("sum = %v, want %v", s.Int64(), 1366)
	}
}
