/*

Sub-string divisibility

The number, 1406357289, is a 0 to 9 pandigital number because it is made up of
each of the digits 0 to 9 in some order, but it also has a rather interesting
sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on.
In this way, we note the following:

d2d3d4=406 is divisible by 2
d3d4d5=063 is divisible by 3
d4d5d6=635 is divisible by 5
d5d6d7=357 is divisible by 7
d6d7d8=572 is divisible by 11
d7d8d9=728 is divisible by 13
d8d9d10=289 is divisible by 17

Find the sum of all 0 to 9 pandigital numbers with this property.

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"github.com/zeroshirts/challenges/projecteuler/go/permgen"
)

func pandigital_0to9(c chan<- []int) {
	gen := permgen.New([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for {
		perm, last := gen.Next()
		// ignore numbers with leading zeros (not considered pandigital?)
		if perm[0] == 0 {
			continue
		}
		c <- perm
		if last {
			break
		}
	}
	close(c)
	return
}

func specialPropertyPassFilter(in <-chan []int, out chan<- uint64) {
	for {
		perm, ok := <-in
		if !ok {
			break
		}
		// d2d3d4 divisible by 2
		subset := perm[1:4]
    if x, _ := misc.DigitsToUint64(subset, 10); x%2 != 0 {
			continue
		}
		// d3d4d5 divisible by 3
		subset = perm[2:5]
		if x, _ := misc.DigitsToUint64(subset, 10); x%3 != 0 {
			continue
		}
		// d4d5d6 divisible by 5
		subset = perm[3:6]
		if x, _ := misc.DigitsToUint64(subset, 10); x%5 != 0 {
			continue
		}
		// d5d6d7 divisible by 7
		subset = perm[4:7]
		if x, _ := misc.DigitsToUint64(subset, 10); x%7 != 0 {
			continue
		}
		// d6d7d8 divisible by 11
		subset = perm[5:8]
		if x, _ := misc.DigitsToUint64(subset, 10); x%11 != 0 {
			continue
		}
		// d7d8d9 divisible by 13
		subset = perm[6:9]
		if x, _ := misc.DigitsToUint64(subset, 10); x%13 != 0 {
			continue
		}
		// d8d9d10 divisible by 17
		subset = perm[7:10]
		if x, _ := misc.DigitsToUint64(subset, 10); x%17 != 0 {
			continue
		}
		x, _ := misc.DigitsToUint64(perm, 10)
		out <- x
	}
	close(out)
}

func sumSpecialPandigitals() (sum uint64) {
	pandigital := make(chan []int, 16)
	property := make(chan uint64, 16)

	go pandigital_0to9(pandigital)
	go specialPropertyPassFilter(pandigital, property)

	for num := range property {
		//fmt.Printf("got %v\n", num)
		sum += num
	}
	return
}

func main() {
	fmt.Println(sumSpecialPandigitals())
}
