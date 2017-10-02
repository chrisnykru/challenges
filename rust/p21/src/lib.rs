/*
 * Amicable numbers
 *
 * Let d(n) be defined as the sum of proper divisors of n (numbers less than n
 * which divide evenly into n).  If d(a) = b and d(b) = a, where a != b,
 * then a and b are an amicable pair and each of a and b are called amicable numbers.
 *
 * For example, the proper divisors of 220 are:
 * 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
 * Therefore d(220) = 284
 *
 * The proper divisors of 284 are:
 * 1, 2, 4, 71 and 142;
 * so d(284) = 220.
 *
 * Evaluate the sum of all the amicable numbers under 10000.
 */

extern crate p3;

use std::collections::HashMap;

pub fn d(x: usize) -> usize {
    let h = p3::proper_divisors(x as u64);
    let mut sum: usize = 0;
    for x in &h {
        // note: read http://stackoverflow.com/questions/34467506/how-to-cast-u8-to-usize
        sum = sum + (*x as usize);
    }
    sum
}

pub fn sum_of_amicable_numbers(under: usize) -> usize {
    let mut amicable_pairs = HashMap::new(); 
    for a in 1..under {
        let d_a = d(a);
        if a == d_a {
            continue;
        }

        // not amicable?
        let d_b = d(d_a);
        if d_b != a {
            continue;
        }

        // amicable pair
        if !amicable_pairs.contains_key(&a) && !amicable_pairs.contains_key(&d_a) {
            amicable_pairs.insert(a, d_a);
        }
    }

    let mut sum = 0;
    for (d_a, d_b) in &amicable_pairs {
        sum = sum + d_a + d_b;
    }
    sum
}

#[test]
fn it_works() {
    assert_eq!(31626, sum_of_amicable_numbers(10000));
}
