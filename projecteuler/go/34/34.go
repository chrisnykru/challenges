/*

Digit factorials

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial
of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.

*/
package main

import (
	"fmt"
)

// 123 --> [3 2 1]
func reversedDigits(x uint) []uint {
	d := make([]uint, 0)
	for {
		digit := x % 10
		d = append(d, digit)
		x /= 10
		if x == 0 {
			break
		}
	}
	return d
}

func factorialDigitLut() []uint {
	lut := make([]uint, 10)
	lut[0] = 1
	lut[1] = 1

	f := uint(1)
	for i := uint(2); i < uint(10); i++ {
		f *= i
		lut[i] = f
	}
	return lut
}

/*

need 2 digits, so minimum = 10

9! = 362880

7 digit sum of factorials max = 7 * 362880 = 2540160
7 digit max = 9999999
2540160 < 9999999

--> stop iterating at 2540160

*/
const (
	max_num = 2540160
)

func factorialOfDigits() (all []uint) {
	fByNum := factorialDigitLut()

	for i := uint(10); i <= max_num; i++ {
		d := reversedDigits(i)
		sum := uint(0)
		for _, digit := range d {
			sum += fByNum[digit]
		}
		if sum == i {
			fmt.Printf("i = %v, sum = %v\n", i, sum)
			all = append(all, i)
		}
	}
	return
}

func sumAll(all []uint) uint64 {
	s := uint64(0)
	for _, v := range all {
		s += uint64(v)
	}
	return s
}

func main() {
	all := factorialOfDigits()
	fmt.Println("all:", all)
	fmt.Println("sum:", sumAll(all))
}
