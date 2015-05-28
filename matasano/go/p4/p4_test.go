package p4

import (
	"testing"
	"os"
)

func Test(t *testing.T) {
	f, err := os.Open("4.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	line, plaintext := FindXorLine(f)
	t.Logf("line ciphertext = %v", line)
	t.Logf("plaintext = %v", plaintext)
}
