package main

import (
	"fmt"
)

func main() {
	n := loopFact(4)
	fmt.Println(n)
}

func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
