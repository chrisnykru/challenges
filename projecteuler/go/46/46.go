/*

Goldbach's other conjecture

It was proposed by Christian Goldbach that every odd composite number can be
written as the sum of a prime and twice a square.

9 = 7 + 2*1^2
15 = 7 + 2*2^2
21 = 3 + 2*3^2
25 = 7 + 2*3^2
27 = 19 + 2*2^2
33 = 31 + 2*1^2

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a prime
and twice a square?

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"github.com/jbarham/primegen.go"
)

const (
	// smallest possible prime; odd composite != even prime + 2*x**2 for all x
	pmin = 3
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

func smallest() uint64 {
	lut := newPrimeLut()

	// 9 is known odd composite satisfying conjecture
	for oc := uint64(9); ; oc += 2 /* keep odd */ {
		if lut.isPrime(oc) {
			// primes are not composites
			continue
		}

		// Strategy:
		//   Assuming the answer isn't massive, the primeLut makes checking
		//   whether a number is prime efficient. Given the growth rate of
		//   squared term s, it's faster to sweep (oc,s) and derive p
		//   than it is to sweep (oc,p) and derive s.
		ok := false
		for s := misc.ISqrt((oc - pmin) / 2); s >= 1; s-- {
			p := oc - 2*s*s
			if p >= oc {
				// underflow
				continue
			}
			if lut.isPrime(p) {
				ok = true
				break
			}
		}
		if !ok {
			return oc
		}
	}
	panic("unexpected")
}

func main() {
	fmt.Println(smallest())
}
