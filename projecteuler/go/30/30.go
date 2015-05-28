/*

Digit fifth powers

Surprisingly there are only three numbers that can be written as the sum of
fourth powers of their digits:

1634 = 1^4 + 6^4 + 3^4 + 4^4
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4

As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of
fifth powers of their digits.

*/
package main

import (
	"fmt"
)

func sumOfDigitsNthPower(x, power uint) (sum uint) {
	if power < 1 {
		panic("bad power")
	}

	// process digits right-to-left
	for x > 0 {
		digit := x % 10
		digitToNth := digit
		for i := uint(2); i <= power; i++ {
			digitToNth *= digit
		}
		sum += digitToNth
		x /= 10
	}
	return
}

func writeableAsSumOfNthPower(power uint) (numbers []uint) {
	// need at least 2 digits, so start at 10
	// TODO XXX what is cutoff point??????
	cutoff := uint(10000000)
	for i := uint(10); i < cutoff; i++ {
		sum := sumOfDigitsNthPower(i, power)
		if sum == i {
			numbers = append(numbers, i)
		}
	}
	return
}

func sum(x []uint) uint {
	s := uint(0)
	for i := range x {
		s += x[i]
	}
	return s
}

func main() {
	num := writeableAsSumOfNthPower(5)
	fmt.Println(num)
	fmt.Println(sum(num))
}
