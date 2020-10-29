/*

Longest Collatz sequence

The following iterative sequence is defined for the set of positive integers:
  n -> n/2 (n is even)
  n -> 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
  13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10
terms. Although it has not been proved yet (Collatz Problem), it is thought that
all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.

*/

func seqNext(_ x: Int) -> Int {
    if x % 2 == 0 { // x is even
        return x / 2
    } else { // x is odd
        return 3 * x + 1
    }
}

func seqLength(_ x: Int) -> Int {
    var seqLen = 1
    var tmp = x
    while tmp != 1 {
        tmp = seqNext(tmp)
        seqLen += 1
    }
    return seqLen
}

func findLargestSeq(_ under: Int) -> (Int, Int) {
    var maxSl = 1
    var maxSlNum = 1
    for i in stride(from: 1, to: under, by: 1) {
        let sl = seqLength(i)
        if sl > maxSl {
            maxSl = sl
            maxSlNum = i
        }
    }
    return (maxSlNum, maxSl)
}
