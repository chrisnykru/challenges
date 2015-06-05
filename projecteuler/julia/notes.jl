
## XXX question: if I do integer/integer in Julia,
## and the first int is larger than 2^53, does it get inaccurately
## converted to float??...YES.

#=

julia> 2^56
72057594037927936

julia> Float64(2^56)
7.205759403792794e16

julia> (2^56)/2
3.602879701896397e16

julia> div(2^56,2)
36028797018963968

=#