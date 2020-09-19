/*

Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91 * 99.

Find the largest palindrome made from the product of two 3-digit numbers.

*/

import Darwin

func isPalindrome(_ x: Int) -> Bool {
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

struct PalindromeProduct: Equatable {
    var i = 0
    var j = 0
    var product = 0
}

func findLargestPalindrome(_ numDigits: Int) -> PalindromeProduct {
    var largest = PalindromeProduct(i: 0, j: 0, product: 0)
    
    let max = Int(pow(Double(10), Double(numDigits))) - 1
    // remember, communitive property of multiplication
    // don't test 9*8 and 8*9
    var i = max
    while i > 0 {
        var j = i
        while j > 0 {
            let ij = i * j
            if isPalindrome(ij) {
                if ij > largest.product {
                    largest.i = i
                    largest.j = j
                    largest.product = ij
                }
            }
            j -= 1
        }
        i -= 1
    }
    return largest
}

