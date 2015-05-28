package pandigital

import (
	"testing"
)

func Test(t *testing.T) {
	var p Pandigital

	// initial status
	status := p.Vet(1, 9)
	if status != Indeterminate {
		t.Fatalf("status = %v, want %v", status, Indeterminate)
	}

	p.AddDigits(1234)
	status = p.Vet(1, 9)
	if status != Indeterminate {
		t.Fatalf("status = %v, want %v", status, Indeterminate)
	}

	p.AddDigits(56789)
	status = p.Vet(1, 9)
	if status != Valid {
		t.Fatalf("status = %v, want %v", status, Valid)
	}

	x := p.Uint64()
	expectedX := uint64(123456789)
	if x != expectedX {
		t.Fatalf("x = %v, want %v", x, expectedX)
	}

	// duplicate 2 --> invalid
	p.AddDigits(2)
	status = p.Vet(1, 9)
	if status != Invalid {
		t.Fatalf("status = %v, want %v", status, Invalid)
	}
}
