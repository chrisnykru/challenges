/*

Large non-Mersenne prime

The first known prime found to exceed one million digits was discovered in 1999,
and is a Mersenne prime of the form 2^6972593-1; it contains exactly 2,098,960
digits. Subsequently other Mersenne primes, of the form 2^p-1, have been found
which contain more digits.

However, in 2004 there was found a massive non-Mersenne prime which contains
2,357,207 digits: 28433 * 2^7830457+1.

Find the last ten digits of this prime number.

*/
package main

import (
	"fmt"
	"math/big"
)

func fast() *big.Int {
	// XXX assumption: mod(x*y,z) === mod(mod(x,z) * mod(y,z), z)

	x := new(big.Int)
	x.Exp(big.NewInt(2), big.NewInt(7830457), big.NewInt(10000000000))
	x.Mul(x, big.NewInt(28433))
	x.Add(x, big.NewInt(1))
	return x
}

func slow() *big.Int {
	x := big.NewInt(28433)
	y := new(big.Int).Exp(big.NewInt(2), big.NewInt(7830457), nil)
	x.Mul(x, y)
	x.Add(x, big.NewInt(1))
	return x
}

func main() {
	s := fast().String()
	//fmt.Println(s)
	fmt.Println("last 10:", s[len(s)-10:])
}
