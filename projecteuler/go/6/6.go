/*

Sum square difference

The sum of the squares of the first ten natural numbers is,
12 + 22 + ... + 102 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.

*/
package main

import (
	"fmt"
)

func naive(n int) int {
	sumofsquares := 0
	for i := 1; i <= n; i++ {
		sumofsquares += i * i
	}
	squareofsums := 0
	for i := 1; i <= n; i++ {
		squareofsums += i
	}
	squareofsums *= squareofsums
	return sumofsquares - squareofsums
}

func smarter(n int) int {
	return -1*n*n*n*n/4 - n*n*n/6 + n*n/4 + n/6
}

func main() {
	// Sum(0,n) i   = n(n+1)/2
	// Sum(0,n) i^2 = n^3/3 + n^2/2 + n/6
	// TODO: derivation of above

	// Answer = n^3/3 + n^2/2 + n/6 - (n(n+1)/2)^2
	//        = -n^4/4 - n^3/6 + n^2/4 + n/6
	// XXX annoying, we must do abs() when reporting answer...

	fmt.Printf("naive(100) = %v\n", naive(100))
	fmt.Printf("smarter(100) = %v\n", smarter(100))
}
