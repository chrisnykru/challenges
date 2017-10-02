/*
 * 1000-digit Fibonacci number
 *
 * The Fibonacci sequence is defined by the recurrence relation:
 *
 * F_n = F_(n-1) + F_(n-2), where F_1 = 1 and F_2 = 1.
 * Hence the first 12 terms will be:
 *
 * F_1 = 1
 * F_2 = 1
 * F_3 = 2
 * F_4 = 3
 * F_5 = 5
 * F_6 = 8
 * F_7 = 13
 * F_8 = 21
 * F_9 = 34
 * F_10 = 55
 * F_11 = 89
 * F_12 = 144
 *
 * The 12th term, F_12, is the first term to contain three digits.
 *
 * What is the first term in the Fibonacci sequence to contain 1000 digits?
 */

extern crate num;
use num::{BigUint};
use num::bigint::{ToBigUint};

// Similar to p2::FibIter, but with BigUint
pub struct BigFibIter {
    i: BigUint,
    i_next: BigUint,
}

impl BigFibIter {
    fn new() -> BigFibIter {
        BigFibIter{
            i: 0.to_biguint().unwrap(),
            i_next: 1.to_biguint().unwrap(),
        }
    }
}

impl Iterator for BigFibIter {
    type Item=BigUint;

    fn next(&mut self) -> Option<BigUint> {
        let old_i = self.i.clone();
        self.i = self.i_next.clone();
        self.i_next = self.i_next.clone() + old_i.clone();
        Some(old_i)
    }
}

pub fn solve(strlen: usize) -> usize {
    let mut it = BigFibIter::new(); 
    let mut term_num = 0;
    while let Some(term) = it.next() {
        let str = term.to_string();
        if str.len() >= strlen {
            return term_num;
        }
        term_num += 1;
    }
    panic!("no solution");
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(solve(1000), 4782);
    }
}
