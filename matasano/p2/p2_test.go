package p2

import (
	"testing"
)

func TestXOR_hexStrings(t *testing.T) {
	out, err := XOR_hexStrings(knownA, knownB)
	if err != nil {
		t.Fatal(err)
	}
	if out != knownOutput {
		t.Fatalf("out = %v, want %v", out, knownOutput)
	}
}
