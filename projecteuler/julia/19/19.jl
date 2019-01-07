#=

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

=#

using Dates
using Test

function numSundaysOnFirstOfMonth()
  d = Date(1901, 1, 1)
  stopAtDate = Date(2000, 12, 31)
  numSundays = 0
  while true
    if d > stopAtDate
	  break
	end
    if Dates.dayofweek(d) == 7 && Dates.dayofmonth(d) == 1  
	  numSundays += 1
	end
	d += Dates.Day(1)
  end
  return numSundays
end

# verify stdlib date function is consistent with problem
@test Dates.dayofweek(Date(1900, 1, 1)) == 1

@time @test numSundaysOnFirstOfMonth() == 171
