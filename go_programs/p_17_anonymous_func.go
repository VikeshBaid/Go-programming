package main

import (
	"fmt"
)

var (
	mul = func(op0, op1 int) int {
		return op0 * op1
	}

	sqr = func(op0 int) int {
		return mul(op0, op0)
	}
)

func main() {
	fmt.Printf("mul(25,7) = %d \n", mul(25, 7))
	fmt.Printf("sqr(23) = %d \n", sqr(23))
}
