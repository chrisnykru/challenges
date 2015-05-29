include("1.jl")

using Base.Test

function test()
	@test sumOfMultiples(10) == 23
	@test sumOfMultiples(1000) == 233168
end

test()