/*

Coin sums

In England the currency is made up of pound, £, and pence, p, and there are
eight coins in general circulation:

  1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).

It is possible to make £2 in the following way:

  1*£1 + 1*50p + 2*20p + 1*5p + 1*2p + 3*1p

How many different ways can £2 be made using any number of coins?

*/
package main

import (
	"fmt"
)

var coinTypes []uint = []uint{1, 2, 5, 10, 20, 50, 100, 200}

type counter struct {
	ways uint64
}

// I don't like this approach... but it works.
func (c *counter) Sol(n uint, coinTypeIdx int) {
	if n == 0 {
		c.ways++
	}
	for i := coinTypeIdx; i < len(coinTypes); i++ {
		if n < coinTypes[i] {
			break
		}
		c.Sol(n-coinTypes[i], i)
	}
}

func main() {
	fmt.Println(coinTypes)

	var c counter
	c.Sol(200, 0)

	fmt.Println("ways:", c.ways)
}
