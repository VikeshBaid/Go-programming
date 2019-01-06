package main

import (
	"fmt"
)

func main() {
	// for init; conditional; post {}
	for i := 1; i <= 5 ; i++ {
		//fmt.Printf("Outer loop: %d\n", i)
		for j := 0; j < i; j++ {
			fmt.Printf("%d", i)
			//fmt.Printf("*") if want ot print "*" pattern
			}
		fmt.Printf("\n")
	}
}
