/*

Double-case palindromes

The decimal number, 585 = 1001001001 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"math"
)

/*

Worry about base10 first, then filter out non-base-2-palindromes.

[1,999999], 6 digits max

1 digit..  x    // x==[0..9], except 0 won't work, is leading zero
  short: x
2 digits.. xx
  short: x
3 digits.. xyx  // y can equal x
  short: xy
4 digits.. xyyx
  short: xy
5 digits.. xyzyx
  short: xyz
6 digits.. xyzzyx
  short: xyz

*/

func genBase10Palindromes(maxDigits int, out chan<- int) {
	for numDigits := 1; numDigits <= maxDigits; numDigits++ {
		shortNumDigits := (numDigits + 1) / 2

		// generate all short digits; right-most is least significant
		limit := math.Pow(10, float64(shortNumDigits))
		for i := 0; i < int(limit); i++ {
			digits, _ := misc.Uint64ToDigits(uint64(i), 10)

			// derive palindrome
			// 1 digit:  x   --> x
			// 2 digits: x   --> xx
			// 3 digits: xy  --> xyx
			// 4 digits: xy  --> xyyx
			// 5 digits: xyz --> xyzyx
			// 6 digits: xyz --> xyzzyx
			mirrored := make([]int, numDigits)
			for i := range digits {
				offset := shortNumDigits - len(digits) + i
				mirrored[offset] = digits[i]
				mirrored[len(mirrored)-1-offset] = digits[i]
			}

			// output
			x, _ := misc.DigitsToUint64(mirrored, 10)
			out <- int(x)
		}
	}
	close(out)
}

func base2Palindrome(x int) bool {
	s := fmt.Sprintf("%b", x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func base2_and_10_palindromes() (sum int) {
	c := make(chan int, 4)
	// problem space < 1e6; thus, 6 digits max
	go genBase10Palindromes(6, c)

	for x := range c {
		if base2Palindrome(x) {
			////fmt.Println("base10 and base2 palindrome:", x)
			sum += x
		}
	}
	return
}

func main() {
	fmt.Println(base2_and_10_palindromes())
}
