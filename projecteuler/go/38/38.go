/*

Pandigital multiples

Take the number 192 and multiply it by each of 1, 2, and 3:
  192 * 1 = 192
  192 * 2 = 384
  192 * 3 = 576

By concatenating each product we get the 1 to 9 pandigital, 192384576.
We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, 5,
giving the pandigital, 918273645, which is the concatenated product
of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the
concatenated product of an integer with (1,2, ... , n) where n > 1?

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"github.com/zeroshirts/challenges/projecteuler/go/pandigital"
	"github.com/zeroshirts/challenges/projecteuler/go/permgen"
)

func concatProductPandigital_1to9(x uint) (num uint64, n uint, ok bool) {
	var p pandigital.Pandigital
	// x * 1
	p.AddDigits(x)
	// x * 2..3..4..
	for n = 2; p.Vet(1, 9) == pandigital.Indeterminate; n++ {
		p.AddDigits(x * n)
	}

	if p.Vet(1, 9) == pandigital.Valid && n > 2 {
		ok = true
	}

	return p.Uint64(), n - 1, ok
}

// generates 1,2,3,4,5,6,7,8,12,13,14,15,...,78,...,5678,...,12345678
func genPandigitalFragments_1to8(c chan<- []int) {
	lut := make(map[uint64]bool)

	gen := permgen.New([]int{1, 2, 3, 4, 5, 6, 7, 8})

	for {
		perm, last := gen.Next()
		for i := 1; i <= len(perm); i++ {
			fragment := make([]int, i)
			copy(fragment, perm[0:i])

			x, _ := misc.DigitsToUint64(fragment, 10)
			if _, ok := lut[x]; !ok {
				// new pandigital fragment/whole
				lut[x] = false // arbitrary
				c <- fragment
			}
		}

		if last {
			break
		}
	}
	close(c)
	return
}

// largest 1 to 9 pandigital 9-digit number that can be formed as the
// concatenated product of an integer with (1,2, ..., n) where n > 1?
func largest() uint64 {
	pandigital := make(chan []int, 16)

	// concatendated product of 9 and (1,2,3,4,5) = 918273645
	// first digit must be 9
	go genPandigitalFragments_1to8(pandigital)

	highNum := uint64(0)
	for x := range pandigital {
		x2 := make([]int, len(x)+1)
		// first digit hardcoded to 9
		x2[0] = 9
		copy(x2[1:], x)
		x3, err := misc.DigitsToUint64(x2, 10)
		if err != nil {
			panic(err)
		}

		num, n, ok := concatProductPandigital_1to9(uint(x3))
		if ok && num > highNum {
			fmt.Printf("x3 = %v, num = %v, n = %v\n", x3, num, n)
			highNum = num
		}
	}
	return highNum
}

func main() {
	fmt.Println(largest())
}
