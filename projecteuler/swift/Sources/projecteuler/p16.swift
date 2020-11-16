/*

Power digit sum

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?

*/




/*
package main

import (
        "fmt"
        "math/big"
)

func sumOfDigits(base, exponent *big.Int) *big.Int {
        x := new(big.Int)
        x.Exp(base, exponent, nil)

        sum := big.NewInt(0)
        for _, digit := range x.String() {
                sum.Add(sum, big.NewInt(int64(digit-'0')))
        }
        return sum
}

func main() {
        s := sumOfDigits(big.NewInt(2), big.NewInt(1000))
        fmt.Printf("sum of digits = %s\n", s)
}
*/
