/*
 * Coin sums
 *
 * In England the currency is made up of pound, £, and pence, p, and there are
 * eight coins in general circulation:
 * 1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
 *
 * It is possible to make £2 in the following way:
 * 1*£1 + 1*50p + 2*20p + 1*5p + 1*2p + 3*1p
 *
 * How many different ways can £2 be made using any number of coins?
 */

// dynamic programming
pub fn solve2(target: usize) -> usize {
    let coin_types = vec![1,2,5,10,20,50,100,200];
    let mut ways = vec![0; target + 1];
    ways[0] = 1; // we define it so
    // ways[2] corresponds to 2p (i.e. ways[x-1] is for x-1, not x)
    for x in &coin_types {
        for i in *x..target+1 {
            ways[i] += ways[i-x];
        }
        println!("coin={}: {:?}", x, ways);
    }
    ways[target]
}

#[test]
fn it_works() {
    assert_eq!(solve2(200), 73682);
}
