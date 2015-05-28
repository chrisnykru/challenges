package misc

import (
	"math"
	"testing"
)

// XXX need more testing/validation
func TestISqrt(t *testing.T) {
	stim := []struct{ n, fn uint64 }{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 2},
		{7, 2},
		{8, 2},
		{9, 3},
		{10, 3},
		{15, 3},
		{16, 4},
		{17, 4},
		{1e10 - 1, 1e5 - 1},
		{1e10, 1e5},
		{1e10 + 1, 1e5},
		{1<<53 - 1, 94906265},
		{1 << 53, 94906265},
		{1<<53 + 1, 94906265},
		{math.MaxUint32 - 1, math.MaxUint16},
		{math.MaxUint32, math.MaxUint16},
		{math.MaxUint64 - 1, math.MaxUint32},
		{math.MaxUint64, math.MaxUint32},
	}
	for _, s := range stim {
		if fn := ISqrt(s.n); fn != s.fn {
			t.Errorf("ISqrt(%v)=%v, want %v", s.n, fn, s.fn)
		}
	}
}
