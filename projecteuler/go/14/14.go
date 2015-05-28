/*

Longest Collatz sequence

The following iterative sequence is defined for the set of positive integers:
  n -> n/2 (n is even)
  n -> 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
  13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10
terms. Although it has not been proved yet (Collatz Problem), it is thought that
all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.

*/
package main

import (
	"fmt"
)

func sequenceNext(x int64) int64 {
	if x%2 == 0 {
		// x is even
		return x / 2
	}
	// x is odd
	return 3*x + 1
}

// sequence ends at 1
func sequenceLength(x int64) (seqlen int64) {
	for seqlen = int64(1); x != 1; seqlen++ {
		x = sequenceNext(x)
	}
	return
}

func startWithLongestSequence(startMax int64) (maxLenStartIdx, maxLen int64) {
	for i := int64(1); i < startMax; i++ {
		if seqlen := sequenceLength(i); seqlen > maxLen {
			maxLen = seqlen
			maxLenStartIdx = i
		}
	}
	return
}

func main() {
	start, maxLen := startWithLongestSequence(1000000)
	fmt.Printf("max sequence length of %v @ start=%v\n", maxLen, start)
}
