package main

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if r := IsPalindrome(99); r != true {
		t.Errorf("IsPalindrome(99)=%v, want %v", r, true)
	}

	if r := IsPalindrome(101); r != true {
		t.Errorf("IsPalindrome(101)=%v, want %v", r, true)
	}

	if r := IsPalindrome(999); r != true {
		t.Errorf("IsPalindrome(999)=%v, want %v", r, true)
	}

	if r := IsPalindrome(1); r != true {
		t.Errorf("IsPalindrome(1)=%v, want %v", r, true)
	}

	if r := IsPalindrome(102); r != false {
		t.Errorf("IsPalindrome(102)=%v, want %v", r, false)
	}

	if r := IsPalindrome(1212); r != false {
		t.Errorf("IsPalindrome(1212)=%v, want %v", r, false)
	}

	if r := IsPalindrome(929); r != true {
		t.Errorf("IsPalindrome(929)=%v, want %v", r, true)
	}
}

func TestLargestPP(t *testing.T) {
	pp := largestProductPalindrome(999)
	expected_pp := productPalindrome{i: 993, j: 913, ij: 906609}
	if !reflect.DeepEqual(pp, expected_pp) {
		t.Errorf("pp = %v, want %v", pp, expected_pp)
	}
}
