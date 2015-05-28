package misc

import (
	"math/big"
)

// Panics if x < 0
func Factorial(x int64) *big.Int {
	if x < 0 {
		panic("x < 0")
	}
	f := big.NewInt(1)
	for i := x; i > 1; i-- {
		f.Mul(f, big.NewInt(i))
	}
	return f
}
