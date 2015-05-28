/*

Consecutive prime sum

The prime 41, can be written as the sum of six consecutive primes:
  41 = 2 + 3 + 5 + 7 + 11 + 13

This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime,
contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?

*/
package main

import (
	"fmt"
	"github.com/jbarham/primegen.go"
)

func primeWithMostConsecutiveTerms(max uint64) (prime uint64, numTerms int) {
	gen := primegen.New()
	primes := make([]uint64, 0)
	primeMap := make(map[uint64]bool)
	for {
		prime := gen.Next()
		if prime >= uint64(max) {
			break
		}
		primes = append(primes, prime)
		primeMap[prime] = false // arbitrary
	}

	for i := range primes {
		sum, count := primes[i], 1
		for j := i + 1; j < len(primes); j++ {
			sum += primes[j]
			// stop summing at max
			if sum >= max {
				break
			}

			count++
			if _, ok := primeMap[sum]; ok {
				if count > numTerms {
					numTerms = count
					prime = sum
				}
			}
		}
	}
	return
}

func main() {
	/*

	  primes must be consecutive

	  approx 79k primes below 1e6

	*/

	prime, numTerms := primeWithMostConsecutiveTerms(1000000)
	fmt.Printf("prime = %v, numTerms = %v\n", prime, numTerms)
}
