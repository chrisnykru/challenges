/*

Triangle, pentagonal, and hexagonal

Triangle, pentagonal, and hexagonal numbers are generated by the following
formulae:

Triangle
  Tn=n(n+1)/2
  1, 3, 6, 10, 15, ...

Pentagonal
  Pn=n(3n-1)/2
  1, 5, 12, 22, 35, ...

Hexagonal
  Hn=n(2n-1)
  1, 6, 15, 28, 45, ...

It can be verified that T285 = P165 = H143 = 40755.

Find the next triangle number that is also pentagonal and hexagonal.


*/
package main

import (
	"fmt"
)

func t(n uint) uint {
	return (n * (n + 1)) / 2
}

func p(n uint) uint {
	return (n * (3*n - 1)) / 2
}

func h(n uint) uint {
	return n * (2*n - 1)
}

func find() uint {
	pinverse := make(map[uint]uint)
	hinverse := make(map[uint]uint)

	// populate for n <= 285
	for n := uint(1); n <= 285; n++ {
		pinverse[p(n)] = n
		hinverse[h(n)] = n
	}

	// for n > 285, p(n) > t(n), and h(n) > t(n)
	for n := uint(286); ; n++ {
		pinverse[p(n)] = n
		hinverse[h(n)] = n

		tn := t(n)
		var (
			ok       bool
			n_p, n_h uint
		)
		if n_p, ok = pinverse[tn]; !ok {
			continue
		}

		if n_h, ok = hinverse[tn]; ok {
			fmt.Printf("t(%d) = p(%d) = h(%d) = %d\n", n, n_p, n_h, tn)
			return tn
		}
	}
	panic("unexpected")
}

func main() {
	fmt.Println(find())
}
