/*

Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91  99.

Find the largest palindrome made from the product of two 3-digit numbers.

*/
package main

import (
	"fmt"
	"strconv"
)

// base 10 palindrome
// e.g., 99 = true, 101 = true, 201 = false
func IsPalindrome(x int64) bool {
	s := strconv.FormatInt(x, 10)
	// len(s) >= 1

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

type productPalindrome struct {
	i, j, ij int64
}

// [1,max]
func largestProductPalindrome(max int64) (pp productPalindrome) {
	// search from high to low
	for i := max; i > 0; i-- {
		for j := i; j > 0; j-- {
			ij := i * j
			if IsPalindrome(ij) {
				if ij > pp.ij {
					pp = productPalindrome{i, j, ij}
				}
			}
		}
	}
	return
}

func main() {
	fmt.Printf("found %#v\n", largestProductPalindrome(999))
}
