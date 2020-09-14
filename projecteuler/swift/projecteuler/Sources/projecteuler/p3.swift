/*

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

func getFactors(_ x: Int) -> [Int] {
    var tmp = [Int:Bool]()
    
    let sqrt_x = Int(Double(x).squareRoot())
    var i = 1
    while i <= sqrt_x {
        if x % i == 0 {
            tmp[i] = true
            tmp[x/i] = true
        }
        i += 1
    }
    print(tmp)
    
    //// append x?
    return Array(tmp.keys)
}

/*
func primeFactors(_ x: Int) -> [Int] {
    var res: [Int] = []
    
    let sqrt_x = Int(x.squareRoot())
    var i = 1
    while i < sqrt_x {
        if x % i == 0 {
            
        }
        i++
    }
    
    /*
     
     there's a sqrt thing here
     take 30
     sqrt = 5.something
     so 1 to 5
     1 & 30, ok
     2 & 15; 2*15 mod 30 = 0, ok

     3 & 10; 3*10 mod 30 = 0, ok;
       isPrime on 3...easy..another sqrt search?
     
     */
    
    return res
}
 */

/*
func largestPrimeFactor(_ x: Int) -> Int {
    let pf = primeFactors(x)
    var largest = -1
    for pf_i in pf {
        if pf_i > largest {
            largest = pf_i
        }
    }
    return largest
}
 */

/*
package main

import (
        "fmt"
        "github.com/zeroshirts/challenges/projecteuler/go/misc"
)

func primeFactors(x uint64) (pf []uint64) {
        for div := range misc.Divisors(x) {
                if misc.IsPrime(div) {
                        pf = append(pf, div)
                }
        }
        return pf
}

func largest() uint64 {
        pf := primeFactors(600851475143)
        fmt.Println("prime factors:", pf)
        var x uint64
        for i := range pf {
                if pf[i] > x {
                        x = pf[i]
                }
        }
        return x
}

func main() {
        // XXX ??? Why is it that prime factors, when multiplied together, yield target?
        fmt.Println(largest())
}
*/
