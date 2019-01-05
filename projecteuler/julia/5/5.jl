#=

Smallest multiple

2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of
the numbers from 1 to 20?

=#

using Test

# all numbers in [1,maxval]
function findSmallest(maxval)
  if maxval < 1
    error("maxval < 1")
  end
  x = maxval
  while true
    ok = true
	  for i in 1:maxval
	    # note: % is potentially unsafe when x > 2^53?
	    if x % i != 0
	      ok = false
		    break
	    end
	  end
	  if ok
	    break
	  end
	  x = x + 1
  end
  return x
end

@time @test findSmallest(10) == 2520
@time @test findSmallest(20) == 232792560
