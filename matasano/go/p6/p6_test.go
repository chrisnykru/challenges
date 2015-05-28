package p6

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
  s1 := "this is a test"
	s2 := "wokka wokka!!!"
	hd := HammingDistance(s1, s2)
	if hd != 37 {
		t.Errorf("hd = %v, want %v", hd, 37)
	}
}

func Test(t *testing.T) {
	s := readFileAndDecodeBase64()
  ks := findLikelyKeysize(s)
	t.Log(ks)

}
