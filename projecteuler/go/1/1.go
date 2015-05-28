/*

Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

*/
package main

import (
	"fmt"
)

// find multiples of 3 and 5 from domain [1, max)
func sumOfMultiples(max int) int {
	sum := 0
	for i := 1; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

// find multiples of 3 and 5 from domain [1, max)
func fasterSumOfMultiples(max int) int {
	sum := 0
	for i := 3; i < max; i += 3 {
		sum += i
	}
	for i := 5; i < max; i += 5 {
		sum += i
	}
	for i := 15; i < max; i += 15 {
		sum -= i
	}
	return sum
}

func main() {
	sum := sumOfMultiples(1000)
	fmt.Printf("multiples <1000: %v\n", sum)
}
