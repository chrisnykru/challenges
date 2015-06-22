#=

Lexicographic permutations

A permutation is an ordered arrangement of objects. For example, 3124 is one
possible permutation of the digits 1, 2, 3 and 4. If all of the permutations
are listed numerically or alphabetically, we call it lexicographic order.

The lexicographic permutations of 0, 1 and 2 are:
  012   021   102   120   201   210

What is the millionth lexicographic permutation of the
digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

=#

using Base.Test

function solve()
  p = permutations([0 1 2 3 4 5 6 7 8 9])
  state = start(p)
  item = []
  for i in 1:1000000
    # next will throw exception after done() is true
	item, state = next(p, state)
  end
  return item
end

@time @test solve() == [2,7,8,3,9,1,5,4,6,0]