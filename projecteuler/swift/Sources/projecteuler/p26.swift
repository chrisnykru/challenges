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
    // if d=13: pow10LeftShift = 10**13
    let pow10LeftShift = BigInt(10).power(d)
    
    // if d=13: m = 10**13 - (13 * 769230769230) = 10
    let (_, m) = pow10LeftShift.quotientAndRemainder(dividingBy: BigInt(d))
    if m == 0 {
        return 0 // modulus is zero, no cycle
    }
     
    // eliminates any non-repeating prefix
    // if d=13: num = 10 * 10**13 = 10**14
    let num = m * pow10LeftShift
    
    // if d=13: q = 10**14 / 13 = 7692307692307
    let (q, _) = num.quotientAndRemainder(dividingBy: BigInt(d))
    
    let q_array = String(q).unicodeScalars.map { $0.value }
    for sublen in stride(from: 1, through: q_array.count/2, by: 1) {
        var match = 0
        for i in stride(from: 0, to: sublen, by: 1) {
            if q_array[i] == q_array[i + sublen] {
                match += 1
            }
        }
        if match == sublen {
            return match
        }
    }
    return q_array.count - 1
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
