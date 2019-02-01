// watch the code from video for more clearity
package main

import (
	"fmt"
)

func main() {

	x := 23
	fmt.Println("x inside main", x)
	y := foo()
	fmt.Println("x inside foo", y)
	fmt.Println("x inside main", x)
}

func foo() int {
	x := 0
	return x
}
