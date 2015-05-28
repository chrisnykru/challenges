package fibonacci

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	gen := NewGen()
	seq := make([]uint64, 10)
	for i := range seq {
		seq[i] = gen()
	}

	expected := []uint64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	if !reflect.DeepEqual(seq, expected) {
		t.Errorf("seq = %v, want %v", seq, expected)
	}
}

func TestBigInt(t *testing.T) {
	gen := NewBigIntGen()
	seq := make([]int64, 10)
	for i := range seq {
		seq[i] = gen().Int64()
	}

	expected := []int64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	if !reflect.DeepEqual(seq, expected) {
		t.Errorf("seq = %v, want %v", seq, expected)
	}
}
