package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 5; i++ {
		for j := 4; j >= i; j-- {
			fmt.Printf(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Printf("%d",i)
		}
		fmt.Printf("\n")
	}
}
