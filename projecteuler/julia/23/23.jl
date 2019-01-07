#=

Non-abundant sums

A perfect number is a number for which the sum of its proper divisors is exactly
equal to the number. For example, the sum of the proper divisors of 28 would be
1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n
and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number
that can be written as the sum of two abundant numbers is 24. By mathematical
analysis, it can be shown that all integers greater than 28123 can be written as
the sum of two abundant numbers. However, this upper limit cannot be reduced any
further by analysis even though it is known that the greatest number that cannot
be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of
two abundant numbers.

=#

using Test

include("../misc/misc.jl")

function SumOfProperDivisors(x)
  sum = 0
  for pd in properDivisors(x)
    sum += pd
  end
  return sum
end

function solve()
  abundant = Int[]
  for i = 12:28123
    if SumOfProperDivisors(i) > i
	  push!(abundant, i)
	end
  end

  canBeWritten = Dict()
  for i in abundant
    for j in abundant
	  if i+j > 28123
	    break
	  end
	  canBeWritten[i+j] = true
	end
  end
  
  # smallest number that can be written as sum of two abundant numbers is 24
  cannotBeWrittenSum = sum(1:23)
  for i in 24:28123
    if !haskey(canBeWritten, i)
	  cannotBeWrittenSum += i
	end
  end
  
  return cannotBeWrittenSum
end

@time @test solve() == 4179871
