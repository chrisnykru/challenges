#=

Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

=#

using Base.Test


function findTripletProduct(sum)
  #=
  a^2 + b^2 = c^2
  find triplet such that a + b = c = 1000

  c = (a^2 + b^2)^.5
  a + b + (a^2 + b^2)^.5 = 1000
  =#
  for b in 1:1000
    for a in 1:1000-b
	  cFloat = sqrt(a^2 + b^2)
	  if cFloat != floor(cFloat)
	    continue
	  end
	  c = round(Int, floor(cFloat))
	  if a+b+c == sum
	    return a*b*c
	  end
	end
  end
  error("unexpected")
end

@time @test findTripletProduct(1000) == 31875000