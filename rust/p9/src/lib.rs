/*
 * Special Pythagorean triplet
 *
 * A Pythagorean triplet is a set of three natural numbers, a < b < c, for which, a^2 + b^2 = c^2
 * For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
 *
 * There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 * Find the product abc.
 */


#![feature(test)]
extern crate test;

/*
 * a^2 + b^2 = c^2
 * find triplet such that a + b = c = 1000
 *
 * c = (a^2 + b^2)^.5
 * a + b + (a^2 + b^2)^.5 = 1000
 */
pub fn solve() -> u64 {
    for b in 1..1000 {
        let mut a = 1;
        while a < b {
            let cc = (a as f64) * (a as f64) + (b as f64) * (b as f64);
            let c = cc.sqrt();
            if c.floor() == c {
                if a + b + (c as u64) == 1000 {
                    return a * b * (c as u64);
                }
            }
            a = a + 1;
        }
    }
    panic!();
}

#[test]
fn it_works() {
    assert_eq!(31875000, solve())
}
