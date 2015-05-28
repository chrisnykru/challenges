/*

Pandigital products

We shall say that an n-digit number is pandigital if it makes use of all the
digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1
through 5 pandigital.

The product 7254 is unusual, as the identity, 39 * 186 = 7254, containing
multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity
can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only
include it once in your sum.

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"github.com/zeroshirts/challenges/projecteuler/go/pandigital"
)

// let mmp = concat(multiplicand, multiplier, product)
// check whether there exists any mmp that is is 1-9 pandigital
func identityPandigital_1to9(product uint) bool {
	// generate all multiplicand, multiplier pairs
	// we ignore divisor 1, as '1*x=x' has duplicate digits (i.e. not pandigital)
	for div := range misc.ProperDivisors(uint64(product)) {
		if div == 1 {
			continue
		}
		var p pandigital.Pandigital
		p.AddDigits(product)
		p.AddDigits(uint(div))
		p.AddDigits(product / uint(div))
		if p.Vet(1, 9) == pandigital.Valid {
			return true
		}
	}
	return false

	/*

	  // skip first divisor--it's always 1
	  for i := 1; i < len(div); i++ {
	    var p pandigital.Pandigital
	    p.AddDigits(product)
	    p.AddDigits(uint(div[i]))
	    p.AddDigits(product / uint(div[i]))
	    if p.Vet(1,9) == pandigital.Valid {
	      return true
	    }
	  }
	  return false
	*/
}

// TODO Performance:
//   We could speed up pandigital1to9 factorization by using permutations,
//   slicing each permutation into (multiplicand, multiplier, product) tuples,
//   and checking whether multiplicand * multiplier = product holds.
func pandigital_1to9_identityProducts() (p []uint) {
	// set product lower bound to 4 digits (i.e. 1234)
	// 3 digit product requires 6 digits across multiplicand and multiplier--won't work
	//
	// set upper bound to 98765; leaves 4 digits across multiplicand and multiplier
	// e.g., 12345 > 6 * 789... but it's less than 10x too high
	for i := uint(1234); i < 12345; i++ {
		if identityPandigital_1to9(i) {
			p = append(p, i)
		}
	}
	return
}

func sum(x []uint) (s uint) {
	for i := range x {
		s += x[i]
	}
	return
}

func main() {
	p := pandigital_1to9_identityProducts()
	fmt.Println(p)
	fmt.Println(sum(p))
}
