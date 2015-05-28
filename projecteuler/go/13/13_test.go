package main

import (
	"testing"
)

func Test(t *testing.T) {
	s10 := firstTenDigits(sumNumberStrings(numbers))
	expected := "5537376230"
	if s10 != expected {
		t.Errorf("s10 = %v, want %v", s10, expected)
	}
}
