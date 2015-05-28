package permgen

import (
	"sort"
)

func reverse(a []int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
}

type PermGen struct {
	perm      []int
	countdown int
}

func factorial(x int) int {
	f := 1
	for i := x; i > 1; i-- {
		f *= i
	}
	return f
}

func New(x []int) *PermGen {
	if len(x) == 0 {
		panic("empty set")
	}

	xiCount := make(map[int]int)
	for i := range x {
		xiCount[x[i]]++
	}

	gen := &PermGen{
		perm:      make([]int, len(x)),
		countdown: factorial(len(x)),
	}
	copy(gen.perm, x)
	sort.Ints(gen.perm)

	// deal with any overcounting
	for _, count := range xiCount {
		gen.countdown /= factorial(count)
	}
	return gen
}

func (gen *PermGen) current() (perm []int, last bool) {
	perm = make([]int, len(gen.perm))
	copy(perm, gen.perm)
	return perm, gen.countdown == 1
}

func (gen *PermGen) Next() (perm []int, last bool) {
	perm, last = gen.current()
	for j := len(gen.perm) - 2; j >= 0; j-- {
		for i := len(gen.perm) - 1; i > j; i-- {
			if gen.perm[i] > gen.perm[j] {
				// swap
				gen.perm[i], gen.perm[j] = gen.perm[j], gen.perm[i]
				// reverse
				reverse(gen.perm[j+1:])
				gen.countdown--
				return
			}
		}
	}
	return
}
