/*

   Smallest multiple

   2520 is the smallest number that can be divided by each of the numbers
   from 1 to 10 without any remainder.

   What is the smallest positive number that is evenly divisible by all of
   the numbers from 1 to 20?

*/

#![feature(test)]
extern crate test;

// [1, x] inclusive
pub fn solve(x: u64) -> u64 {
    // get start val
    let mut i = 1;
    loop {
        if i % x == 0 {
            break;
        }
        i = i + 1;
    }

    'outer: loop {
        for j in 1..x+1 {
            if i % j != 0 {
                i = i + x;
                continue 'outer;
            }
        }
        break
    }
    i
}

#[test]
fn it_works() {
    assert_eq!(2520, solve(10));
    assert_eq!(232792560, solve(20));
}
