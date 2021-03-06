/*

Highly divisible triangular number

The sequence of triangle numbers is generated by adding the natural numbers.
So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28.

The first ten terms would be:
1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

Let us list the factors of the first seven triangle numbers:

 1: 1
 3: 1,3
 6: 1,2,3,6
10: 1,2,5,10
15: 1,3,5,15
21: 1,3,7,21
28: 1,2,4,7,14,28

We can see that 28 is the first triangle number to have over five divisors.

What is the value of the first triangle number to have over five hundred divisors?

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
)

func triangleNumGen() func() int64 {
	a := int64(1)
	sum := int64(0)
	return func() int64 {
		sum += a
		a += 1
		return sum
	}
}

func divisorCount(x int64) int {
	if x < 0 {
		panic("x < 0")
	}
	sqrt_x := int64(misc.ISqrt(uint64(x)))

	divisors := make(map[int64]bool)
	divisors[1] = false // arbitrary
	divisors[x] = false

	for i := int64(2); i <= sqrt_x; i++ {
		if x%i == 0 {
			divisors[i] = false
			divisors[x/i] = false
		}
	}

	return len(divisors)
}

func findTriNumWithOver500Divisors() int64 {
	gen := triangleNumGen()
	for {
		tnum := gen()
		if count := divisorCount(tnum); count > 500 {
			return tnum
		}
	}
	panic("!!!")
}

func main() {
	fmt.Println(">500 divisors:", findTriNumWithOver500Divisors())
}
