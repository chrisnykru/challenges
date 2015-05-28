package misc

// Trial Divison method.  1 is not Prime.
func IsPrime(x uint64) bool {
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	// even?
	if x%2 == 0 {
		return false
	}

	if len(ProperDivisors(x)) > 1 {
		return false
	}
	return true
}
