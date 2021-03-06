/*

Even Fibonacci numbers

Each new term in the Fibonacci sequence is generated by adding the previous two
terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/fibonacci"
)

func sumFibUpTo(max uint64) (sum uint64) {
	fgen := fibonacci.NewGen()

	// skip first '1' term to be consistent with problem sequence
	_ = fgen()

	// find sum of even fib numbers up to max of 4,000,000
	for {
		v := fgen()
		if v > max {
			break
		}
		// filter out odd numbers
		if v%2 != 0 {
			continue
		}
		//fmt.Printf("term = %v\n", v)
		sum += v
	}
	return sum
}

func main() {
	fmt.Printf("sum = %d\n", sumFibUpTo(4000000))
}
