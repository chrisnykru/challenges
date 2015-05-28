package main

import (
	"testing"
)

func Test(t *testing.T) {
	s := reduce(matrix5x5())
	if s != 2427 {
		t.Errorf("5x5 minimal path = %v, want %v", s, 2427)
	}

	s = reduce(matrix80x80())
	if s != 427337 {
		t.Errorf("80x80 minimal path = %v, want %v", s, 427337)
	}
}
