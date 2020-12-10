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
        print("**", self.remaining)
        
        // deal with overcounting from any duplicates
        for (_, count) in numCount {
            self.remaining /= factorial(count)
            print("  **", self.remaining)
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
    
    /*
     
     func (gen *PermGen) Next() (perm []int, last bool) {
         perm, last = gen.current()
         for j := len(gen.perm) - 2; j >= 0; j-- {
             for i := len(gen.perm) - 1; i > j; i-- {
                 if gen.perm[i] > gen.perm[j] {
                     // swap
                     gen.perm[i], gen.perm[j] = gen.perm[j], gen.perm[i]
                     // reverse
                     reverse(gen.perm[j+1:])
                     gen.countdown--
                     return
                 }
             }
         }
         return
     }
     */
}

/*
func NewPermGen(_ x: [Int]) throws -> PermGen {
    if x.count == 0 {
        throw ProjectEulerError.outOfRange
    }
    
    let numCount = x.reduce(into: [:]) { counts, elements in
        counts[elements, default: 1] += 1
    }
    
    var pg: PermGen = PermGen()
    pg.perm = x
    pg.perm.sort()
    
    // deal with overcounting from any duplicates
    for (_, count) in numCount {
        pg.remaining /= factorial(count)
    }
    
    return pg
}
 */


/*
 
 func New(x []int) *PermGen {
     if len(x) == 0 {
         panic("empty set")
     }

     xiCount := make(map[int]int)
     for i := range x {
         xiCount[x[i]]++
     }

     gen := &PermGen{
         perm:      make([]int, len(x)),
         countdown: factorial(len(x)),
     }
     copy(gen.perm, x)
     sort.Ints(gen.perm)

     // deal with any overcounting
     for _, count := range xiCount {
         gen.countdown /= factorial(count)
     }
     return gen
 }

 */


/*
 
 
 func (gen *PermGen) current() (perm []int, last bool) {
     perm = make([]int, len(gen.perm))
     copy(perm, gen.perm)
     return perm, gen.countdown == 1
 }

 func (gen *PermGen) Next() (perm []int, last bool) {
     perm, last = gen.current()
     for j := len(gen.perm) - 2; j >= 0; j-- {
         for i := len(gen.perm) - 1; i > j; i-- {
             if gen.perm[i] > gen.perm[j] {
                 // swap
                 gen.perm[i], gen.perm[j] = gen.perm[j], gen.perm[i]
                 // reverse
                 reverse(gen.perm[j+1:])
                 gen.countdown--
                 return
             }
         }
     }
     return
 }

 
 
 */
