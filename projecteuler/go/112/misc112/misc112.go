package misc112

import "math/big"

func IsBouncy(x int) bool {
	lastDigit := -1
	dec, inc := false, false
	for ; x != 0; x /= 10 {
		digit := x%10
		if lastDigit != -1 {
			if digit < lastDigit {
				dec = true
			}
			if digit > lastDigit {
				inc = true
			}
		}
		lastDigit = digit
	}
	return dec && inc
}

var (
	negative_one = big.NewInt(-1)
	zero = big.NewInt(0)
	ten = big.NewInt(10)
)

func IsBouncyBigInt(orig *big.Int) bool {
	x := new(big.Int).Set(orig) // make copy that we can mutate
	lastDigit := new(big.Int).Set(negative_one)
	dec, inc := false, false
	for ; x.Cmp(zero) != 0; x = x.Div(x, ten) {
		digit := new(big.Int).Mod(x, ten)
		if lastDigit.Cmp(negative_one) != 0 {
			if digit.Cmp(lastDigit) < 0 { // less than
				dec = true
			}
			if digit.Cmp(lastDigit) > 0 { // greater than
				inc = true
			}
		}
		lastDigit.Set(digit)
	}
	return dec && inc
}
