/*

Non-bouncy numbers

XXX TODO unsolved

Working from left-to-right if no digit is exceeded by the digit to its left it
is called an increasing number; for example, 134468.

Similarly if no digit is exceeded by the digit to its right it is called a
decreasing number; for example, 66420.

We shall call a positive integer that is neither increasing nor decreasing a
"bouncy" number; for example, 155349.

As n increases, the proportion of bouncy numbers below n increases such that
there are only 12951 numbers below one-million that are not bouncy and only
277032 non-bouncy numbers below 10**10.

How many numbers below a googol (10**100) are not bouncy?

*/
package main

import (
	"fmt"
	"math/big"
	"github.com/zeroshirts/challenges/projecteuler/go/112/misc112"
)

var (
	zero = big.NewInt(0)
	one = big.NewInt(1)
	ten = big.NewInt(10)
)

func googol() *big.Int {
	return big.NewInt(0).Exp(big.NewInt(10), big.NewInt(100), nil)
}

/*
// [start, stopBefore)
func notBouncyUnder2(start, stopBefore, ceil *big.Int) *big.Int {
	notBouncy := make([]*big.Int, 0)
	if stopBefore.Cmp(ceil) >= 0 {
		stopBefore = ceil
	}
	for i := new(big.Int).Set(start); i.Cmp(stopBefore) < 0; i = i.Add(i, one) {
		if !misc112.IsBouncyBigInt(i) {
			icopy := new(big.Int).Set(i)
			notBouncy = append(notBouncy, icopy)
		}
	}

	//fmt.Printf("\nfor start=%v, stopBefore=%v, ceil=%v\n  notBouncy=%v\n", start, stopBefore, ceil, notBouncy)

	count := new(big.Int)
	for i := 0; i < len(notBouncy); i++ {
		count = count.Add(count, one) // accounts for notBouncy[i]
		start := new(big.Int).Mul(notBouncy[i], ten)
		if start.Cmp(ceil) <= 0 {
			var stopBefore *big.Int
      if i == len(notBouncy)-1 {
				stopBefore = new(big.Int).Set(notBouncy[i])
				stopBefore = stopBefore.Add(stopBefore, one)
				stopBefore = new(big.Int).Mul(stopBefore, ten)
			} else {
				stopBefore = new(big.Int).Mul(notBouncy[i+1], ten)
			}
			tmp := notBouncyUnder2(start, stopBefore, ceil)
			count = count.Add(count, tmp)
		}
	}
	return count
}
*/

func notBouncyUnder2(start, ceil *big.Int) *big.Int {
  stopBefore := new(big.Int).Mul(start, ten)
	if stopBefore.Cmp(ceil) > 0 {
		stopBefore = ceil
	}
	notBouncy := make([]*big.Int, 0)

	for i := new(big.Int).Set(start); i.Cmp(stopBefore) < 0; i = i.Add(i, one) {
		if !misc112.IsBouncyBigInt(i) {
			icopy := new(big.Int).Set(i)
			notBouncy = append(notBouncy, icopy)
		}
	}



	count := big.NewInt(int64(len(notBouncy)))
	return count
}

// [ [1,10) ]
//   vet everything in ranges --> everything is 'not-bouncy'
//   construct new ranges based on contiguous values * 10
//   1*10=10, 10*10=100 // note, do we just do (9+1)*10 here?
// [ [10, 100) ]
//   --> everything is still 'not-bouncy'
//   construct new ranges
//   10*10=100, 100*10=1000
// [ [100, 1000) ]
//   100, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 122, 
//   [ 100, 120) // 119+1...
//   [122, 130) // 129+1
// 
// 
// 
func notBouncyUnder2(powOfTen *big.Int, ranges [][]*big.Int) *big.Int

func notBouncyUnder(ceil *big.Int) *big.Int {
  for i := new(big.Int).Set(zero); i.Cmp(stopBefore) < 0; i = i.Add(i, one) {
		notBouncy = append(notBouncy, icopy)
	}
	notBouncyUnder2(one

	x := notBouncyUnder2(one, ceil)
	return x.Add(x, one)
}

func main() {
	ceil := big.NewInt(1000)
	//ceil := googol()
	x := notBouncyUnder(ceil)
	fmt.Printf("Not bouncy under %v = %v\n", ceil, x)
}
