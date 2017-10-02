/*
 * Longest Collatz sequence
 *
 * The following iterative sequence is defined for the set of positive integers:
 *   n -> n/2 (n is even)
 *   n -> 3n + 1 (n is odd)
 *
 * Using the rule above and starting with 13, we generate the following sequence:
 * 13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1
 *
 * It can be seen that this sequence (starting at 13 and finishing at 1) contains 10
 * terms. Although it has not been proved yet (Collatz Problem), it is thought that
 * all starting numbers finish at 1.
 *
 * Which starting number, under one million, produces the longest chain?
 *
 * NOTE: Once the chain starts the terms are allowed to go above one million.
 */

pub struct CollatzSeq {
    n: Option<usize>,
}

impl Iterator for CollatzSeq {
    type Item = usize;

    fn next(&mut self) -> Option<usize> {
        let res = self.n;
        if self.n == None {
        } else if self.n.unwrap() == 1 {
            self.n = None;
        } else if self.n.unwrap() % 2 == 0 {
            self.n = Some(self.n.unwrap() / 2);
        } else {
            // odd
            self.n = Some(3 * self.n.unwrap() + 1);
        }
        res
    }
}

fn collatz_seq(start: usize) -> CollatzSeq {
    CollatzSeq { n: Some(start) }
}

pub fn find_longest_seq(startmax: usize) -> (usize, usize) {
    let mut maxlen: usize = 0;
    let mut maxlenstart: usize = 0;
    for i in 1..startmax {
        let seqlen = collatz_seq(i).count();
        if seqlen > maxlen {
            maxlen = seqlen;
            maxlenstart = i;
        }
    }
    return (maxlen, maxlenstart);
}

#[test]
fn test_collatz_seq() {
    let mut cs = collatz_seq(13);
    assert_eq!(13, cs.next().unwrap());
    assert_eq!(40, cs.next().unwrap());
    assert_eq!(20, cs.next().unwrap());
    assert_eq!(10, cs.next().unwrap());
    assert_eq!(5, cs.next().unwrap());
    assert_eq!(16, cs.next().unwrap());
    assert_eq!(8, cs.next().unwrap());
    assert_eq!(4, cs.next().unwrap());
    assert_eq!(2, cs.next().unwrap());
    assert_eq!(1, cs.next().unwrap());
    assert_eq!(None, cs.next());
}

#[test]
fn it_works() {
    assert_eq!((525, 837799), find_longest_seq(1000000));
}
