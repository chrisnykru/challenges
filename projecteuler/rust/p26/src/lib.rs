/*
 * Reciprocal cycles
 *
 * A unit fraction contains 1 in the numerator. The decimal representation of the
 * unit fractions with denominators 2 to 10 are given:
 *
 * 1/2	= 	0.5
 * 1/3	= 	0.(3)
 * 1/4	= 	0.25
 * 1/5	= 	0.2
 * 1/6	= 	0.1(6)
 * 1/7	= 	0.(142857)
 * 1/8	= 	0.125
 * 1/9	= 	0.(1)
 * 1/10	= 	0.1
 *
 * Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be
 * seen that 1/7 has a 6-digit recurring cycle.
 *
 * Find the value of d<1000 for which 1/d contains the longest recurring cycle
 * in its decimal fraction part.
 */

extern crate num;
use num::bigint::{ToBigUint};
use num::*;

pub fn cycle(d: usize) -> usize {
    /*
     * strategy by example:
     *   d = 7
     *   pow = 10**(7*2) = 100000000000000
     *   m = pow % d = 2
     *   x = m * pow / d = 28571428571428
     *   cycle length == 6
     */

    // num::pow is weakened by the fact that exp must be usize
    let pow: BigUint = num::pow(10.to_biguint().unwrap(), d * 2);
    let d = d.to_biguint().unwrap();
    let m = &pow % &d;
    if m == 0.to_biguint().unwrap() {
        return 0;
    }

    // Eliminate non-repreating prefix
    // question: how does multiplying by 'm' help us eliminate the non-repeating prefix in the
    // fraction?  I can't remember why this works.
    let x = m * pow / &d;
    let x_str = x.to_string();
    //println!("x_str = {}", x_str);

    // start substring matching from the right
    // this let's us skip any non-repeating 'prefix' on the left, e.g., .1666666
    for sublen in 1..x_str.len() / 2 + 1 {
        let mut ok = true;
        let sub1 = &x_str[x_str.len() - sublen .. x_str.len()];
        // pattern must extend throughout otherwise we DO get false positives
        for subgroup in 2..x_str.len() / sublen + 1 {
            let sub2 = &x_str[x_str.len() - subgroup * sublen .. x_str.len() - (subgroup-1) * sublen];
            if sub1 != sub2 {
                ok = false;
                break;
            }
        }
        if ok {
            //println!("  --> d={}  cycle_len={}", d.to_string(), sublen);
            return sublen;
        }
    }
    panic!("unsolved");
}

pub fn solve(dmax: usize) -> usize {
    let mut cycle_max = 0;
    let mut cycle_max_d = None;
    for d in 2..dmax {
        let cycle_len = cycle(d);
        if cycle_len > cycle_max {
            cycle_max = cycle_len;
            cycle_max_d = Some(d);
        }
    }
    cycle_max_d.unwrap()
}

#[cfg(test)]
mod test {
    extern crate num;
    use super::*;

    #[test]
    fn test_cycle() {
        assert_eq!(cycle(1), 0);
        assert_eq!(cycle(2), 0);
        assert_eq!(cycle(3), 1);
        assert_eq!(cycle(4), 0);
        assert_eq!(cycle(5), 0);
        assert_eq!(cycle(6), 1);
        assert_eq!(cycle(7), 6);
        assert_eq!(cycle(8), 0);
        assert_eq!(cycle(9), 1);
        assert_eq!(cycle(10), 0);
        assert_eq!(cycle(990), 2);
    }

    #[test]
    fn test_solve() {
        assert_eq!(solve(1000), 983);
    }
}
