#=

1000-digit Fibonacci number

The Fibonacci sequence is defined by the recurrence relation:

F_n = F_(n-1) + F_(n-2), where F_1 = 1 and F_2 = 1.
Hence the first 12 terms will be:

F_1 = 1
F_2 = 1
F_3 = 2
F_4 = 3
F_5 = 5
F_6 = 8
F_7 = 13
F_8 = 21
F_9 = 34
F_10 = 55
F_11 = 89
F_12 = 144
The 12th term, F_12, is the first term to contain three digits.

What is the first term in the Fibonacci sequence to contain 1000 digits?

=#

using Test

### idea: fibonacci thingy that returns collect()able iteraterable stuff

function fibSolve(numDigits)
  target = big(10)^(numDigits-1)
  f_iMinus1 = big(1)
  f_iMinus2 = big(1)
  term = 2
  while true
    f = f_iMinus1 + f_iMinus2
	term += 1
	#println("term ", term, " f=", f)
	if f >= target
	  return term
	end
	f_iMinus2 = f_iMinus1
    f_iMinus1 = f
  end
end

@time @test fibSolve(3) == 12
@time @test fibSolve(1000) == 4782
