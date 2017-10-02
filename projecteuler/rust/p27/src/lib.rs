/*
 * Quadratic primes
 *
 * Euler published the remarkable quadratic formula:
 * n^2 + n + 41
 *
 * It turns out that the formula will produce 40 primes for the consecutive values
 * n = 0 to 39. However, when n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible
 * by 41, and certainly when n = 41, 41^2 + 41 + 41 is clearly divisible by 41.
 *
 * Using computers, the incredible formula  n^2 - 79n + 1601 was discovered, which
 * produces 80 primes for the consecutive values n = 0 to 79. The product of the
 * coefficients, -79 and 1601, is -126479.
 *
 * Considering quadratics of the form:
 * n^2 + an + b, where |a|<1000 and |b|<1000
 * where |n| is the modulus/absolute value of n
 * e.g. |11| = 11 and |-4| = 4
 *
 * Find the product of the coefficients, a and b, for the quadratic expression that
 * produces the maximum number of primes for consecutive values of n,
 * starting with n = 0.
 */

extern crate primal;

// n^2 + an + b
pub fn quadratic(a: isize, b: isize, n: isize) -> isize {
    n * n + a * n + b
}

fn num_consecutive_n_primes(a: isize, b: isize) -> usize {
    let mut n = 0;
    loop {
        let x = quadratic(a, b, n);
        if x < 0 {
            // negative numbers aren't prime
            break
        }
        if !primal::is_prime(x as u64) {
            break
        }
        n += 1;
    }
    n as usize
}

// Returns a, b, # of primes
pub fn find_a_b_yielding_max_num_primes() -> (isize, isize, usize) {
    let mut best = (0,0,0);
    let mut best_count_so_far = 0;
    for a in -999..1000 {
        for b in -999..1000 {
            let count = num_consecutive_n_primes(a, b);
            if count > best_count_so_far {
                best = (a,b,count);
                best_count_so_far = count;
            }
        }
    }
    best
}

#[test]
fn test_quadratic() {
    assert_eq!(quadratic(2,3,4), 27);
}

#[test]
#[should_panic]
fn test_quadratic_overflow() {
    assert_eq!(quadratic(std::isize::MAX, 2, 3), 123);
}

#[test]
fn test_find_a_b() {
    assert_eq!(find_a_b_yielding_max_num_primes(), (-61, 971, 71));
}
