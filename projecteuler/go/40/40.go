/*

Champernowne's constant

An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the
following expression.

d1 * d10 * d100 * d1000 * d10000 * d100000 * d1000000

*/
package main

import (
	"bytes"
	"fmt"
)

func dSlow() func(n uint) uint {
	var b bytes.Buffer
	for i := uint(1); i <= 1000000; i++ {
		b.WriteString(fmt.Sprintf("%d", i))
	}
	s := b.String()

	return func(n uint) uint {
		if n < 1 {
			panic("bad n")
		}
		// n=1 --> [0]
		return uint(s[n-1] - '0')
	}
}

/*
// XXX
func dFast(n uint) uint {

  length := ceil(log10(n)) 
  l(1) = 1
  l(2) = 1
  l(10) = 1
  l(11) = 2
  l(100) = 2


  if n=101
  l(101) = 3

  smallest(3) = exp(10, 3-1) = 100

  3 * (101-100) = 3
  + 2 * (99 - 10 + 1) = 180
  + 1 * (9 - 1 + 1) = 9 // special case... no zero


  hmm...


}
*/

func expr() uint {
	//d := dFast
	d := dSlow()
	return d(1) * d(10) * d(100) * d(1000) * d(10000) * d(100000) * d(1000000)
}

func main() {
	fmt.Println(expr())
}
