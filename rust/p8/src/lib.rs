/*
 * Largest product in a series
 *
 * Find the greatest product of five consecutive digits in the 1000-digit number.
 * (reproduced below)
 */

#![feature(test)]
extern crate test;

pub fn solve(s: String) -> Result<u64, &'static str> {
    if s.len() < 5 {
        return Err("string too short");
    }

    let mut i5: u64; // declared separately to avoid unused assignment warning
    let (mut i, mut i2, mut i3, mut i4) = (0, 0, 0, 0);

    let mut greatest = 0;
    for x in s.char_indices() {
        if x.1 < '0' || x.1 > '9' {
            return Err("invalid char in string");
        }
        i5 = i4;
        i4 = i3;
        i3 = i2;
        i2 = i;
        i = (x.1 as u64) - ('0' as u64);
        let tmp = i * i2 * i3 * i4 * i5;
        if tmp > greatest {
            greatest = tmp;
        }
    }
    Ok(greatest)
}

#[test]
fn it_works() {
    let mut s = String::from("73167176531330624919225119674426574742355349194934");
    s.push_str("96983520312774506326239578318016984801869478851843");
    s.push_str("85861560789112949495459501737958331952853208805511");
    s.push_str("12540698747158523863050715693290963295227443043557");
    s.push_str("66896648950445244523161731856403098711121722383113");
    s.push_str("62229893423380308135336276614282806444486645238749");
    s.push_str("30358907296290491560440772390713810515859307960866");
    s.push_str("70172427121883998797908792274921901699720888093776");
    s.push_str("65727333001053367881220235421809751254540594752243");
    s.push_str("52584907711670556013604839586446706324415722155397");
    s.push_str("53697817977846174064955149290862569321978468622482");
    s.push_str("83972241375657056057490261407972968652414535100474");
    s.push_str("82166370484403199890008895243450658541227588666881");
    s.push_str("16427171479924442928230863465674813919123162824586");
    s.push_str("17866458359124566529476545682848912883142607690042");
    s.push_str("24219022671055626321111109370544217506941658960408");
    s.push_str("07198403850962455444362981230987879927244284909188");
    s.push_str("84580156166097919133875499200524063689912560717606");
    s.push_str("05886116467109405077541002256983155200055935729725");
    s.push_str("71636269561882670428252483600823257530420752963450");
    assert_eq!(40824, solve(s).unwrap());
}
