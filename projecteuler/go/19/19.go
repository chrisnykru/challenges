/*

Counting Sundays

You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.

Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.

A leap year occurs on any year evenly divisible by 4, but not on a century
unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century
(1 Jan 1901 to 31 Dec 2000)?

*/
package main

import (
	"fmt"
)

func leapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}

func daysInMonth(month, year int) int {
	switch month {
	case 2:
		// February
		if leapYear(year) {
			return 29
		} else {
			return 28
		}

	case 1, 3, 5, 7, 8, 10, 12:
		return 31

	case 9, 4, 6, 11:
		// September, April, June, November
		return 30
	}
	panic("bad month")
}

type date struct {
	year      int
	month     int // 1 = jan
	day       int // day of month
	dayOfWeek int // 1 = monday, 7 = sunday
}

// 1900-Jan-1 was a Monday
func reference() *date {
	return &date{year: 1900, month: 1, day: 1, dayOfWeek: 1}
}

func (d *date) next() *date {
	d2 := new(date)

	// day of week
	d2.dayOfWeek = (d.dayOfWeek % 7) + 1

	// day of month
	totalDaysInMonth := daysInMonth(d.month, d.year)

	// same year+month
	if d.day < totalDaysInMonth {
		d2.year = d.year
		d2.month = d.month
		d2.day = d.day + 1
		return d2
	}
	// same year
	if d.month < 12 {
		d2.year = d.year
		d2.month = d.month + 1
		d2.day = 1
		return d2
	}
	// everything wraps
	d2.year = d.year + 1
	d2.month = 1
	d2.day = 1
	return d2
}

// intentionally not using any standard library stuff
func sundaysOnFirstOfMonth() (count int) {
	d1 := reference()
	for d1.year != 1901 || d1.month != 1 || d1.day != 1 {
		d1 = d1.next()
	}

	// d1 is now at 1901-Jan-1
	fmt.Printf("start date = %v\n", d1)

	for d1.year != 2000 || d1.month != 12 || d1.day != 31 {
		if d1.day == 1 && d1.dayOfWeek == 7 {
			count++
		}
		d1 = d1.next()
	}
	return
}

// intentionally not using any standard library stuff
func main() {
	count := sundaysOnFirstOfMonth()
	fmt.Printf("sundays on first of month = %v\n", count)
}
