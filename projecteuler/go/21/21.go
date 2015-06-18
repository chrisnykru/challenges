/*

Amicable numbers

Let d(n) be defined as the sum of proper divisors of n (numbers less than n
which divide evenly into n).  If d(a) = b and d(b) = a, where a != b,
then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are:
  1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
  Therefore d(220) = 284

The proper divisors of 284 are:
  1, 2, 4, 71 and 142;
  so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
)

func d(x uint64) (sum uint64) {
	for div := range misc.ProperDivisors(x) {
		sum += div
	}
	return
}

func sumOfAmicableNumbersUnder(max uint64) uint64 {
	amicablePairs := make(map[uint64]uint64)
	for a := uint64(1); a < max; a++ {
		d_a := d(a)
		// a != b
		if a == d_a {
			continue
		}

		// not amicable?
		if d_b := d(d_a); d_b != a {
			continue
		}

		// amicable pair
		if _, ok := amicablePairs[a]; !ok {
			if _, ok = amicablePairs[d_a]; !ok {
				fmt.Printf("a = %v   d_a = %v\n", a, d_a)
				amicablePairs[a] = d_a
			}
		}
	}

	sum := uint64(0)
	for d_a, d_b := range amicablePairs {
		//fmt.Printf("d_a = %v, d_b = %v\n", d_a, d_b)
		sum += d_a + d_b
	}
	return sum
}

func main() {
	sum := sumOfAmicableNumbersUnder(10000)
	fmt.Println("sum of amicable nums", sum)
}
