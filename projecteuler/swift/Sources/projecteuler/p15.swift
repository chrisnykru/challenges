/*

Lattice paths

Starting in the top left corner of a 2x2 grid, there are 6 routes
(without backtracking) to the bottom right corner.

How many routes are there through a 20x20 grid?

*/


func factorial(_ n: Int) -> Int{
    return n < 2 ? 1 : n * factorial(n-1)
}

// Square grid
func numRoutes(squareSide: Int) -> Int {
    // Note: r1,r2,d1,d2 is the same route as r2,r1,d2,d1
    
    let n = squareSide * 2
    let k = squareSide
    
    // TODO: get factorial!
    let nfac = factorial(n)
    let kfac = factorial(k)
    let n_minus_k_fac = factorial(n - k)
    
    return nfac / (kfac * n_minus_k_fac)
}

/*
package main

import (
        "fmt"
        "github.com/zeroshirts/challenges/projecteuler/go/misc"
        "math/big"
)

// square grid
// no backtracking allowed
func numRoutes(size int64) int64 {
        // note: r1,r2,d1,d2 is same route as r2,r1,d2,d1

        n := size + size
        k := size

        nfactorial := misc.Factorial(n)
        kfactorial := misc.Factorial(k)
        nMinuskfactorial := misc.Factorial(n - k)

        C := new(big.Int).Div(nfactorial, new(big.Int).Mul(kfactorial, nMinuskfactorial))
        return C.Int64()
}

func main() {
        fmt.Println(numRoutes(20))
}
*/
