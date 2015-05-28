package main

import "testing"

func Test(t *testing.T) {
	r := sumOfSixCyclic()
	if r != 28684 {
		t.Errorf("r = %v, want %v", r, 28684)
	}
}
