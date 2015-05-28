/*

Circular primes

The number, 197, is called a circular prime because all rotations
of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100:
  2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?

*/
package main

import (
	"fmt"
	"github.com/jbarham/primegen.go"
	"strconv"
)

func leftRotateStr(s string) string {
	if len(s) < 2 {
		return s
	}
	s2 := s[1:] + string(s[0])
	return s2
}

func circularPrimesBelow1e6() []int64 {
	gen := primegen.New()

	primeMap := make(map[int64]bool)
	for {
		prime := int64(gen.Next())
		if prime >= 1e6 {
			break
		}
		primeMap[prime] = false // arbitrary
	}

	circularPrimes := make([]int64, 0)
	for prime := range primeMap {
		primeStr := strconv.FormatInt(prime, 10)
		circular := true

		for i := 1; i < len(primeStr); i++ {
			primeStr = leftRotateStr(primeStr)
			// convert primeStr back to integer
			rprime, _ := strconv.ParseInt(primeStr, 10, 64)
			// rotated prime guaranteed to be < 1e6 (but not necessarily under an arbitrary max value)
			if _, ok := primeMap[rprime]; !ok {
				circular = false
				break
			}
		}
		if circular {
			circularPrimes = append(circularPrimes, prime)
		}
	}
	return circularPrimes
}

func main() {
	c := circularPrimesBelow1e6()
	fmt.Println("circular primes:", c)
	fmt.Println("count:", len(c))
}
