/*
 * Digit fifth powers
 *
 * Surprisingly there are only three numbers that can be written as the sum of
 * fourth powers of their digits:
 *
 * 1634 = 1^4 + 6^4 + 3^4 + 4^4
 * 8208 = 8^4 + 2^4 + 0^4 + 8^4
 * 9474 = 9^4 + 4^4 + 7^4 + 4^4
 *
 * As 1 = 1^4 is not a sum it is not included.
 *
 * The sum of these numbers is 1634 + 8208 + 9474 = 19316.
 *
 * Find the sum of all the numbers that can be written as the sum of
 * fifth powers of their digits.
 */
extern crate num;

fn sum_digits_to_pow(num: usize, pow: usize) -> usize {
    if pow < 1 {
        panic!("bad pow");
    }
    let mut sum = 0;
    let mut num = num;
    while num > 0 {
        let digit = num % 10;
        let x = num::pow(digit, pow);
        sum += x;
        num /= 10;
    }
    sum
}

/*
 * 1^5 + 1^5 + 1^5 + 1^5 = 4
 * 2^5 (32) + 2^
 *
 *
 */
pub fn solve(pow: usize) -> usize {
    let mut sum = 0;
    // sum requires 2 digits, so start at 10
    // from experimentation, we stop at 1000000
    // need insight on when to stop iterating
    for i in 10..1000000 {
        let x = sum_digits_to_pow(i, pow);
        if x == i {
            println!("i={}", i);
            sum += x;
        }
    }
    sum
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(solve(4), 19316);
        assert_eq!(solve(5), 443839);
    }
}
