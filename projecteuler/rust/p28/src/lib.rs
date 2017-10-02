/*
 * Number spiral diagonals
 *
 * Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral
 * is formed as follows:
 *     
 * 21 22 23 24 25
 * 20  7  8  9 10
 * 19  6  1  2 11
 * 18  5  4  3 12
 * 17 16 15 14 13
 *
 * It can be verified that the sum of the numbers on the diagonals is 101.
 *
 * What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same
 * way?
 */

pub fn solve(edge: usize) -> usize {
    // notice that up-and-right diag follows pattern: 1**2, 3**2, 5**2, 7**2, ...
    // let ring 1 = the center, which is 1
    // ring 2 is 2,3,4,5,6,7,8,9
    // etc.
    
    // edge length must be odd
    if edge % 2 == 0 {
        panic!("edge len must be odd");
    }

    let mut sum = 1; // accounts for ring 1
    let end_ring = (edge + 1) / 2;
    for ring in 2..end_ring + 1 {
        // ring 2 -> edge=3, ring 3 -> edge=5, ...
        let ring_edge = ring * 2 - 1;
        // up-and-right diag
        let mut x = ring_edge * ring_edge;
        sum += x;
        // up-and-left diag
        x -= ring_edge - 1;
        sum += x;
        // down-and-left diag
        x -= ring_edge - 1;
        sum += x;
        // down-and-right diag
        x -= ring_edge - 1;
        sum += x;
    }
    sum
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_5x5() {
        assert_eq!(solve(5), 101);
    }

    #[test]
    fn test_1001x1001() {
        assert_eq!(solve(1001), 669171001);
    }
}
