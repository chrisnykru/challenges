/*
 * Pandigital products
 *
 * We shall say that an n-digit number is pandigital if it makes use of all the
 * digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1
 * through 5 pandigital.
 *
 * The product 7254 is unusual, as the identity, 39 * 186 = 7254, containing
 * multiplicand, multiplier, and product is 1 through 9 pandigital.
 *
 * Find the sum of all products whose multiplicand/multiplier/product identity
 * can be written as a 1 through 9 pandigital.
 *
 * HINT: Some products can be obtained in more than one way so be sure to only
 * include it once in your sum.
 */
extern crate p3;

pub struct Pandigital {
    digit_counts: Vec<usize>,
}

impl Pandigital {
    pub fn new() -> Pandigital {
        Pandigital{
            digit_counts: vec![0;10],
        }
    }

    pub fn take(&mut self, x: usize) {
        let mut x = x;
        while x > 0 {
            let digit = x % 10;
            self.digit_counts[digit] += 1;
            x /= 10;
        }
    }

    pub fn check(self) -> Option<usize> {
        if self.digit_counts[0] != 0 || self.digit_counts[1] != 1 {
            return None;
        }
        let mut high = 1;
        for i in 1..10 {
            match self.digit_counts[i] {
                0 => {},
                1 => high = i,
                _ => return None // > 1
            }
        }

        let mut count_changes = 0; // 0 to 1, 1 to 0
        for i in 2..10 {
            if self.digit_counts[i] != self.digit_counts[i-1] {
                count_changes += 1;
            }
        }
        if count_changes > 1 {
            return None;
        }
        Some(high)
    }
}

fn identity_pan1to9(x: usize) -> bool {
    let divisors = p3::proper_divisors(x as u64);

    for d in divisors {
        let d = d as usize;
        let mut p = Pandigital::new();
        p.take(x);
        p.take(d);
        p.take(x / d);
        match p.check() {
            Some(y) => if y == 9 { return true; },
            _ => {},
        }
    }
    return false;
}

pub fn solve() -> usize {
    let mut sum = 0;
    for i in 1234..(12345+1) {
        if identity_pan1to9(i) {
            sum += i;
        }
    }
    sum
}

#[test]
fn test_pandigital() {
    let mut p = Pandigital::new();
    p.take(123);
    assert_eq!(p.check().unwrap(), 3);

    let mut p = Pandigital::new();
    p.take(1230);
    assert_eq!(p.check(), None);

    let mut p = Pandigital::new();
    p.take(1223);
    assert_eq!(p.check(), None);

    let mut p = Pandigital::new();
    p.take(1235);
    assert_eq!(p.check(), None);

    let mut p = Pandigital::new();
    p.take(1223);
    assert_eq!(p.check(), None);

    let mut p = Pandigital::new();
    p.take(123456789);
    assert_eq!(p.check().unwrap(), 9);

    let mut p = Pandigital::new();
    p.take(15234);
    assert_eq!(p.check().unwrap(), 5);
}

#[test]
fn test_identity_pan1to9() {
    assert_eq!(identity_pan1to9(7254), true);
}

#[test]
fn test_solve() {
    assert_eq!(solve(), 45228);
}
