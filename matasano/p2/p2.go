// Fixed XOR
// Write a function that takes two equal-length buffers and produces their XOR combination.
//
// If your function works properly, then when you feed it the string:
//
// 1c0111001f010100061a024b53535009181c
// ... after hex decoding, and when XOR'd against:
//
// 686974207468652062756c6c277320657965
// ... should produce:
//
// 746865206b696420646f6e277420706c6179
package p2

import (
	"encoding/hex"
	"errors"
)

const (
	knownA      = "1c0111001f010100061a024b53535009181c"
	knownB      = "686974207468652062756c6c277320657965"
	knownOutput = "746865206b696420646f6e277420706c6179"
)

// I'm sure this will come up later... this is not a constant time implementation.
func XOR_hexStrings(a, b string) (string, error) {
	if len(a) != len(b) {
		return "", errors.New("inputs are not equal length")
	}
	araw, err := hex.DecodeString(a)
	if err != nil {
		return "", err
	}
	braw, err := hex.DecodeString(b)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(araw); i++ {
		// overwrite araw in place
		araw[i] = araw[i] ^ braw[i]
	}
	return hex.EncodeToString(araw), nil
}
