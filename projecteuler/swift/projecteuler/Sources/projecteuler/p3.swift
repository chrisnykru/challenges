/*

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

enum ProjectEulerError: Error {
    case outOfRange
}

// Results are not sorted
func factors(_ x: Int) throws -> [Int] {
    guard x > 0 else {
        throw ProjectEulerError.outOfRange
    }
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

// Results are not sorted
func primeFactors(_ x: Int) throws -> [Int] {
    var pf: [Int] = []
    let allf = try factors(x)
    for factor in allf {
        let tmp = try factors(factor)
        if tmp.count <= 2 {
            pf.append(factor)
        }
    }
    return pf
}

func largestPrimeFactor(_ x: Int) throws -> Int {
    return try primeFactors(x).max()!
}
