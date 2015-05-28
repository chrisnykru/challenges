/*

Distinct primes factors

The first two consecutive numbers to have two distinct prime factors are:

14 = 2 * 7
15 = 3 * 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2**2 * 7 * 23
645 = 3 * 5 * 43
646 = 2 * 17 * 19.

Find the first four consecutive integers to have four distinct primes factors.
What is the first of these numbers?

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"github.com/jbarham/primegen.go"
)

type primeLut struct {
	lut     map[uint64]bool
	highest uint64
	gen     *primegen.Primegen
}

func newPrimeLut() *primeLut {
	lut := &primeLut{
		lut: make(map[uint64]bool),
		gen: primegen.New(),
	}
	return lut
}

func (lut *primeLut) isPrime(x uint64) bool {
	if x <= lut.highest {
		_, ok := lut.lut[x]
		return ok
	}
	// expand LUT
	var p uint64
	for p < x {
		p = lut.gen.Next()
		lut.lut[p] = false // arbitrary
	}
	lut.highest = p
	return p == x && p != 0
}

func primeFactors(lut *primeLut, x uint64) []uint64 {
	divisors := misc.Divisors(x)
	pf := make([]uint64, 0, len(divisors))
	for div := range divisors {
		if lut.isPrime(div) {
			pf = append(pf, div)
		}
	}
	return pf
}

func find() uint64 {
	lut := newPrimeLut()
	// start off at 644;
	// perhaps 647 is 4th consecutive integer to have 4 _distinct_ prime factors
	consecutive := 0
	for i := uint64(644); ; i++ {

		// new thinking...take 644...get prime factors: [2 7 23]
		// start with: 2
		//   644/2 = 322; prime factors: [2 7 23]
		//   etc.. recursively do this
		// actually, lazy approach works here
		pf := primeFactors(lut, i)
		if len(pf) < 4 {
			consecutive = 0
			continue
		}
		consecutive++
		//fmt.Printf("i=%v  pf=%v\n", i, pf)

		if consecutive == 4 {
			// first entry
			return i - 3
		}
	}
	panic("unexpected")
}

func main() {
	fmt.Println(find())
}
