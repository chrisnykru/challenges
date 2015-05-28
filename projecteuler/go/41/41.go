/*

Pandigital prime

We shall say that an n-digit number is pandigital if it makes use of all the
digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and
is also prime.

What is the largest n-digit pandigital prime that exists?

*/
package main

import (
	"fmt"
	"github.com/jbarham/primegen.go"
)

const (
	largest9digitPandigitalPrime = 987654321
)

func nPandigital(x uint64) (n uint, ok bool) {
	var digitCount [10]int8

	for x > 0 {
		digit := x % 10
		if digit == 0 {
			return 0, false
		}
		if digitCount[digit] > 0 {
			return 0, false
		}
		digitCount[digit]++
		x /= 10
	}

	//fmt.Printf("digitCount=%v\n", digitCount)

	// at this point, digit counts are 0 or 1
	// hack digitCount[0] to make this easy
	digitCount[0] = 1
	for i := 1; i <= 9; i++ {
		if digitCount[i-1] == 0 && digitCount[i] != 0 {
			return 0, false
		}
		if digitCount[i] == 1 {
			n = uint(i)
		}
	}
	ok = true
	return
}

// 1 <= 'n' <= 9
func find() uint64 {
	//ndigitPrime := uint64(0)
	gen := primegen.New()

	var (
		x, y uint64
	)
	for x <= largest9digitPandigitalPrime {
		x = gen.Next()

		if n, ok := nPandigital(x); ok {
			y = x
			fmt.Printf("x = %v, n = %v\n", x, n)
		}
	}
	return y
}

func main() {
	fmt.Println(find())
}
