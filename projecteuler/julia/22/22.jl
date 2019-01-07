#=

Names scores

Using names.txt, a 46K text file containing over five-thousand first names,
begin by sorting it into alphabetical order. Then working out the alphabetical
value for each name, multiply this value by its alphabetical position in the
list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is
worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN
would obtain a score of 938 * 53 = 49714.

What is the total of all the name scores in the file?

=#

using DelimitedFiles
using Test

function nameScore(name::String)
  score = 0
  for c in name
    score += c - 'A' + 1
  end
  return score
end

function totalNameScores()
  # specify type to avoid "NAN" name from getting intepreted as a float64 NaN
  names = readdlm("names.txt", ',') #, ASCIIString)
  # vectorize so we can sort
  names = vec(names)
  sort!(names)
  totalScore = 0
  i = 1
  for name in names
    totalScore += i * nameScore(name)
    i += 1
  end
  return totalScore
end

@test nameScore("COLIN") == 53
@time @test totalNameScores() == 871198282
