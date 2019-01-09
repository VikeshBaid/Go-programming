package main

import (
	"fmt"
)
const rows = 5

func main() {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Printf("\n")
	}
}
