#=

Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

=#

using Base.Test

function sumOfPrimesBelow(ceil)
  sum = 0
  for i in 1:ceil-1
    if isprime(i)
	    sum += i
	  end
  end
  return sum
end

@time @test sumOfPrimesBelow(10) == 17
@time @test sumOfPrimesBelow(2000000) == 142913828922