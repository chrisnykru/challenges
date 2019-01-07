#=

Number letter counts

If the numbers 1 to 5 are written out in words: one, two, three, four, five,
then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out
in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
forty-two) contains 23 letters and 115 (one hundred and fifteen) contains
20 letters. The use of "and" when writing out numbers is in compliance with
British usage.

=#

using Test

function numLetters(numberStr::String)
  count = 0
  for c in numberStr
    if c >= 'a' && 'c' <= 'z'
	  count += 1
	end
  end
  return count
end

function digitToWord(x)
  if x == 9
    return "nine"
  elseif x == 8
    return "eight"
  elseif x == 7
    return "seven"
  elseif x == 6
    return "six"
  elseif x == 5
    return "five"
  elseif x == 4
    return "four"
  elseif x == 3
    return "three"
  elseif x == 2
    return "two"
  elseif x == 1
    return "one"
  elseif x == 0
    return "zero"
  else
    error("bad digit: ", x)
  end
end

function numToString(x)
  if x < 1 || x > 1000
    error("out of range")
  end
  
  str = ""
  spacer = " "
  while x > 0
    ##println("str: ", str, " x=", x)
    if length(str) > 0
	  str *= spacer
	end
	
    if x == 1000
	  str *= digitToWord(div(x, 1000)) * " thousand"
	  x = mod(x, 1000)
	  
	elseif x >= 100 # [100,999]
	  str *= digitToWord(div(x, 100)) * " hundred"
	  x = mod(x, 100)
	  if x > 0
	    str *= " and"
	  end
	  
	elseif x >= 20 # [20,99]
	  y = div(x, 10)
	  x = mod(x, 10)
	  spacer = "-"
	  if y == 9
	    str *= "ninety"
	  elseif y == 8
	    str *= "eighty"
	  elseif y == 7
	    str *= "seventy"
	  elseif y == 6
	    str *= "sixty"
	  elseif y == 5
	    str *= "fifty"
	  elseif y == 4
	    str *= "forty"
	  elseif y == 3
	    str *= "thirty"
	  elseif y == 2
	    str *= "twenty"
	  else
	    error("unexpected")
	  end
	  
	elseif x >= 10
	  if x == 19
	    str *= "nineteen"
	  elseif x == 18
	    str *= "eighteen"
	  elseif x == 17
	    str *= "seventeen"
	  elseif x == 16
	    str *= "sixteen"
	  elseif x == 15
	    str *= "fifteen"
	  elseif x == 14
	    str *= "fourteen"
	  elseif x == 13
	    str *= "thirteen"
	  elseif x == 12
	    str *= "twelve"
	  elseif x == 11
	    str *= "eleven"
	  elseif x == 10
	    str *= "ten"
	  else
	    error("unexpected")
	  end
	  break # done
	 
	else # [1,9]
	  str *= digitToWord(x)
	  break # done
	end
  end
  return str
end

function letterCount()
  count = 0
  for i in 1:1000
    count += numLetters(numToString(i))
  end
  return count
end

@time @test numLetters("three hundred and forty-two") == 23
@time @test numLetters("one hundred and fifteen") == 20

@time @test numToString(342) == "three hundred and forty-two"
@time @test numToString(115) == "one hundred and fifteen"

@time @test letterCount() == 21124
