/*

Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

*/
package main

import (
	"fmt"
	"math"
)

func pythagoreanTripletOf1000() [3]int {
	/*

	  a^2 + b^2 = c^2
	  find triplet such that a + b = c = 1000

	  c = (a^2 + b^2)^.5
	  a + b + (a^2 + b^2)^.5 = 1000

	*/

	for b := 1; b < 1000; b++ {
		for a := 1; a < b; a++ {
			cf := math.Sqrt(float64(a*a + b*b))
			if math.Floor(cf) != cf {
				// c is not an integer
				continue
			}

			c := int(cf)
			if c <= b {
				continue
			}

			if a+b+c == 1000 {
				return [3]int{a, b, c}
			}
		}
	}
	panic("fail")
}

func main() {
	triplet := pythagoreanTripletOf1000()
	fmt.Printf("a=%v, b=%v, c=%v\n", triplet[0], triplet[1], triplet[2])
	fmt.Printf("a*b*c = %v\n", triplet[0]*triplet[1]*triplet[2])
}
