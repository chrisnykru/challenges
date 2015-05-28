package main

import (
	"testing"
)

func TestLeftRotateStr(t *testing.T) {
	s := "a"
	s2 := leftRotateStr(s)
	if s2 != "a" {
		t.Errorf("s2 = %v, want %v", s2, "a")
	}

	s = "abc"
	s2 = leftRotateStr(s)
	if s2 != "bca" {
		t.Errorf("s2 = %v, want %v", s2, "bca")
	}
	s2 = leftRotateStr(s2)
	if s2 != "cab" {
		t.Errorf("s2 = %v, want %v", s2, "cab")
	}
}

func TestBelow1e6(t *testing.T) {
	c := circularPrimesBelow1e6()
	if len(c) != 55 {
		t.Errorf("len(c) = %v, want %v", len(c), 55)
	}
}
