/*

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

func factors(_ x: Int) -> [Int] {
    var res = [Int:Bool]()
    
    let sqrt_x = Int(Double(x).squareRoot())
    var i = 1
    while i <= sqrt_x {
        if x % i == 0 {
            res[i] = true
            res[x/i] = true
        }
        i += 1
    }
    return Array(res.keys)
}

func primeFactors(_ x: Int) -> [Int] {
    var pf: [Int] = []
    let allf = factors(x)
    for factor in allf {
        let tmp = factors(factor)
        if tmp.count <= 2 {
            pf.append(factor)
        }
    }
    return pf
}

func largestPrimeFactor(_ x: Int) -> Int {
    var max = -1
    for pf in primeFactors(x) {
        if pf > max {
            max = pf
        }
    }
    return max
}
