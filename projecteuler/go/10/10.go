/*

Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

*/
package main

import (
	"fmt"
	"github.com/jbarham/primegen.go"
	"math/big"
)

// < ceil
func sumOfPrimesBelow(ceil int64) *big.Int {
	sum := big.NewInt(0)

	gen := primegen.New()
	for {
		// unsafe conversion
		prime := int64(gen.Next())
		if prime >= ceil {
			break
		}
		sum = sum.Add(sum, big.NewInt(prime))
	}

	return sum
}

func main() {
	ceil := int64(2000000)
	sum := sumOfPrimesBelow(ceil)
	fmt.Printf("sum of primes below %v = %s\n", ceil, sum)
}
