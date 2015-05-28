/*

Number letter counts

If the numbers 1 to 5 are written out in words: one, two, three, four, five,
then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out
in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
forty-two) contains 23 letters and 115 (one hundred and fifteen) contains
20 letters. The use of "and" when writing out numbers is in compliance with
British usage.

*/
package main

import (
	"fmt"
)

func numLetters(s string) int {
	count := 0
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			count++
		}
	}
	return count
}

func digitToWord(x int) string {
	if x < 0 || x > 9 {
		panic("bad digit")
	}
	switch x {
	case 9:
		return "nine"
	case 8:
		return "eight"
	case 7:
		return "seven"
	case 6:
		return "six"
	case 5:
		return "five"
	case 4:
		return "four"
	case 3:
		return "three"
	case 2:
		return "two"
	case 1:
		return "one"
	case 0:
		return "zero"
	}
	panic("!!!")
}

func numToString(x int) (s string) {
	if x < 1 || x > 1000 {
		panic("bad number")
	}

	for x > 0 {
		if x >= 1000 {
			if len(s) > 0 {
				s += " "
			}
			s += digitToWord(x/1000) + " thousand"
			x %= 1000
			// XXX would need 'and' in some cases, except 1000 is max for this problem
			continue
		}

		if x >= 100 {
			if len(s) > 0 {
				s += " "
			}
			s += digitToWord(x/100) + " hundred"
			x %= 100
			if x > 0 {
				s += " and"
			}
			continue
		}

		// [20,99]
		if x >= 20 {
			var s2 string
			switch x / 10 {
			case 9:
				s2 = "ninety"
			case 8:
				s2 = "eighty"
			case 7:
				s2 = "seventy"
			case 6:
				s2 = "sixty"
			case 5:
				s2 = "fifty"
			case 4:
				s2 = "forty"
			case 3:
				s2 = "thirty"
			case 2:
				s2 = "twenty"
			}
			if len(s) > 0 {
				s += " "
			}
			s += s2
			x %= 10
			continue
		}

		// [10,19]
		if x >= 10 {
			var s2 string
			switch x {
			case 19:
				s2 = "nineteen"
			case 18:
				s2 = "eighteen"
			case 17:
				s2 = "seventeen"
			case 16:
				s2 = "sixteen"
			case 15:
				s2 = "fifteen"
			case 14:
				s2 = "fourteen"
			case 13:
				s2 = "thirteen"
			case 12:
				s2 = "twelve"
			case 11:
				s2 = "eleven"
			case 10:
				s2 = "ten"
			}
			if len(s) > 0 {
				s += " "
			}
			s += s2
			// done
			break
		}

		// [1,9]
		if len(s) > 0 {
			s += " "
		}
		s += digitToWord(x)
		break
	}
	return
}

func letterCount_1_to_1000() int {
	count := 0
	for i := 1; i <= 1000; i++ {
		s := numToString(i)
		count += numLetters(s)
	}
	return count
}

func main() {
	count := letterCount_1_to_1000()
	fmt.Printf("letter count of [1,1000] = %v\n", count)
}
