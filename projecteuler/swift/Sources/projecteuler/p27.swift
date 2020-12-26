/*

Quadratic primes

Euler published the remarkable quadratic formula:
  n^2 + n + 41

It turns out that the formula will produce 40 primes for the consecutive values
n = 0 to 39. However, when n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible
by 41, and certainly when n = 41, 41^2 + 41 + 41 is clearly divisible by 41.

Using computers, the incredible formula  n^2 - 79n + 1601 was discovered, which
produces 80 primes for the consecutive values n = 0 to 79. The product of the
coefficients, -79 and 1601, is -126479.

Considering quadratics of the form:
  n^2 + an + b, where |a|<1000 and |b|<1000
  where |n| is the modulus/absolute value of n
  e.g. |11| = 11 and |-4| = 4

Find the product of the coefficients, a and b, for the quadratic expression that
produces the maximum number of primes for consecutive values of n,
starting with n = 0.

*/

func primeyQuadratic(a: Int, b: Int, n: Int) -> Int {
    return n * n + a * n + b
}

func primesForConsecutiveN(a: Int, b: Int) -> [Int] {
    // precompute list of primes
    // note: approx 79k primes from [1,1e6]
    // we don't know how many consecutive values of n will generate primes,
    // so we just pick 1e6 as upper bound for our prime check lookup table.
    let primes = eratosthenesSieve(to: 1000000)
    let primeMap: [Int: Bool] = primes.reduce(into: [:], { result, next in
        result[next] = true
    })
    
    var n = 0
    var consecutiveN_primes: [Int] = []
    while true {
        let pq = primeyQuadratic(a: a, b: b, n: n)
        
        //let keyExists = dict[key] != nil
        if primeMap[pq] == nil {
            break // not prime -> consecutive sequence broken
        }
        consecutiveN_primes.append(pq)
        n += 1
    }
    return consecutiveN_primes
}

func optimalAB() -> (optimalA: Int, optimalB: Int, primes: [Int]) {
    var optimalA: Int = 0
    var optimalB: Int = 0
    var primes: [Int] = []
    
    for a in stride(from: -999, through: 999, by: 1) {
        for b in stride(from: -999, through: 999, by: 1) {
            let p = primesForConsecutiveN(a: a, b: b)
            if p.count > primes.count {
                primes = p
                optimalA = a
                optimalB = b
            }
        }
    }
    return (optimalA, optimalB, primes)
}


/*

 func optimalAB() (optimalA, optimalB int, primes []uint64) {
     for a := -999; a <= 999; a++ {
         for b := -999; b <= 999; b++ {
             p := primesForConsecutiveN(a, b)
             if len(p) > len(primes) {
                 primes = p
                 optimalA = a
                 optimalB = b
             }
         }
     }
     return
 }
 
 func Test(t *testing.T) {
     
 }


 
 */
