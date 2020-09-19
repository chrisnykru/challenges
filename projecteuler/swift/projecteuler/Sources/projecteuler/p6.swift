/*

Sum square difference

The sum of the squares of the first ten natural numbers is,
12 + 22 + ... + 102 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.

*/

func sumSquareAndSquareOfSumDifferenceNaive(_ num: Int) -> Int {
    var sumOfSquares = 0
    var i = 1
    while i <= num {
        sumOfSquares += i * i
        i += 1
    }
    var squareOfSums = 0
    i = 1
    while i <= num {
        squareOfSums += i
        i += 1
    }
    squareOfSums *= squareOfSums
    return abs(sumOfSquares - squareOfSums)
}

func sumSquareAndSquareOfSumDifference(_ num: Int) -> Int {
    // Sum(0,n) i   = n(n+1)/2
    // Sum(0,n) i^2 = n^3/3 + n^2/2 + n/6
    //
    // Answer = n^3/3 + n^2/2 + n/6 - (n(n+1)/2)^2
    //        = -n^4/4 - n^3/6 + n^2/4 + n/6
    return abs(-1 * num * num * num * num / 4 - num * num * num / 6 + num * num / 4 + num / 6)
}
