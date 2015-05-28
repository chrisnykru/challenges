package main

import (
	"math/big"
	"testing"
)

func TestPalindrome(t *testing.T) {
	x := palindrome(big.NewInt(123))
	if x.Int64() != 321 {
		t.Errorf("x = %v, want %v", x, 321)
	}
	x = palindrome(big.NewInt(4))
	if x.Int64() != 4 {
		t.Errorf("x = %v, want %v", x, 4)
	}
}

func TestBelow10000(t *testing.T) {
	ll := new(lychrelLut).Init()
	count := ll.countBelow(10000)
	t.Log(count)
}
