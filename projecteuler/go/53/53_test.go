package main

import (
	"testing"
)

func TestChoose(t *testing.T) {
	c5_3 := choose(5, 3).Int64()
	if c5_3 != 10 {
		t.Errorf("c(5,3)=%v, want %v", c5_3, 10)
	}

	c23_10 := choose(23, 10).Int64()
	if c23_10 != 1144066 {
		t.Errorf("c(23,10)=%v, want %v", c23_10, 1144066)
	}
}

func Test(t *testing.T) {
	num := count()
	if num != 4075 {
		t.Errorf("num = %v, want %v", num, 4075)
	}
}
