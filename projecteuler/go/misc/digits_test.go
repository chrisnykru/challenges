package misc

import (
	"reflect"
	"testing"
)

func TestUint64ToDigits(t *testing.T) {
	stim := []struct {
		x     uint64
		radix int
		d     []int
		err   error
	}{
		{8, 2, []int{1, 0, 0, 0}, nil},
		{0, 2, []int{0}, nil},
		{4, 20, []int{4}, nil},
		{0, 0, nil, ErrBadRadix},
	}
	for _, s := range stim {
		d, err := Uint64ToDigits(s.x, s.radix)
		if !reflect.DeepEqual(d, s.d) {
			t.Errorf("Uint64ToDigits(%v,%v) = %v, want %v", s.x, s.radix, d, s.d)
		}
		if err != s.err {
			t.Errorf("Uint64ToDigits(%v,%v) err = %v, want %v", s.x, s.radix, err, s.err)
		}
	}
}

func TestUint64ToFixedWidthDigits(t *testing.T) {
	stim := []struct {
		x            uint64
		radix, width int
		d            []int
		err          error
	}{
		{8, 2, 4, []int{1, 0, 0, 0}, nil},
		{0, 2, 1, []int{0}, nil},
		{4, 20, 1, []int{4}, nil},
		{8, 2, 3, nil, ErrOverflow},
	}
	for _, s := range stim {
		d, err := Uint64ToFixedWidthDigits(s.x, s.radix, s.width)
		if !reflect.DeepEqual(d, s.d) {
			t.Errorf("Int64ToFixedWidthDigits(%v,%v,%v) = %v, want %v", s.x, s.radix, s.width, d, s.d)
		}
		if err != s.err {
			t.Errorf("Int64ToFixedWidthDigits(%v,%v,%v) err = %v, want %v", s.x, s.radix, s.width, err, s.err)
		}
	}
}

func TestDigitsToUInt64(t *testing.T) {
	stim := []struct {
		d     []int
		radix int
		x     uint64
		err   error
	}{
		{[]int{0, 0, 8}, 10, 8, nil}, // zero padding
		{[]int{1, 0, 0, 0}, 2, 8, nil},
		{[]int{0}, 2, 0, nil},
		{[]int{2}, -2, 0, ErrBadRadix},
		{[]int{2}, 2, 0, ErrBadDigit},
		{[]int{9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9}, 10, 0, ErrOverflow},
	}
	for _, s := range stim {
    x, err := DigitsToUint64(s.d, s.radix)
		if !reflect.DeepEqual(x, s.x) {
			t.Errorf("DigitsToInt64(%v,%v) = %v, want %v", s.d, s.radix, x, s.x)
		}
		if err != s.err {
			t.Errorf("DigitsToInt64(%v,%v) err = %v, want %v", s.d, s.radix, err, s.err)
		}
	}
}
