/*

Lattice paths

Starting in the top left corner of a 22 grid, there are 6 routes
(without backtracking) to the bottom right corner.

How many routes are there through a 2020 grid?

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"math/big"
)

// square grid
// no backtracking allowed
func numRoutes(size int64) int64 {
	// note: r1,r2,d1,d2 is same route as r2,r1,d2,d1

	n := size + size
	k := size

	nfactorial := misc.Factorial(n)
	kfactorial := misc.Factorial(k)
	nMinuskfactorial := misc.Factorial(n - k)

	C := new(big.Int).Div(nfactorial, new(big.Int).Mul(kfactorial, nMinuskfactorial))
	return C.Int64()
}

func main() {
	fmt.Println(numRoutes(20))
}
