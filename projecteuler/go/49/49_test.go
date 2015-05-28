package main

import (
	"testing"
)

func Test(t *testing.T) {
	s := concatSeq(sequence())
	expected_s := "296962999629"
	if s != expected_s {
		t.Errorf("s = %q, want %q", s, expected_s)
	}
}
