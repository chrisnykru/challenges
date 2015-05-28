/*

Integer right triangles

If p is the perimeter of a right angle triangle with integral length sides,
{a,b,c}, there are exactly three solutions for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p <= 1000, is the number of solutions maximised?

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
)

func solutions(p int) (sol [][3]int) {
	// p = a+b+c
	// right triangle: a**2 + b**2 = c**2
	// c = (a**2 + b**2)**.5
	//
	// p = a + b + (a**2 + b**2)**.5
	//
	// let a = 1..2..3..
	// solve for b

	for a := 1; a < p; a++ {
		// filter out similar triangles
		for b := a; b >= 1; b-- {
			cSquared := a*a + b*b
			c := int(misc.ISqrt(uint64(cSquared)))
			if c*c != cSquared {
				// ignore
				continue
			}
			// a,b too small?
			if a+b+c < p {
				break
			}
			if a+b+c == p {
				sol = append(sol, [3]int{a, b, c})
			}

		}
	}
	return
}

func maxSol() (solCount, p int) {
	maxSolCount, maxSolP := 0, 0
	for p := 8; p <= 1000; p++ {
		sol := solutions(p)
		if len(sol) > maxSolCount {
			maxSolCount = len(sol)
			maxSolP = p
		}
	}
	return maxSolCount, maxSolP
}

func main() {
	//fmt.Println("p=120:", solutions(120))
	solCount, p := maxSol()
	fmt.Printf("solCount=%v for p=%v\n", solCount, p)
}
