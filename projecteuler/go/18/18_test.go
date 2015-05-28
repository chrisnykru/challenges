package main

import (
	"testing"
)

func TestSmall(t *testing.T) {
	sum := small()
	if sum != 23 {
		t.Errorf("sum = %v, want %v", sum, 23)
	}
}

func TestBig(t *testing.T) {
	sum := big()
	if sum != 1074 {
		t.Errorf("sum = %v, want %v", sum, 1074)
	}
}
