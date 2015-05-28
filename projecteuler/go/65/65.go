/*

Convergents of e

The square root of 2 can be written as an infinite continued fraction.

√2 = 1 + 1 / (2 + 1 / (2 + 1 / (2 + 1 / (2 + ...))))

√2 = 1 +              1
         ----------------------------
				 2 +            1
				     ------------------------
						 2 +          1
						     --------------------
								 2 +        1
								     ----------------
										 2 + ...


The infinite continued fraction can be written, √2 = [1;(2)], (2) indicates
that 2 repeats ad infinitum. In a similar way, √23 = [4;(1,3,1,8)].

It turns out that the sequence of partial values of continued fractions for
square roots provide the best rational approximations. Let us consider the
convergents for √2.


1 + 1 / 2 = 3 / 2


1 +     1      = 7/5
    ---------
		2 + 1 / 2


1 +        1        = 17/12
     -------------
		 2 +     1
		     ---------
				 2 + 1/2


1 +        1            = 41/29
    ---------------
    2 +      1
		    -----------
				2 +    1
				    -------
						2 + 1/2


Hence the sequence of the first ten convergents for √2 are:
1, 3/2, 7/5, 17/12, 41/29, 99/70, 239/169, 577/408, 1393/985, 3363/2378, ...

What is most surprising is that the important mathematical constant,
e = [2; 1,2,1, 1,4,1, 1,6,1 , ... , 1,2k,1, ...].

The first ten terms in the sequence of convergents for e are:
2, 3, 8/3, 11/4, 19/7, 87/32, 106/39, 193/71, 1264/465, 1457/536, ...
The sum of digits in the numerator of the 10th convergent is 1+4+5+7=17.

Find the sum of digits in the numerator of the 100th convergent of the
continued fraction for e.

*/
package main

import (
	"fmt"
	"math/big"
)

func solve(atConvergent, stopConvergent int64) *big.Rat {
	// not sure what to call 'x' here
	var x int64
	switch (atConvergent - 2) % 3 {
	case 0:
		x = 1
	case 1:
		x = 2 * ((atConvergent-2)/3 + 1)
	case 2:
		x = 1
	default:
		panic("impossible!")
	}
	fmt.Printf("  at=%v\t\tx = %v\n", atConvergent, x)
	denom := big.NewRat(x, 1)
	if atConvergent == stopConvergent {
		return denom
	}
	denom.Add(denom, solve(atConvergent+1, stopConvergent))
	denom.Inv(denom)
	return denom
}

func eApproximation(convergent int64) *big.Rat {
	r := big.NewRat(2, 1)
	r.Add(r, solve(2, convergent))
	return r
}

func sumOfNumeratorDigits(convergent int64) int {
	r := eApproximation(convergent)
	num := r.Num()
	numStr := num.String()
	sum := 0
	for i := 0; i < len(numStr); i++ {
		if numStr[i] < '0' || numStr[i] > '9' {
			panic("bad digit")
		}
		sum += int(numStr[i]) - '0'
	}
	return sum
}

func main() {
	fmt.Println(sumOfNumeratorDigits(100))
}
