/*

Lexicographic permutations

A permutation is an ordered arrangement of objects. For example, 3124 is one
possible permutation of the digits 1, 2, 3 and 4. If all of the permutations
are listed numerically or alphabetically, we call it lexicographic order.

The lexicographic permutations of 0, 1 and 2 are:
  012   021   102   120   201   210

What is the millionth lexicographic permutation of the
digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

*/

struct PermGen {
    var perm: [Int] = []
    var remaining: Int = 0
    
    init(_ x: [Int]) throws {
        if x.count == 0 {
            throw ProjectEulerError.outOfRange
        }
        
        let numCount = x.reduce(into: [:]) { counts, elements in
            counts[elements, default: 0] += 1
        }
        self.perm = x
        self.perm.sort()
        self.remaining = factorial(x.count)
        
        // deal with overcounting from any duplicates
        for (_, count) in numCount {
            self.remaining /= factorial(count)
        }
    }
    
    func get() -> ([Int], Bool) {
        return (self.perm, self.remaining == 1)
    }
    
    mutating func next() -> ([Int], Bool) {
        let (perm, last) = self.get()
        for j in stride(from: self.perm.count - 2, through: 0, by: -1) {
            for i in stride(from: self.perm.count - 1, to: j, by: -1) {
                if self.perm[i] > self.perm[j] {
                    self.perm.swapAt(i, j)
                    self.perm[j+1..<self.perm.count].reverse()
                    self.remaining -= 1
                    return (perm, last)
                }
            }
        }
        return (perm, last)
    }
}

func perm1e6() throws -> [Int] {
    var x = try PermGen([0,1,2,3,4,5,6,7,8,9])
    var perm: [Int] = []
    for _ in stride(from: 0, to: 1e6, by: 1) {
        (perm, _) = x.next()
    }
    return perm
}
