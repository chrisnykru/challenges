#=

Even Fibonacci numbers

Each new term in the Fibonacci sequence is generated by adding the previous two
terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.

=#

## short-hand:
## fib(n) = n < 2 ? n : fib(n - 1) + fib(n - 2)

function sumOfEvenFibonacci(maxval)
	# according to problem convention
	f_iMinus1 = 2
	f_iMinus2 = 1
	sum = 2 # precompute terms 1 and 2
	while true
		f = f_iMinus1 + f_iMinus2
		##println(f, " ", f_iMinus1, " ", f_iMinus2)
		if f > maxval
			break
		elseif f % 2 == 0
			sum += f
		end
		f_iMinus2 = f_iMinus1
		f_iMinus1 = f
	end
	return sum
end