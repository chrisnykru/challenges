package main

import (
	"testing"
)

func Test(t *testing.T) {
	tprimes := truncatablePrimes()
	t.Log(tprimes)

	tprimeSum := sum(tprimes)
	expected := uint64(748317)
	if tprimeSum != expected {
		t.Errorf("tprime sum = %v, want %v", tprimeSum, expected)
	}
}
