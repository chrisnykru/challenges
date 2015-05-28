package main

import (
	"testing"
)

func Test(t *testing.T) {
	count := sundaysOnFirstOfMonth()
	if count != 171 {
		t.Errorf("count = %v, want %v", count, 171)
	}
}
