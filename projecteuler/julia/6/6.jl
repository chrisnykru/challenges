#=

Sum square difference

The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.

=#

using Base.Test

function naive(n)
  sumofsquares = 0
  squareofsums = 0
  for i in 1:n
    sumofsquares += i^2
	  squareofsums += i
  end
  squareofsums *= squareofsums
  return abs(sumofsquares - squareofsums)
end

function smarter(n)
  #=
  Sum(0,n) i = n(n+1)/2
  Sum(0,n) i^2 = n^3/3 + n^2/2 + n/6
  
  answer = n^3/3 + n^2/2 + n/6 - (n(n+1)/2)^2
  answer = -n^4/4 - n^3/6 + n^2/4 + n/6
  =#
  return abs(-1*n^4/4 - n^3/6 + n^2/4 + n/6)
end

@time @test naive(10) == 2640
@time @test naive(100) == 25164150
@time @test smarter(10) == 2640
@time @test smarter(100) == 25164150