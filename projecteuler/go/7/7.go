/*

10001st prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
we can see that the 6th prime is 13.

What is the 10 001st prime number?

*/
package main

import (
	"fmt"
	"github.com/jbarham/primegen.go"
)

func main() {
	nth_prime := 10001

	gen := primegen.New()
	for i := 0; i < nth_prime; i++ {
		prime := gen.Next()
		fmt.Println(prime)
	}
}
