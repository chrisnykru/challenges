package fibonacci

import (
	"math/big"
)

// derived from golang fibonacci closure demo
func NewGen() func() uint64 {
	var a, b uint64
	a, b = 0, 1
	return func() uint64 {
		a, b = b, a+b
		return a
	}
}

func NewBigIntGen() func() *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)

	return func() *big.Int {
		b2 := new(big.Int).Add(big.NewInt(0), b)
		b = b.Add(b, a)
		a = b2
		return a
	}
}
