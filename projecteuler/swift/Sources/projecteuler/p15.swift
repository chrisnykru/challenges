/*

Lattice paths

Starting in the top left corner of a 2x2 grid, there are 6 routes
(without backtracking) to the bottom right corner.

How many routes are there through a 20x20 grid?

https://projecteuler.net/project/images/p015.png
*/

import BigInt

func factorial(_ n: BigInt) -> BigInt{
    return n < 2 ? 1 : n * factorial(n-1)
}

// Square grid
func numRoutes(squareSide: Int) -> BigInt {
    // Note: r1,r2,d1,d2 is the same route as r2,r1,d2,d1
    
    let n = squareSide * 2
    let k = squareSide
    
    // TODO: get factorial!
    let nfac = factorial(BigInt(n))
    let kfac = factorial(BigInt(k))
    let n_minus_k_fac = factorial(BigInt(n - k))
    
    // C(n,k) = n! / (k! * (n-k)!
    return nfac / (kfac * n_minus_k_fac)
}
