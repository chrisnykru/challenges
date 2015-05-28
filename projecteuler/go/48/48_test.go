package main

import (
	"testing"
)

func Test(t *testing.T) {
	s := fast().String()
	last10 := s[len(s)-10:]
	expected := "9110846700"
	if last10 != expected {
		t.Errorf("last10 = %v, want %v", last10, expected)
	}
}
