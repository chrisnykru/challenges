/*
 * Power digit sum
 *
 * 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
 *
 * What is the sum of the digits of the number 2^1000?
 */

#![feature(test)]
extern crate test;
extern crate num;

use num::bigint::{ToBigUint};

pub fn solve2(base: usize, exp: usize) -> usize {
    if exp == 0 {
        return 1;
    }

    let base = base.to_biguint().unwrap();

    let mut x = base.clone();
    for _ in 2..(exp+1) {
        x = x * &base;
    }
    let xstr = x.to_str_radix(10);
    let mut sum = 0;
    for c in xstr.chars() {
        let digit = (c as usize) - ('0' as usize);
        sum += digit;
    }
    sum
}

pub fn solve(base: usize, exp: usize) -> usize {
    // TODO really need a bignum lib w/ Exp...
    // maybe I should write it... do a git pull?
    
    if exp == 0 {
        return 1;
    }

    let base = base.to_biguint().unwrap();

    let mut x = base.clone();
    for _ in 2..(exp+1) {
        x = x * base.clone();
    }
    let xstr = x.to_str_radix(10);
    let mut sum = 0;
    for c in xstr.chars() {
        let digit = (c as usize) - ('0' as usize);
        sum += digit;
    }
    sum
}

#[test]
fn it_works() {
    assert_eq!(26, solve(2, 15));
    assert_eq!(1366, solve(2, 1000));

    assert_eq!(26, solve2(2, 15));
    assert_eq!(1366, solve2(2, 1000));
}

#[cfg(test)]
mod tests {
    use super::*;
    use test::Bencher;

    #[bench]
    fn bench_solve(b: &mut Bencher) {
        b.iter(|| solve(2, 1000));
    }

    #[bench]
    fn bench_solve2(b: &mut Bencher) {
        b.iter(|| solve2(2, 1000));
    }
}

