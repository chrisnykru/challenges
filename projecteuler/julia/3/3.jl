#=

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

=#

using Base.Test

function primeFactors(x)
  d = Dict(1 => true)
  if isprime(x)
    d[x] = true
  end
  for i in 1:round(Int, floor(sqrt(x)))
    if x%i == 0 && isprime(i)
	    d[i] = true
	  end
  end
  return d
end

function largestPrimeFactor(x)
  largest = 0
  for k in keys(primeFactors(x))
    if k > largest
      largest = k
    end
  end
  return largest
end

@time @test primeFactors(13195) == Dict(1=>true, 5=>true, 7=>true,13=>true,29=>true)
@time @test largestPrimeFactor(600851475143) == 6857