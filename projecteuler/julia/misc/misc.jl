#=

misc. utility functions

=#

using Base.Test

# XXX does this work w/ big shit?
function divisors(x)
  d = Dict(1 => true, x => true)
  for i in 2:round(typeof(x), floor(sqrt(x)))
    if mod(x, i) == 0
      d[i] = true
      d[div(x, i)] = true
    end
  end
  return d
end

function properDivisors(x)
  d = divisors(x)
  pop!(d, x)
  return d
end

@test divisors(28) == Dict(1 => true, 28 => true, 2 => true, 4 => true, 7 => true, 14 => true)
@test divisors(big(24)) == Dict(1 => true, 24 => true, 2 => true, 3 => true, 4 => true, 6 => true, 8 => true, 12 => true)

@test properDivisors(28) == Dict(1 => true, 2 => true, 4 => true, 7 => true, 14 => true)