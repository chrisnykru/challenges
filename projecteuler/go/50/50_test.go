package main

import (
	"testing"
)

func Test(t *testing.T) {
	prime, numTerms := primeWithMostConsecutiveTerms(1000000)
	if prime != uint64(997651) {
		t.Errorf("prime = %v, want %v", prime, 997651)
	}
	if numTerms != 543 {
		t.Errorf("numTerms = %v, want %v", numTerms, 543)
	}
}
