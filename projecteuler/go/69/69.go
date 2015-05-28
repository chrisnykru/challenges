/*

Totient maximum

Euler's Totient function, φ(n) [sometimes called the phi function], is used to
determine the number of numbers less than n which are relatively prime to n.
For example, as 1,2,4,5,7,8, are all less than nine and relatively prime
to nine, φ(9)=6.

n    Relatively Prime   φ(n)   n/φ(n)
2    1                  1      2
3    1,2                2      1.5
4    1,3                2      2
5    1,2,3,4            4      1.25
6    1,5                2      3
7    1,2,3,4,5,6        6      1.1666...
8    1,3,5,7            4      2
9    1,2,4,5,7,8        6      1.5
10   1,3,7,9            4      2.5

It can be seen that n=6 produces a maximum n/φ(n) for n <= 10.

Find the value of n <= 1,000,000 for which n/φ(n) is a maximum.


-------------------------------------------------------------------------------


definition: relatively prime
  Two integers are relatively prime if they share no common positive factors

*/
package main

import (
	"fmt"
	"math"
)

/*

n <= 1,000,000

1000 * 1000 = 1,000,000

space complexity of factors for numbers [1,1000] ~= 1000 * 33 * 8 bytes ~= 256,000 bytes?


*/

// assumes n < 2**53 for floating point accuracy
func totient(n int) int {
	divisors := make(map[int]bool)
	sqrt_n := int(math.Sqrt(float64(n)))
	for i := 2; i < sqrt_n; i++ {
		if n % i == 0 {
			divisors[i] = true
			divisors[n/i] = true
		}
	}
	fmt.Printf("n=%v divisors=%v\n", n, divisors)
	return n - len(divisors)
}

/*
type factorLut struct {
  lut map[uint][]uint
  highest uint
}

func newFactorLut() *factorLut {
  return &factorLut{lut: make(map[uint][]bool)}
}
*/

func findMax() uint64 {
	for n := uint64(2); n <= 1000000; n++ {

	}
	return 12345
}

func main() {
	//fmt.Println(findMax())

  fmt.Println(totient(10))


}
