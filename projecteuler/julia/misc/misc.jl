#=

misc. utility functions

=#

using Base.Test

# XXX does this work w/ big shit?
function divisors(x)
  div = Dict(1 => true, x => true)
  for i in 2:round(Int, floor(sqrt(x)))
    if x % i == 0
      div[i] = true
      div[x/i] = true
    end
  end
  return div
end

function properDivisors(x)
  div = divisors(x)
  pop!(div, x)
  return div
end

@test divisors(28) == Dict(1 => true, 28 => true, 2 => true, 4 => true, 7 => true, 14 => true)
@test divisors(big(24)) == Dict(1 => true, 24 => true, 2 => true, 4 => true, 6 => true, 12 => true)
@test properDivisors(28) == Dict(1 => true, 2 => true, 4 => true, 7 => true, 14 => true)