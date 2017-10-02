/*
 * Lexicographic permutations
 *
 * A permutation is an ordered arrangement of objects. For example, 3124 is one
 * possible permutation of the digits 1, 2, 3 and 4. If all of the permutations
 * are listed numerically or alphabetically, we call it lexicographic order.
 *
 * The lexicographic permutations of 0, 1 and 2 are:
 * 012   021   102   120   201   210
 *
 * What is the millionth lexicographic permutation of the
 * digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
 */

extern crate permutohedron;
use permutohedron::{LexicalPermutation};

pub fn solve(stop_at: usize) -> Vec<usize> {
    let mut data = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];

    for _ in 1..stop_at {
        data.next_permutation();
    }
    data.iter().cloned().collect()
}

#[cfg(test)]
mod test {
    extern crate permutohedron;
    use super::*;
    use permutohedron::{LexicalPermutation};

    #[test]
    fn demo_and_vet_permutohedron_behavior() {
        let mut data = [0, 1, 2];
        data.next_permutation();
        assert_eq!(&data, &[0, 2, 1]);
        data.next_permutation();
        assert_eq!(&data, &[1, 0, 2]);
        data.next_permutation();
        assert_eq!(&data, &[1, 2, 0]);
        data.next_permutation();
        assert_eq!(&data, &[2, 0, 1]);
        data.next_permutation();
        assert_eq!(&data, &[2, 1, 0]);
    }

    #[test]
    fn it_works() {
        assert_eq!(solve(1000000), [2, 7, 8, 3, 9, 1, 5, 4, 6, 0]);
    }
}
