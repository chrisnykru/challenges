package main

import (
	"testing"
)

func TestNumLetters(t *testing.T) {
	c1 := numLetters("three hundred and forty-two")
	if c1 != 23 {
		t.Errorf("c1 = %v, want %v", c1, 23)
	}

	c1 = numLetters("one hundred and fifteen")
	if c1 != 20 {
		t.Errorf("c1 = %v, want %v", c1, 20)
	}
}

func TestNumToString(t *testing.T) {
	s1 := numToString(342)
	expStr := "three hundred and forty two" // forty two ~= forty-two for our purposes
	if s1 != expStr {
		t.Errorf("s1 = %v, want %v", s1, expStr)
	}

	s1 = numToString(115)
	expStr = "one hundred and fifteen"
	if s1 != expStr {
		t.Errorf("s1 = %v, want %v", s1, expStr)
	}
}

func TestCount_1_to_1000(t *testing.T) {
	count := letterCount_1_to_1000()
	if count != 21124 {
		t.Errorf("count = %v, want %v", count, 21124)
	}
}
