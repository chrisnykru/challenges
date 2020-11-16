/*

10001st prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
we can see that the 6th prime is 13.

What is the 10,001st prime number?

*/

import Darwin // just for log()

func getPrime(nth n: Int) throws -> Int {
    let numToConsider = 1 << 17
    let primes = eratosthenesSieve(to: numToConsider)
    guard primes.count >= n else {
        throw ProjectEulerError.internalError
    }
    return primes[n - 1]
}

// Code from 'Martin R' on:
// https://codereview.stackexchange.com/questions/192021/prime-number-generator-in-swift
func eratosthenesSieve(to n: Int) -> [Int] {
    var composite = Array(repeating: false, count: n + 1) // The sieve
    var primes: [Int] = []

    if n >= 150 {
        // Upper bound for the number of primes up to and including `n`,
        // from https://en.wikipedia.org/wiki/Prime_number_theorem#Non-asymptotic_bounds_on_the_prime-counting_function :
        let d = Double(n)
        let upperBound = Int(d / (log(d) - 4))
        primes.reserveCapacity(upperBound)
    } else {
        primes.reserveCapacity(n)
    }

    let squareRootN = Int(Double(n).squareRoot())
    var p = 2
    while p <= squareRootN {
        if !composite[p] {
            primes.append(p)
            for q in stride(from: p * p, through: n, by: p) {
                composite[q] = true
            }
        }
        p += 1
    }
    while p <= n {
        if !composite[p] {
            primes.append(p)
        }
        p += 1
    }
    return primes
}

