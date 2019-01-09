package main

import (
	"fmt"
)

func main() {
	x := 33
	for x <= 122 {
		fmt.Printf("%d\t\t%q\n", x, x)
		x++
	}
}
