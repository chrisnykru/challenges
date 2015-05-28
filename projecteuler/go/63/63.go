/*

Powerful digit counts

The 5-digit number, 16807=7**5, is also a fifth power. Similarly, the 9-digit
number, 134217728=8**9, is a ninth power.

How many n-digit positive integers exist which are also an nth power?

*/
package main

import (
	"fmt"
	"math/big"
)

func solve() int {
	targetCount := 0
	for n := int64(1); ; n++ {
		nBig := big.NewInt(n)
		foundAnything := false
		for base := int64(1); ; base++ {
			powResult := new(big.Int).Exp(big.NewInt(base), nBig, nil)
			numDigits := int64(len(powResult.String()))

			if numDigits == n {
				fmt.Printf("n=%v\t%v**%v=%v (numDigits=%v)\n", n, base, n, powResult, numDigits)
				foundAnything = true
				targetCount++
			} else if numDigits > n {
				// too big
				break
			}
		}
		if !foundAnything {
			// stop heuristic
			break
		}
	}
	return targetCount
}

func main() {
	fmt.Println(solve())
}
