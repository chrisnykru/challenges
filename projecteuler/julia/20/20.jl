#=

Factorial digit sum

n! means n * (n  1) * ... * 3 * 2 * 1

For example, 10! = 10 * 9 * ... * 3 * 2 * 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!

=#

using Test

function sumDigits()
  x = factorial(big(100))
  sum = 0
  while x > 0
    sum += mod(x, 10)
    x = div(x, 10)
  end
  return sum
end

@time @test sumDigits() == 648
