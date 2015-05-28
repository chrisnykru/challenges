package main

import (
	"testing"
)

func Test(t *testing.T) {
	count := triangleWordCount()
	if count != 162 {
		t.Errorf("count = %v, want %v", count, 162)
	}
}
