/*

Permuted multiples

It can be seen that the number, 125874, and its double, 251748, contain exactly
the same digits, but in a different order.

Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x,
contain the same digits.

*/
package main

import (
	"fmt"
)

type digitsLut struct {
	lut map[uint][10]uint
}

func newDigitsLut() *digitsLut {
	return &digitsLut{
		lut: make(map[uint][10]uint),
	}
}

func (lut *digitsLut) Digits(x uint) [10]uint {
	if d, ok := lut.lut[x]; ok {
		return d
	}
	var digits [10]uint
	for {
		digit := x % 10
		digits[digit]++
		x /= 10
		if x == 0 {
			break
		}
	}
	return digits
}

func find() uint {
	lut := newDigitsLut()
	for i := uint(1); ; i++ {
		d := lut.Digits(i)

		ok := true
		for mul := uint(2); mul <= 6 && ok; mul++ {
			if d2 := lut.Digits(i * mul); d2 != d {
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
	panic("unexpected")
}

func main() {
	fmt.Println(find())
}
