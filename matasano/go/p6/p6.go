// Break repeating-key XOR
//
// It is officially on, now.
// This challenge isn't conceptually hard, but it involves actual error-prone
// coding. The other challenges in this set are there to bring you up to speed.
// This one is there to qualify you. If you can do this one, you're probably
// just fine up to Set 6.
//
// There's a file here. It's been base64'd after being encrypted with
// repeating-key XOR.
//
// Decrypt it.
//
// Here's how:
//
// Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.
// Write a function to compute the edit distance/Hamming distance between two
// strings. The Hamming distance is just the number of differing bits.
//
// The distance between:
//   this is a test
// and
//   wokka wokka!!!
// is 37. Make sure your code agrees before you proceed.
//
// For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second
// KEYSIZE worth of bytes, and find the edit distance between them. Normalize
// this result by dividing by KEYSIZE.
//
// The KEYSIZE with the smallest normalized edit distance is probably the key.
// You could proceed perhaps with the smallest 2-3 KEYSIZE values. Or take 4
// KEYSIZE blocks instead of 2 and average the distances.
//
// Now that you probably know the KEYSIZE: break the ciphertext into blocks
// of KEYSIZE length.
// Now transpose the blocks: make a block that is the first byte of every block,
// and a block that is the second byte of every block, and so on.
// Solve each block as if it was single-character XOR. You already have code to
// do this.
//
// For each block, the single-byte XOR key that produces the best looking
// histogram is the repeating-key XOR key byte for that block. Put them together
// and you have the key.
//
// This code is going to turn out to be surprisingly useful later on. Breaking
// repeating-key XOR ("Vigenere") statistically is obviously an academic
// exercise, a "Crypto 101" thing. But more people "know how" to break it than
// can actually break it, and a similar technique breaks something much more
// important.
//
// No, that's not a mistake.
// We get more tech support questions for this challenge than any of the other
// ones. We promise, there aren't any blatant errors in this text. In
// particular: the "wokka wokka!!!" edit distance really is 37.
//
package p6

import (
	"os"
	"io/ioutil"
	"encoding/base64"
	"math"
	"fmt"
)

// Returns edit distance/Hamming distance
func HammingDistance(s1, s2 string) int {
	if len(s1) != len(s2) {
		panic("len(s1) != len(s2)")
	}
	hd := 0
	for i := 0; i < len(s1); i++ {
		x := s1[i]
		y := s2[i]
		// assume runes are ASCII
		for j := 0; j < 8; j++ {
			hd += int((x & 0x01) ^ (y & 0x01))
      x = x >> 1
			y = y >> 1
		}
	}
	return hd
}

func readFileAndDecodeBase64() string {
	f, err := os.Open("6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b64, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	data, err := base64.StdEncoding.DecodeString(string(b64))
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Returns likely keysize
func findLikelyKeysize(data string) int {
	fmt.Printf("len(data)=%v\n", len(data))
	minhd := float64(math.MaxFloat64)
	var minhdks int
	for ks := 2; ks <= 40; ks++ {
		var hd float64
		numHd := 0
		for i := 0; i < len(data) - ks - ks; i += ks {
			//fmt.Printf("HammingDistance(data[%d:%d], data[%v:%v]\n", i, i+ks, i+ks, i+2*ks)
			hd += float64(HammingDistance(data[i:i+ks], data[i+ks:i+2*ks]))
			numHd++
		}
		if len(data) % (2 * ks) == 0 {
			//fmt.Printf("HammingDistance(data[%d:%d], data[%v:%v]\n", len(data)-2*ks, len(data)-ks, len(data)-ks, len(data))
			hd += float64(HammingDistance(data[len(data)-2*ks:len(data)-ks], data[len(data)-ks:len(data)]))
			numHd++
		}
		// normalize
		hd /= float64(numHd)
		hd /= float64(ks)
		if hd < minhd {
			minhd = hd
			minhdks = ks
			fmt.Printf("hd = %v for ks = %v\n", hd, ks)
		}
	}
	return minhdks
}
