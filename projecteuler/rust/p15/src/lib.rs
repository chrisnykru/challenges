/*
 * Lattice paths
 *
 * Starting in the top left corner of a 2x2 grid, there are 6 routes
 * (without backtracking) to the bottom right corner.
 *
 *   @@@@@@@    @@@@```    @@@@```    @``````    @``````    @``````
 *   `  `  @    `  @  `    `  @  `    @  `  `    @  `  `    @  `  `
 *   ``````@    ```@@@@    ```@```    @@@@@@@    @@@@```    @``````
 *   `  `  @    `  `  @    `  @  `    `  `  @    `  @  `    @  `  `
 *   ``````@    ``````@    ```@```    ``````@    ```@@@@    @@@@@@@
 *
 * How many routes are there through a 20x20 grid?
 */

extern crate num;

use num::{BigUint, One};
use num::bigint::{ToBigUint};

pub fn factorial(n: usize) -> BigUint {
    let mut res: BigUint = One::one();
    // XXX wonderful, factorial(usize_max) will fail, n+1 overflows...
    for i in 2..(n+1) {
        let mul = i.to_biguint().unwrap();
        res = res * mul;
    }
    res
}

// grids are square
// no backtracking involved in routes
// consider for 2x2 grid: route r1 r2 d1 d2 == r2 r1 d1 d2 == r2 r1 d2 d1
pub fn solve(grid_size: usize) -> BigUint {
    let n = grid_size + grid_size;
    let k = grid_size;

    let nfac = factorial(n);
    let kfac = factorial(k);
    let n_minus_k_fac = factorial(n-k);
    let c = nfac / (kfac * n_minus_k_fac);
    c
}

#[test]
fn test_factorial() {
    assert_eq!(1.to_biguint().unwrap(), factorial(0));
    assert_eq!(1.to_biguint().unwrap(), factorial(1));
    assert_eq!(2.to_biguint().unwrap(), factorial(2));
    assert_eq!(6.to_biguint().unwrap(), factorial(3));
    assert_eq!(24.to_biguint().unwrap(), factorial(4));
    assert_eq!(120.to_biguint().unwrap(), factorial(5));
    assert_eq!(720.to_biguint().unwrap(), factorial(6));
    assert_eq!(BigUint::parse_bytes(b"12146304367025329675766243241881295855454217088483382315328918161829235892362167668831156960612640202170735835221294047782591091570411651472186029519906261646730733907419814952960000000000000000000000000000", 10).unwrap(), factorial(123));
}

#[test]
fn it_works() {
    assert_eq!(BigUint::parse_bytes(b"6", 10).unwrap(), solve(2));
    assert_eq!(BigUint::parse_bytes(b"137846528820", 10).unwrap(), solve(20));
}
