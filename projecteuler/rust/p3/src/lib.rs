/*

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

#![feature(test)]
extern crate test;

extern crate num;
use std::collections::HashSet;

const F64_AS_U64_MAX: u64 = (1 << 53) - 1;

// Uses IEEE-754 sqrt() for x < F64_AS_U64_MAX
pub fn isqrt(n: u64) -> u64 {
    // fast path
    if n < F64_AS_U64_MAX {
        return (n as f64).sqrt() as u64;
    }
    // slow path
    if n == 0 || n == std::u64::MAX {
        // prevent divide by zero
        return n >> 32;
    }
    // Newton's method
    // TODO: talk to nisha
    let mut x_k = n;
    loop {
        let x_k_plus_1 = (x_k + n/x_k) / 2;
        if x_k_plus_1 >= x_k {
            return x_k_plus_1;	
        }
        x_k = x_k_plus_1;
    }
}

// A positive proper divisor is a positive divisor of a number n, excluding n itself.
pub fn proper_divisors(x: u64) -> HashSet<u64> {
    let mut div = HashSet::new();
    div.insert(1);

    let ubound = isqrt(x);
    for i in 2..ubound+1 {
        if x % i == 0 {
            div.insert(i);
            div.insert(x/i);
        }
    }

    return div
}

pub fn solve(x: u64) -> u64 {
    let mut largest = 0;
    for i in proper_divisors(x) {
        let tmp = proper_divisors(i);
        if tmp.len() == 1 { // implies prime
            if i > largest {
                largest = i;
            }
        }
    }
    assert!(largest != 0);
	largest
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::collections::HashSet;
    use test::Bencher;

    #[test]
    fn isqrt_works() {
        assert_eq!(1, isqrt(1));
        assert_eq!(1, isqrt(2));
        assert_eq!(1, isqrt(3));
        assert_eq!(2, isqrt(4));
        assert_eq!(2, isqrt(5));
        assert_eq!(3, isqrt(15));	
        assert_eq!(4, isqrt(16));
        assert_eq!(4, isqrt(17));

        assert_eq!(99, isqrt(9999));
        assert_eq!(100, isqrt(10000));
        assert_eq!(100, isqrt(10001));

        assert_eq!(94906265, isqrt((1<<53) - 1));
        assert_eq!(94906265, isqrt(1<<53));
        assert_eq!(94906265, isqrt((1<<53) + 1));

        assert_eq!(::std::u16::MAX as u64, isqrt(::std::u32::MAX as u64 - 1));
        assert_eq!(::std::u16::MAX as u64, isqrt(::std::u32::MAX as u64));
        assert_eq!(::std::u32::MAX as u64, isqrt(::std::u64::MAX - 1));
        assert_eq!(::std::u32::MAX as u64, isqrt(::std::u64::MAX));
    }

    #[test]
    fn proper_divisors_works() {
        let want: HashSet<_> = [1,2].iter().cloned().collect();
        assert_eq!(want, proper_divisors(4));

        // shadows previous def.
        let want: HashSet<_> = [1,2,3].iter().cloned().collect();
        assert_eq!(want, proper_divisors(6));

        let want: HashSet<_> = [1,3].iter().cloned().collect();
        assert_eq!(want, proper_divisors(9));
    }

    #[test]
    fn solve_works() {
        assert_eq!(29, solve(13195));
        assert_eq!(6857, solve(600851475143));
    }

    #[bench]
    fn bench_solve(b: &mut Bencher) {
        b.iter(|| solve(600851475143));
    }
}
