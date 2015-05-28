/*

Lychrel numbers

If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.

Not all numbers produce palindromes so quickly. For example,

349 + 943 = 1292,
1292 + 2921 = 4213
4213 + 3124 = 7337

That is, 349 took three iterations to arrive at a palindrome.

Although no one has proved it yet, it is thought that some numbers, like 196,
never produce a palindrome. A number that never forms a palindrome through the
reverse and add process is called a Lychrel number. Due to the theoretical
nature of these numbers, and for the purpose of this problem, we shall assume
that a number is Lychrel until proven otherwise. In addition you are given
that for every number below ten-thousand, it will either (i) become a
palindrome in less than fifty iterations, or, (ii) no one, with all the
computing power that exists, has managed so far to map it to a palindrome.
In fact, 10677 is the first number to be shown to require over fifty iterations
before producing a palindrome:
4668731596684224866951378664 (53 iterations, 28-digits).

Surprisingly, there are palindromic numbers that are themselves Lychrel
numbers; the first example is 4994.

How many Lychrel numbers are there below ten-thousand?

NOTE: Wording was modified slightly on 24 April 2007 to emphasise the
theoretical nature of Lychrel numbers.

*/
package main

import (
	"fmt"
	"math/big"
)

// base 10 palindrome
// e.g., 99 = true, 101 = true, 201 = false
func IsPalindrome(x *big.Int) bool {
	s := x.String()
	// len(s) >= 1

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func palindrome(x *big.Int) *big.Int {
	b := []byte(x.String())
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	x2, ok := new(big.Int).SetString(string(b), 10)
	if !ok {
		panic("SetString() failed")
	}
	return x2
}

type lychrelLut struct {
	lut map[int64]bool
}

func (ll *lychrelLut) Init() *lychrelLut {
	ll.lut = make(map[int64]bool)
	return ll
}

func (ll *lychrelLut) IsLychrel(x int64) bool {
	x2 := big.NewInt(x)
	x2.Add(x2, palindrome(big.NewInt(x)))

	try := 0
	for {
		////fmt.Printf("try=%2d  x = %v, x2 = %s\n", try, x, x2)
		try++
		if try > 50 {
			// reminder: failure means it IS lychrel
			ll.lut[x] = true
			return true
		}

		if IsPalindrome(x2) {
			return false
		}
		x2.Add(x2, palindrome(x2))
	}
	panic("!!")
}

func (ll *lychrelLut) countBelow(ceiling int) int {
	count := 0
	for i := 0; i < ceiling; i++ {
		if ll.IsLychrel(int64(i)) {
			count++
		}
	}
	return count
}

func main() {
	ll := new(lychrelLut).Init()
	ceiling := 10000
	fmt.Printf("countBelow(%v) = %v\n", ceiling, ll.countBelow(ceiling))
}
