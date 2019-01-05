#=

10001st prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
we can see that the 6th prime is 13.

What is the 10001st prime number?

=#

using Primes
using Test

function findPrime(nth)
  x = 0
  ith = 0
  while true
    if isprime(x)
	  ith += 1
	  if ith == nth
	    return x
	  end
	end
	x += 1
  end
end

@time @test findPrime(6) == 13
@time @test findPrime(10001) == 104743
