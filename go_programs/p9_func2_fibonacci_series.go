package main

import (
	"fmt"
)

func fib(n int) {

	var p0, p1 int = 0, 1
	for i := 2; i <= n; i++ {
		p0, p1 = p1, p0+p1
		fmt.Printf("%d ", p1)
	}
}

func main() {
	fib(41)
}
