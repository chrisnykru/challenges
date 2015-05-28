package main

import (
	"testing"
	"math/big"
)

func TestNonBouncyBelow1e6(t *testing.T) {
  r := nonBouncyUnder(big.NewInt(1000000))
	if r != 12951 {
		t.Errorf("r = %v, want %v", r, 12951)
	}
}
