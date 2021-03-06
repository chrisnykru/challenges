/*

Lexicographic permutations

A permutation is an ordered arrangement of objects. For example, 3124 is one
possible permutation of the digits 1, 2, 3 and 4. If all of the permutations
are listed numerically or alphabetically, we call it lexicographic order.

The lexicographic permutations of 0, 1 and 2 are:
  012   021   102   120   201   210

What is the millionth lexicographic permutation of the
digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/permgen"
)

func perm1e6() []int {
	gen := permgen.New([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	var perm []int
	for i := 0; i < 1000000; i++ {
		perm, _ = gen.Next()
	}
	return perm
}

func main() {
	fmt.Println(perm1e6())
}
