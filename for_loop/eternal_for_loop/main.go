package main

import (
	"fmt"
)

func main() {
	x := 1
	for {
		x++
		if x >= 10 {
			break
		}

		// continue statement take the iterator to the next iterator
		// it won't allow next statment(s) to execute
		if x%2 != 0 {
			continue
		}

		fmt.Println(x)
	}
	fmt.Println(x)
}
