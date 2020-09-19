/*

Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91  99.

Find the largest palindrome made from the product of two 3-digit numbers.

*/

func isPalindrome(x: Int) -> Bool {
    let xstr = String(x)
    let x = xstr.unicodeScalars.map { $0.value }
    
    var i = 0
    while i < x.count {
        if x[i] != x[x.count - 1 - i] {
            return false
        }
        i += 1
    }
    return true
}

struct PalindromeProduct {
    var i = 0
    var j = 0
    var product = 0
}


/*

type productPalindrome struct {
        i, j, ij int64
}

// [1,max]
func largestProductPalindrome(max int64) (pp productPalindrome) {
        // search from high to low
        for i := max; i > 0; i-- {
                for j := i; j > 0; j-- {
                        ij := i * j
                        if IsPalindrome(ij) {
                                if ij > pp.ij {
                                        pp = productPalindrome{i, j, ij}
                                }
                        }
                }
        }
        return
}

func main() {
        fmt.Printf("found %#v\n", largestProductPalindrome(999))
}
*/
