/*
 * Number letter counts
 *
 * If the numbers 1 to 5 are written out in words: one, two, three, four, five,
 * then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
 *
 * If all the numbers from 1 to 1000 (one thousand) inclusive were written out
 * in words, how many letters would be used?
 *
 * NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
 * forty-two) contains 23 letters and 115 (one hundred and fifteen) contains
 * 20 letters. The use of "and" when writing out numbers is in compliance with
 * British usage.
 */

fn num_letters(s: String) -> usize {
    let mut count = 0;
    for c in s.chars() {
        if c >= 'a' && c <= 'z' {
            count += 1;
        }
    }
    count
}

fn digit_to_string(x: usize) -> &'static str {
    match x {
        0 => "zero",
        1 => "one",
        2 => "two",
        3 => "three",
        4 => "four",
        5 => "five",
        6 => "six",
        7 => "seven",
        8 => "eight",
        9 => "nine",
        _ => panic!("bad digit")
    }
}

fn num_to_string(x: usize) -> String {
    if x < 1 || x > 1000 {
        panic!("bad number");
    }
    let mut x = x;
    let mut s = String::new();
    while x > 0 {
        if x >= 1000 {
            if s.len() > 0 {
                s.push_str(" ");
            }
            s.push_str(digit_to_string(x / 1000));
            s.push_str(" thousand");
            x = x % 1000;
            continue;
        }

        if x >= 100 {
            if s.len() > 0 {
                s.push_str(" ");
            }
            s.push_str(digit_to_string(x / 100));
            s.push_str(" hundred");
            x = x % 100;
            if x > 0 {
                s.push_str(" and");
            }
            continue;
        }

        // [20,99]
        if x >= 20 {
            if s.len() > 0 {
                s.push_str(" ");
            }
            s.push_str(
                match x / 10 {
                    9 => "ninety",
                    8 => "eighty",
                    7 => "seventy",
                    6 => "sixty",
                    5 => "fifty",
                    4 => "forty",
                    3 => "thirty",
                    2 => "twenty",
                    _ => panic!("!!!")
                });
            x = x % 10;
            continue;
        }

        // [10,19]
        if x >= 10 {
            if s.len() > 0 {
                s.push_str(" ");
            }
            s.push_str(
                match x {
                    19 => "nineteen",
                    18 => "eighteen",
                    17 => "seventeen",
                    16 => "sixteen",
                    15 => "fifteen",
                    14 => "fourteen",
                    13 => "thirteen",
                    12 => "twelve",
                    11 => "eleven",
                    10 => "ten",
                    _ => panic!("!!!")
                });
            break // done
        }

        // [1,9]
        if s.len() > 0 {
            s.push_str(" ");
        }
        s.push_str(digit_to_string(x));
        break;
    }
    s
}

pub fn solve() -> usize {
    let mut count = 0;
    for i in 1..1001 {
        let s = num_to_string(i);
        count = count + num_letters(s);
    }
    count
}

#[test]
fn it_works() {
    assert_eq!(4, num_letters(String::from("five")));
    assert_eq!(9, num_letters(String::from("twenty-two")));
    assert_eq!(17, num_letters(String::from("one hundred and five")));

    assert_eq!(String::from("one hundred and fifteen"), num_to_string(115));
    assert_eq!(String::from("three hundred and forty two"), num_to_string(342));


    assert_eq!(21124, solve());
}
