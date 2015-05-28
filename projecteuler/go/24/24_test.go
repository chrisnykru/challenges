package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	perm := perm1e6()
	expected := []int{2, 7, 8, 3, 9, 1, 5, 4, 6, 0}
	if !reflect.DeepEqual(perm, expected) {
		t.Errorf("perm = %v, want %v", perm, expected)
	}
}
