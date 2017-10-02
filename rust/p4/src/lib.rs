/*

Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91 * 99.

Find the largest palindrome made from the product of two 3-digit numbers.

*/

#![feature(test)]
extern crate test;
extern crate num;

pub fn is_palindrome(x: u64) -> bool {
	let mut v: Vec<u64> = Vec::new();
	let mut tmp = x;
	loop {
		let m = tmp % 10;
		v.push(m);
		tmp = tmp / 10;
		if tmp == 0 {
			break;
		}
	}

	let mut i = 0;
	while i < v.len() / 2 {
		if v[i] != v[v.len() - 1 - i] {
			return false;
		}
		i += 1;
	}
	return true;
}

pub fn solve(num_digits: usize) -> u64 {
	let upper = num::pow(10, num_digits) - 1; // 3 digits -> 999

	/*
	obviously a*b = b*a
	let's do better than trying all 999 * 999 -> 998801 values
	*/
	
	let mut largest = 0;
	let mut i = upper;
	while i > 0 {
		let mut j = i;
		while j > 0 {
			let ij = i * j;
			if ij > largest && is_palindrome(ij) {
				largest = ij;
			}
			j -= 1;	
		}
		i -= 1;
	}
	largest
}

#[cfg(test)]
mod tests {
    use super::*;
    use test::Bencher;

    #[test]
	fn is_palindrome_works() {
		assert_eq!(false, is_palindrome(100));
		assert_eq!(true, is_palindrome(101));
		assert_eq!(true, is_palindrome(999));
		assert_eq!(false, is_palindrome(1212));
		assert_eq!(false, is_palindrome(9008));
		assert_eq!(true, is_palindrome(9009));
		assert_eq!(true, is_palindrome(906609));
	}

	#[test]
	fn solve_works() {
		assert_eq!(9009, solve(2));
		assert_eq!(906609, solve(3));
	}

	#[bench]
    fn bench_solve(b: &mut Bencher) {
        b.iter(|| solve(3));
    }
}	
