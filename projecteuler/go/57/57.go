/*

Square root convergents

It is possible to show that the square root of two can be expressed as an
infinite continued fraction.

sqrt(2) = 1 + 1/(2 + 1/(2 + 1/(2 + ... ))) = 1.414213...

By expanding this for the first four iterations, we get:

1 + 1/2 = 3/2 = 1.5
1 + 1/(2 + 1/2) = 7/5 = 1.4
1 + 1/(2 + 1/(2 + 1/2)) = 17/12 = 1.41666...
1 + 1/(2 + 1/(2 + 1/(2 + 1/2))) = 41/29 = 1.41379...

The next three expansions are 99/70, 239/169, and 577/408, but the eighth expansion,
1393/985, is the first example where the number of digits in the numerator exceeds
the number of digits in the denominator.

In the first one-thousand expansions, how many fractions contain a numerator with
more digits than denominator?

-----------------------------------------

1+...

1/2 = 1/2
1/(2 + 1/2) = 2/5
1/(2 + 1/(2 + 1/2)) = 5/12
1/(2 + 1/(2 + 1/(2 + 1/2))) = 12/27

*/
package main

import (
	"fmt"
	"math/big"
)

type sqrt2Approximator struct {
  frac *big.Rat
}

func (a *sqrt2Approximator) Next() {
	if a.frac == nil {
		// init
		a.frac = big.NewRat(1,2)
		return
	}
	denom := big.NewRat(2,1)
	denom.Add(denom, a.frac)
	a.frac = denom.Inv(denom)
}

func (a *sqrt2Approximator) Rat() *big.Rat {
	x := big.NewRat(1,1)
	return x.Add(x, a.frac)
}

func (a *sqrt2Approximator) String() string {
	return a.Rat().String()
}

func (a *sqrt2Approximator) IsNumMoreDigitsThanDenom() bool {
	x := big.NewRat(1,1)
	x.Add(x, a.frac)
	numStr := x.Num().String()
	denomStr := x.Denom().String()
	return len(numStr) > len(denomStr)
}

func find(expansionMax int) int {
	count := 0
	var a sqrt2Approximator
	for i := 1; i <= expansionMax; i++ {
		a.Next()
		if a.IsNumMoreDigitsThanDenom() {
			//fmt.Printf("!! i=%v  %s\n", i, a.String())
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(find(1000))
}
