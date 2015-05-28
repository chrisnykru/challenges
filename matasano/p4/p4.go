// Detect single-character XOR
// One of the 60-character strings in this file has been encrypted by single-character XOR.
//
// Find it.
//
// (Your code from #3 should help.)
package p4

import (
	"fmt"
	"bufio"
	"github.com/cd5/crypto-challenges/matasano/p3"
	"math"
	"io"
)

// Return original line ciphertext and decrypted plaintext
func FindXorLine(r io.Reader) (string, string) {
	english_lf := p3.GetEnglishLetterFrequencies()

	bestScore := float64(math.MaxFloat64)
	bestLine := ""
	bestGuess := ""

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := string(scanner.Bytes())
		guess, key, score := p3.BruteForceHexStr(line, english_lf)
		fmt.Printf("guess=%q key=%q score=%v\n", guess, key, score)
		if score < bestScore {
			bestScore = score
			bestLine = line
			bestGuess = guess
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return bestLine, bestGuess
}
