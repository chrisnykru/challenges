#=

Longest Collatz sequence

The following iterative sequence is defined for the set of positive integers:
  n -> n/2 (n is even)
  n -> 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
  13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10
terms. Although it has not been proved yet (Collatz Problem), it is thought that
all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.

=#

using Test

function nextCollatzTerm(x)
  if x % 2 == 0
    return div(x,2)
  else
    return 3*x + 1
  end
end

function sequenceLength(x)
  seqLen = 1
  while x != 1
    x = nextCollatzTerm(x)
	seqLen += 1
  end
  return seqLen
end

function findLongestSequence(stopAt)
  longestSeqNum = 0
  longestSeq = 0
  for i in 1:stopAt
    seqLen = sequenceLength(i)
	if seqLen > longestSeq
	  longestSeqNum = i
	  longestSeq = seqLen
	end
  end
  return (longestSeqNum, longestSeq)
end 

@time @test findLongestSequence(1000000) == (837799, 525)
