/*

Reciprocal cycles

A unit fraction contains 1 in the numerator. The decimal representation of the
unit fractions with denominators 2 to 10 are given:

1/2    =     0.5
1/3    =     0.(3)
1/4    =     0.25
1/5    =     0.2
1/6    =     0.1(6)
1/7    =     0.(142857)
1/8    =     0.125
1/9    =     0.(1)
1/10    =     0.1

Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be
seen that 1/7 has a 6-digit recurring cycle.

Find the value of d<1000 for which 1/d contains the longest recurring cycle
in its decimal fraction part.

*/

import BigInt

func cycle(_ d: Int) -> Int {
    // assumption: worst-case cycle length for 1/d is d-1
    // pow==10**d sufficient to determine second worst-case
    //            cycle length of floor(d/2)
    //
    // if d=13: 1/13 --> 10**13/13 = 769230769230; cycle len = 6
    let x = BigInt(10).power(d)
    
    // if d=13: m = 10**13 - (13 * 769230769230) = 10
    let (_, m) = x.quotientAndRemainder(dividingBy: BigInt(d))
    if m == 0 {
        return 0 // modulus is zero, no cycle
    }
    
    // eliminates any non-repeating prefix
    // e.g., 1/6=.166666; after elimination: 666
    //
    // if d=13: num = 10 * 10000000000000 = 100000000000000
    let num = m * x
    
    // if d=13: x2 = 100000000000000 / 13 = 7692307692307
    //          m2 = 100000000000000 - (7692307692307 * 13) = 9
    let (x2, _) = num.quotientAndRemainder(dividingBy: BigInt(d))
    
    //print(d, x2, m2)
    
    let x2_array = String(x2).unicodeScalars.map { $0.value }
    for sublen in stride(from: 1, through: x2_array.count/2, by: 1) {
        var match = 0
        for i in stride(from: 0, to: sublen, by: 1) {
            if x2_array[i] == x2_array[i + sublen] {
                match += 1
            }
        }
        if match == sublen {
            return match
        }
    }
    return x2_array.count - 1
}

func findLongestCycle(_ dMax: Int) -> (d: Int, cycleLen: Int) {
    var longestD = 0
    var longestCycleLen = 0
    for d in stride(from: 2, to: dMax, by: 1) {
        let cycleLen = cycle(d)
        if cycleLen > longestCycleLen {
            longestD = d
            longestCycleLen = cycleLen
        }
    }
    return (longestD, longestCycleLen)
}
