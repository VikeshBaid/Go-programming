package main

import (
	"fmt"
)

func main() {
	x := foo(bar, []int{1, 2, 3, 4, 5})
	fmt.Println(x)
}

func bar(xi []int) int {
	return xi[0] + xi[len(xi)-1]
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	return n
}
