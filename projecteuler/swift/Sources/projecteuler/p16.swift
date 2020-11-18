/*

Power digit sum

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?

*/

import BigInt

func sumOfDigits(base: BigInt, exponent: Int) -> Int {
    var x = base.power(exponent)
    var sum = 0
    while x > 0 {
        let r: BigInt
        (x, r) = x.quotientAndRemainder(dividingBy: 10)
        sum += Int(r)
    }
    return sum
}
