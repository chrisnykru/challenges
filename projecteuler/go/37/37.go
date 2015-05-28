/*

Truncatable primes

The number 3797 has an interesting property. Being prime itself, it is possible
to continuously remove digits from left to right, and remain prime at each stage: 
3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right
and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

*/
package main

import (
	"fmt"
	"github.com/jbarham/primegen.go"
	"strconv"
)

func truncatablePrimes() []uint64 {
	gen := primegen.New()
	primeMap := make(map[uint64]bool)
	tprimes := make([]uint64, 0)
	for len(tprimes) < 11 {
		prime := gen.Next()
		// memoize prime
		primeMap[prime] = false // arbitrary

		// single digit primes are not considered truncatable
		if prime < 10 {
			continue
		}

		primeStr := strconv.FormatInt(int64(prime), 10)

		// right-to-left truncation
		tmpStr := primeStr[0 : len(primeStr)-1]
		truncatable := true
		for len(tmpStr) > 0 {
			// convert back to integer
			tprime, _ := strconv.ParseInt(tmpStr, 10, 64)
			if _, ok := primeMap[uint64(tprime)]; !ok {
				truncatable = false
				break
			}
			tmpStr = tmpStr[0 : len(tmpStr)-1]
		}
		if !truncatable {
			continue
		}

		// left-to-right truncation
		// invariant: len(primestr) > 1
		tmpStr = primeStr[1:len(primeStr)]
		for len(tmpStr) > 0 {
			// convert back to integer
			tprime, _ := strconv.ParseInt(tmpStr, 10, 64)
			if _, ok := primeMap[uint64(tprime)]; !ok {
				truncatable = false
				break
			}
			tmpStr = tmpStr[1:len(tmpStr)]
		}

		if truncatable {
			tprimes = append(tprimes, prime)
		}
	}
	return tprimes
}

func sum(a []uint64) uint64 {
	sum := uint64(0)
	for i := range a {
		sum += a[i]
	}
	return sum
}

func main() {
	tprimes := truncatablePrimes()
	fmt.Println(tprimes)
	fmt.Println(sum(tprimes))
}
