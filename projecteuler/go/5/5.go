/*

Smallest multiple

2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of
the numbers from 1 to 20?

*/
package main

import (
	"fmt"
)

// [1,max] inclusive
func smallestNumberEvenlyDivisible(max int64) int64 {
	if max < 1 {
		panic("bad max")
	}

	// given a number n evenly divisible by each number in [floor(max/2)+1, max],
	// it is evenly divisible by each number in [1,floor(max/2)] as well.
	//
	// e.g., let max=20
	//   if n is divisible by each number in [11,20], it is clearly
	//   divisible by each number in [1,10]
	//
	// e.g., let max=19
	//   if n is divisible by each number in [floor(19/2)+1, 19] = [10,19]
	//   it is also divisible by [1,9]
	lowestUniqueDivisor := (max / 2) + 1

	// find first value >= max that is evenly divisible by lowestUniqueDivisor
	startVal := max
	for {
		if startVal%lowestUniqueDivisor == 0 {
			break
		}
		startVal++
	}

	for i := startVal; ; i += lowestUniqueDivisor {
		ok := true
		for j := lowestUniqueDivisor + 1; j <= max; j++ {
			if i%j != 0 {
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
	panic("!!!")
}

func main() {
	// 20! fits in int64 domain.. but not by much
	fmt.Println(smallestNumberEvenlyDivisible(20))
}
