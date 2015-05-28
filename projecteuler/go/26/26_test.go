package main

import (
	"testing"
)

func TestCycle(t *testing.T) {
	clen := cycle(2)
	if clen != 0 {
		t.Errorf("cycle(%v) = %v, want %v", 2, clen, 0)
	}

	clen = cycle(3)
	if clen != 1 {
		t.Errorf("cycle(%v) = %v, want %v", 3, clen, 1)
	}

	clen = cycle(4)
	if clen != 0 {
		t.Errorf("cycle(%v) = %v, want %v", 4, clen, 0)
	}

	clen = cycle(6)
	if clen != 1 {
		t.Errorf("cycle(%v) = %v, want %v", 6, clen, 1)
	}

	clen = cycle(7)
	if clen != 6 {
		t.Errorf("cycle(%v) = %v, want %v", 7, clen, 6)
	}

	clen = cycle(13)
	if clen != 6 {
		t.Errorf("cycle(%v) = %v, want %v", 13, clen, 6)
	}

	clen = cycle(990)
	if clen != 2 {
		t.Errorf("cycle(%v) = %v, want %v", 990, clen, 2)
	}
}

func Test(t *testing.T) {
	d, dlen := longestCycle(1000)
	if d != 983 {
		t.Errorf("d=%v, want %v", d, 983)
	}
	if dlen != 981 {
		t.Errorf("dlen=%v, want %v", dlen, 981)
	}
}
