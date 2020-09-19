/*

   Smallest multiple

   2520 is the smallest number that can be divided by each of the numbers
   from 1 to 10 without any remainder.

   What is the smallest positive number that is evenly divisible by all of
   the numbers from 1 to 20?

*/

// Keep incrementing y by 'x' amount until it's divisible by
// everything in [1..x]
func smallestMultipleEvenlyDivisible(_ x: Int) -> Int {
    var y = x;
    while true {
        var z = x - 1
        var ok = true
        while z > 1 {
            //print("\(x) \(y) \(z)")
            if y % z != 0 {
                ok = false
                break
            }
            z -= 1
        }
        if ok {
            return y
        }
        y += x
    }
}
