package main

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	out, err := HexToBase64(knownInput)
	if err != nil {
		t.Fatal(err)
	}
	if out != knownOutput {
		t.Fatalf("out = %v, want %v", out, knownOutput)
	}
}
