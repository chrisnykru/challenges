#=

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

=#

using Primes
using Test

function largestPrimeFactor(x)
  largest = 0
  for k in keys(factor(x))
    if k > largest
      largest = k
    end
  end
  return largest
end

@time @test factor(13195) == Dict(7=>1, 13 => 1, 29 => 1, 5 => 1)
@time @test largestPrimeFactor(600851475143) == 6857