package main

import (
	"testing"
	"math/big"
	"reflect"
)

func TestApproximator(t *testing.T) {
	stim := []struct{
		iterations int
		rat *big.Rat
	}{
		{1, big.NewRat(3,2)},
		{2, big.NewRat(7,5)},
		{3, big.NewRat(17,12)},
		{4, big.NewRat(41,29)},
	}
	for _, s := range stim {
		var a sqrt2Approximator
    for i := 0; i < s.iterations; i++ {
			a.Next()
		}
		if !reflect.DeepEqual(a.Rat(), s.rat) {
			t.Errorf("rat = %v, want %v", a.Rat(), s.rat)
		}
	}
}

func TestFindFirstIteration(t *testing.T) {
	if r := find(8); r != 1 {
		t.Errorf("r = %v, want %v", r, 1)
	}
}

func TestFindIterationCount1000(t *testing.T) {
	expected := 153
	if r := find(1000); r != expected {
		t.Errorf("r = %v, want %v", r, expected)
	}
}
