package main

import (
	"testing"
)

func TestIdentityPandigital_1to9(t *testing.T) {
	r := identityPandigital_1to9(7254)
	if r != true {
		t.Errorf("identity(%v) = %v, want %v", 7254, r, true)
	}
}

func Test(t *testing.T) {
	p := pandigital_1to9_identityProducts()
	if sum(p) != 45228 {
		t.Errorf("sum = %v, want %v", sum(p), 45228)
	}
}
