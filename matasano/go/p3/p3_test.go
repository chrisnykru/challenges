package p3

import (
	"testing"
)

func Test(t *testing.T) {
	english_lf := GetEnglishLetterFrequencies()
	//for l, f := range english_lf {
	//	fmt.Printf("%c frequency is %f\n", rune(l), f)
	//}

	cipherText := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	plainText, key, _ := BruteForceHexStr(cipherText, english_lf)
	t.Logf("plainText = %v", plainText)
	t.Logf("key = %v", key)

	wantPlainText := "Cooking MC's like a pound of bacon"
	wantKey := "58585858585858585858585858585858585858585858585858585858585858585858"
	if plainText != wantPlainText {
		t.Errorf("plainText = %v, want %v", plainText, wantPlainText)
	}
	if key != wantKey {
		t.Errorf("key = %v, want %v", key, wantKey)
	}
}
