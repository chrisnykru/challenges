/*

Prime pair sets

The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes and
concatenating them in any order the result will always be prime. For example,
taking 7 and 109, both 7109 and 1097 are prime. The sum of these four primes, 792,
represents the lowest sum for a set of four primes with this property.

Find the lowest sum for a set of five primes for which any two primes concatenate
to produce another prime.

*/
package main

import (
	"fmt"
	"github.com/jbarham/primegen.go"
	"strconv"
)

type primeLut struct {
	lut     map[uint64]bool
	highest uint64
	gen     *primegen.Primegen
}

func newPrimeLut() *primeLut {
	lut := &primeLut{
		lut: make(map[uint64]bool),
		gen: primegen.New(),
	}
	return lut
}

func (lut *primeLut) IsPrime(x uint64) bool {
	if x <= lut.highest {
		_, ok := lut.lut[x]
		return ok
	}
	// expand LUT
	var p uint64
	for p < x {
		p = lut.gen.Next()
		lut.lut[p] = false // arbitrary
	}
	lut.highest = p
	return p == x && p != 0
}

func isRemarkable(primeLut *primeLut, x, y uint64) bool {
	sxy := fmt.Sprintf("%d%d", x, y)
	cxy, err := strconv.ParseUint(sxy, 10, 64)
	if err != nil {
		panic(err)
	}
	if !primeLut.IsPrime(cxy) {
		return false
	}
	syx := fmt.Sprintf("%d%d", y, x)
	cyx, err := strconv.ParseUint(syx, 10, 64)
	if err != nil {
		panic(err)
	}
	return primeLut.IsPrime(cyx)
}

type node struct {
	val               uint64
	remarkableBuddies []*node
}

func (n *node) BuddiesWith(n2 *node) bool {
	for _, buddy := range n.remarkableBuddies {
		if buddy == n2 {
			return true
		}
	}
	return false
}

type setCandidates struct {
	sets [][]*node
}

func (sc *setCandidates) crawlNode(x *node, set []*node, targetSetSize int) {
	set = append(set, x)
	if len(set) == targetSetSize {
		sc.sets = append(sc.sets, set)
		return
	}

	for _, buddyNode := range x.remarkableBuddies {
		alreadyInSet := false
		for i := range set {
			if set[i] == buddyNode {
				alreadyInSet = true
				break
			}
		}
		if alreadyInSet {
			continue
		}
		fullyConnected := true // w.r.t. subgraph
		for _, setNode := range set {
			if !setNode.BuddiesWith(buddyNode) {
				fullyConnected = false
				break
			}
		}
		if fullyConnected {
			sc.crawlNode(buddyNode, set, targetSetSize)
		}
	}
}

func findLowestSetSum(targetSetSize int) uint64 {
	var nodes []*node
	var candidates setCandidates

	primeLut := newPrimeLut()
	gen := primegen.New()
	for {
		prime := gen.Next() // 3,5,7,11,13,17,19,...
		x := &node{val: prime}

		// connect node X to existing nodes
		for i := range nodes {
			if isRemarkable(primeLut, x.val, nodes[i].val) {
				x.remarkableBuddies = append(x.remarkableBuddies, nodes[i])
				nodes[i].remarkableBuddies = append(nodes[i].remarkableBuddies, x)
			}
		}
		nodes = append(nodes, x)

		// look for target set
		if len(x.remarkableBuddies) >= targetSetSize-1 {
			candidates.crawlNode(x, nil, targetSetSize)
		}
		if len(candidates.sets) > 0 {
			break
		}
	}

	// hack: only examine first candidate set
	sum := candidates.sets[0][0].val
	for i := 1; i < len(candidates.sets[0]); i++ {
		sum += candidates.sets[0][i].val
	}
	return sum
}

func main() {
	fmt.Println(findLowestSetSum(5))
}
