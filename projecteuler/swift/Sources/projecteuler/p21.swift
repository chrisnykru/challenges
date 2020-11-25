/*

Amicable numbers

Let d(n) be defined as the sum of proper divisors of n (numbers less than n
which divide evenly into n).  If d(a) = b and d(b) = a, where a != b,
then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are:
  1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
  Therefore d(220) = 284

The proper divisors of 284 are:
  1, 2, 4, 71 and 142;
  so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

*/

func d(_ x: Int) throws -> Int {
    let divisors = try factors(x)
    // Sum up divisors, and subtract 'x', which is not a proper divisor
    return divisors.reduce(0, +) - x
}

func sumOfAmicableNumbers(under: Int) throws -> Int {
    var amicablePairs = [Int:Int]()
    for a in stride(from: 1, to: under, by: 1) {
        let b = try d(a)
        guard a != b else {
            // constraint: a != b
            continue
        }
        print("-->", a, b)
        let a_maybe = try d(b)
        print("-->  ", a_maybe)
        guard a_maybe == a else {
            // not amicably paired
            continue
        }
        
        // amicable pair
        print("a. pair", a, b)
        amicablePairs[a] = b
        amicablePairs[b] = a
    }
    
    var sum = 0
    for (d_a, d_b) in amicablePairs {
        sum += d_a + d_b
    }
    // Overcounted each number twice
    sum /= 2
    return sum
}


/*
func d(x uint64) (sum uint64) {
        for div := range misc.ProperDivisors(x) {
                sum += div
        }
        return
}

func sumOfAmicableNumbersUnder(max uint64) uint64 {
        amicablePairs := make(map[uint64]uint64)
        for a := uint64(1); a < max; a++ {
                d_a := d(a)
                // a != b
                if a == d_a {
                        continue
                }

                // not amicable?
                if d_b := d(d_a); d_b != a {
                        continue
                }

                // amicable pair
                if _, ok := amicablePairs[a]; !ok {
                        if _, ok = amicablePairs[d_a]; !ok {
                                fmt.Printf("a = %v   d_a = %v\n", a, d_a)
                                amicablePairs[a] = d_a
                        }
                }
        }

        sum := uint64(0)
        for d_a, d_b := range amicablePairs {
                //fmt.Printf("d_a = %v, d_b = %v\n", d_a, d_b)
                sum += d_a + d_b
        }
        return sum
}
*/
