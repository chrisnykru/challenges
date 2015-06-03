include("2.jl")

using Base.Test

function test()
	@test sumOfEvenFibonacci(4e6) == 4613732
end

test()