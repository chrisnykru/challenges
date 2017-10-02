/*
 * Non-abundant sums
 *
 * A perfect number is a number for which the sum of its proper divisors is exactly
 * equal to the number. For example, the sum of the proper divisors of 28 would be
 * 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
 *
 * A number n is called deficient if the sum of its proper divisors is less than n
 * and it is called abundant if this sum exceeds n.
 *
 * As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number
 * that can be written as the sum of two abundant numbers is 24. By mathematical
 * analysis, it can be shown that all integers greater than 28123 can be written as
 * the sum of two abundant numbers. However, this upper limit cannot be reduced any
 * further by analysis even though it is known that the greatest number that cannot
 * be expressed as the sum of two abundant numbers is less than this limit.
 *
 * Find the sum of all the positive integers which cannot be written as the sum of
 * two abundant numbers.
 */

extern crate p21;

use std::collections::HashSet;

pub fn solve() -> usize {
	// everything over 28123 can be written as sum of two abundant numbers
    let mut abundant: Vec<usize> = vec![];
    for i in 12..28124 {
        // we've already implemented fn to compute sum of proper divisors
        // p21::d(x: usize) -> usize
        let sum = p21::d(i);
        if sum > i {
            abundant.push(i);
        }
    }

    let mut can_be_written: HashSet<usize> = HashSet::new();
    for i in 0..abundant.len() {
        // a + b == b + a, so start j iteration at i to reduce work
        for j in i..abundant.len() {
            if abundant[i] + abundant[j] <= 28123 {
                can_be_written.insert(abundant[i] + abundant[j]);
            } else {
                break;
            }
        }
    }

    let mut sum = 0;
    for i in 1..28124 {
        if ! can_be_written.contains(&i) {
            sum += i;
        }
    }
    sum
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(solve(), 4179871);
    }
}
