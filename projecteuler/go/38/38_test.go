package main

import (
	"testing"
)

func TestConcatProductPandigital_1to9(t *testing.T) {
	// known pandigital
	num, n, ok := concatProductPandigital_1to9(192)
	if ok != true {
		t.Fatalf("ok = %v, want %v", ok, true)
	}
	expectedNum := uint64(192384576)
	if num != expectedNum {
		t.Errorf("num = %v, want %v", num, expectedNum)
	}
	if n != 3 {
		t.Errorf("n = %v, want %v", n, 3)
	}

	// not pandigital (191 * 1 = 191; duplicate 1s)
	_, _, ok = concatProductPandigital_1to9(191)
	if ok != false {
		t.Fatalf("ok = %v, want %v", ok, true)
	}

	// not pandigital (n < 2)
	_, _, ok = concatProductPandigital_1to9(98764321)
	if ok != false {
		t.Fatalf("ok = %v, want %v", ok, true)
	}
}

func Test(t *testing.T) {
	expected := uint64(932718654)
	l := largest()
	if l != expected {
		t.Errorf("l = %v, want %v", l, expected)
	}
}
