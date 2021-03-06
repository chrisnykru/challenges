/*

Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

*/


func pythagoreanTripletSumEquals1000() -> Int {
    /*
      a^2 + b^2 = c^2
      find triplet such that a + b + c = 1000, and: a < b < c

      c = (a^2 + b^2)^.5
      a + b + (a^2 + b^2)^.5 = 1000
    */
    
    for b in stride(from: 2, to: 998, by: 1) {
        for a in stride(from: 1, to: b, by: 1) {
            let cf = Double(a * a + b * b).squareRoot()
            // check if it's convertable to integer without precision loss
            var cfTmp = cf
            cfTmp.round(.down)
            if cf != cfTmp {
                continue // precision loss
            }
            let c = Int(cf)
            if b >= c {
                continue // violates a < b < c
            }
            if a + b + c == 1000 {
                return a * b * c
            }
        }
    }
    return -1
}
