package permgen

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	// odd length
	a := []int{1, 2, 3}
	reverse(a)
	if !reflect.DeepEqual(a, []int{3, 2, 1}) {
		t.Errorf("a=%v, want %v", a, []int{3, 2, 1})
	}

	// even length
	a = []int{1, 2, 3, 4}
	reverse(a)
	if !reflect.DeepEqual(a, []int{4, 3, 2, 1}) {
		t.Errorf("a=%v, want %v", a, []int{4, 3, 2, 1})
	}
}

func testPermGen(t *testing.T, gen *PermGen, expected [][]int) {
	for i := 0; i < len(expected); i++ {
		perm, last := gen.Next()
		if !reflect.DeepEqual(perm, expected[i]) {
			t.Fatalf("i=%v: perm = %v, want %v", i, perm, expected[i])
		}
		expectedLast := i == len(expected)-1
		if last != expectedLast {
			t.Fatalf("i=%v: last = %v, want %v", i, last, expectedLast)
		}
	}

	// verify last permutation is sticky
	perm, last := gen.Next()
	if !reflect.DeepEqual(perm, expected[len(expected)-1]) {
		t.Fatalf("last perm = %v, want %v", perm, expected[len(expected)-1])
	}
	if last != true {
		t.Errorf("last = %v, want %v", last, true)
	}
}

func TestPermGen_012(t *testing.T) {
	gen := New([]int{0, 1, 2})
	expected := [][]int{
		[]int{0, 1, 2},
		[]int{0, 2, 1},
		[]int{1, 0, 2},
		[]int{1, 2, 0},
		[]int{2, 0, 1},
		[]int{2, 1, 0},
	}
	testPermGen(t, gen, expected)
}

func TestPermGen_5569(t *testing.T) {
	gen := New([]int{5, 5, 6, 9})
	expected := [][]int{
		[]int{5, 5, 6, 9},
		[]int{5, 5, 9, 6},
		[]int{5, 6, 5, 9},
		[]int{5, 6, 9, 5},
		[]int{5, 9, 5, 6},
		[]int{5, 9, 6, 5},
		[]int{6, 5, 5, 9},
		[]int{6, 5, 9, 5},
		[]int{6, 9, 5, 5},
		[]int{9, 5, 5, 6},
		[]int{9, 5, 6, 5},
		[]int{9, 6, 5, 5},
	}
	testPermGen(t, gen, expected)
}

func TestPermGen_5666(t *testing.T) {
	gen := New([]int{5, 6, 6, 6})
	expected := [][]int{
		[]int{5, 6, 6, 6},
		[]int{6, 5, 6, 6},
		[]int{6, 6, 5, 6},
		[]int{6, 6, 6, 5},
	}
	testPermGen(t, gen, expected)
}

func TestPermGen_5566(t *testing.T) {
	gen := New([]int{5, 5, 6, 6})
	expected := [][]int{
		[]int{5, 5, 6, 6},
		[]int{5, 6, 5, 6},
		[]int{5, 6, 6, 5},
		[]int{6, 5, 5, 6},
		[]int{6, 5, 6, 5},
		[]int{6, 6, 5, 5},
	}
	testPermGen(t, gen, expected)
}
