package misc

import (
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	stim := []struct {
		x  int64
		fx *big.Int
	}{
		{0, big.NewInt(1)},
		{1, big.NewInt(1)},
		{2, big.NewInt(2)},
		{3, big.NewInt(6)},
	}

	for _, s := range stim {
		if fx := Factorial(s.x); fx.Cmp(s.fx) != 0 {
			t.Errorf("Factorial(%v) = %v, want %v", s.x, fx, s.fx)
		}
	}
}
