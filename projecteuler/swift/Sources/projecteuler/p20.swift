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
