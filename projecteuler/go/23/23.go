/*

Non-abundant sums

A perfect number is a number for which the sum of its proper divisors is exactly
equal to the number. For example, the sum of the proper divisors of 28 would be
1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n
and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number
that can be written as the sum of two abundant numbers is 24. By mathematical
analysis, it can be shown that all integers greater than 28123 can be written as
the sum of two abundant numbers. However, this upper limit cannot be reduced any
further by analysis even though it is known that the greatest number that cannot
be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of
two abundant numbers.

*/
package main

import (
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
)

func SumProperDivisors(x uint64) (sum uint64) {
	for div := range misc.ProperDivisors(x) {
		sum += div
	}
	return
}

func Sum() int {
	// everything over 28123 can be written as sum of two abundant numbers
	abundant := make([]int, 0, 28123)
	for i := 12; i <= 28123; i++ {
		if sum := int(SumProperDivisors(uint64(i))); sum > i {
			abundant = append(abundant, i)
		}
	}

	// a+b == b+a
	m := make(map[int]bool)
	for i := range abundant {
		for j := i; j < len(abundant) && abundant[i]+abundant[j] <= 28123; j++ {
			m[abundant[i]+abundant[j]] = false //arbitrary
		}
	}

	sum := 0
	for i := 1; i <= 28123; i++ {
		if _, ok := m[i]; !ok {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(Sum())
}
