#=

Amicable numbers

Let d(n) be defined as the sum of proper divisors of n (numbers less than n
which divide evenly into n).  If d(a) = b and d(b) = a, where a != b,
then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are:
  1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
  Therefore d(220) = 284

The proper divisors of 284 are:
  1, 2, 4, 71 and 142;
  so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

=#

using Base.Test

include("../misc/misc.jl")

function d(x)
  sum = 0
  for pd in properDivisors(x)
    sum += pd
  end
  return sum
end

function amicable(a)
  b = d(a)
  isAmicable = (d(b) == a && a != b)
  return (b, isAmicable) 
end


function sumOfAmicableNumbersLessThan(lessThanVal)
  sum = 0
  amicablePairs = Dict()
  for a in 1:lessThanVal-1
    b, isAmicable = amicable(a)
    if !isAmicable
      continue
    end
    if !haskey(amicablePairs, a) && !haskey(amicablePairs, b)
      amicablePairs[a] = b
    end
  end
  
  for a in keys(amicablePairs)
    sum += a + amicablePairs[a] 
  end
  return sum
end

@test d(220) == 284
@test d(284) == 220
@time @test sumOfAmicableNumbersLessThan(10000) == 31626