package main

import "testing"

func TestCount89(t *testing.T) {
	count := count89()
	expected := 8581146
	if count != expected {
		t.Errorf("count = %v, want %v", count, expected)
	}
}
