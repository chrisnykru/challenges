package misc

import (
	"errors"
	"math"
)

var (
	ErrBadRadix = errors.New("radix < 2")
	ErrBadDigit = errors.New("bad digit")
	ErrOverflow = errors.New("overflow")
)

// Big-Endian; Error if radix < 2
func Uint64ToDigits(x uint64, radix int) ([]int, error) {
	if radix < 2 {
		return nil, ErrBadRadix
	}
	var d []int
	for {
		d = append(d, int(x%uint64(radix)))
		x /= uint64(radix)
		if x == 0 {
			break
		}
	}
	// reverse digits
	for i := 0; i < len(d)/2; i++ {
		d[i], d[len(d)-1-i] = d[len(d)-1-i], d[i]
	}
	return d, nil
}

// Big-Endian; Error if radix < 2 or width < log_radix(x)
func Uint64ToFixedWidthDigits(x uint64, radix, width int) ([]int, error) {
	d, err := Uint64ToDigits(x, radix)
	if err != nil {
		return nil, err
	}
	if len(d) > width {
		return nil, ErrOverflow
	}
	d2 := make([]int, width)
	copy(d2[len(d2)-len(d):], d)
	return d2, nil
}

// Big-Endian; Error if radix < 2, d contains invalid digit, or conversion overflows
func DigitsToUint64(d []int, radix int) (uint64, error) {
	if radix < 2 {
		return 0, ErrBadRadix
	}
	sum := uint64(0)
	pow := uint64(1)
	for i := len(d) - 1; i >= 0; i-- {
		if d[i] < 0 || d[i] >= radix {
			return 0, ErrBadDigit
		}
		if uint64(d[i]) > math.MaxUint64/pow {
			return 0, ErrOverflow
		}
		v := uint64(d[i]) * pow
		if sum > math.MaxInt64-v {
			return 0, ErrOverflow
		}
		sum += v
		if pow > math.MaxInt64/uint64(radix) {
			return 0, ErrOverflow
		}
		pow *= uint64(radix)
	}
	return sum, nil
}
