package main

import (
	"testing"
)

func Test(t *testing.T) {
	oc := smallest()
	if oc != 5777 {
		t.Errorf("oc = %v, want %v", oc, 5777)
	}
}
