package misc

import (
	"math"
)

const (
	// IEEE-754 float64: 53 bits of precision
	Float64ExactIntMax = (1 << 53) - 1
)

// Uses IEEE-754 float64 and math.Sqrt() for n<=FloatMaxExactInt
func ISqrt(n uint64) uint64 {
	// fast path
	if n <= Float64ExactIntMax {
		return uint64(math.Floor(math.Sqrt(float64(n))))
	}
	// slow path; prevent divide by zero
	if n == 0 || n == math.MaxUint64 {
		return n >> 32
	}
	x := n
	for {
		y := (x + n/x) / 2
		if y >= x {
			break
		}
		x = y
	}
	return x
}
