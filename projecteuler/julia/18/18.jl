#=

Maximum path sum I

By starting at the top of the triangle below and moving to adjacent numbers
on the row below, the maximum total from top to bottom is 23.

   3
  7 4
 2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

                            75
                          95  64
                        17  47  82
                      18  35  87  10
                    20  04  82  47  65
                  19  01  23  75  03  34
                88  02  77  73  07  63  67
              99  65  04  28  06  16  70  92
            41  41  26  56  83  40  80  70  33
          41  48  72  33  47  32  37  16  94  29
        53  71  44  65  25  43  91  52  97  51  14
      70  11  33  28  77  73  17  78  39  68  17  57
    91  71  52  38  17  14  91  43  58  50  27  29  48
  63  66  04  68  89  53  67  30  73  16  69  87  40  31
04  62  98  27  23  09  70  98  73  93  38  53  60  04  23

NOTE: As there are only 16384 routes, it is possible to solve this problem
by trying every route. However, Problem 67, is the same challenge with a
triangle containing one-hundred rows; it cannot be solved by brute force,
and requires a clever method! ;o)

=#

using Base.Test

function triSmall()
  s = """3
        7 4
       2 4 6
      8 5 9 3"""
  return IOBuffer(s)
end

function triBig()
  s = """                   75
                          95  64
                        17  47  82
                      18  35  87  10
                    20  04  82  47  65
                  19  01  23  75  03  34
                88  02  77  73  07  63  67
              99  65  04  28  06  16  70  92
            41  41  26  56  83  40  80  70  33
          41  48  72  33  47  32  37  16  94  29
        53  71  44  65  25  43  91  52  97  51  14
      70  11  33  28  77  73  17  78  39  68  17  57
    91  71  52  38  17  14  91  43  58  50  27  29  48
  63  66  04  68  89  53  67  30  73  16  69  87  40  31
04  62  98  27  23  09  70  98  73  93  38  53  60  04  23"""
  return IOBuffer(s)
end

# Returns NxN Array{Any,2}
# Values are Int64
# Empties are ""
function ParseTriangleNumbers(src::IO)
  m = readdlm(src)
  for i in 1:size(m)[1] # rows
    for j in 1:size(m)[2] # cols
      if i >= j
        if typeof(m[i,j]) != Int64
          error(@sprintf("typeof(m[%d,%d]) != Int64", i, j))
          error(@sprintf("typeof(m[%d,%d]) != Int64", i, j))
        end
      else # i < j
        if typeof(m[i,j]) != SubString{ASCIIString}
          error(@sprintf("typeof(m[%d,%d]) != SubString{ASCIIString}", i, j))
        end
        if m[i,j] != ""
          error(@sprintf("m[%d,%d] != \"\"", i, j))
        end
      end
    end
  end
  return m
end

function TriangleMergeUpAndReduce(m::Array{Any,2})
  if size(m)[1] != size(m)[2]
    error("expect square matrix, got ", size(m))
  end
  
  while size(m)[1] > 1
    rows = size(m)[1]
    cols = size(m)[2]
    for i in 1:cols-1
      if m[rows,i] > m[rows,i+1]
        m[rows-1,i] += m[rows, i]
      else
        m[rows-1,i] += m[rows, i+1]
      end
    end
    m = m[1:size(m)[1]-1, 1:size(m)[2]-1]
  end
  return m[1,1]
end

# Matrix numbers are not triangle shaped
@test_throws ErrorException ParseTriangleNumbers(IOBuffer("1\n2 3 4"))
@test_throws ErrorException ParseTriangleNumbers(IOBuffer("1\n2\n3 4 5"))
@test_throws ErrorException ParseTriangleNumbers(IOBuffer("1\n2 3 4\n5 6 7"))

# XXX test edge cases for mergeUpAndReduce
#@time @test_throws ErrorException TriangleMergeUpAndReduce(Array{Any, 3})

@time @test TriangleMergeUpAndReduce(ParseTriangleNumbers(triSmall())) == 23
@time @test TriangleMergeUpAndReduce(ParseTriangleNumbers(triBig())) == 1074