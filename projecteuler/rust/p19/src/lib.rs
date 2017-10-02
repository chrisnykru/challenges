/*
 * Counting Sundays
 *
 * You are given the following information, but you may prefer to do some research for yourself.
 *
 * 1 Jan 1900 was a Monday.
 * Thirty days has September,
 * April, June and November.
 * All the rest have thirty-one,
 * Saving February alone,
 * Which has twenty-eight, rain or shine.
 * And on leap years, twenty-nine.
 *
 * A leap year occurs on any year evenly divisible by 4, but not on a century
 * unless it is divisible by 400.
 *
 * How many Sundays fell on the first of the month during the twentieth century
 * (1 Jan 1901 to 31 Dec 2000)?
 */

extern crate time;

// not used in p19, might be useful later
pub fn leap_year(year: usize) -> bool {
    if year % 400 == 0 {
        return true;
    } else if year % 100 == 0 {
        return false;
    } else if year % 4 == 0 {
        return true;
    }
    return false;
}

// not used in p19, might be useful later
pub fn days_in_month(month: usize, year: usize) -> usize {
    match month {
        2 => if leap_year(year) { 29 } else { 28 },
        1 | 3 | 5 | 7 | 8 | 10 | 12 => 31,
        9 | 4 | 6 | 11 => 30, // Sep, Apr, Jun, Nov
        _ => panic!("bad month!")
    }
}

pub fn solve() -> usize {
    let mut t = time::Tm {
        tm_sec: 0,
        tm_hour: 0,
        tm_year: 1,
        tm_isdst: 0,
        tm_mday: 1,
        tm_nsec: 0,
        tm_yday: 1,
        tm_min: 0,
        tm_mon: 0,
        tm_wday: 1,
        tm_utcoff: 0,
    };

    let mut count = 0;
    loop {
        if t.tm_wday == 6 && t.tm_mday == 1 {
            count = count + 1;
            let tm = t.asctime();
            println!("{}", tm);
        }
        if t.tm_mday == 31 && t.tm_year == 100 && t.tm_mon == 11 {
            break
        }
        t = t + time::Duration::days(1);
    }
    count
}

#[test]
fn it_works() {
    assert_eq!(171, solve());
}
