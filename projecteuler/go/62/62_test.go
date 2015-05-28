package main

import "testing"

func Test3Perm(t *testing.T) {
	r := smallestCubeWith5Perm(3)
	expected := 41063625
	if r != expected {
		t.Errorf("r = %v, want %v", r, expected)
	}
}

func Test5Perm(t *testing.T) {
	r := smallestCubeWith5Perm(5)
	expected := 127035954683
	if r != expected {
		t.Errorf("r = %v, want %v", r, expected)
	}
}

func TestLazyDeconstruct(t *testing.T) {
	var ld lazyDeconstruct
	ld.Set(0)
	expected := lazyDeconstruct{[10]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	if ld != expected {
		t.Errorf("ld = %v, want %v", ld, expected)
	}

	ld.Set(123)
	expected = lazyDeconstruct{[10]int{0, 1, 1, 1, 0, 0, 0, 0, 0, 0}}
	if ld != expected {
		t.Errorf("ld = %v, want %v", ld, expected)
	}

	ld.Set(10)
	expected = lazyDeconstruct{[10]int{1, 1, 0, 0, 0, 0, 0, 0, 0, 0}}
	if ld != expected {
		t.Errorf("ld = %v, want %v", ld, expected)
	}

	ld.Set(100)
	expected = lazyDeconstruct{[10]int{2, 1, 0, 0, 0, 0, 0, 0, 0, 0}}
	if ld != expected {
		t.Errorf("ld = %v, want %v", ld, expected)
	}

	ld.Set(5567)
	expected = lazyDeconstruct{[10]int{0, 0, 0, 0, 0, 2, 1, 1, 0, 0}}
	if ld != expected {
		t.Errorf("ld = %v, want %v", ld, expected)
	}
}
