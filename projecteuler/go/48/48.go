/*

Self powers

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.

*/
package main

import (
	"fmt"
	"math/big"
)

func fast() *big.Int {
	last10mod := big.NewInt(10000000000)
	sum := big.NewInt(0)
	for i := int64(1); i <= 1000; i++ {
		big_i := big.NewInt(i)
		term := new(big.Int).Exp(big_i, big_i, last10mod)
		sum.Add(sum, term)
	}
	return sum
}

func main() {
	s := fast().String()
	fmt.Println(s)
	last10 := s[len(s)-10:]
	fmt.Println("last 10:", last10)
}
