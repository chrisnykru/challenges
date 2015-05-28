/*

Diophantine equation

Consider quadratic Diophantine equations of the form:

x**2 – Dy**2 = 1

For example, when D=13, the minimal solution in x is
649**2 – 13×180**2 = 1.

It can be assumed that there are no solutions in positive integers
when D is square.

By finding minimal solutions in x for D = {2, 3, 5, 6, 7},
we obtain the following:

3**2 – 2×2**2 = 1
2**2 – 3×1**2 = 1
9**2 – 5×4**2 = 1
5**2 – 6×2**2 = 1
8**2 – 7×3**2 = 1

Hence, by considering minimal solutions in x for D ≤ 7, the largest x is obtained
when D=5.

Find the value of D ≤ 1000 in minimal solutions of x for which the largest
value of x is obtained.

*/
package main

import (
	"fmt"
)

type squareLut struct {
	lut map[int64]bool
	x, xx int64
}

func newSquareLut() *squareLut {
	return &squareLut{
		lut: make(map[int64]bool),
	}
}

func (lut *squareLut) IsSquare(xx int64) bool {
	if xx <= lut.xx {
		return lut.lut[xx]
	}
	for {
		lut.x++
		lut.xx = lut.x * lut.x
		lut.lut[lut.xx] = true
		if lut.xx >= xx {
			return lut.lut[xx]
		}
	}
}

/*

Given: D is not square

x**2 - D * y**2 = 1


Square numbers are always EVEN, except for 1 (is 1 square?)



For D=13, the minimal solution in x is 649**2 - 13×180**2 = 1.
x=649, y=180





Given that D is not square

SquareNum * SquareNum always == SquareNum?
NonSquareNum * SquareValue = NonSquareNum
+ 1... maybe it's square again



64 is square
64 = 16 * 4, both of which are square





*/


func solve(d int64) int64 {
	x := int64(1)
	y := int64(1)

	for {
		rhs := d * y * y + 1

		stop := false
		for !stop {
			lhs := x * x
			if lhs < rhs {
				// increment x, recalcuate lhs
				x++
				continue
			} else if lhs > rhs {
				// increment y, ...
				y++
				stop = true
				break
			} else {
				// lhs == rhs
				fmt.Printf("for D = %v, x = %v, y = %v\n", d, x, y)
				return x
			}
		}
	}

	panic("X!")
}

func search(dmax int64) int64 {
	lut := newSquareLut()
	xmax := int64(-1)
	d_with_largest_minimal_x := int64(-1)

	for d := int64(1); d <= dmax; d++ {
		if lut.IsSquare(d) {
			continue
		}
		fmt.Printf("d = %v\n", d)

		x := solve(d)
		if x > xmax {
			xmax = x
			d_with_largest_minimal_x = d
		}
	}
	return d_with_largest_minimal_x
}

func main() {
	fmt.Println(search(1000))
}
