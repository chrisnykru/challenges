#=

misc. utility functions

=#

using Base.Test

function divisors(x)
  d = Set([1, x])
  for i in 2:round(typeof(x), floor(sqrt(x)))
    if mod(x, i) == 0
      push!(d, i)
      push!(d, div(x, i))
    end
  end
  return d
end

function properDivisors(x)
  d = divisors(x)
  pop!(d, x)
  return d
end

@test divisors(28) == Set([1, 28, 2, 4, 7, 14])
@test divisors(big(24)) == Set([1, 24, 2, 3, 4, 6, 8, 12])

@test properDivisors(28) == Set([1, 2, 4, 7, 14])