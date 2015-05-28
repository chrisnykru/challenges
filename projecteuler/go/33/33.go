/*

Digit cancelling fractions

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in
attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is
correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than
one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms,
find the value of the denominator.

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
)

// notation: n/d
func reduce(n, d uint) (nreduced, dreduced uint) {
	ddiv := misc.Divisors(uint64(d))
	biggestDiv := uint(1)
	for div := range misc.Divisors(uint64(n)) {
		if _, ok := ddiv[div]; ok {
			if uint(div) > biggestDiv {
				biggestDiv = uint(div)
			}
		}
	}
	return n / biggestDiv, d / biggestDiv
}

func solve(cf [][2]uint) (n, d uint) {
	n, d = 1, 1
	for i := range cf {
		n *= cf[i][0]
		d *= cf[i][1]
	}
	return reduce(n, d)
}

func findCuriousFractions() (cf [][2]uint) {
	for n := uint(10); n <= 99; n++ {
		for d := n + 1; d <= 99; d++ {
			// skip "trivial examples"
			if (n%10 == 0) && (d%10 == 0) {
				continue
			}

			ndigits := []uint{n / 10, n % 10}
			ddigits := []uint{d / 10, d % 10}
			commonDigits := make(map[uint]bool)
			for i := range ndigits {
				for j := range ddigits {
					if ndigits[i] == ddigits[j] {
						commonDigits[ndigits[i]] = false // arbitrary
					}
				}
			}
			// skip if numerator and denominator have no common digits
			if len(commonDigits) == 0 {
				continue
			}

			nr, dr := reduce(n, d)

			// do incorrect cancellations
			curious := false
			for cdigit := range commonDigits {
				var naive_n, naive_d uint
				// naively cancel out common digit
				for _, digit := range ndigits {
					if digit != cdigit {
						naive_n = digit
					}
				}
				for _, digit := range ddigits {
					if digit != cdigit {
						naive_d = digit
					}
				}
				// compare reduced naive fraction to legit reduced fraction
				naive_nr, naive_dr := reduce(naive_n, naive_d)
				if naive_nr == nr && naive_dr == dr {
					curious = true
				}
			}
			if curious {
				cf = append(cf, [2]uint{nr, dr})
			}
		}
	}
	return
}

func main() {
	cf := findCuriousFractions()
	fmt.Println(cf)
	fmt.Println(solve(cf))
}
