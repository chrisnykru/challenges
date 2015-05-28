/*

Square digit chains

A number chain is created by continuously adding the square of the digits in a
number to form a new number until it has been seen before.

For example,
44 --> 32 --> 13 --> 10 --> 1 --> 1
85 --> 89 --> 145 --> 42 --> 20 --> 4 --> 16 --> 37 --> 58 --> 89

Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop.
What is most amazing is that EVERY starting number will eventually arrive at 1 or 89.

How many starting numbers below ten million will arrive at 89?

*/
package main

import (
	"fmt"
)

func sumSquareDigits(x int) (sum int) {
  for {
		d := x % 10
		sum += (d * d)
		x /= 10
		if x == 0 {
			break
		}
	}
	return
}

func count89() (count int) {
	lut := map[int]int{1: 1, 89: 89}
	for i := 1; i < 10e6; i++ {
    x := i
		for {
			//fmt.Printf("i=%v  x = %v\n", i, x)
			if result, ok := lut[x]; ok {
				if result == 89 {
					//fmt.Printf("i=%v arrives at 89", i)
					count++
				}
				lut[i] = result
				break
			}
			x = sumSquareDigits(x)
		}
	}
	//fmt.Printf("lut = %v\n", lut)
	return
}

func main() {
	fmt.Println(count89())
}
