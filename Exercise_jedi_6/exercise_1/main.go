package main

import (
	"fmt"
)

func main() {
	x := foo()
	y, z := bar()

	fmt.Println(x, y, z)
}

func foo() int {
	return 2
}

func bar() (int, string) {
	a := 1
	b := "hello bar"
	return a, b
}
