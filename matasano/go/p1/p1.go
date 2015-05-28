// Convert hex to base64
// The string:
//
// 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
// Should produce:
//
// SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
// So go ahead and make that happen. You'll need to use this code for the rest of the exercises.
//
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const (
	knownInput  = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	knownOutput = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
)

func HexToBase64(in string) (string, error) {
	b, err := hex.DecodeString(in)
	if err != nil {
		return "", err
	}
	out := base64.StdEncoding.EncodeToString(b)
	return out, nil
}

func main() {
	fmt.Println(HexToBase64(knownInput))
}
