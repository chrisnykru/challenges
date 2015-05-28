/*

Spiral primes

Starting with 1 and spiralling anticlockwise in the following way, a square
spiral with side length 7 is formed.

37 36 35 34 33 32 31
38 17 16 15 14 13 30
39 18  5  4  3 12 29
40 19  6  1  2 11 28
41 20  7  8  9 10 27
42 21 22 23 24 25 26
43 44 45 46 47 48 49

It is interesting to note that the odd squares lie along the bottom
right diagonal, but what is more interesting is that 8 out of the 13
numbers lying along both diagonals are prime; that is, a ratio of 8/13 ≈ 62%.

If one complete new layer is wrapped around the spiral above, a square spiral
with side length 9 will be formed. If this process is continued, what is the
side length of the square spiral for which the ratio of primes along both
diagonals first falls below 10%?

---------------------------------------------------

side lengths:
1
3
5
7
9...

side=1
  center = 1

side=3
  top-right = 1 + 3-1 = 3
  top-left = 3+3-1 = 5
	bottom-left = 5+3-1 = 7
	bottom-right = 7+3-1 = 9
	shift-out!

side=5
  top-right = 9+5-1 = 13 
	top-left = 13+5-1 = 17
	bottom-left = 17+5-1 = 21
	bottom-right = 21+5-1 = 25
	shift-out!

side=7
  top-right = 25+7-1 = 31
	top-left: 31+7-1 = 37
	etc...




*/
package main

import (
	"fmt"
	"github.com/jbarham/primegen.go"
)

type primeChecker struct {
	gen   *primegen.Primegen
	lastp  uint64
}

func newPrimeChecker() *primeChecker {
	return &primeChecker{
		gen: primegen.New(),
		lastp: 0,
	}
}

// Returns 0 if x is not prime, 1 if x is prime
func (pc *primeChecker) Check(x uint64) int {
  if x < pc.lastp {
		// Expectation is that x increases with each invocation
		pc.gen = primegen.New()
		pc.lastp = 0
		return pc.Check(x)
	}
	for {
		pc.lastp = pc.gen.Next()
		if pc.lastp == x {
			return 1
		} else if pc.lastp > x {
			return 0
		}
	}
}


func solve() uint64 {
	pc := newPrimeChecker()
	lastcell := uint64(1) // center
	diagcount := 1 // center
	primediagcount := 0

	for side := uint64(3);; side += 2 {
		topright := lastcell + side-1
		topleft := topright + side-1
		bottomleft := topleft + side-1
		bottomright := bottomleft + side-1
		lastcell = bottomright
		diagcount += 4

		// determine whether new diags are prime
		primediagcount += pc.Check(topright)
		primediagcount += pc.Check(topleft)
		primediagcount += pc.Check(bottomleft)
		primediagcount += pc.Check(bottomright)

		if primediagcount * 100 / diagcount < 10 {
			return side
		}
	}
	panic("!")
}

func main() {
	fmt.Println(solve())
}
