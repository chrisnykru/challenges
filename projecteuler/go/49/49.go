/*

Prime permutations

The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases
by 3330, is unusual in two ways: (i) each of the three terms are prime, and,
(ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes,
exhibiting this property, but there is one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three terms in this sequence?

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"github.com/zeroshirts/challenges/projecteuler/go/permgen"
	"github.com/jbarham/primegen.go"
	"sort"
)

// x is 4 digits (1000 <= x <= 9999)
func permutations(x int64) []uint64 {
	digits := make([]int, 0, 4)
	for {
		digits = append(digits, int(x%10))
		x /= 10
		if x == 0 {
			break
		}
	}

	pg := permgen.New(digits)
	perm := make([]uint64, 0)
	for {
		p, last := pg.Next()
		x, _ := misc.DigitsToUint64(p, 10)
		perm = append(perm, x)
		if last {
			break
		}
	}
	return perm
}

// e.g., [1487 1847 4817 4871 7481 7841 8147 8741]
//       sequence=[1487 4817 8147]
func arithSeq3(x []int) (seq []int) {
	sort.Ints(x)

	m := make(map[int]bool)
	for i := range x {
		m[x[i]] = false // arbitrary
	}

	for i := 0; i < len(x)-1; i++ {
		for j := i + 1; j < len(x)-1; j++ {
			delta := x[j] - x[i]
			if _, ok := m[x[j]+delta]; ok {
				return []int{x[i], x[j], x[j] + delta}
			}
		}
	}
	return []int{}
}

func sequence() []int {
	primeGen := primegen.New()

	// skip past 1487 (first 3 term arithmetic sequence)
	// generate primes in range [1488,9999]
	primeGen.SkipTo(1488)
	primeMap := make(map[uint64]bool)
	for {
		p := primeGen.Next()
		if p > 9999 {
			break
		}
		primeMap[p] = false // arbitrary
	}

	// iteration order undefined
	for p, _ := range primeMap {
		perm := permutations(int64(p))
		primePerm := make([]int, 0, len(perm))
		for i := range perm {
			if _, ok := primeMap[uint64(perm[i])]; ok {
				primePerm = append(primePerm, int(perm[i]))
			}
		}

		// arithmetric sequence = constant delta
		// primePerm may be 8 elements long, but somewhere inside we have have a 3 arith-seq
		seq := arithSeq3(primePerm)
		if len(seq) > 0 {
			return seq
		}
	}
	panic("unexpected")
}

func concatSeq(x []int) string {
	var s string
	for i := range x {
		s += fmt.Sprintf("%d", x[i])
	}
	return s
}

func main() {
	fmt.Println(concatSeq(sequence()))
}
