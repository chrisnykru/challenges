#=

Power digit sum

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?

=#

using Test

function sumOfDigits(base, exponent)
  x = big(base)^exponent
  sum = 0
  while x != 0
    sum += x % 10
	x = div(x, 10)
  end
  return sum
end

@time @test sumOfDigits(2,15) == 26
@time @test sumOfDigits(2,1000) == 1366
