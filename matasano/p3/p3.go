// Single-byte XOR cipher
// The hex encoded string:
//
// 1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
// ... has been XOR'd against a single character. Find the key, decrypt the message.
//
// You can do this by hand. But don't: write code to do it for you.
//
// How? Devise some method for "scoring" a piece of English plaintext. Character
// frequency is a good metric. Evaluate each output and choose the one with the
// best score.
//
// Achievement Unlocked
// You now have our permission to make "ETAOIN SHRDLU" jokes on Twitter.
//
// TODO: should probably use n-grams for this
// see: http://people.seas.harvard.edu/~jones/cscie129/papers/stanford_info_paper/entropy_of_english_9.htm
package p3

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/cd5/crypto-challenges/matasano/p2"
	"io"
	"math"
	"os"
	"strings"
)

func getRuneFrequencies(r io.Reader) map[rune]float64 {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	totalLetters := 0
	lf := make(map[rune]float64)
	for scanner.Scan() {
		// deal with word, carve it up into letters and count each letter...
		word := strings.ToLower(string(scanner.Bytes()))
		for i := 0; i < len(word); i++ {
			totalLetters++
			lf[rune(word[i])]++
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	for letter, count := range lf {
		lf[letter] = count / float64(totalLetters)
	}
	return lf
}

func scoreTextAsEnglish(text string, english_lf map[rune]float64) float64 {
	text = strings.ToLower(text)
	text_lf := getRuneFrequencies(strings.NewReader(text))
	score := float64(0)
	for l, f := range english_lf {
		f2 := text_lf[l] // defaults to 0 for non-existant keys
		score += math.Abs(f2 - f)
	}
	for l2, f2 := range text_lf {
		if _, ok := english_lf[l2]; !ok {
			score += f2
		}
	}
	return score
}

func BruteForceStr(s string, lf map[rune]float64) (string, string, float64) {
	hexStr := hex.EncodeToString([]byte(s))
	return BruteForceHexStr(hexStr, lf)
}

// Returns decrypted message, key, and english score
func BruteForceHexStr(hexStr string, lf map[rune]float64) (string, string, float64) {
	bestScore := math.MaxFloat64
	bestGuess := ""
	bestKey := ""

	for char := 0; char <= 255; char++ {
		b := make([]byte, len(hexStr))
		tmp := fmt.Sprintf("%02x", char)
		for i := 0; i < len(b); i += 2 {
			b[i] = tmp[0]
			b[i+1] = tmp[1]
		}
		key := string(b)
		xhex, err := p2.XOR_hexStrings(hexStr, key)
		if err != nil {
			panic(err)
		}
		b, err = hex.DecodeString(xhex)
		if err != nil {
			panic(err)
		}
		guess := string(b)
		score := scoreTextAsEnglish(guess, lf)
		if score < bestScore {
			bestScore = score
			bestGuess = guess
			bestKey = key
			fmt.Printf("new score = %v, guess = %q\n", bestScore, bestGuess)
		}
	}
	return bestGuess, bestKey, bestScore
}

func GetEnglishLetterFrequencies() map[rune]float64 {
	f, err := os.Open("/usr/share/dict/words")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lf := getRuneFrequencies(f)
	// Hack in some decent frequencies for punctuation
	lf[rune(' ')] = float64(.1)
	return lf
}
