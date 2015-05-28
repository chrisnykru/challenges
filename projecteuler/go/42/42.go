/*

Coded triangle numbers

The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1);
so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its
alphabetical position and adding these values we form a word value. For example,
the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a
triangle number then we shall call the word a triangle word.

Using words.txt, a 16K text file containing nearly two-thousand common English
words, how many are triangle words?

*/
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var (
	words []string
)

func init() {
	f, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}

	csvr := csv.NewReader(f)
	records, err := csvr.ReadAll()
	if err != nil {
		panic(err)
	}
	words = records[0]
}

func t(n uint) uint {
	return (n * (n + 1)) / 2
}

func tinverse() func(x uint) (n uint, ok bool) {
	// t(n) --> n
	inverseMap := make(map[uint]uint)
	// initial entry
	highN, highT := uint(1), t(1)
	inverseMap[highT] = highN

	return func(x uint) (n uint, ok bool) {
		for highT < x {
			highN++
			highT = t(highN)
			inverseMap[highT] = highN
		}
		n, ok = inverseMap[x]
		return
	}
}

// similar to 22/22.go namescore()
func wordVal(word string) uint {
	score := uint(0)
	for _, v := range word {
		letterScore := uint(v - 'A' + 1)
		score += letterScore
	}
	return score
}

func triangleWordCount() uint {
	count := uint(0)

	tinv := tinverse()
	for _, word := range words {
		v := wordVal(word)
		if n, ok := tinv(v); ok {
			fmt.Printf("word %q = %v, a.k.a. t(%d)\n", word, v, n)
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(triangleWordCount())
}
