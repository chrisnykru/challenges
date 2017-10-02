/*
 * Summation of primes
 *
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 *
 * Find the sum of all the primes below two million.
 */

#![feature(test)]
extern crate test;
extern crate primal;

pub fn solve(max: usize) -> usize {
    let mut sum = 0;
    for p in primal::Primes::all() {
        if p >= max {
            break
        }
        sum = sum + p;
    }
    return sum;
}

#[test]
fn it_works() {
    assert_eq!(17, solve(10));
    assert_eq!(142913828922, solve(2000000));
}

#[bench]
fn bench_solve(b: &mut test::Bencher) {
    b.iter(|| solve(2000000));
}
