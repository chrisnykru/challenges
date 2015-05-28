/*

Counting summations

It is possible to write five as a sum in exactly six different ways:

4 + 1
3 + 2
3 + 1 + 1
2 + 2 + 1
2 + 1 + 1 + 1
1 + 1 + 1 + 1 + 1

How many different ways can one hundred be written as a sum of at least two positive integers?

*/
package main

import (
	//"fmt"
)

/*

expect problem is intractable.. probably want to decompose
something like 6 into multiples...

6 = 1*6
6 = 2*3

1: 1 (1 term, not a sum)
6: 6 (1 term, not a sum)
2: 1+1
3: 2+1
   1+1+1

------------

Target = 6

5+1
4+2
4+1+1
3+2+1
3+1+1+1
2+2+2
2+2+1+1
2+1+1+1+1
1+1+1+1+1+1


Target = 3
2+1
1+1+1
--> 2 ways... target**(#ways) = 3**2 = 9.. likely bogus

*/

func main() {
}
