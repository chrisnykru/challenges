/*

Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

*/

//#![feature(test)]
//#![feature(step_by)]

//extern crate test;

pub fn sum_of_multiples(max: i64) -> i64 {
  let mut sum = 0;
  for i in 1..max {
  	//println!("{}", i);
  	if i % 3 == 0 || i % 5 == 0 {
  		sum += i;
  	}
  }
  return sum;
}

/* see: https://github.com/rust-lang/rust/issues/27741
pub fn sum_of_multiples2(max: i64) -> i64 {
  let mut sum = 0;
  for i in (0..max).step_by(3) {
    sum += i
  }
  for i in (0..max).step_by(5) {
    if i % 3 != 0 {
      sum += i
    }
  }
  return sum;
}
*/

pub fn sum_of_multiples3(max: i64) -> i64 {
  let mut sum = 0;
  let mut i = 0;
  while i < max {
    sum += i;
    i += 3;
  }
  i = 0;
  while i < max {
    if i % 3 != 0 {
      sum += i;
    }
    i += 5;
  }
  return sum;
}


#[cfg(test)]
mod tests {
    use super::*;
    //use test::Bencher;

    #[test]
    fn it_works() {
      assert_eq!(23, sum_of_multiples(10));
      assert_eq!(233168, sum_of_multiples(1000));
      //assert_eq!(233168, sum_of_multiples2(1000));
      assert_eq!(233168, sum_of_multiples3(1000));
    }

    /*
    #[bench]
    fn bench_sum_of_multiples(b: &mut Bencher) {
        b.iter(|| sum_of_multiples(1000));
    }
    
    //#[bench]
    //fn bench_sum_of_multiples2(b: &mut Bencher) {
    //    b.iter(|| sum_of_multiples2(1000));
    //}
    
    #[bench]
    fn bench_sum_of_multiples3(b: &mut Bencher) {
        b.iter(|| sum_of_multiples3(1000));
    }
    */
}