package main

import (
	"fmt"
)


func main() {
	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}
}
