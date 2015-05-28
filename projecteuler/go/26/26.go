/*

Reciprocal cycles

A unit fraction contains 1 in the numerator. The decimal representation of the
unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1

Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be
seen that 1/7 has a 6-digit recurring cycle.

Find the value of d<1000 for which 1/d contains the longest recurring cycle
in its decimal fraction part.

*/
package main

import (
	"fmt"
	"math/big"
)

func cycle(d uint) uint {
	// assumption: worst-case cycle length for 1/d is d-1
	// pow==10**d sufficient to determine second worst-case 
	//            cycle length of floor(d/2)
	// 1/13 --> 1*10**13/13 = 769230769230; cycle len = 6
	pow := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(d)), nil)

	numerator := new(big.Int).Mul(big.NewInt(1), pow)
	m := new(big.Int)
	_, m = new(big.Int).DivMod(numerator, big.NewInt(int64(d)), m)
	if m.Cmp(big.NewInt(0)) == 0 {
		return 0
	}

	// eliminates any non-repeating prefix
	// e.g., 1/6=.166666; after elimination: 666
	numerator = new(big.Int).Mul(m, pow)
	x2, m := new(big.Int).DivMod(numerator, big.NewInt(int64(d)), m)

	x2s := x2.String()
	for sublen := 1; sublen <= len(x2s)/2; sublen++ {
		match := 0
		for i := 0; i < sublen; i++ {
			if x2s[i] == x2s[i+sublen] {
				match++
			}
		}
		if match == sublen {
			return uint(sublen)
		}
	}

	return uint(len(x2s) - 1)
}

func longestCycle(dMax uint) (denom, length uint) {
	for d := uint(2); d < dMax; d++ {
		cycleLen := cycle(d)
		if cycleLen > length {
			denom = d
			length = cycleLen
		}
	}
	return
}

func main() {
	d, dlen := longestCycle(1000)
	fmt.Printf("d=%v, length=%v\n", d, dlen)
}
