/*

Cubic permutations

The cube, 41063625 (345**3), can be permuted to produce two other cubes:
56623104 (384**3) and 66430125 (405**3). In fact, 41063625 is the smallest
cube which has exactly three permutations of its digits which are also cube.

Find the smallest cube for which exactly five permutations of its digits are
cube.

*/
package main

import (
	"fmt"
)

type lazyDeconstruct struct {
	numDigits [10]int // 0-9
}

func (ld *lazyDeconstruct) Set(x int) {
	if x < 0 {
		panic("x < 0")
	}
	// clear
	for i := 0; i < len(ld.numDigits); i++ {
		ld.numDigits[i] = 0
	}
	for {
		digit := x % 10
		ld.numDigits[digit]++
		x /= 10
		if x == 0 {
			break
		}
	}
}

type cubeVals struct {
	vals []int
}

func smallestCubeWith5Perm(numPermutations int) int {
	m := make(map[lazyDeconstruct]*cubeVals)
	var ld lazyDeconstruct

	for i := 1; ; i++ {
		cube := i * i * i
		ld.Set(cube)

		if cv, ok := m[ld]; !ok {
			cv = &cubeVals{[]int{cube}}
			m[ld] = cv
		} else {
			cv.vals = append(cv.vals, cube)
			if len(cv.vals) == numPermutations {
				fmt.Printf("cv = %v\n", cv)
				smallest := cv.vals[0]
				for j := 1; j < len(cv.vals); j++ {
					if cv.vals[j] < smallest {
						smallest = cv.vals[j]
					}
				}
				return smallest
			}
		}
	}
	panic("!")
}

func main() {
	fmt.Println(smallestCubeWith5Perm(5))
}
