include("1.jl")

using Base.Test

function test()
	@test sumOfMultiples(10-1) == 23
	@test sumOfMultiples(1000-1) == 233168
end

test()