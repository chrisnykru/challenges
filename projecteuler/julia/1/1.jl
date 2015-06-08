#=

Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

=#

using Base.Test

# Find multiples of 3 and 5 in domain [1,maxval]
# Note convention: maxval is inclusive to domain!
function sumOfMultiples(maxval)
	sum = 0
	for i = 1:maxval
		if i % 3 == 0 || i % 5 == 0
		  sum += i
		end
	end
	return sum
end

@time @test sumOfMultiples(10-1) == 23
@time @test sumOfMultiples(1000-1) == 233168