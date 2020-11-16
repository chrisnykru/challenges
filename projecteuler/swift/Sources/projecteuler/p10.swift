/*

Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

*/

func findSumOfPrimesBelow2Mill() -> Int {
    let x = eratosthenesSieve(to: 2000000)
    return x.reduce(0, +)
}
