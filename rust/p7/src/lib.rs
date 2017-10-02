/*
 * 10001st prime
 *
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime
 * is 13.
 *
 * What is the 10,001st prime number?
 */

#![feature(test)]
extern crate test;
extern crate primal;

pub fn solve(x: usize) -> usize {
    // nth term is n-1 for this API
    primal::Primes::all().nth(x-1).unwrap()
}

#[test]
fn it_works() {
    assert_eq!(13, solve(6));
    assert_eq!(104743, solve(10001));
}
