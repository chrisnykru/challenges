/*

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
)

func primeFactors(x uint64) (pf []uint64) {
	for div := range misc.Divisors(x) {
		if misc.IsPrime(div) {
			pf = append(pf, div)
		}
	}
	return pf
}

func largest() uint64 {
	pf := primeFactors(600851475143)
	fmt.Println("prime factors:", pf)
	var x uint64
	for i := range pf {
		if pf[i] > x {
			x = pf[i]
		}
	}
	return x
}

func main() {
	// XXX ??? Why is it that prime factors, when multiplied together, yield target?
	fmt.Println(largest())
}
