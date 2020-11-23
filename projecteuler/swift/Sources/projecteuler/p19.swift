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

func leapYear(_ year: Int) -> Bool {
    if year % 400 == 0 {
        return true
    } else if year % 100 == 0 {
        return false
    } else if year % 4 == 0 {
        return true
    }
    return false
}

func daysInMonth(month: Int, year: Int) throws -> Int {
    switch month {
    case 2: // Febuarary
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
    default:
        // bad month
        throw ProjectEulerError.internalError
    }
}

struct date: Equatable {
    var year: Int = 1900
    var month: Int = 1 // 1 = Jan
    var dayOfMonth: Int = 1
    // 1900-Jan-1 was a Monday
    var dayOfWeek: Int = 1 // 1 = monday, 7 = sunday
    
    func next() throws -> date {
        var d2: date = date()
        
        d2.dayOfWeek = (dayOfWeek % 7) + 1
        let totalDaysInCurrentMonth = try daysInMonth(month: month, year: year)
        // same year and month?
        if dayOfMonth < totalDaysInCurrentMonth {
            d2.year = year
            d2.month = month
            d2.dayOfMonth = dayOfMonth + 1
        }
        // same year
        else if month < 12 {
            d2.year = year
            d2.month = month + 1
            d2.dayOfMonth = 1
        }
        // everything wraps
        else {
            d2.year = year + 1
            d2.month = 1
            d2.dayOfMonth = 1
        }
        return d2
    }
}

// intentionally not using any standard library stuff
func sundaysOnFirstOfMonth() throws -> Int {
    var d1 = date()
    
    // iterate until we reach start of 20th century
    while d1.year != 1901 || d1.month != 1 || d1.dayOfMonth != 1 {
        d1 = try d1.next()
    }
    // d1 is now at 1901-Jan-1
    var sundayOnFirstCount = 0
    while d1.year != 2000 || d1.month != 12 || d1.dayOfMonth != 31 {
        if d1.dayOfMonth == 1 && d1.dayOfWeek == 7 {
            sundayOnFirstCount += 1
        }
        d1 = try d1.next()
    }
    return sundayOnFirstCount
}
