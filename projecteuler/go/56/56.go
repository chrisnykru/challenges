/*

Powerful digit sum

A googol (10**100) is a massive number: one followed by one-hundred zeros;
100**100 is almost unimaginably large: one followed by two-hundred zeros.
Despite their size, the sum of the digits in each number is only 1.

Considering natural numbers of the form, ab, where a, b < 100,
what is the maximum digital sum?

*/
package main

import (
	"fmt"
	"math/big"
)

func digitalSum(x *big.Int) int64 {
	xs := x.String()
	sum := int64(0)
	for _, c := range xs {
		if c < '0' || c > '9' {
			panic(fmt.Sprintf("bad %c in %q", c, xs))
		}
		sum += int64(c - '0')
	}
	return sum
}

func find() int64 {
	var maxSum int64
	c := new(big.Int)

	for a := int64(1); a < 100; a++ {
		for b := int64(1); b < 100; b++ {
			c.Exp(big.NewInt(a), big.NewInt(b), nil)
			sum := digitalSum(c)
			if sum > maxSum {
				//fmt.Printf("%s has digital sum %d\n", c, sum)
				maxSum = sum
			}
		}
	}
	return maxSum
}

func main() {
	fmt.Println(find())
}
