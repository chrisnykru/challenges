/*

Combinatoric selections

There are exactly ten ways of selecting three from five, 12345:
123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

In combinatorics, we use the notation, C(5,3) = 10.

In general,
C(n,r) = n!/(r!*(n-r)!) where r <= n, n! = n*(n-1)...*3*2*1, and 0!=1.

It is not until n = 23, that a value exceeds one-million: C(23,10) = 1144066.

How many, not necessarily distinct, values of C(n,r), for 1 <= n <= 100,
are greater than one-million?

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"math/big"
)

// note: unsimplified C(23,10) overflows in uint domain
func choose(n, k uint) *big.Int {
	if k > n {
		panic("k > n")
	}
	// C(n,k) = n!/(k!*(n-k)!)
	//        = Product(n,n-1,...,k+1)/(n-k)!
	p1 := big.NewInt(int64(n))
	for i := n - 1; i > k; i-- {
		p1.Mul(p1, big.NewInt(int64(i)))
	}
	return p1.Div(p1, misc.Factorial(int64(n-k)))
}

// brute force method; not exploiting any symmetry
func count() uint {
	over1e6 := uint(0)
	int_1e6 := big.NewInt(1000000)

	for n := uint(1); n <= 100; n++ {
		for k := uint(1); k <= n; k++ {
			c := choose(n, k)
			if c.Cmp(int_1e6) == 1 {
				// C > 1e6
				over1e6++
			}
		}
	}
	return over1e6
}

func main() {
	fmt.Println(count())
}
