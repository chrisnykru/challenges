package misc112

import (
	"testing"
	"math/big"
)

func TestIsBouncy(t *testing.T) {
	stim := []struct {
		x      int
		bouncy bool
	}{
		{155349, true},
		{123, false},
		{321, false},
		{23, false},
	}
	for _, s := range stim {
		r := IsBouncy(s.x)
		if r != s.bouncy {
			t.Errorf("IsBouncy(%d) = %v, want %v", s.x, r, s.bouncy)
		}
		r = IsBouncyBigInt(big.NewInt(int64(s.x)))
		if r != s.bouncy {
			t.Errorf("IsBouncy(%d) = %v, want %v", s.x, r, s.bouncy)
		}
	}
}
