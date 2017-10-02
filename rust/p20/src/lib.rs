/*
 * Factorial digit sum
 *
 * n! means n * (n-1) * ... * 3 * 2 * 1
 *
 * For example, 10! = 10 * 9 * ... * 3 * 2 * 1 = 3628800,
 * and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
 *
 * Find the sum of the digits in the number 100!
 */

extern crate num;
extern crate p15;

use num::bigint::ToBigUint;
use num::ToPrimitive;

pub fn solve(x: usize) -> usize {
    let mut f = p15::factorial(x);
    let ten = 10.to_biguint().unwrap();
    let zero = 0.to_biguint().unwrap();

    let mut count = 0;
    while f > zero {
        let digit = (f.clone() % ten.clone()).to_usize().unwrap();
        count = count + digit;
        f = f / ten.clone();
    }
    count
}

#[test]
fn it_works() {
    assert_eq!(648, solve(100));
}
