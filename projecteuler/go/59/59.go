/*

XOR decryption

Each character on a computer is assigned a unique code and the preferred
standard is ASCII (American Standard Code for Information Interchange). For
example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.

A modern encryption method is to take a text file, convert the bytes to ASCII,
then XOR each byte with a given value, taken from a secret key. The advantage
with the XOR function is that using the same encryption key on the cipher text,
restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.

For unbreakable encryption, the key is the same length as the plain text
message, and the key is made up of random bytes. The user would keep the
encrypted message and the encryption key in different locations, and without
both "halves", it is impossible to decrypt the message.

Unfortunately, this method is impractical for most users, so the modified method
is to use a password as a key. If the password is shorter than the message,
which is likely, the key is repeated cyclically throughout the message. The
balance for this method is using a sufficiently long password key for security,
but short enough to be memorable.

Your task has been made easy, as the encryption key consists of three lower case
characters. Using cipher1.txt (right click and 'Save Link/Target As...'), a file
containing the encrypted ASCII codes, and the knowledge that the plain text must
contain common English words, decrypt the message and find the sum of the ASCII
values in the original text.

-------------------------------------------------------------------------------

Keyspace = 26**3 = 17576

*/
package main

import (
	"encoding/csv"
	"fmt"
	"github.com/zeroshirts/challenges/projecteuler/go/misc"
	"os"
	"strconv"
	"strings"
	"math"
)

func loadCipher() []byte {
	f, err := os.Open("cipher1.txt")
	if err != nil {
		panic(err.Error())
	}

	csvr := csv.NewReader(f)
	records, err := csvr.ReadAll()
	if err != nil {
		panic(err.Error())
	}

	var b []byte
	for _, rec := range records {
		for _, numStr := range rec {
			num, err := strconv.ParseInt(numStr, 10, 0)
			if err != nil {
				panic(err.Error())
			}
			b = append(b, byte(num))
		}
	}
	return b
}

func testPlaintext(s string) (ok bool) {
	// guess: legit ASCII code range for msg [32,126]
	nonlegitCount := 0
	for i := range s {
		if s[i] < 32 || s[i] > 126 {
			nonlegitCount++
		}
	}

	if nonlegitCount > 0 {
		return false
	}

	// look for common words
	common := []string{"the ", "a ", "an ", "to ", "and ", "is ", "are "}
	commonCount := 0
	for _, word := range common {
		if strings.Contains(s, word) {
			commonCount++
		}
	}

	// arbitrary guess
	if commonCount >= 6 {
		return true
	}

	return false
}

func decrypt(cipher []byte, key []byte) (plaintext []byte, ok bool) {
	b := make([]byte, len(cipher))
	for i := range cipher {
		b[i] = cipher[i] ^ key[i%len(key)]
	}
	return b, testPlaintext(string(b))
}

// key is known to be three lower case characters
const ascii_a = 97

func bruteForce(cipher []byte) (plaintext []byte) {
	limit := uint64(math.Pow(26, 3))
	for i := uint64(0); i < limit; i++ {
		d, _ := misc.Uint64ToFixedWidthDigits(i, 26, 3)

		key := []byte{
			byte(ascii_a + d[0]),
			byte(ascii_a + d[1]),
			byte(ascii_a + d[2]),
		}

		if plaintext, ok := decrypt(cipher, key); ok {
			fmt.Printf("plaintext = %s\n", plaintext)
			return plaintext
		}
	}
	panic("no valid plaintext found!")
}

func sum(b []byte) int {
	var s int
	for i := range b {
		s += int(b[i])
	}
	return s
}

func main() {
	//fmt.Println("cipher: ", loadCipher())
	plaintext := bruteForce(loadCipher())
	fmt.Println("plaintext:", string(plaintext))
	fmt.Println("sum:", sum(plaintext))
}
