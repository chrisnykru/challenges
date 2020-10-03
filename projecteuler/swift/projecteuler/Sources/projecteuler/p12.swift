/*

Highly divisible triangular number

The sequence of triangle numbers is generated by adding the natural numbers.
So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28.

The first ten terms would be:
1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

Let us list the factors of the first seven triangle numbers:

 1: 1
 3: 1,3
 6: 1,2,3,6
10: 1,2,5,10
15: 1,3,5,15
21: 1,3,7,21
28: 1,2,4,7,14,28

We can see that 28 is the first triangle number to have over five divisors.

What is the value of the first triangle number to have over five hundred divisors?

*/

class TriangleNumGen {
    var a = 1
    var sum = 0
    
    func next() -> Int {
        sum += a
        a += 1
        return sum
    }
}

func findTriNumWithOver500Divisors() -> Int {
    return -1
    
    let g = TriangleNumGen()
    while true {
        let x = g.next()
        do {
            let f = try factors(x)
            if f.count > 500 {
                return x
            }
        } catch {
            return -1
        }
    }
    return -1
}
