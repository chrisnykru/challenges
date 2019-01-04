/*

   Smallest multiple

   2520 is the smallest number that can be divided by each of the numbers
   from 1 to 10 without any remainder.

   What is the smallest positive number that is evenly divisible by all of
   the numbers from 1 to 20?

*/

#![feature(test)]
extern crate test;

// [1, x] inclusive
pub fn solve(x: u64) -> u64 {
    let mut i = x;

    'outer: loop {
        for j in 1..x+1 {
            if i % j != 0 {
                i = i + x;
                //println!("{}", i);
                continue 'outer;
            }
        }
        break
    }
    i
}

// [1, x] inclusive
pub fn solvefaster(x: u64) -> u64 {
    /*

    2017-10-10: reduce field
    1..20 is same as 11..20

    start with highest value
    20; 20 % 19 == 0?  false
    40; 40 % 19 == 0?  false
    ...
    380; 380 % 19 == 0?  true
         380 % 18 == 0?  false
    400; 400 % 19 == 0?  false
    ...
    and so on


    Future optimization: reduce the field
    1,2,3,4,5,6,7,8,9,10
    is the same as
    6,7,8,9,10; // 5 collapses into 10, 4 to 8, 3 to 9, 2 to lots, 1, ...

    */
    
    let mut y = x;
    loop {
        let mut z = x - 1;
        let mut ok = true;
        //println!("{}", y);
        while z > 1 {
            if y % z != 0 {
                ok = false;
                break;
            }
            z = z - 1;
        }
        if ok {
            return y;
        }
        y = y + x;
    }
}


#[cfg(test)]
mod tests {
    use super::*;
    use test::Bencher;

    #[test]
    fn solve_works() {
        assert_eq!(2520, solve(10));
        assert_eq!(232792560, solve(20));
    }

    #[test]
    fn solvefaster_works() {
        assert_eq!(2520, solvefaster(10));
        assert_eq!(232792560, solvefaster(20));
    }

    #[bench]
    fn bench_solve(b: &mut Bencher) {
        b.iter(|| solve(20));
    }

    #[bench]
    fn bench_solvefaster(b: &mut Bencher) {
        b.iter(|| solvefaster(20));
    }
}
