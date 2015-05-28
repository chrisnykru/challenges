package main

import (
	"testing"
)

func Test(t *testing.T) {
	var c counter
	c.Sol(200, 0)
	if c.ways != 73682 {
		t.Errorf("ways = %v, want %v", c.ways, 73682)
	}
}
