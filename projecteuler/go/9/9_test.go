package main

import (
	"testing"
)

func Test(t *testing.T) {
	triplet := pythagoreanTripletOf1000()
	abc := triplet[0] * triplet[1] * triplet[2]
	if abc != 31875000 {
		t.Errorf("abc = %v, want %v", abc, 31875000)
	}
}
