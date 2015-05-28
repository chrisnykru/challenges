package misc

func ProperDivisors(x uint64) map[uint64]bool {
	sqrt_x := ISqrt(x)

	// divisor set
	div := make(map[uint64]bool)
	div[1] = true

	for i := uint64(2); i <= sqrt_x; i++ {
		if x%i == 0 {
			div[i] = true
			div[x/i] = true
		}
	}
	return div
}

func Divisors(x uint64) map[uint64]bool {
	div := ProperDivisors(x)
	div[x] = true
	return div
}
