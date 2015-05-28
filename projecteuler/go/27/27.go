/*

Quadratic primes

Euler published the remarkable quadratic formula:
  n^2 + n + 41

It turns out that the formula will produce 40 primes for the consecutive values
n = 0 to 39. However, when n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible
by 41, and certainly when n = 41, 41^2 + 41 + 41 is clearly divisible by 41.

Using computers, the incredible formula  n^2 - 79n + 1601 was discovered, which
produces 80 primes for the consecutive values n = 0 to 79. The product of the
coefficients, -79 and 1601, is -126479.

Considering quadratics of the form:
  n^2 + an + b, where |a|<1000 and |b|<1000
  where |n| is the modulus/absolute value of n
  e.g. |11| = 11 and |-4| = 4

Find the product of the coefficients, a and b, for the quadratic expression that
produces the maximum number of primes for consecutive values of n,
starting with n = 0.

*/
package main

import (
	"fmt"
	"github.com/jbarham/primegen.go"
)

const (
	// note: approx 79k primes from [1,1e6]
	// we don't know how many consecutive values of n will generate primes,
	// so we just pick 1e6 as upper bound for our prime check lookup table.
	primeLutCeil = 1000000
)

var (
	primeMap map[uint64]bool
)

func init() {
	primeMap = make(map[uint64]bool)
	gen := primegen.New()
	for {
		prime := gen.Next()
		if prime >= primeLutCeil {
			break
		}
		primeMap[prime] = false // arbitrary
	}
}

func primeyQuadratic(a, b, n int) int {
	ans := n*n + a*n + b
	return ans
}

func primesForConsecutiveN(a, b int) (primes []uint64) {
	primes = make([]uint64, 0)

	for n := 0; ; n++ {
		px := primeyQuadratic(a, b, n)
		// hmm: are negative values prime? let's say no
		if px < 0 {
			break
		}
		// is px prime?
		if px >= primeLutCeil {
			// expensive check...
			panic("TBD")
		}
		if _, ok := primeMap[uint64(px)]; !ok {
			// not consecutive
			break
		}
		primes = append(primes, uint64(px))
	}
	return primes
}

func optimalAB() (optimalA, optimalB int, primes []uint64) {
	for a := -999; a <= 999; a++ {
		for b := -999; b <= 999; b++ {
			p := primesForConsecutiveN(a, b)
			if len(p) > len(primes) {
				primes = p
				optimalA = a
				optimalB = b
			}
		}
	}
	return
}

func main() {
	a, b, primes := optimalAB()
	fmt.Println(primes)
	fmt.Println("len(primes):", len(primes))
	fmt.Printf("a=%v, b=%v, a*b=%v\n", a, b, a*b)
}
