package main

import (
	"math/big"
	"testing"
)

func TestBigExp(t *testing.T) {
	// 2**1024
	x := new(big.Int).Exp(big.NewInt(2), big.NewInt(1024), nil)

	expected_xstr := "179769313486231590772930519078902473361797697894230657273430081157732675805500963132708477322407536021120113879871393357658789768814416622492847430639474124377767893424865485276302219601246094119453082952085005768838150682342462881473913110540827237163350510684586298239947245938479716304835356329624224137216"
	if x.String() != expected_xstr {
		t.Errorf("x=%s\nwant %s", x.String(), expected_xstr)
	}

	// 2**1024 mod 1000000
	x2 := new(big.Int).Exp(big.NewInt(2), big.NewInt(1024), big.NewInt(1000000))

	expected_x2str := "137216"
	if x2.String() != expected_x2str {
		t.Errorf("x2=%s\nwant %s", x2.String(), expected_x2str)
	}
}
