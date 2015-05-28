/*

Number spiral diagonals

Starting with the number 1 and moving to the right in a clockwise direction
a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral
formed in the same way?

*/
package main

import (
	"fmt"
)

func sumAllDiags(edge uint) (sum uint64) {
	if edge%2 == 0 {
		panic("edge length must be odd")
	}

	// diagonal terms:
	// 4 terms... in a 7x7
	// 3 terms in a 5x5
	// 2 terms in a 3x3
	// 1 term in a 1x1
	maxTerm := (edge + 1) / 2
	// term 1 is 1; 1^2 = 1, which we reflect into sum
	sum = 1

	for term := uint(2); term <= maxTerm; term++ {
		// top-right diagonal
		termVal := 2*term - 1
		termVal *= termVal
		sum += uint64(termVal)

		// top-left diagonal
		counterClockwiseDiagDelta := 2 * (term - 1)
		termVal = termVal - counterClockwiseDiagDelta
		sum += uint64(termVal)

		// bottom-left diagonal
		termVal = termVal - counterClockwiseDiagDelta
		sum += uint64(termVal)

		// bottom-right diagonal
		termVal = termVal - counterClockwiseDiagDelta
		sum += uint64(termVal)
	}
	return
}

func main() {
	sum := sumAllDiags(1001)
	fmt.Println(sum)
}
