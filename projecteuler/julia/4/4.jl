#=

Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91 * 99.

Find the largest palindrome made from the product of two 3-digit numbers.

=#

using Test

function IsPalindrome(x)
  xstr = string(x)
  for i in 1:div(length(xstr),2) # integer divison
    if xstr[i] != xstr[length(xstr)+1-i]
	    return false
	  end
  end
  return true
end

function largestPalindromeFromProduct(maxval)
  besti = 0
  bestj = 0
  bestij = 0
  # search from high to low
  for i in maxval:-1:1
    for j in maxval:-1:1
      ij = i * j
      if IsPalindrome(ij)
        if ij > bestij
          bestij = ij
          besti = i
          bestj = j
        end
      end
    end
  end
  return (besti, bestj, bestij)
end

@test IsPalindrome(9) == true
@test IsPalindrome(42) == false
@test IsPalindrome(44) == true
@test IsPalindrome(102) == false
@test IsPalindrome(929) == true
@test IsPalindrome(1212) == false

@time @test largestPalindromeFromProduct(999) == (993, 913, 906609)