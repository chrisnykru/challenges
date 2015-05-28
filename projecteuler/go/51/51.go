/*

Prime digit replacements

By replacing the 1st digit of *3, it turns out that six of the nine possible
values: 13, 23, 43, 53, 73, and 83, are all prime.

By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit
number is the first example having seven primes among the ten generated numbers,
yielding the family: 56003, 56113, 56333, 56443, 56663, 56773, and 56993.
Consequently 56003, being the first member of this family, is the smallest
prime with this property.

Find the smallest prime which, by replacing part of the number (not necessarily
adjacent digits) with the same digit, is part of an eight prime value family.

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"github.com/jbarham/primegen.go"
	"strconv"
	"math"
)

// "part of the number" construed to mean not-nothing and not-everything
func partialMasks(x int) (num []uint8, m [][]int) {
	s := strconv.Itoa(x)

	limit := uint64(math.Pow(2, float64(len(s))))
	// skip first--mask == "nothing"
	// skip last--mask == "everything"
	for i := uint64(1); i < limit-1; i++ {
		mask, _ := misc.Uint64ToFixedWidthDigits(i, 2, len(s))
		m = append(m, mask)
	}
	return []uint8(s), m
}

func applyMask(digits []uint8, mask []int, overlayDigit uint8) int {
	tmp := append([]uint8{}, digits...)

	for i := 0; i < len(tmp); i++ {
		if mask[i] == 1 {
			tmp[i] = '0' + overlayDigit
		}
	}

	x, err := strconv.ParseInt(string(tmp), 10, 0)
	if err != nil {
		panic(err)
	}
	return int(x)
}

func npvf(n int) []int {
	if n < 1 {
		panic("n < 1")
	}

	primeMap := make(map[int]bool)

	gen := primegen.New()
	for {
		// uint64
		prime := int(gen.Next())
		primeMap[prime] = true

		////fmt.Printf("prime = %v\n", prime)

		digits, masks := partialMasks(prime)

		for _, mask := range masks {

			family := make([]int, 0)
			for i := uint8(0); i < 10; i++ {
				if mask[0] == 1 && i == 0 {
					// skip... would turn '13' into '03'-->3...
					continue
				}
				x2 := applyMask(digits, mask, i)

				// XXX if x2 is prime.... this means we'll detect family at the END
				if _, ok := primeMap[x2]; ok {
					family = append(family, x2)
				}
			}
			if len(family) >= n {
				return family
			}

		}
	}

	panic("!!")
}

func main() {
	n := 8 // 8
	family := npvf(n)
	fmt.Printf("npvf(%v) = %v\n", n, family)
}
