package main

import (
	"fmt"
	"math/rand"
)

var list1 = []string{"break", "lake", "go", "right", "strong", "kite", "hello"}
var list2 = []string{"fix", "river", "stop", "left", "weak", "flight", "bye"}

func main() {
	// random seeding is a way through which we can get fix result instead of using random number generators
	// the particular value of seed gives particular result
	rand.Seed(2)
	for w1, w2 := nextPair(); w1 != "go" && w2 != "stop"; w1, w2 = nextPair() {
		fmt.Printf("Word Pair -> [%s, %s]\n", w1, w2)
	}
}

func nextPair() (w1, w2 string) {
	// rand.Intn is generating non-negative integer
	pos := rand.Intn(len(list1))
	return list1[pos], list2[pos]
}
