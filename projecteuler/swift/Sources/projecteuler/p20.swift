/*

Factorial digit sum

n! means n * (n  1) * ... * 3 * 2 * 1

For example, 10! = 10 * 9 * ... * 3 * 2 * 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!

*/

import BigInt

func sumDigits(_ x: BigInt) -> Int {
    var sum = 0
    var tmp = x
    while tmp > 0 {
        sum += Int(tmp % 10)
        tmp /= 10
    }
    return sum
}

func sumOfDigitsOfFactorial(_ n: Int) -> Int {
    let x = factorial(BigInt(n))
    return sumDigits(x)
}


/*
func sumDigits(x *big.Int) int64 {
    sum := big.NewInt(0)
    for _, digit := range x.String() {
        sum.Add(sum, big.NewInt(int64(digit)-'0'))
    }
    return sum.Int64()
}

func main() {
    f := misc.Factorial(100)
    fmt.Printf("factorial = %s\n", f)
    fmt.Printf("sum = %v\n", sumDigits(f))
}
*/
