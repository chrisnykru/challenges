/*
 * Sum square difference
 *
 * The sum of the squares of the first ten natural numbers is,
 * 1^2 + 2^2 + ... + 10^2 = 385
 *
 * The square of the sum of the first ten natural numbers is,
 * (1 + 2 + ... + 10)^2 = 55^2 = 3025
 *
 * Hence the difference between the sum of the squares of the first ten natural
 * numbers and the square of the sum is 3025 - 385 = 2640.
 *
 * Find the difference between the sum of the squares of the first one hundred
 * natural numbers and the square of the sum.
 */

#![feature(test)]
extern crate test;

pub fn solve(x: i64) -> i64 {
    /*
     * TODO: re-do derivations by hand
     * Sum(1,n) k = k*(k+1)/2
     * Sum(1,n) k^2 = k^3/3 + k^2/2 + k/6
     * 
     * Sum of squares - square of sum:
     * = k^3/3 + k^2/2 + k/6 - (k^2/2 + k/2)^2
     * = k^3/3 + k^2/2 + k/6 - (k^4/4 + k^3/2 + k^2/4)
     * = k^3/3 + k^2/2 + k/6 - k^4/4 - k^3/2 - k^2/4
     * = -k^4/4 - k^3/6 + k^2/4 + k/6
     */
    (-1 * x * x * x * x / 4 - x * x * x / 6 + x * x / 4 + x / 6).abs()
}

#[test]
fn it_works() {
    assert_eq!(2640, solve(10));
    assert_eq!(25164150, solve(100));
}
