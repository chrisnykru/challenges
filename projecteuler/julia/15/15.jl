#=

Lattice paths

Starting in the top left corner of a 2x2 grid, there are 6 routes
(without backtracking) to the bottom right corner.

(note: see image p015.gif)

How many routes are there through a 20x20 grid?

=#

using Test

# square grid, no backtracking
function numRoutes(gridSize)
  #=
  n = gridSize + gridSize
  n! permutations
  
  Overcounting by gridSize! * gridSize!
  e.g., r1 r2 d1 d2 === r2 r1 d2 d1
  and r1 r2 d1 d2 === r2 r1 d1 d2
  and r1 r2 d1 d2 === r1 r2 d2 d1
  
  =#
  n = big(gridSize + gridSize)
  p = factorial(n)
  return div(div(p, factorial(gridSize)), factorial(gridSize))
end

@time @test numRoutes(2) == 6
@time @test numRoutes(20) == 137846528820
