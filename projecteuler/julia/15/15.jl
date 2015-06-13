#=

Lattice paths

Starting in the top left corner of a 2x2 grid, there are 6 routes
(without backtracking) to the bottom right corner.

(note: see image p015.gif)

How many routes are there through a 20x20 grid?

=#

using Base.Test

# square grid, no backtracking
function numRoutes(gridSize)
  #=
  note: r1 r2 d1 d2 === r2 r1 d2 d1
  and r1 r2 d1 d2 === r2 r1 d1 d2
  and r1 r2 d1 d2 === r1 r2 d2 d1
  
  n P k = n!/(k!(n-k)!)
  
  XXX I don't understand my own logic anymore...
  why does this work?
  
  =#
  
  # need BigInt here to avoid factorial overflow error!
  n = big(gridSize + gridSize)
  k = big(gridSize)
  return div(factorial(n), factorial(k) * factorial(n-k))
end

@time @test numRoutes(2) == 6
@time @test numRoutes(20) == 137846528820