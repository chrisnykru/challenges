package main

import (
	"testing"
)

func TestFast(t *testing.T) {
	s := fast().String()
	last10 := s[len(s)-10:]
	expected := "8739992577"
	if last10 != expected {
		t.Errorf("last10 = %s, want %s", last10, expected)
	}
}
